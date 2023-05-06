package mritunjay_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "mritunjay/testutil/keeper"
	"mritunjay/testutil/nullify"
	"mritunjay/x/mritunjay"
	"mritunjay/x/mritunjay/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MritunjayKeeper(t)
	mritunjay.InitGenesis(ctx, *k, genesisState)
	got := mritunjay.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
