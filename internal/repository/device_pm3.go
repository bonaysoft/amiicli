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

func (d *PM3Device) Restore(ctx context.Context, amiibo *entity.Amiibo, password string) error {
	distBinFilePath := filepath.Join(d.tmpDir, "new.bin")
	err := d.resetUID(ctx, amiibo.Path, distBinFilePath)
	if err != nil {
		return err
	}

	exeCmd := exec.CommandContext(ctx, "pm3", "-c", fmt.Sprintf("script run hf_mfu_amiibo_restore -f %s -k %s", distBinFilePath, password))
	_, err = exeCmd.CombinedOutput()
	return err
}

func (d *PM3Device) Simulate(ctx context.Context, amiibo *entity.Amiibo) error {
	distBinFilePath := filepath.Join(d.tmpDir, "new.bin")
	err := d.resetUID(ctx, amiibo.Path, distBinFilePath)
	if err != nil {
		return err
	}

	exeCmd := exec.CommandContext(ctx, "pm3", "-c", "script run hf_mfu_amiibo_sim -f "+distBinFilePath)
	exeCmd.Stdout = os.Stdout
	exeCmd.Stderr = os.Stderr
	return exeCmd.Run()
}

func (d *PM3Device) resetUID(ctx context.Context, srcBinFilePath string, distBinFilePath string) error {
	uid := "04" + hexuid.RandomText(12)
	exeCmd := exec.CommandContext(ctx, "pm3", "-c", fmt.Sprintf("script run amiibo_change_uid %s %s %s %s", uid, srcBinFilePath, distBinFilePath, d.keyRetail))
	_, err := exeCmd.CombinedOutput()
	return err
}
