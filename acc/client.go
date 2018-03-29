package acc

import (
	"google.golang.org/grpc"

	pbConfig "github.com/kenticny/go-acc-client/pb/configuration"
)

const (
	DefaultAccServer = ""

	ENV_PRODUCTION    = pbConfig.Environment_PRODUCTION
	ENV_PREPRODUCTION = pbConfig.Environment_PREPRODUCTION
	ENV_DEVELOPMENT   = pbConfig.Environment_DEVELOPMENT
	ENV_LOCAL         = pbConfig.Environment_LOCAL
)

type Client struct {
	conn *grpc.ClientConn
	rpc  *pbConfig.ConfigurationClient
	env  pbConfig.Environment
}

func (c *Client) Pull(keys []string) *Configuration {
	return nil
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
		rpc:  &rpc,
		env:  opts.Environment,
	}
	return client
}
