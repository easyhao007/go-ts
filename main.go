package main

import (
	"fmt"
	"go-ts/mpegts"
	"os"
	"encoding/hex"
)

func main() {
	fmt.Println("for mpegts info")
	packet := new(mpegts.TsPacket)
	buf := make([]byte , 188)

	filename:="D:\\video\\4k_264.ts"

	pfile , err := os.Open(filename)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	index :=0
	for{
		fmt.Println("pk-index:" , index)
		nread , err := pfile.Read(buf)
		if err != nil{
			break
		}else {
			fmt.Println(hex.Dump(buf))
			if nread != 188{
				fmt.Println("读取的数据不是188大小")
				break
			}else {
				if err = packet.Demux(buf) ; err != nil{
					fmt.Println("demux ts packet error , error info is " , err.Error())
					return
				}
			}
		}
		index++

	}
}
