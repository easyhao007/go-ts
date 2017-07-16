package mpegts

import (
	"go-ts/bitbuffer"
	"fmt"
)

type TsAdaptationField struct {
	AdaptationFieldLength                  uint8  //8bit
	DiscontinuityIndicator                 uint8  //1bit
	RandomAccessIndicator                  uint8  //1bit
	ElementaryStreamPriorityIndicator      uint8  //1bit
	PcrFlag                                uint8  //1bit
	OpcrFlag                               uint8  //1bit
	SplicingPointFlag                      uint8  //1bit
	TransportPrivateDataFlag               uint8  //1bit
	AdaptationFieldExtensionFlag           uint8  //1bit
	ProgramClockReferenceBase              uint64 //33bit
	ProgramClockReferenceExtension         uint16 //9bit
	OriginalProgramClockReferenceBase      uint64 //33bit
	OriginalProgramClockReferenceExtension uint16 //9bit
	SpliceCountdown                        uint8  //8bit
	TransportPrivateDataLength             uint8  // 8bit
	TransportPrivateData                   []byte
	AdaptationFieldExtensionLength         uint8  //8bit
	LtwFlag                                uint8  //1bit
	PiecewiseRateFlag                      uint8  //1bit -> 分段
	SeamlessSpliceFlag                     uint8  //1bit
	LtwValidFlag                           uint8  //1bit
	LtwOffset                              uint16 //15bit
	PiecewiseRate                          uint32 // 2bit
	SpliceType                             uint8  //4bit
	DtsNextAu0                             uint8  // 3bit
	MarkerBit0                             uint8  // 1bit
	DtsNextAu1                             uint16 // 15bit
	MarkerBit1                             uint8  // 1bit
	DtsNextAu2                             uint16 // 15bit
	MarkerBit2                             uint8  // 1bit
	AfExtReserved                          []byte
	AfReserved                             []byte
}

func (ad *TsAdaptationField) Demux(buf []byte) (err error) {
	//设置buf到bitbuffer
	bb := new(bitbuffer.BitBuffer)
	bb.Set(buf)

	if ad.AdaptationFieldLength, err = bb.PeekUint8(8); err != nil {
		fmt.Println("AdaptationFieldLength")
		return err
	}
	if ad.AdaptationFieldLength <= 0{
		return nil
	}

	if ad.DiscontinuityIndicator, err = bb.PeekUint8(1); err != nil {
		fmt.Println("DiscontinuityIndicator")
		return err
	}

	if ad.RandomAccessIndicator, err = bb.PeekUint8(1); err != nil {
		fmt.Println("RandomAccessIndicator")
		return err
	}

	if ad.ElementaryStreamPriorityIndicator, err = bb.PeekUint8(1); err != nil {
		fmt.Println("ElementaryStreamPriorityIndicator")
		return err
	}

	if ad.PcrFlag, err = bb.PeekUint8(1); err != nil {
		fmt.Println("PcrFlag")
		return err
	}

	if ad.OpcrFlag, err = bb.PeekUint8(1); err != nil {
		fmt.Println("OpcrFlag")
		return err
	}

	if ad.SplicingPointFlag, err = bb.PeekUint8(1); err != nil {
		fmt.Println("SplicingPointFlag")
		return err
	}

	if ad.TransportPrivateDataFlag, err = bb.PeekUint8(1); err != nil {
		fmt.Println("TransportPrivateDataFlag")
		return err
	}

	if ad.AdaptationFieldExtensionFlag, err = bb.PeekUint8(1); err != nil {
		fmt.Println("AdaptationFieldExtensionFlag")
		return err
	}

	if ad.PcrFlag == 1 {
		//33bit
		if ad.ProgramClockReferenceBase, err = bb.PeekUint64(33); err != nil {
			fmt.Println("ProgramClockReferenceBase")
			return err
		}
		//reserved 6bit
		if err = bb.Skip(6); err != nil {
			return err
		}
		//9bit
		if ad.ProgramClockReferenceExtension, err = bb.PeekUint16(9); err != nil {
			fmt.Println("ProgramClockReferenceExtension")
			return err
		}
	}

	if ad.OpcrFlag == 1 {
		//33bit
		if ad.OriginalProgramClockReferenceBase, err = bb.PeekUint64(33); err != nil {
			fmt.Println("OriginalProgramClockReferenceBase")
			return err
		}
		//reserved 6bit
		if err = bb.Skip(6); err != nil {
			return err
		}
		//9bit
		if ad.OriginalProgramClockReferenceExtension, err = bb.PeekUint16(9); err != nil {
			fmt.Println("OriginalProgramClockReferenceExtension")
			return err
		}
	}

	if ad.SplicingPointFlag == 1 {
		if ad.SpliceCountdown, err = bb.PeekUint8(8); err != nil {
			fmt.Println("SpliceCountdown")
			return err
		}
	}

	if ad.TransportPrivateDataFlag == 1 {
		if ad.TransportPrivateDataLength, err = bb.PeekUint8(8); err != nil {
			fmt.Println("TransportPrivateDataLength")
			return err
		}

		for i := uint8(0); i < ad.TransportPrivateDataLength; i++ {
			if chunk, err := bb.PeekUint8(8); err != nil {
				fmt.Println("TransportPrivateData")
				return err
			} else {
				ad.TransportPrivateData = append(ad.TransportPrivateData, chunk)
			}
		}
	}

	if ad.AdaptationFieldExtensionFlag == 1 {
		if ad.AdaptationFieldExtensionLength, err = bb.PeekUint8(8); err != nil {
			fmt.Println("AdaptationFieldExtensionLength")
			return err
		}

		if ad.LtwFlag, err = bb.PeekUint8(1); err != nil {
			fmt.Println("LtwFlag")
			return err
		}

		if ad.PiecewiseRateFlag, err = bb.PeekUint8(1); err != nil {
			fmt.Println("PiecewiseRateFlag")
			return err
		}

		if ad.SeamlessSpliceFlag, err = bb.PeekUint8(1); err != nil {
			fmt.Println("SeamlessSpliceFlag")
			return err
		}

		if err = bb.Skip(5); err != nil {
			return err
		}

		if ad.LtwFlag == 1 {
			if ad.LtwValidFlag, err = bb.PeekUint8(1); err != nil {
				fmt.Println("LtwValidFlag")
				return err
			}

			if ad.LtwOffset, err = bb.PeekUint16(15); err != nil {
				fmt.Println("LtwOffset")
				return err
			}
		}

		if ad.PiecewiseRateFlag == 1 {
			if err = bb.Skip(2); err != nil {
				return err
			}
			if ad.PiecewiseRate, err = bb.PeekUint32(22); err != nil {
				fmt.Println("PiecewiseRate")
				return err
			}
		}

		if ad.SeamlessSpliceFlag == 1 {
			if ad.SpliceType, err = bb.PeekUint8(4); err != nil {
				fmt.Println("SpliceType")
				return err
			}
			if ad.DtsNextAu0, err = bb.PeekUint8(3); err != nil {
				fmt.Println("DtsNextAu0")
				return err
			}
			if err = bb.Skip(1); err != nil {
				return err
			}
			if ad.DtsNextAu1, err = bb.PeekUint16(15); err != nil {
				fmt.Println("DtsNextAu1")
				return err
			}
			if err = bb.Skip(1); err != nil {
				return err
			}
			if ad.DtsNextAu2, err = bb.PeekUint16(15); err != nil {
				fmt.Println("DtsNextAu2")
				return err
			}
			if err = bb.Skip(1); err != nil {
				return err
			}

		}
	}
	return nil
}
