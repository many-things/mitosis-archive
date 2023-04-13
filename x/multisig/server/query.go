package server

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/exported"
	"github.com/many-things/mitosis/x/multisig/keeper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type queryServer struct {
	baseKeeper keeper.Keeper
}

func NewQueryServer(keeper keeper.Keeper) QueryServer {
	return queryServer{keeper}
}

func (k queryServer) Params(goCtx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &QueryParamsResponse{Params: k.baseKeeper.GetParams(ctx)}, nil
}

func (k queryServer) Keygen(ctx context.Context, msg *QueryKeygen) (*QueryKeygenResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	keygen, err := k.baseKeeper.QueryKeygen(wctx, msg.Chain, msg.Id)
	if err != nil {
		return nil, err
	}

	return &QueryKeygenResponse{
		Keygen: keygen,
	}, nil
}

func (k queryServer) KeygenList(ctx context.Context, msg *QueryKeygenList) (*QueryKeygenListResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	kvKeygen, page, err := k.baseKeeper.QueryKeygenList(wctx, msg.Chain, msg.Pagination)
	if err != nil {
		return nil, err
	}

	return &QueryKeygenListResponse{
		List: mitotypes.Values(kvKeygen),
		Page: page,
	}, nil
}

func (k queryServer) KeygenResult(ctx context.Context, msg *QueryKeygenResult) (*QueryKeygenResultResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	keyID := exported.KeyID(msg.KeyId)
	chainID, id, err := keyID.ToInternalVariables()
	if err != nil {
		return nil, fmt.Errorf("keyId has invalid format")
	}

	pubKey, err := k.baseKeeper.QueryKeygenResult(wctx, chainID, id)
	if err != nil {
		return nil, err
	}

	return &QueryKeygenResultResponse{
		Result: pubKey,
	}, nil
}

func (k queryServer) KeygenResultList(ctx context.Context, msg *QueryKeygenResultList) (*QueryKeygenResultListResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	kvPubkeys, page, err := k.baseKeeper.QueryKeygenResultList(wctx, msg.ChainId, msg.Pagination)
	if err != nil {
		return nil, err
	}

	return &QueryKeygenResultListResponse{
		List: mitotypes.Values(kvPubkeys),
		Page: page,
	}, err
}

func (k queryServer) Sign(ctx context.Context, msg *QuerySign) (*QuerySignResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	sign, err := k.baseKeeper.QuerySign(wctx, msg.Chain, msg.Id)
	if err != nil {
		return nil, err
	}

	return &QuerySignResponse{Sign: sign}, nil
}

func (k queryServer) SignList(ctx context.Context, msg *QuerySignList) (*QuerySignListResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	kvSigns, page, err := k.baseKeeper.QuerySignList(wctx, msg.Chain, msg.Pagination)
	if err != nil {
		return nil, err
	}

	return &QuerySignListResponse{List: mitotypes.Values(kvSigns), Page: page}, nil
}

func (k queryServer) SignResult(ctx context.Context, msg *QuerySignResult) (*QuerySignResultResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	sigID := exported.SigID(msg.SigId)
	chainID, id, err := sigID.ToInternalVariables()
	if err != nil {
		return nil, fmt.Errorf("sigId has invalid format: must be chainId-sequence")
	}

	signature, err := k.baseKeeper.QuerySignResult(wctx, chainID, id)
	if err != nil {
		return nil, err
	}

	return &QuerySignResultResponse{Signature: signature}, nil
}

func (k queryServer) SignResultList(ctx context.Context, msg *QuerySignResultList) (*QuerySignResultListResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)
	kvSignature, page, err := k.baseKeeper.QuerySignResultList(wctx, msg.ChainId, msg.Pagination)
	if err != nil {
		return nil, err
	}

	return &QuerySignResultListResponse{List: mitotypes.Values(kvSignature), Page: page}, nil
}
