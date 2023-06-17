package repository

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bonaysoft/amiicli/internal/entity"
	"github.com/bonaysoft/amiicli/pkg/hexuid"
)

var _ Device = (*PM3Device)(nil)

type PM3Device struct {
	tmpDir    string
	keyRetail string
}

func NewPM3Device(keyRetail string) (*PM3Device, error) {
	tmpDir, err := os.MkdirTemp("", "amiicli")
	if err != nil {
		return nil, fmt.Errorf("error creating temporary: %v", err)
	}

	if _, err := os.Stat(keyRetail); err != nil {
		return nil, fmt.Errorf("keyRetail: %v", err)
	}

	return &PM3Device{
		tmpDir:    tmpDir,
		keyRetail: keyRetail,
	}, nil
}

func (d *PM3Device) Clone(ctx context.Context, amiibo *entity.Amiibo) error {
	// TODO implement me
	panic("implement me")
}

func (d *PM3Device) Simulate(ctx context.Context, amiibo *entity.Amiibo) error {
	uid := "04" + hexuid.RandomText(12)
	srcBinFilePath := amiibo.Path
	distBinFilePath := filepath.Join(d.tmpDir, "new.bin")
	ec := exec.CommandContext(ctx, "pm3", "-c", fmt.Sprintf("script run amiibo_change_uid %s %s %s %s", uid, srcBinFilePath, distBinFilePath, d.keyRetail))
	ec.Stdout = os.Stdout
	ec.Stderr = os.Stderr
	if err := ec.Run(); err != nil {
		return err
	}

	exeCmd := exec.CommandContext(ctx, "pm3", "-c", "script run hf_mfu_amiibo_sim -f "+distBinFilePath)
	exeCmd.Stdout = os.Stdout
	exeCmd.Stderr = os.Stderr
	return exeCmd.Run()
}
