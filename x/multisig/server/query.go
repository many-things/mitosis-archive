package server

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/multisig/keeper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type queryServer struct {
	keeper.Keeper
}

func NewQueryServer(keeper keeper.Keeper) QueryServer {
	return queryServer{keeper}
}

func (k queryServer) Params(goCtx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

func (k queryServer) Keygen(ctx context.Context, keygen *QueryKeygen) (*QueryKeygenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k queryServer) KeygenList(ctx context.Context, list *QueryKeygenList) (*QueryKeygenListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k queryServer) PubKey(ctx context.Context, key *QueryPubKey) (*QueryPubKeyResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k queryServer) PubKeyList(ctx context.Context, list *QueryPubKeyList) (*QueryPubKeyListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k queryServer) Sign(ctx context.Context, sign *QuerySign) (*QuerySignResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k queryServer) SignList(ctx context.Context, list *QuerySignList) (*QuerySignListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k queryServer) Signature(ctx context.Context, signature *QuerySignature) (*QuerySignatureResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k queryServer) SignatureList(ctx context.Context, list *QuerySignatureList) (*QuerySignatureListResponse, error) {
	//TODO implement me
	panic("implement me")
}
