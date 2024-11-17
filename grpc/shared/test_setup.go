package shared

import (
	"crypto/ed25519"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/stretchr/testify/require"
	"testing"
)

func SetupSharedService(t *testing.T, noOfBlocksToGenerate int) (*eth.Ethereum, *SharedServiceContainer, ed25519.PrivateKey) {
	t.Helper()
	genesis, blocks, bridgeAddress, feeCollectorKey, trustedBuilderPrivkey := GenerateMergeChain(noOfBlocksToGenerate, true)
	ethservice := StartEthService(t, genesis)

	sharedService, err := NewSharedServiceContainer(ethservice)
	require.Nil(t, err, "can't create shared service")

	feeCollector := crypto.PubkeyToAddress(feeCollectorKey.PublicKey)
	require.Equal(t, feeCollector, sharedService.NextFeeRecipient(), "nextFeeRecipient not set correctly")

	bridgeAsset := genesis.Config.AstriaBridgeAddressConfigs[0].AssetDenom
	_, ok := sharedService.BridgeAllowedAssets()[bridgeAsset]
	require.True(t, ok, "bridgeAllowedAssetIDs does not contain bridge asset id")

	_, ok = sharedService.BridgeAddresses()[bridgeAddress]
	require.True(t, ok, "bridgeAddress not set correctly")

	_, err = ethservice.BlockChain().InsertChain(blocks)
	require.Nil(t, err, "can't insert blocks")

	// FIXME - this interface isn't right for the tests, we shouldn't be exposing the trusted builder priv key like this
	// we should instead allow the test to create it and pass it to the shared service container in the constructor
	// but that can make the codebase a bit weird, so we can leave it like this for now
	return ethservice, sharedService, trustedBuilderPrivkey
}
