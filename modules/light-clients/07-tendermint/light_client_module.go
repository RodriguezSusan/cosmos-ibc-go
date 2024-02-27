package tendermint

import (
	errorsmod "cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	ibcerrors "github.com/cosmos/ibc-go/v8/modules/core/errors"
	"github.com/cosmos/ibc-go/v8/modules/core/exported"
	"github.com/cosmos/ibc-go/v8/modules/light-clients/07-tendermint/internal/keeper"
)

var _ exported.LightClientModule = (*LightClientModule)(nil)

type LightClientModule struct {
	keeper        keeper.Keeper
	storeProvider exported.ClientStoreProvider
}

func NewLightClientModule(cdc codec.BinaryCodec, authority string) LightClientModule {
	return LightClientModule{
		keeper: keeper.NewKeeper(cdc, authority),
	}
}

func (lcm *LightClientModule) RegisterStoreProvider(storeProvider exported.ClientStoreProvider) {
	lcm.storeProvider = storeProvider
}

// Initialize checks that the initial consensus state is an 07-tendermint consensus state and
// sets the client state, consensus state and associated metadata in the provided client store.
func (lcm LightClientModule) Initialize(ctx sdk.Context, clientID string, clientStateBz, consensusStateBz []byte) error {
	clientType, _, err := clienttypes.ParseClientIdentifier(clientID)
	if err != nil {
		return err
	}

	if clientType != exported.Tendermint {
		return errorsmod.Wrapf(clienttypes.ErrInvalidClientType, "expected: %s, got: %s", exported.Tendermint, clientType)
	}

	var clientState ClientState
	if err := lcm.keeper.Codec().Unmarshal(clientStateBz, &clientState); err != nil {
		return err
	}

	if err := clientState.Validate(); err != nil {
		return err
	}

	var consensusState ConsensusState
	if err := lcm.keeper.Codec().Unmarshal(consensusStateBz, &consensusState); err != nil {
		return err
	}

	if err := consensusState.ValidateBasic(); err != nil {
		return err
	}

	clientStore := lcm.storeProvider.ClientStore(ctx, clientID)

	return clientState.Initialize(ctx, lcm.keeper.Codec(), clientStore, &consensusState)
}

func (lcm LightClientModule) VerifyClientMessage(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) error {
	clientType, _, err := clienttypes.ParseClientIdentifier(clientID)
	if err != nil {
		return err
	}

	if clientType != exported.Tendermint {
		return errorsmod.Wrapf(clienttypes.ErrInvalidClientType, "expected: %s, got: %s", exported.Tendermint, clientType)
	}

	clientStore := lcm.storeProvider.ClientStore(ctx, clientID)
	cdc := lcm.keeper.Codec()

	clientState, found := getClientState(clientStore, cdc)
	if !found {
		return errorsmod.Wrap(clienttypes.ErrClientNotFound, clientID)
	}

	return clientState.VerifyClientMessage(ctx, cdc, clientStore, clientMsg)
}

func (lcm LightClientModule) CheckForMisbehaviour(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) bool {
	clientType, _, err := clienttypes.ParseClientIdentifier(clientID)
	if err != nil {
		panic(err)
	}

	if clientType != exported.Tendermint {
		panic(errorsmod.Wrapf(clienttypes.ErrInvalidClientType, "expected: %s, got: %s", exported.Tendermint, clientType))
	}

	clientStore := lcm.storeProvider.ClientStore(ctx, clientID)
	cdc := lcm.keeper.Codec()

	clientState, found := getClientState(clientStore, cdc)
	if !found {
		panic(errorsmod.Wrap(clienttypes.ErrClientNotFound, clientID))
	}

	return clientState.CheckForMisbehaviour(ctx, cdc, clientStore, clientMsg)
}

func (lcm LightClientModule) UpdateStateOnMisbehaviour(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) {
	clientType, _, err := clienttypes.ParseClientIdentifier(clientID)
	if err != nil {
		panic(err)
	}

	if clientType != exported.Tendermint {
		panic(errorsmod.Wrapf(clienttypes.ErrInvalidClientType, "expected: %s, got: %s", exported.Tendermint, clientType))
	}

	clientStore := lcm.storeProvider.ClientStore(ctx, clientID)
	cdc := lcm.keeper.Codec()

	clientState, found := getClientState(clientStore, cdc)
	if !found {
		panic(errorsmod.Wrap(clienttypes.ErrClientNotFound, clientID))
	}

	clientState.UpdateStateOnMisbehaviour(ctx, cdc, clientStore, clientMsg)
}

func (lcm LightClientModule) UpdateState(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) []exported.Height {
	clientType, _, err := clienttypes.ParseClientIdentifier(clientID)
	if err != nil {
		panic(err)
	}

	if clientType != exported.Tendermint {
		panic(errorsmod.Wrapf(clienttypes.ErrInvalidClientType, "expected: %s, got: %s", exported.Tendermint, clientType))
	}
	clientStore := lcm.storeProvider.ClientStore(ctx, clientID)
	cdc := lcm.keeper.Codec()

	clientState, found := getClientState(clientStore, cdc)
	if !found {
		panic(errorsmod.Wrap(clienttypes.ErrClientNotFound, clientID))
	}

	return clientState.UpdateState(ctx, cdc, clientStore, clientMsg)
}

func (lcm LightClientModule) VerifyMembership(
	ctx sdk.Context,
	clientID string,
	height exported.Height,
	delayTimePeriod uint64,
	delayBlockPeriod uint64,
	proof []byte,
	path exported.Path,
	value []byte,
) error {
	clientType, _, err := clienttypes.ParseClientIdentifier(clientID)
	if err != nil {
		return err
	}

	if clientType != exported.Tendermint {
		return errorsmod.Wrapf(clienttypes.ErrInvalidClientType, "expected: %s, got: %s", exported.Tendermint, clientType)
	}

	clientStore := lcm.storeProvider.ClientStore(ctx, clientID)
	cdc := lcm.keeper.Codec()

	clientState, found := getClientState(clientStore, cdc)
	if !found {
		return errorsmod.Wrap(clienttypes.ErrClientNotFound, clientID)
	}

	return clientState.VerifyMembership(ctx, clientStore, cdc, height, delayTimePeriod, delayBlockPeriod, proof, path, value)
}

func (lcm LightClientModule) VerifyNonMembership(
	ctx sdk.Context,
	clientID string,
	height exported.Height, // TODO: change to concrete type
	delayTimePeriod uint64,
	delayBlockPeriod uint64,
	proof []byte,
	path exported.Path, // TODO: change to conrete type
) error {
	clientType, _, err := clienttypes.ParseClientIdentifier(clientID)
	if err != nil {
		return err
	}

	if clientType != exported.Tendermint {
		return errorsmod.Wrapf(clienttypes.ErrInvalidClientType, "expected: %s, got: %s", exported.Tendermint, clientType)
	}

	clientStore := lcm.storeProvider.ClientStore(ctx, clientID)
	cdc := lcm.keeper.Codec()

	clientState, found := getClientState(clientStore, cdc)
	if !found {
		return errorsmod.Wrap(clienttypes.ErrClientNotFound, clientID)
	}

	return clientState.VerifyNonMembership(ctx, clientStore, cdc, height, delayTimePeriod, delayBlockPeriod, proof, path)
}

func (lcm LightClientModule) Status(ctx sdk.Context, clientID string) exported.Status {
	clientType, _, err := clienttypes.ParseClientIdentifier(clientID)
	if err != nil {
		return exported.Unknown
	}

	if clientType != exported.Tendermint {
		return exported.Unknown
	}

	clientStore := lcm.storeProvider.ClientStore(ctx, clientID)
	cdc := lcm.keeper.Codec()

	clientState, found := getClientState(clientStore, cdc)
	if !found {
		return exported.Unknown
	}

	return clientState.Status(ctx, clientStore, cdc)
}

func (lcm LightClientModule) TimestampAtHeight(
	ctx sdk.Context,
	clientID string,
	height exported.Height,
) (uint64, error) {
	clientType, _, err := clienttypes.ParseClientIdentifier(clientID)
	if err != nil {
		return 0, err
	}

	if clientType != exported.Tendermint {
		return 0, errorsmod.Wrapf(clienttypes.ErrInvalidClientType, "expected: %s, got: %s", exported.Tendermint, clientType)
	}

	clientStore := lcm.storeProvider.ClientStore(ctx, clientID)
	cdc := lcm.keeper.Codec()

	clientState, found := getClientState(clientStore, cdc)
	if !found {
		return 0, errorsmod.Wrap(clienttypes.ErrClientNotFound, clientID)
	}

	return clientState.GetTimestampAtHeight(ctx, clientStore, cdc, height)
}

func (lcm LightClientModule) RecoverClient(ctx sdk.Context, clientID, substituteClientID string) error {
	clientType, _, err := clienttypes.ParseClientIdentifier(clientID)
	if err != nil {
		return err
	}

	if clientType != exported.Tendermint {
		return errorsmod.Wrapf(clienttypes.ErrInvalidClientType, "expected: %s, got: %s", exported.Tendermint, clientType)
	}

	substituteClientType, _, err := clienttypes.ParseClientIdentifier(substituteClientID)
	if err != nil {
		return err
	}

	if substituteClientType != exported.Tendermint {
		return errorsmod.Wrapf(clienttypes.ErrInvalidClientType, "expected: %s, got: %s", exported.Tendermint, substituteClientType)
	}

	cdc := lcm.keeper.Codec()

	clientStore := lcm.storeProvider.ClientStore(ctx, clientID)
	clientState, found := getClientState(clientStore, cdc)
	if !found {
		return errorsmod.Wrap(clienttypes.ErrClientNotFound, clientID)
	}

	substituteClientStore := lcm.storeProvider.ClientStore(ctx, substituteClientID)
	substituteClient, found := getClientState(substituteClientStore, cdc)
	if !found {
		return errorsmod.Wrap(clienttypes.ErrClientNotFound, substituteClientID)
	}

	if clientState.GetLatestHeight().GTE(substituteClient.GetLatestHeight()) {
		return errorsmod.Wrapf(clienttypes.ErrInvalidHeight, "subject client state latest height is greater or equal to substitute client state latest height (%s >= %s)", clientState.GetLatestHeight(), substituteClient.GetLatestHeight())
	}

	return clientState.CheckSubstituteAndUpdateState(ctx, cdc, clientStore, substituteClientStore, substituteClient)
}

// VerifyUpgradeAndUpdateState checks if the upgraded client has been committed by the current client
// It will zero out all client-specific fields (e.g. TrustingPeriod) and verify all data
// in client state that must be the same across all valid Tendermint clients for the new chain.
// VerifyUpgrade will return an error if:
// - the upgradedClient is not a Tendermint ClientState
// - the latest height of the client state does not have the same revision number or has a greater
// height than the committed client.
//   - the height of upgraded client is not greater than that of current client
//   - the latest height of the new client does not match or is greater than the height in committed client
//   - any Tendermint chain specified parameter in upgraded client such as ChainID, UnbondingPeriod,
//     and ProofSpecs do not match parameters set by committed client
func (lcm LightClientModule) VerifyUpgradeAndUpdateState(
	ctx sdk.Context,
	clientID string,
	newClient []byte,
	newConsState []byte,
	upgradeClientProof,
	upgradeConsensusStateProof []byte,
) error {
	var (
		cdc               = lcm.keeper.Codec()
		newClientState    exported.ClientState
		newConsensusState exported.ConsensusState
	)

	if err := cdc.UnmarshalInterface(newClient, &newClientState); err != nil {
		return err
	}
	newTmClientState, ok := newClientState.(*ClientState)
	if !ok {
		return errorsmod.Wrapf(clienttypes.ErrInvalidClient, "expected client state type %T, got %T", (*ClientState)(nil), newClientState)
	}

	if err := cdc.UnmarshalInterface(newConsState, &newConsensusState); err != nil {
		return err
	}
	newTmConsensusState, ok := newConsensusState.(*ConsensusState)
	if !ok {
		return errorsmod.Wrapf(clienttypes.ErrInvalidConsensus, "expected consensus state type %T, got %T", (*ConsensusState)(nil), newConsensusState)
	}

	clientStore := lcm.storeProvider.ClientStore(ctx, clientID)
	clientState, found := getClientState(clientStore, cdc)
	if !found {
		return errorsmod.Wrap(clienttypes.ErrClientNotFound, clientID)
	}

	// last height of current counterparty chain must be client's latest height
	lastHeight := clientState.GetLatestHeight()
	if !newClientState.GetLatestHeight().GT(lastHeight) {
		return errorsmod.Wrapf(ibcerrors.ErrInvalidHeight, "upgraded client height %s must be at greater than current client height %s", newClientState.GetLatestHeight(), lastHeight)
	}

	return clientState.VerifyUpgradeAndUpdateState(ctx, cdc, clientStore, newTmClientState, newTmConsensusState, upgradeClientProof, upgradeConsensusStateProof)
}
