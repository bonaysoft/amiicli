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

func RandomBytes(count int) []byte {
	output := make([]byte, count)

	for i := 0; i < count; i++ {
		output[i] = byte(rand.Intn(256))
	}

	return output
}
