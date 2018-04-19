package acc

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/grpc"

	pbConfig "github.com/kenticny/go-acc-client/pb/configuration"
)

const (
	DefaultAccServer = "127.0.0.1:6767"

	ENV_PRODUCTION  = pbConfig.Environment_PRODUCTION
	ENV_STAGING     = pbConfig.Environment_STAGING
	ENV_DEVELOPMENT = pbConfig.Environment_DEVELOPMENT
	ENV_LOCAL       = pbConfig.Environment_LOCAL
)

type Client struct {
	conn      *grpc.ClientConn
	rpc       pbConfig.ConfigurationClient
	env       pbConfig.Environment
	namespace string
}

func (c *Client) Pull(keys ...string) *Configuration {
	config := new(Configuration)
	if len(keys) == 0 {
		return config
	}

	// get context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer func() {
		cancel()
	}()

	configMap := make(map[string]interface{})

	for _, key := range keys {
		req := &pbConfig.RequestByKey{
			Key: &pbConfig.Key{Key: key},
			Env: c.env,
		}
		cfg, err := c.rpc.Get(ctx, req)
		if err != nil {
			continue
		}
		var tpCfgMap map[string]interface{}
		if err := json.Unmarshal([]byte(cfg.Config), &tpCfgMap); err != nil {
			continue
		}
		for k, v := range tpCfgMap {
			var sk string
			if c.namespace == key {
				sk = k
			} else {
				sk = fmt.Sprintf("%s.%s", key, k)
			}
			configMap[sk] = v
		}
	}
	config.cfg = configMap
	return config
}

// NewClient create a acc client
func NewClient(opts AccOptions) *Client {

	// check parameters
	if opts.Appid == "" || opts.Appsecret == "" {
		panic("appid and appsecret are required")
	}
	if opts.Host == "" {
		opts.Host = DefaultAccServer
	}

	// connect grpc server
	conn, err := grpc.Dial(opts.Host, grpc.WithInsecure())
	if err != nil {
		panic("connect server failed")
	}

	// rpc configuration client
	rpc := pbConfig.NewConfigurationClient(conn)

	client := &Client{
		conn: conn,
		rpc:  rpc,
		env:  opts.Environment,
	}
	if opts.Namespace != "" {
		client.namespace = opts.Namespace
	}
	return client
}
