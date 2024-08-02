package keeper

import (
	"github.com/0glabs/0g-chain/x/cdp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Implements StakingHooks interface
var _ types.CDPHooks = Keeper{}

// AfterCDPCreated - call hook if registered
func (k Keeper) AfterCDPCreated(ctx sdk.Context, cdp types.CDP) {
	if k.hooks != nil {
		k.hooks.AfterCDPCreated(ctx, cdp)
	}
}

// BeforeCDPModified - call hook if registered
func (k Keeper) BeforeCDPModified(ctx sdk.Context, cdp types.CDP) {
	if k.hooks != nil {
		k.hooks.BeforeCDPModified(ctx, cdp)
	}
}
