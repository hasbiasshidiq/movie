package movie

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"movie/testutil/sample"
	moviesimulation "movie/x/movie/simulation"
	"movie/x/movie/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = moviesimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateMovie = "op_weight_msg_movie"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMovie int = 100

	opWeightMsgUpdateMovie = "op_weight_msg_movie"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMovie int = 100

	opWeightMsgDeleteMovie = "op_weight_msg_movie"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteMovie int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	movieGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		MovieList: []types.Movie{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		MovieCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&movieGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateMovie int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateMovie, &weightMsgCreateMovie, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMovie = defaultWeightMsgCreateMovie
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMovie,
		moviesimulation.SimulateMsgCreateMovie(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateMovie int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateMovie, &weightMsgUpdateMovie, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMovie = defaultWeightMsgUpdateMovie
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateMovie,
		moviesimulation.SimulateMsgUpdateMovie(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteMovie int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteMovie, &weightMsgDeleteMovie, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMovie = defaultWeightMsgDeleteMovie
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteMovie,
		moviesimulation.SimulateMsgDeleteMovie(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateMovie,
			defaultWeightMsgCreateMovie,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				moviesimulation.SimulateMsgCreateMovie(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateMovie,
			defaultWeightMsgUpdateMovie,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				moviesimulation.SimulateMsgUpdateMovie(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteMovie,
			defaultWeightMsgDeleteMovie,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				moviesimulation.SimulateMsgDeleteMovie(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
