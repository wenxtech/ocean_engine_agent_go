/*
巨量方舟

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ocean_engine_agent_go

import (
	"github.com/wenxtech/ocean_engine_agent_go/api"
	"github.com/wenxtech/ocean_engine_agent_go/middleware"
)

type Client struct {
	ApiClient *api.APIClient
}

func Init(cfg *api.Configuration) *Client {
	client := &Client{
		ApiClient: api.NewAPIClient(cfg),
	}
	client.ApiClient.Use(
		middleware.AuthMiddleware,
		middleware.HeaderMiddleware,
	)
	if cfg.UseLogMw {
		client.ApiClient.Use(middleware.LogMiddleware)
	}
	return client
}
// API Services
func (c *Client) AgentTicketListV2API() *api.AgentTicketListV2APIService {
	return c.ApiClient.AgentTicketListV2API
}

func (c *Client) CommonApi() *api.CommonApiService {
	return c.ApiClient.CommonApi
}

func (c *Client) SetHost(host string) {
	c.ApiClient.Cfg.Host = host
}

func (c *Client) Use(mws ...api.Middleware) {
	c.ApiClient.Use(mws...)
}

func (c *Client) AddDefaultHeader(key string, value string) {
	c.ApiClient.Cfg.AddDefaultHeader(key, value)
}

func (c *Client) SetLogEnable(enable bool) {
	c.ApiClient.Cfg.LogEnable = enable
}