package keeper

import (
	metrics "github.com/hashicorp/go-metrics"

	errorsmod "cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v8/modules/core/exported"
)

// CreateClient generates a new client identifier and invokes the associated light client module in order to
// initialize a new client. An isolated prefixed store will be reserved for this client using the generated
// client identifier. The light client module is responsible for setting any client-specific data in the store
// via the Initialize method. This includes the client state, initial consensus state and any associated
// metadata. The generated client identifier will be returned if a client was successfully initialized.
func (k *Keeper) CreateClient(ctx sdk.Context, clientType string, clientState, consensusState []byte) (string, error) {
	if clientType == exported.Localhost {
		return "", errorsmod.Wrapf(types.ErrInvalidClientType, "cannot create client of type: %s", clientType)
	}

	params := k.GetParams(ctx)
	if !params.IsAllowedClient(clientType) {
		return "", errorsmod.Wrapf(
			types.ErrInvalidClientType,
			"client state type %s is not registered in the allowlist", clientType,
		)
	}

	clientID := k.GenerateClientIdentifier(ctx, clientType)

	clientModule, found := k.router.GetRoute(clientID)
	if !found {
		return "", errorsmod.Wrap(types.ErrRouteNotFound, clientID)
	}

	if err := clientModule.Initialize(ctx, clientID, clientState, consensusState); err != nil {
		return "", err
	}

	if status := k.GetClientStatus(ctx, clientID); status != exported.Active {
		return "", errorsmod.Wrapf(types.ErrClientNotActive, "cannot create client (%s) with status %s", clientID, status)
	}

	initialHeight := clientModule.LatestHeight(ctx, clientID)
	k.Logger(ctx).Info("client created at height", "client-id", clientID, "height", initialHeight.String())

	defer telemetry.IncrCounterWithLabels(
		[]string{"ibc", "client", "create"},
		1,
		[]metrics.Label{telemetry.NewLabel(types.LabelClientType, clientType)},
	)

	emitCreateClientEvent(ctx, clientID, clientType, initialHeight)

	return clientID, nil
}

func (k *Keeper) getClientTypeAndModule(clientID string) (clientType string, clientModule exported.LightClientModule, err error) {
	clientType, _, err = types.ParseClientIdentifier(clientID)
	if err != nil {
		return clientType, clientModule, errorsmod.Wrapf(err, "unable to parse client identifier %s", clientID)
	}

	clientModule, found := k.router.GetRoute(clientID)
	if !found {
		return "", clientModule, errorsmod.Wrap(types.ErrRouteNotFound, clientID)
	}
	return clientType, clientModule, nil
}

// UpdateClient updates the consensus state and the state root from a provided header.
func (k *Keeper) UpdateClient(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) error {
	if status := k.GetClientStatus(ctx, clientID); status != exported.Active {
		return errorsmod.Wrapf(types.ErrClientNotActive, "cannot update client (%s) with status %s", clientID, status)
	}

	clientType, clientModule, err := k.getClientTypeAndModule(clientID)
	if err != nil {
		return err
	}

	if err := clientModule.VerifyClientMessage(ctx, clientID, clientMsg); err != nil {
		return err
	}

	foundMisbehaviour := clientModule.CheckForMisbehaviour(ctx, clientID, clientMsg)
	if foundMisbehaviour {
		clientModule.UpdateStateOnMisbehaviour(ctx, clientID, clientMsg)

		k.Logger(ctx).Info("client frozen due to misbehaviour", "client-id", clientID)

		defer telemetry.IncrCounterWithLabels(
			[]string{"ibc", "client", "misbehaviour"},
			1,
			[]metrics.Label{
				telemetry.NewLabel(types.LabelClientType, clientType),
				telemetry.NewLabel(types.LabelClientID, clientID),
				telemetry.NewLabel(types.LabelMsgType, "update"),
			},
		)

		emitSubmitMisbehaviourEvent(ctx, clientID, clientType)

		return nil
	}

	consensusHeights := clientModule.UpdateState(ctx, clientID, clientMsg)

	k.Logger(ctx).Info("client state updated", "client-id", clientID, "heights", consensusHeights)

	defer telemetry.IncrCounterWithLabels(
		[]string{"ibc", "client", "update"},
		1,
		[]metrics.Label{
			telemetry.NewLabel(types.LabelClientType, clientType),
			telemetry.NewLabel(types.LabelClientID, clientID),
			telemetry.NewLabel(types.LabelUpdateType, "msg"),
		},
	)

	// emitting events in the keeper emits for both begin block and handler client updates
	emitUpdateClientEvent(ctx, clientID, clientType, consensusHeights, k.cdc, clientMsg)

	return nil
}

func (k *Keeper) CheckTxUpdateClient(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) error {
	if status := k.GetClientStatus(ctx, clientID); status != exported.Active {
		return errorsmod.Wrapf(types.ErrClientNotActive, "cannot update client (%s) with status %s", clientID, status)
	}

	_, clientModule, err := k.getClientTypeAndModule(clientID)
	if err != nil {
		return err
	}

	if !ctx.IsReCheckTx() {
		if err := clientModule.VerifyClientMessage(ctx, clientID, clientMsg); err != nil {
			return err
		}
	}

	_ = clientModule.UpdateState(ctx, clientID, clientMsg)

	return nil
}

// UpgradeClient upgrades the client to a new client state if this new client was committed to
// by the old client at the specified upgrade height
func (k *Keeper) UpgradeClient(
	ctx sdk.Context,
	clientID string,
	upgradedClient, upgradedConsState, upgradeClientProof, upgradeConsensusStateProof []byte,
) error {
	if status := k.GetClientStatus(ctx, clientID); status != exported.Active {
		return errorsmod.Wrapf(types.ErrClientNotActive, "cannot upgrade client (%s) with status %s", clientID, status)
	}

	clientType, _, err := types.ParseClientIdentifier(clientID)
	if err != nil {
		return errorsmod.Wrapf(err, "unable to parse client identifier %s", clientID)
	}

	clientModule, found := k.router.GetRoute(clientID)
	if !found {
		return errorsmod.Wrap(types.ErrRouteNotFound, clientID)
	}

	if err := clientModule.VerifyUpgradeAndUpdateState(ctx, clientID, upgradedClient, upgradedConsState, upgradeClientProof, upgradeConsensusStateProof); err != nil {
		return errorsmod.Wrapf(err, "cannot upgrade client with ID %s", clientID)
	}

	latestHeight := clientModule.LatestHeight(ctx, clientID)
	k.Logger(ctx).Info("client state upgraded", "client-id", clientID, "height", latestHeight.String())

	defer telemetry.IncrCounterWithLabels(
		[]string{"ibc", "client", "upgrade"},
		1,
		[]metrics.Label{
			telemetry.NewLabel(types.LabelClientType, clientType),
			telemetry.NewLabel(types.LabelClientID, clientID),
		},
	)

	emitUpgradeClientEvent(ctx, clientID, clientType, latestHeight)

	return nil
}

// RecoverClient will invoke the light client module associated with the subject clientID requesting it to
// recover the subject client given a substitute client identifier. The light client implementation
// is responsible for validating the parameters of the substitute (ensuring they match the subject's parameters)
// as well as copying the necessary consensus states from the substitute to the subject client store.
// The substitute must be Active and the subject must not be Active.
func (k *Keeper) RecoverClient(ctx sdk.Context, subjectClientID, substituteClientID string) error {
	if status := k.GetClientStatus(ctx, subjectClientID); status == exported.Active {
		return errorsmod.Wrapf(types.ErrInvalidRecoveryClient, "cannot recover %s subject client", exported.Active)
	}

	if status := k.GetClientStatus(ctx, substituteClientID); status != exported.Active {
		return errorsmod.Wrapf(types.ErrClientNotActive, "substitute client is not %s, status is %s", exported.Active, status)
	}

	clientType, _, err := types.ParseClientIdentifier(subjectClientID)
	if err != nil {
		return errorsmod.Wrapf(types.ErrClientNotFound, "clientID (%s)", subjectClientID)
	}

	clientModule, found := k.router.GetRoute(subjectClientID)
	if !found {
		return errorsmod.Wrap(types.ErrRouteNotFound, subjectClientID)
	}

	subjectLatestHeight := clientModule.LatestHeight(ctx, subjectClientID)
	substituteLatestHeight := clientModule.LatestHeight(ctx, substituteClientID)
	if subjectLatestHeight.GTE(substituteLatestHeight) {
		return errorsmod.Wrapf(types.ErrInvalidHeight, "subject client state latest height is greater or equal to substitute client state latest height (%s >= %s)", subjectLatestHeight, substituteLatestHeight)
	}

	if err := clientModule.RecoverClient(ctx, subjectClientID, substituteClientID); err != nil {
		return err
	}

	k.Logger(ctx).Info("client recovered", "client-id", subjectClientID)

	defer telemetry.IncrCounterWithLabels(
		[]string{"ibc", "client", "update"},
		1,
		[]metrics.Label{
			telemetry.NewLabel(types.LabelClientType, clientType),
			telemetry.NewLabel(types.LabelClientID, subjectClientID),
			telemetry.NewLabel(types.LabelUpdateType, "recovery"),
		},
	)

	// emitting events in the keeper for recovering clients
	emitRecoverClientEvent(ctx, subjectClientID, clientType)

	return nil
}
