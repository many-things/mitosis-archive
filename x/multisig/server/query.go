package server

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/keeper"
	"github.com/many-things/mitosis/x/multisig/types"
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

func (k queryServer) PubKey(ctx context.Context, msg *QueryPubKey) (*QueryPubKeyResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	keyID := types.KeyID(msg.KeyId)
	chainID, id, err := keyID.ToInternalVariables()
	if err != nil {
		return nil, fmt.Errorf("keyId has invalid format")
	}

	pubKey, err := k.baseKeeper.QueryPubKey(wctx, chainID, id, msg.Validator)
	if err != nil {
		return nil, err
	}

	return &QueryPubKeyResponse{
		PubKey: pubKey,
	}, nil
}

func (k queryServer) PubKeyList(ctx context.Context, msg *QueryPubKeyList) (*QueryPubKeyListResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	keyID := types.KeyID(msg.KeyId)
	chainID, id, err := keyID.ToInternalVariables()
	if err != nil {
		return nil, fmt.Errorf("keyId has invalid format: chainId-sequence")
	}

	kvPubkeys, page, err := k.baseKeeper.QueryPubKeyList(wctx, chainID, id, msg.Pagination)
	if err != nil {
		return nil, err
	}

	return &QueryPubKeyListResponse{
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

func (k queryServer) Signature(ctx context.Context, msg *QuerySignature) (*QuerySignatureResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	sigID := types.SigID(msg.SigId)
	chainID, id, err := sigID.ToInternalVariables()
	if err != nil {
		return nil, fmt.Errorf("sigId has invalid format: must be chainId-sequence")
	}

	signature, err := k.baseKeeper.QuerySignature(wctx, chainID, id, msg.Validator)
	if err != nil {
		return nil, err
	}

	return &QuerySignatureResponse{Signature: signature}, nil
}

func (k queryServer) SignatureList(ctx context.Context, msg *QuerySignatureList) (*QuerySignatureListResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	sigID := types.SigID(msg.SigId)
	chainID, id, err := sigID.ToInternalVariables()
	if err != nil {
		return nil, fmt.Errorf("sigId has invalid foramt: must be chainId-sequence")
	}

	kvSignature, page, err := k.baseKeeper.QuerySignatureList(wctx, chainID, id, msg.Pagination)
	if err != nil {
		return nil, err
	}

	return &QuerySignatureListResponse{List: mitotypes.Values(kvSignature), Page: page}, nil
}
