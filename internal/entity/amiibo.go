package entity

import (
	"encoding/hex"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/bonaysoft/amiicli/pkg/amiibo"
	"github.com/bonaysoft/amiicli/pkg/hexuid"
	"github.com/spf13/viper"
)

type Amiibo struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
	Path string `json:"path,omitempty"`
}

func NewAmiibo(name string) *Amiibo {
	return &Amiibo{Name: name}
}

func (a *Amiibo) Build() error {
	if a.Path != "" {
		return nil
	}

	encryptedData, err := amiibo.NewAmiiTool(viper.GetString("key-retail")).Encrypt(a.generate())
	if err != nil {
		return err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	repoPath := path.Join(homeDir, ".local/share/amiicli/remote")
	a.Path = filepath.Join(repoPath, fmt.Sprintf("%s.bin", a.ID))
	_ = os.MkdirAll(repoPath, 0755)
	return os.WriteFile(a.Path, encryptedData, 0644)
}

func (a *Amiibo) generate() []byte {
	arr := make([]byte, 540)
	copy(arr[0x1D4:], []byte{0x04, 0xC0, 0x0A, 0x46, 0x61, 0x6B, 0x65, 0x0A}) // Set UID
	copy(arr[:8], []byte{0x65, 0x48, 0x0F, 0xE0, 0xF1, 0x10, 0xFF, 0xEE})     // Set BCC, Internal, Static Lock, and CC
	copy(arr[0x28:], []byte{0xA5, 0x00, 0x00, 0x00})                          // Set 0xA5, Write Counter, and Unknown
	copy(arr[0x208:], []byte{0x01, 0x00, 0x0F, 0xBD})                         // Set Dynamic Lock, and RFUI
	copy(arr[0x20C:], []byte{0x00, 0x00, 0x00, 0x04})                         // Set CFG0
	copy(arr[0x210:], []byte{0x5F, 0x00, 0x00, 0x00})                         // Set CFG1
	copy(arr[0x1E8:], hexuid.RandomBytes(32))                                 // Set Keygen Salt

	// write key/amiibo num in big endian as a 64 bit value starting from offset off
	id := a.ID[2:]
	off := 0x1DC
	for i := 0; i < 16; i += 2 {
		b, _ := hex.DecodeString(id[i : i+2])
		arr[off] = b[0]
		off += 1
	}
	return arr
}
