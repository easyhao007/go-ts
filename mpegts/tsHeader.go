package mpegts

import (
	"errors"
	"fmt"
	"go-ts/bitbuffer"
)

const (
	TsPidTablePAT  = 0x00
	TsPidTableCAT  = 0x01
	TsPidTableTSDT = 0x02
	TsPidTableNULL = 0x01FF
)

const (
	TsAdaptationTypeReserved       = 0x00
	TsAdaptationTypePayloadOnly    = 0x01
	TsAdaptationTypeAdaptationOnly = 0x02
	TsAdaptationTypeBoth           = 0x03
)

//ts包的包头
type TsHeader struct {
	SyncByte                   uint8  //8bit
	TransportErrorIndicator    uint8  //1bit
	PayloadUintStartIndicator  uint8  //1bit
	TransportPriority          uint8  //1bit
	TsPidTable                 uint16 //13bit
	TransportScramblingControl uint8  //2bit
	AdaptationFieldControl     uint8  //2bit
	ContinuityCounter          uint8  //4bit
}

func (header *TsHeader) Demux(buf []uint8) (err error) {
	if len(buf) != 4 {
		err = errors.New("ts包头的长度为固定的4字节，传入的buf不正确")
		return err
	}

	bb := new(bitbuffer.BitBuffer)
	bb.Set(buf)

	if header.SyncByte, err = bb.PeekUint8(8); err != nil {
		fmt.Println("SyncByte")
		return err
	}

	if header.TransportErrorIndicator, err = bb.PeekUint8(1); err != nil {
		fmt.Println("TransportErrorIndicator")
		return err
	}

	if header.PayloadUintStartIndicator, err = bb.PeekUint8(1); err != nil {
		fmt.Println("PayloadUintStartIndicator")
		return err
	}

	if header.TransportPriority, err = bb.PeekUint8(1); err != nil {
		fmt.Println("TransportPriority")
		return err
	}

	if header.TsPidTable, err = bb.PeekUint16(13); err != nil {
		fmt.Println("TsPidTable")
		return err
	}

	if header.TransportScramblingControl, err = bb.PeekUint8(2); err != nil {
		fmt.Println("TransportScramblingControl")
		return err
	}

	if header.AdaptationFieldControl, err = bb.PeekUint8(2); err != nil {
		fmt.Println("AdaptationFieldControl")
		return err
	}

	if header.ContinuityCounter, err = bb.PeekUint8(4); err != nil {
		fmt.Println("ContinuityCounter")
		return err
	}

	return nil
}
