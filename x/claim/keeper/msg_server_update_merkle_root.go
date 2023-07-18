package keeper

import (
	"context"

	"github.com/munblockchain/mun-node/x/claim/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	AdminAddress = "mun173na0dpu7kj79na4ctgg7s58ws60n4w0gd2vkd"
)

func (k msgServer) UpdateMerkleRoot(goCtx context.Context, msg *types.MsgUpdateMerkleRoot) (*types.MsgUpdateMerkleRootResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if AdminAddress != msg.Sender {
		return nil, types.ErrInvalidAdminAddress
	}

	k.Keeper.SetMerkleRoot(ctx, msg.RootValue)

	return &types.MsgUpdateMerkleRootResponse{}, nil
}
