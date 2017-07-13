package mpegts

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

func (ad *TsAdaptationField)Demux(buf []byte) (err error){
	return nil
}