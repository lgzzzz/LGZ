package pkg

import (
	"math/rand"
	"time"
)

func NewPassword() string {
	p := make([]byte, 10)
	rand.Seed(time.Now().Unix())
	for i := range p {
		for {
			ch := rand.Intn(255)
			if ch >= int('a') && ch <= int('z') {
				p[i] = byte(ch)
				break
			}
			if ch >= int('A') && ch <= int('Z') {
				p[i] = byte(ch)
				break
			}
			if ch >= int('0') && ch <= int('9') {
				p[i] = byte(ch)
				break
			}
		}
	}
	return string(p)
}
