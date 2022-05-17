package client_test

import (
	imvclient "github.com/imversed/relayer/chain_client/imversed"
	imvprovider "github.com/imversed/relayer/relayer/provider/imversed"

	"os"
	"testing"

	"github.com/imversed/relayer/chain_client/imversed/crypto/hd"

	"go.uber.org/zap/zaptest"
)

// TestKeyRestore restores a test mnemonic
func TestKeyRestore(t *testing.T) {
	keyName := "test_key"
	mnemonic := "empower bounce moon grain plug brisk anchor breeze van submit task develop blush lottery border dish card moment drum pull glass must vicious main"
	expectedAddress := "imv1x70p5l9ntluxu6kmkj6wg44mlkzku70e8fx222"
	var coinType uint32
	coinType = 60 // Ethermint coin type used in address derivation

	homepath := t.TempDir()

	config := imvprovider.ImversedProviderConfig{
		Key:            "imversed-local-key",
		ChainID:        "imversed_1234-1",
		RPCAddr:        "http://localhost:26657",
		AccountPrefix:  "imv",
		KeyringBackend: "test",
		GasAdjustment:  1.3,
		GasPrices:      "0.001nimv",
		Debug:          true,
		Timeout:        "10s",
		OutputFormat:   "json",
		SignModeStr:    "direct",
	}

	cl, err := imvclient.NewChainClient(
		zaptest.NewLogger(t),
		imvprovider.ImvChainClientConfig(&config),
		homepath,
		os.Stdin,
		os.Stdout,
		hd.EthSecp256k1Option(),
	)

	if err != nil {
		t.Fatal(err)
	}

	_ = cl.DeleteKey(keyName) // Delete if test is being run again
	address, err := cl.RestoreKey(keyName, mnemonic, coinType)
	if err != nil {
		t.Fatalf("Error while restoring mnemonic: %v", err)
	}
	if address != expectedAddress {
		t.Fatalf("Restored address: %s does not match expected: %s", address, expectedAddress)
	}
}
