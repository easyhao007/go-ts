package main

import (
	"fmt"
	"go-ts/mpegts"
)

func main() {
	fmt.Println("for mpegts info")
	header := new(mpegts.TsHeader)

	head := []byte{0, 1, 3, 4}
	header.Demux(head)
}
