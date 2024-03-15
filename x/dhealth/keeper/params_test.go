package keeper_test

import (
	"testing"

	testkeeper "github.com/dhealthproject/dhealth/testutil/keeper"
	"github.com/dhealthproject/dhealth/x/dhealth/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DhealthKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
