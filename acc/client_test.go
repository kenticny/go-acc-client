package acc

import (
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient(AccOptions{
		Appid:       "Testappid",
		Appsecret:   "Testappsecret",
		Environment: ENV_LOCAL,
	})
	fmt.Println(client)
}
