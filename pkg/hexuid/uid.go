package hexuid

import (
	"math/rand"
	"time"

	"github.com/saltbo/gopkg/strutil"
)

const (
	letterBytes = "ABCDEF123456789"
)

var letterLength = len(letterBytes)

func init() {
	rand.Seed(time.Now().Unix())
}

func RandomText(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int()%letterLength]
	}

	return strutil.B2s(b)
}
