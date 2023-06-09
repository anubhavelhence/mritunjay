package keeper


import (
	"fmt"
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mritunjay/x/mritunjay/types"
)

func (k Keeper) SayBye(goCtx context.Context, req *types.QuerySayByeRequest) (*types.QuerySayByeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QuerySayByeResponse{Name: fmt.Sprintf("Bye, %s!",req.Name)}, nil
}
