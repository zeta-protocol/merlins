package wasmbinding

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/merlins-labs/merlins/v15/app"
)

func CreateTestInput() (*app.MerlinsApp, sdk.Context) {
	merlins := app.Setup(false)
	ctx := merlins.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "merlins-1", Time: time.Now().UTC()})
	return merlins, ctx
}

func FundAccount(t *testing.T, ctx sdk.Context, merlins *app.MerlinsApp, acct sdk.AccAddress) {
	err := simapp.FundAccount(merlins.BankKeeper, ctx, acct, sdk.NewCoins(
		sdk.NewCoin("ufury", sdk.NewInt(10000000000)),
	))
	require.NoError(t, err)
}

// we need to make this deterministic (same every test run), as content might affect gas costs
func keyPubAddr() (crypto.PrivKey, crypto.PubKey, sdk.AccAddress) {
	key := ed25519.GenPrivKey()
	pub := key.PubKey()
	addr := sdk.AccAddress(pub.Address())
	return key, pub, addr
}

func RandomAccountAddress() sdk.AccAddress {
	_, _, addr := keyPubAddr()
	return addr
}

func RandomBech32AccountAddress() string {
	return RandomAccountAddress().String()
}
