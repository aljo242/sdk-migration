package keeper_test

import (
	"testing"

	testkeeper "github.com/aljo242/sdk-migration/testutil/keeper"
	"github.com/aljo242/sdk-migration/x/sdkmigration/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SdkmigrationKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
