package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSQL(t *testing.T) {
	var b bytes.Buffer

	b.WriteString("G")
	b.WriteString("e")
	b.WriteString("e")
	b.WriteString("k")
	b.WriteString("s")

	fmt.Println("String: ", b.String())

}
