package apimiddleware

import (
	"github.com/pkg/errors"
	"github.com/prysmaticlabs/prysm/api/gateway/apimiddleware"
)

// ValidatorEndpointFactory creates endpoints used for running validator API calls through the API Middleware.
type ValidatorEndpointFactory struct {
}

func (f *ValidatorEndpointFactory) IsNil() bool {
	return f == nil
}

// Paths is a collection of all valid validator API paths.
func (*ValidatorEndpointFactory) Paths() []string {
	return []string{
		"/eth/v1/keystores",
		"/eth/v1/remotekeys",
		"/eth/v1/validator/{pubkey}/feerecipient",
	}
}

// Create returns a new endpoint for the provided API path.
func (*ValidatorEndpointFactory) Create(path string) (*apimiddleware.Endpoint, error) {
	endpoint := apimiddleware.DefaultEndpoint()
	switch path {
	case "/eth/v1/keystores":
		endpoint.GetResponse = &listKeystoresResponseJson{}
		endpoint.PostRequest = &importKeystoresRequestJson{}
		endpoint.PostResponse = &importKeystoresResponseJson{}
		endpoint.DeleteRequest = &deleteKeystoresRequestJson{}
		endpoint.DeleteResponse = &deleteKeystoresResponseJson{}
	case "/eth/v1/remotekeys":
		endpoint.GetResponse = &listRemoteKeysResponseJson{}
		endpoint.PostRequest = &importRemoteKeysRequestJson{}
		endpoint.PostResponse = &importRemoteKeysResponseJson{}
		endpoint.DeleteRequest = &deleteRemoteKeysRequestJson{}
		endpoint.DeleteResponse = &deleteRemoteKeysResponseJson{}
	case "/eth/v1/validator/{pubkey}/feerecipient":
		endpoint.GetResponse = &getFeeRecipientByPubkeyResponseJson{}
		endpoint.PostRequest = &setFeeRecipientByPubkeyRequestJson{}
	default:
		return nil, errors.New("invalid path")
	}
	endpoint.Path = path
	return &endpoint, nil
}
