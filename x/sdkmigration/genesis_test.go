package sdkmigration_test

import (
	"testing"

	keepertest "github.com/aljo242/sdk-migration/testutil/keeper"
	"github.com/aljo242/sdk-migration/testutil/nullify"
	"github.com/aljo242/sdk-migration/x/sdkmigration"
	"github.com/aljo242/sdk-migration/x/sdkmigration/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SdkmigrationKeeper(t)
	sdkmigration.InitGenesis(ctx, *k, genesisState)
	got := sdkmigration.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
