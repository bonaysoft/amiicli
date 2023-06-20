package amiibo

import (
	_ "embed"
)

//go:embed binfile/amiitool_darwin_arm64
var amiitoolBytes []byte
