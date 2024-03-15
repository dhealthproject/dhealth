package dhealth_test

import (
	"testing"

	keepertest "github.com/dhealthproject/dhealth/testutil/keeper"
	"github.com/dhealthproject/dhealth/testutil/nullify"
	"github.com/dhealthproject/dhealth/x/dhealth"
	"github.com/dhealthproject/dhealth/x/dhealth/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DhealthKeeper(t)
	dhealth.InitGenesis(ctx, *k, genesisState)
	got := dhealth.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
