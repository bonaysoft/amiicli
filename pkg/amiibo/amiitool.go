package amiibo

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/saltbo/gopkg/strutil"
)

var _ Client = (*AmiiTool)(nil)

type AmiiTool struct {
	keyRetail string
}

func NewAmiiTool(keyRetail string) Client {
	return &AmiiTool{keyRetail: keyRetail}
}

func (a *AmiiTool) Encrypt(src []byte) ([]byte, error) {
	srcFile, err := a.stash(src)
	if err != nil {
		return nil, err
	}

	outPath := fmt.Sprintf("%s/%s-encrypted.bin", filepath.Dir(srcFile.Name()), strutil.RandomText(10))
	cmd := exec.Command("amiitool", "-e", "-k", a.keyRetail, "-i", srcFile.Name(), "-o", outPath)
	if v, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("%s: %v", err, string(v))
	}

	return os.ReadFile(outPath)
}

func (a *AmiiTool) Decrypt(src []byte) ([]byte, error) {
	srcFile, err := a.stash(src)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("amiitool", "-d", "-k", a.keyRetail, "-i", srcFile.Name())
	return cmd.Output()
}

func (a *AmiiTool) stash(src []byte) (*os.File, error) {
	srcFile, err := os.CreateTemp("", "amiicli-amiitool")
	if err != nil {
		return nil, err
	}
	if _, err := srcFile.Write(src); err != nil {
		return nil, err
	}

	return srcFile, nil
}
