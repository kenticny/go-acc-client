# go-acc-client

ACC client for Golang.

## Installation

    go get -u github.com/kenticny/go-acc-client

## Usage

```go
import "github.com/kenticny/go-acc-client/acc"

client := acc.NewClient(acc.AccOptions{
  Appid:       "Your app id",
  Appsecret:   "Your app secret",
  Environment: acc.ENV_PRODUCTION,
})

config := client.Pull("config1", "config2")

fmt.Println(config)  // config json string

fmt.Println(config.GetString("key1")) // get string value

fmt.Println(config.GetNumber("key2")) // get float64 value

fmt.Println(config.GetBoolean("key3")) // get boolean value
```

## API

### AccOptions

- `Host` Acc server host, optional, default server is `127.0.0.1:6767`.

- `Appid` Application id, required.

- `Appsecret` Application secret, required.

- `Environment` Config environment, optional, default is `acc.ENV_LOCAL`.

    - `acc.ENV_LOCAL`: local

    - `acc.ENV_DEVELOPMENT`: development

    - `acc.ENV_STAGING`: STAGING

    - `acc.ENV_PRODUCTION`: production

- `Namespace` optional, Configurations append in namespace configuration.

### Client

- `Pull(keys ...string) Configuration`

  Fetch configurations from remote by keys.

### Configuration

- `GetString(key string) string`

  Get string value by key.

- `GetNumber(key string) float64`

  Get float64 value by key.

- `GetBoolean(key string) bool`

  Get boolean value by key.
