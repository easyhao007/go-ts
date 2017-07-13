package mpegts

import "github.com/go-oryx-lib/errors"

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
	SyncByte                   uint8 //8bit
	TransportErrorIndicator    uint8 //1bit
	PayloadUintStartIndicator  uint8 //1bit
	TransportPriority          uint8 //1bit
	TsPidTable                 uint  //13bit
	TransportScramblingControl uint8 //2bit
	AdaptationFieldControl     uint8 //2bit
	ContinuityCounter          uint8 //4bit
}

func (header *TsHeader) Demux(buf []uint8) (err error) {
	if len(buf) != 4 {
		err = errors.New("ts包头的长度为固定的4字节，传入的buf不正确")
		return err
	}
	header.SyncByte = buf[0]
	header.TransportErrorIndicator = buf[1] & 0x80
	header.PayloadUintStartIndicator = buf[1] & 0x40
	header.TsPidTable = uint((buf[1]&0xE0)<<8 | buf[2])
	header.TransportScramblingControl = buf[3] & 0xC0
	header.AdaptationFieldControl = buf[3] & 0x30
	header.ContinuityCounter = buf[3] & 0x0F

	return nil
}
