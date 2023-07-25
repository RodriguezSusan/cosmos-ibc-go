package client_test

import (
	"strings"
	"testing"

	testifysuite "github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	abci "github.com/cometbft/cometbft/abci/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	client "github.com/cosmos/ibc-go/v7/modules/core/02-client"
	"github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	ibctm "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
	ibctesting "github.com/cosmos/ibc-go/v7/testing"
)

type ClientTestSuite struct {
	testifysuite.Suite

	coordinator *ibctesting.Coordinator

	chainA *ibctesting.TestChain
	chainB *ibctesting.TestChain
}

func (s *ClientTestSuite) SetupTest() {
	s.coordinator = ibctesting.NewCoordinator(s.T(), 2)

	s.chainA = s.coordinator.GetChain(ibctesting.GetChainID(1))
	s.chainB = s.coordinator.GetChain(ibctesting.GetChainID(2))
}

func TestClientTestSuite(t *testing.T) {
	testifysuite.Run(t, new(ClientTestSuite))
}

func (s *ClientTestSuite) TestBeginBlocker() {
	for i := 0; i < 10; i++ {
		// increment height
		s.coordinator.CommitBlock(s.chainA, s.chainB)

		s.Require().NotPanics(func() {
			client.BeginBlocker(s.chainA.GetContext(), s.chainA.App.GetIBCKeeper().ClientKeeper)
		}, "BeginBlocker shouldn't panic")
	}
}

func (s *ClientTestSuite) TestBeginBlockerConsensusState() {
	plan := &upgradetypes.Plan{
		Name:   "test",
		Height: s.chainA.GetContext().BlockHeight() + 1,
	}
	// set upgrade plan in the upgrade store
	store := s.chainA.GetContext().KVStore(s.chainA.GetSimApp().GetKey(upgradetypes.StoreKey))
	bz := s.chainA.App.AppCodec().MustMarshal(plan)
	store.Set(upgradetypes.PlanKey(), bz)

	nextValsHash := []byte("nextValsHash")
	newCtx := s.chainA.GetContext().WithBlockHeader(tmproto.Header{
		ChainID:            s.chainA.ChainID,
		Height:             s.chainA.GetContext().BlockHeight(),
		NextValidatorsHash: nextValsHash,
	})

	err := s.chainA.GetSimApp().UpgradeKeeper.SetUpgradedClient(newCtx, plan.Height, []byte("client state"))
	s.Require().NoError(err)

	req := abci.RequestBeginBlock{Header: newCtx.BlockHeader()}
	s.chainA.App.BeginBlock(req)

	// plan Height is at ctx.BlockHeight+1
	consState, found := s.chainA.GetSimApp().UpgradeKeeper.GetUpgradedConsensusState(newCtx, plan.Height)
	s.Require().True(found)
	bz, err = types.MarshalConsensusState(s.chainA.App.AppCodec(), &ibctm.ConsensusState{Timestamp: newCtx.BlockTime(), NextValidatorsHash: nextValsHash})
	s.Require().NoError(err)
	s.Require().Equal(bz, consState)
}

func (s *ClientTestSuite) TestBeginBlockerUpgradeEvents() {
	plan := &upgradetypes.Plan{
		Name:   "test",
		Height: s.chainA.GetContext().BlockHeight() + 1,
	}
	// set upgrade plan in the upgrade store
	store := s.chainA.GetContext().KVStore(s.chainA.GetSimApp().GetKey(upgradetypes.StoreKey))
	bz := s.chainA.App.AppCodec().MustMarshal(plan)
	store.Set(upgradetypes.PlanKey(), bz)

	nextValsHash := []byte("nextValsHash")
	newCtx := s.chainA.GetContext().WithBlockHeader(tmproto.Header{
		Height:             s.chainA.GetContext().BlockHeight(),
		NextValidatorsHash: nextValsHash,
	})

	err := s.chainA.GetSimApp().UpgradeKeeper.SetUpgradedClient(newCtx, plan.Height, []byte("client state"))
	s.Require().NoError(err)

	cacheCtx, writeCache := s.chainA.GetContext().CacheContext()

	client.BeginBlocker(cacheCtx, s.chainA.App.GetIBCKeeper().ClientKeeper)
	writeCache()

	s.requireContainsEvent(cacheCtx.EventManager().Events(), types.EventTypeUpgradeChain, true)
}

func (s *ClientTestSuite) TestBeginBlockerUpgradeEventsAbsence() {
	cacheCtx, writeCache := s.chainA.GetContext().CacheContext()
	client.BeginBlocker(s.chainA.GetContext(), s.chainA.App.GetIBCKeeper().ClientKeeper)
	writeCache()
	s.requireContainsEvent(cacheCtx.EventManager().Events(), types.EventTypeUpgradeChain, false)
}

// requireContainsEvent verifies if an event of a specific type was emitted.
func (s *ClientTestSuite) requireContainsEvent(events sdk.Events, eventType string, shouldContain bool) {
	found := false
	var eventTypes []string
	for _, e := range events {
		eventTypes = append(eventTypes, e.Type)
		if e.Type == eventType {
			found = true
			break
		}
	}
	if shouldContain {
		s.Require().True(found, "event type %s was not found in %s", eventType, strings.Join(eventTypes, ","))
	} else {
		s.Require().False(found, "event type %s was found in %s", eventType, strings.Join(eventTypes, ","))
	}
}
