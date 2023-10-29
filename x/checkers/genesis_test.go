package checkers_test

import (
	"testing"

	keepertest "github.com/BenWolfaardt/Checkers/testutil/keeper"
	"github.com/BenWolfaardt/Checkers/testutil/nullify"
	"github.com/BenWolfaardt/Checkers/x/checkers"
	"github.com/BenWolfaardt/Checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CheckersKeeper(t)
	checkers.InitGenesis(ctx, *k, genesisState)
	got := checkers.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}