package mpegts

/*
	doc/ios13818规范（中文）.pdf - 2.4.4.5 节目相关分段中字段的语义定义
*/

// Pat Program Map Table.
type Pat struct {
	tableID                uint8  //1B-固定为0x00 ，标志是该表是PAT表
	sectionSyntaxIndicator uint8  //1bit-段语法标志位，固定为1
	const0Value            uint8  //1bit-0
	reserved1              uint8  //2bit-保留位
	sectionLength          uint16 //12bit-头两比特必为‘00’， 剩余 10 比特指定该分段的字节数，紧随分段长度字段开始，并包括CRC。
	transportStreamID      uint16 //16bit-标识网络内此传输流有别于任何其他多路复用流
	reserved2              uint8  //2bit-保留位
	versionNumber          uint8  //5bit-整个节目相关表的版本号
	currentNextIndicator   uint8  //1bit-指示发送的节目相关表为当前有效
	sectionNumber          uint8  //8bit-给出此分段的编号
	lastSectionNumber      uint8  //8bit-指定完整节目相关表的最后分段编号
	programInfo            []PatProgramInfo
	crc32                  uint32 //32bit
}

// PatProgramInfo Program Info of mpeg.
type PatProgramInfo struct {
	programNumber uint16 //16bit-它指定 program_map_PID 所适用的节目
	reserved      uint8  //3bit-保留位
	networkPid    uint16 //13bit-仅同设置为 0x0000 值的 program_number 一起使用
	programMapPid uint16 //13bit
}

// new pat parse
func NewPatParse() (pat *Pat) {
	return new(Pat)
}

// parse pat
func (pat *Pat) parsePATSection(buf []byte) (err error) {
	pat.tableID = uint8(buf[0])
	pat.sectionSyntaxIndicator = uint8(buf[1] & 0x80)
	pat.const0Value = 0
	pat.sectionLength = uint16((buf[1]&0x0F)<<4 | buf[2])
	pat.transportStreamID = uint16(buf[3]<<8 | buf[4])
	pat.versionNumber = uint8(buf[5] & 0x3E)
	pat.currentNextIndicator = uint8(buf[5] & 0x01)
	pat.sectionNumber = uint8(buf[6])
	pat.lastSectionNumber = uint8(buf[7])

	progSize := int((pat.sectionLength)-9) / 4
	index := 8
	for i := 0; i < progSize; i++ {
		var patProgramInfo PatProgramInfo
		patProgramInfo.programNumber = uint16(buf[index]<<8 | buf[index+1])
		index += 2
		if patProgramInfo.programNumber == 0x00 {
			patProgramInfo.networkPid = uint16((buf[index]&0x1F)<<8 | (buf[index+1]))
		} else {
			patProgramInfo.programMapPid = uint16((buf[index]&0x1F)<<8 | (buf[index+1]))
		}
		index += 2
	}

	return nil
}
