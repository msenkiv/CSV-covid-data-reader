package main

import (
	"fmt"
	"math/rand"

	"github.com/segmentio/ksuid"
)

type UniqueRand struct {
	generated map[int]bool
}

func (u *UniqueRand) Int() int {
	for {
		i := rand.Int()
		if !u.generated[i] {
			u.generated[i] = true
			return i
		}
	}
}

func main() {

}
func genKsuid() {
	id := ksuid.New()

	fmt.Printf("github.com/segmentio/ksuid:     %s\n", id.String())
}
