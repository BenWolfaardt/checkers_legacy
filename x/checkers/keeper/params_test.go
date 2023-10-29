package keeper_test

import (
	"testing"

	testkeeper "github.com/BenWolfaardt/Checkers/testutil/keeper"
	"github.com/BenWolfaardt/Checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CheckersKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
