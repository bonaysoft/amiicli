package amiibo

import (
	"embed"
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata
var testdata embed.FS

func TestEncrypt(t *testing.T) {
	keyRetail := "./key_retail.bin"
	for s, constructor := range registerClients {
		client := constructor(keyRetail)
		walker := func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				return nil
			}

			data, err := testdata.ReadFile(path)
			assert.NoError(t, err)

			decrypted, err := client.Decrypt(data)
			assert.NoError(t, err)

			encrypted, err := client.Encrypt(decrypted)
			assert.NoError(t, err)
			assert.Equal(t, data, encrypted)
			return nil
		}

		t.Run(s, func(t *testing.T) {
			assert.NoError(t, fs.WalkDir(testdata, ".", walker))
		})
	}
}
