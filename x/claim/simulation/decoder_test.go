package simulation_test

import (
	"fmt"
	"testing"

	"github.com/munblockchain/mun-node/app"

	"github.com/munblockchain/mun-node/x/claim/simulation"
	"github.com/munblockchain/mun-node/x/claim/types"

	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/stretchr/testify/require"
)

func TestDecodeStore(t *testing.T) {
	cdc := app.MakeEncodingConfig().Marshaler
	dec := simulation.NewDecodeStore(cdc)

	merkleRoot := "558ad18828f6da6d471cdb1a3443f039a770e03617f163896980d914d643e4bc"
	claimRecord := types.ClaimRecord{
		Address:         "mun1vk2zrdsg74vrztr5pe5qa0h6ljgmfy4sz2apec",
		ActionReady:     make([]bool, 4),
		ActionCompleted: make([]bool, 4),
	}

	kvParis := kv.Pairs{
		Pairs: []kv.Pair{
			{Key: types.ClaimRecordsStorePrefix, Value: cdc.MustMarshal(&claimRecord)},
			{Key: types.MerkleRootStorePrefix, Value: []byte(merkleRoot)},
		},
	}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"ClaimRecord", fmt.Sprintf("%v\n%v", claimRecord, claimRecord)},
		{"MerkleRoot", fmt.Sprintf("%v\n%v", merkleRoot, merkleRoot)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { dec(kvParis.Pairs[i], kvParis.Pairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, dec(kvParis.Pairs[i], kvParis.Pairs[i]), tt.name)
			}
		})
	}
}
