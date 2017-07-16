package mpegts

type TsPacket struct {
	header     TsHeader
	adaptation TsAdaptationField
}

func (packet *TsPacket)Demux(buf []byte)(err error){
	if err = packet.header.Demux(buf[0:4]) ; err != nil{
		return err
	}
	if packet.header.AdaptationFieldControl == TsAdaptationTypeAdaptationOnly || packet.header.AdaptationFieldControl == TsAdaptationTypeBoth{
		if err = packet.adaptation.Demux(buf[4:]) ; err != nil{
			return err
		}
	}
	return nil
}