package equinix

import (
	"github.com/jxoir/equinix-tools/pkg/ecxlib/api/buyer"
	ecxbuyer "github.com/jxoir/equinix-tools/pkg/ecxlib/api/buyer"
	equinixapiclient "github.com/jxoir/equinix-tools/pkg/ecxlib/api/client"
)

// Config struct for Equinix client params
type Config struct {
	AppID           string
	AppSecret       string
	GrantType       string
	UserName        string
	Password        string
	Endpoint        string
	PlaygroundToken string
	Debug           bool
}

type EquinixClient struct {
	apitoken                  string
	client                    *equinixapiclient.EquinixAPIClient
	ECXConnectionsAPI         *buyer.ECXConnectionsAPI
	ECXBuyerSellerServicesAPI *buyer.ECXSellerServicesAPI
}

// Client configures and returns a fully initialized EquinixAPIClient
func (c *Config) Client() (interface{}, error) {
	//var client equinixapiclient.EquinixAPIClient
	// default api host
	ecxapihost := "api.equinix.com"
	if c.Endpoint != "" {
		ecxapihost = c.Endpoint
	}

	clientParams := &equinixapiclient.EquinixAPIParams{
		AppID:           c.AppID,
		AppSecret:       c.AppSecret,
		GrantType:       "client_credentials",
		UserName:        c.UserName,
		UserPassword:    c.Password,
		Endpoint:        ecxapihost,
		PlaygroundToken: c.PlaygroundToken,
		Debug:           c.Debug,
	}

	// create new ecx api client
	ecxclient := equinixapiclient.NewEcxAPIClient(clientParams, ecxapihost, false)
	//client.client = ecxclient
	//client.ECXConnectionsAPI = ecxbuyer.NewECXConnectionsAPI(client.client)
	client := &EquinixClient{
		client:            ecxclient,
		ECXConnectionsAPI: ecxbuyer.NewECXConnectionsAPI(ecxclient),
	}
	return client, nil
}
