package pkg

import (
	"encoding/hex"
	"math/rand"
	"time"
)

var NewPassportName = "newpassport"

func NewPassport() string {
	p := make([]byte, 10)
	rand.Seed(time.Now().Unix())
	rand.Read(p)
	str := hex.EncodeToString(p)
	return str
}
