package amiibo

import _ "embed"

type Client interface {
	Encrypt([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
}

type ClientConstructor func(keyRetail string) Client

var registerClients = map[string]ClientConstructor{
	"amiitool": NewAmiiTool,
}

//go:embed key_retail.bin
var keyRetailBinFile []byte
