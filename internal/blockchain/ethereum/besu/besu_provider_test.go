package besu

import (
	"context"
	"testing"

	"github.com/hyperledger/firefly-cli/pkg/types"
	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/stretchr/testify/assert"
)

func TestNewBesuProvider(t *testing.T) {
	var ctx context.Context
	testCases := []struct {
		Name  string
		Ctx   context.Context
		Stack *types.Stack
	}{
		{
			Name: "testcase1",
			Ctx:  ctx,
			Stack: &types.Stack{
				Name:                   "TestBesuProviderWithEthconnect",
				Members:                []*types.Organization{{OrgName: "Org1"}},
				BlockchainProvider:     fftypes.FFEnumValue("BlockchainProvider", "Ethereum"),
				BlockchainConnector:    fftypes.FFEnumValue("BlockchainConnector", "EthConnect"),
				BlockchainNodeProvider: fftypes.FFEnumValue("BlockchainNodeProvider", "besu"),
			},
		},
		{
			Name: "testcase2",
			Ctx:  ctx,
			Stack: &types.Stack{
				Members:                []*types.Organization{{OrgName: "Org2"}, {OrgName: "org4"}},
				Name:                   "TestBesuProviderWithEvmconnect",
				BlockchainProvider:     fftypes.FFEnumValue("BlockchainProvider", "Ethereum"),
				BlockchainConnector:    fftypes.FFEnumValue("BlockchainConnector", "EvmConnect"),
				BlockchainNodeProvider: fftypes.FFEnumValue("BlockchainNodeProvider", "besu"),
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			besuProvider := NewBesuProvider(tc.Ctx, tc.Stack)
			assert.NotNil(t, besuProvider)
		})
	}
}
