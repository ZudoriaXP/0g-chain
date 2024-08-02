package e2e_test

import (
	"context"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ibctypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	emtypes "github.com/evmos/ethermint/types"

	"github.com/0glabs/0g-chain/app"
	"github.com/0glabs/0g-chain/tests/e2e/testutil"
	"github.com/0glabs/0g-chain/tests/util"
)

var (
	minEvmGasPrice = big.NewInt(1e10) // neuron
)

func a0gi(amt *big.Int) sdk.Coin {
	return sdk.NewCoin("ua0gi", sdkmath.NewIntFromBigInt(amt))
}

type IntegrationTestSuite struct {
	testutil.E2eTestSuite
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

// example test that queries 0gchain via SDK and EVM
func (suite *IntegrationTestSuite) TestChainID() {
	expectedEvmNetworkId, err := emtypes.ParseChainID(suite.ZgChain.ChainID)
	suite.NoError(err)

	// EVM query
	evmNetworkId, err := suite.ZgChain.EvmClient.NetworkID(context.Background())
	suite.NoError(err)
	suite.Equal(expectedEvmNetworkId, evmNetworkId)

	// SDK query
	nodeInfo, err := suite.ZgChain.Tm.GetNodeInfo(context.Background(), &tmservice.GetNodeInfoRequest{})
	suite.NoError(err)
	suite.Equal(suite.ZgChain.ChainID, nodeInfo.DefaultNodeInfo.Network)
}

// example test that funds a new account & queries its balance
func (suite *IntegrationTestSuite) TestFundedAccount() {
	funds := a0gi(big.NewInt(1e3))
	acc := suite.ZgChain.NewFundedAccount("example-acc", sdk.NewCoins(funds))

	// check that the sdk & evm signers are for the same account
	suite.Equal(acc.SdkAddress.String(), util.EvmToSdkAddress(acc.EvmAddress).String())
	suite.Equal(acc.EvmAddress.Hex(), util.SdkToEvmAddress(acc.SdkAddress).Hex())

	// check balance via SDK query
	res, err := suite.ZgChain.Bank.Balance(context.Background(), banktypes.NewQueryBalanceRequest(
		acc.SdkAddress, "ua0gi",
	))
	suite.NoError(err)
	suite.Equal(funds, *res.Balance)

	// check balance via EVM query
	neuronBal, err := suite.ZgChain.EvmClient.BalanceAt(context.Background(), acc.EvmAddress, nil)
	suite.NoError(err)
	suite.Equal(funds.Amount.MulRaw(1e12).BigInt(), neuronBal)
}

// example test that signs & broadcasts an EVM tx
func (suite *IntegrationTestSuite) TestTransferOverEVM() {
	// fund an account that can perform the transfer
	initialFunds := a0gi(big.NewInt(1e6)) // 1 A0GI
	acc := suite.ZgChain.NewFundedAccount("evm-test-transfer", sdk.NewCoins(initialFunds))

	// get a rando account to send 0gchain to
	randomAddr := app.RandomAddress()
	to := util.SdkToEvmAddress(randomAddr)

	// example fetching of nonce (account sequence)
	nonce, err := suite.ZgChain.EvmClient.PendingNonceAt(context.Background(), acc.EvmAddress)
	suite.NoError(err)
	suite.Equal(uint64(0), nonce) // sanity check. the account should have no prior txs

	// transfer a0gi over EVM
	a0giToTransfer := big.NewInt(1e17) // .1 A0GI; neuron has 18 decimals.
	req := util.EvmTxRequest{
		Tx:   ethtypes.NewTransaction(nonce, to, a0giToTransfer, 1e5, minEvmGasPrice, nil),
		Data: "any ol' data to track this through the system",
	}
	res := acc.SignAndBroadcastEvmTx(req)
	suite.Require().NoError(res.Err)
	suite.Equal(ethtypes.ReceiptStatusSuccessful, res.Receipt.Status)

	// evm txs refund unused gas. so to know the expected balance we need to know how much gas was used.
	a0giUsedForGas := sdkmath.NewIntFromBigInt(minEvmGasPrice).
		Mul(sdkmath.NewIntFromUint64(res.Receipt.GasUsed)).
		QuoRaw(1e12) // convert neuron to a0gi

	// expect (9 - gas used) A0GI remaining in account.
	balance := suite.ZgChain.QuerySdkForBalances(acc.SdkAddress)
	suite.Equal(sdkmath.NewInt(9e5).Sub(a0giUsedForGas), balance.AmountOf("ua0gi"))
}

// TestIbcTransfer transfers A0GI from the primary 0g-chain (suite.ZgChain) to the ibc chain (suite.Ibc).
// Note that because the IBC chain also runs 0g-chain's binary, this tests both the sending & receiving.
func (suite *IntegrationTestSuite) TestIbcTransfer() {
	suite.SkipIfIbcDisabled()

	// ARRANGE
	// setup 0g-chain account
	funds := a0gi(big.NewInt(1e5)) // .1 A0GI
	zgChainAcc := suite.ZgChain.NewFundedAccount("ibc-transfer-0g-side", sdk.NewCoins(funds))
	// setup ibc account
	ibcAcc := suite.Ibc.NewFundedAccount("ibc-transfer-ibc-side", sdk.NewCoins())

	gasLimit := int64(2e5)
	fee := a0gi(big.NewInt(200))

	fundsToSend := a0gi(big.NewInt(5e4)) // .005 A0GI
	transferMsg := ibctypes.NewMsgTransfer(
		testutil.IbcPort,
		testutil.IbcChannel,
		fundsToSend,
		zgChainAcc.SdkAddress.String(),
		ibcAcc.SdkAddress.String(),
		ibcclienttypes.NewHeight(0, 0), // timeout height disabled when 0
		uint64(time.Now().Add(30*time.Second).UnixNano()),
		"",
	)
	// initial - sent - fee
	expectedSrcBalance := funds.Sub(fundsToSend).Sub(fee)

	// ACT
	// IBC transfer from 0g-chain -> ibc
	transferTo := util.ZgChainMsgRequest{
		Msgs:      []sdk.Msg{transferMsg},
		GasLimit:  uint64(gasLimit),
		FeeAmount: sdk.NewCoins(fee),
		Memo:      "sent from ZgChain!",
	}
	res := zgChainAcc.SignAndBroadcastZgChainTx(transferTo)

	// ASSERT
	suite.NoError(res.Err)

	// the balance should be deducted from 0g-chain account
	suite.Eventually(func() bool {
		balance := suite.ZgChain.QuerySdkForBalances(zgChainAcc.SdkAddress)
		return balance.AmountOf("ua0gi").Equal(expectedSrcBalance.Amount)
	}, 10*time.Second, 1*time.Second)

	// expect the balance to be transferred to the ibc chain!
	suite.Eventually(func() bool {
		balance := suite.Ibc.QuerySdkForBalances(ibcAcc.SdkAddress)
		found := false
		for _, c := range balance {
			// find the ibc denom coin
			if strings.HasPrefix(c.Denom, "ibc/") {
				suite.Equal(fundsToSend.Amount, c.Amount)
				found = true
			}
		}
		return found
	}, 15*time.Second, 1*time.Second)
}
