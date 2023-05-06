package keeper

import (
	"fmt"
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mritunjay/x/mritunjay/types"
)

func (k Keeper) SayGoodMorning(goCtx context.Context, req *types.QuerySayGoodMorningRequest) (*types.QuerySayGoodMorningResponse, error){
	if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }
    ctx := sdk.UnwrapSDKContext(goCtx)
    // TODO: Process the query
    _ = ctx
    return &types.QuerySayGoodMorningResponse{Name: fmt.Sprintf("Goodmorning %s", req.Name)}, nil

}


