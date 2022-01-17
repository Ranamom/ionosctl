package resources

import (
	"context"

	sdkgoauth "github.com/ionos-cloud/sdk-go-auth"
)

type Token struct {
	sdkgoauth.Token
}

type Jwt struct {
	sdkgoauth.Jwt
}

type Tokens struct {
	sdkgoauth.Tokens
}

type Response struct {
	sdkgoauth.APIResponse
}

type DeleteResponse struct {
	sdkgoauth.DeleteResponse
}

// TokensService is a wrapper around ionoscloud.Token
type TokensService interface {
	List(contractNumber int32) (Tokens, *Response, error)
	Get(tokenId string, contractNumber int32) (*Token, *Response, error)
	Create(contractNumber int32) (*Jwt, *Response, error)
	DeleteByID(tokenId string, contractNumber int32) (*DeleteResponse, *Response, error)
	DeleteByCriteria(criteria string, contractNumber int32) (*DeleteResponse, *Response, error)
}

type tokensService struct {
	client  *Client
	context context.Context
}

var _ TokensService = &tokensService{}

func NewTokenService(client *Client, ctx context.Context) TokensService {
	return &tokensService{
		client:  client,
		context: ctx,
	}
}

func (ts *tokensService) List(contractNumber int32) (Tokens, *Response, error) {
	req := ts.client.TokenApi.GetAllTokens(ts.context)
	if contractNumber != 0 {
		req = req.XContractNumber(contractNumber)
	}
	dcs, res, err := ts.client.TokenApi.GetAllTokensExecute(req)
	return Tokens{dcs}, &Response{*res}, err
}

func (ts *tokensService) Get(tokenId string, contractNumber int32) (*Token, *Response, error) {
	req := ts.client.TokenApi.GetTokenById(ts.context, tokenId)
	if contractNumber != 0 {
		req = req.XContractNumber(contractNumber)
	}
	token, res, err := ts.client.TokenApi.GetTokenByIdExecute(req)
	return &Token{token}, &Response{*res}, err
}

func (ts *tokensService) Create(contractNumber int32) (*Jwt, *Response, error) {
	req := ts.client.TokenApi.CreateToken(ts.context)
	if contractNumber != 0 {
		req = req.XContractNumber(contractNumber)
	}
	token, res, err := ts.client.TokenApi.CreateTokenExecute(req)
	return &Jwt{token}, &Response{*res}, err
}

func (ts *tokensService) DeleteByID(tokenId string, contractNumber int32) (*DeleteResponse, *Response, error) {
	req := ts.client.TokenApi.DeleteTokenById(ts.context, tokenId)
	if contractNumber != 0 {
		req = req.XContractNumber(contractNumber)
	}
	tokenDeleteById, res, err := ts.client.TokenApi.DeleteTokenByIdExecute(req)
	return &DeleteResponse{tokenDeleteById}, &Response{*res}, err
}

// DeleteByCriteria removes all tokens based on criteria: EXPIRED, CURRENT or ALL
func (ts *tokensService) DeleteByCriteria(criteria string, contractNumber int32) (*DeleteResponse, *Response, error) {
	req := ts.client.TokenApi.DeleteTokenByCriteria(ts.context).Criteria(criteria)
	if contractNumber != 0 {
		req = req.XContractNumber(contractNumber)
	}
	tokenDeleteByCriteria, res, err := ts.client.TokenApi.DeleteTokenByCriteriaExecute(req)
	return &DeleteResponse{tokenDeleteByCriteria}, &Response{*res}, err
}