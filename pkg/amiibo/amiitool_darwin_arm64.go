package amiibo

import (
	_ "embed"
)

//go:embed binfile/amiibo_darwin_arm64
var amiitoolBytes []byte
