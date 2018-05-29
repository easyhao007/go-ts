package mpegts

// Pat Program Map Table.
type Pat struct {
	//1B-固定为0x00 ，标志是该表是PAT表
	tableID uint8 `json:table_id`
	//1bit-段语法标志位，固定为1
	sectionSyntaxIndicator uint8
	//1bit-0
	const0Value uint8
	//2bit-保留位
	reserved1 uint8
	//12bit-头两比特必为‘00’， 剩余 10 比特指定该分段的字节数，紧随分段长度字段开始，并包括CRC。
	sectionLength uint16
	//16bit-标识网络内此传输流有别于任何其他多路复用流
	transportStreamID uint16
	//5bit-整个节目相关表的版本号
	versionNumber uint8
	//
	currentNextIndicator uint8
	sectionNumber        uint8
	lastSectionNumber    uint8
	programInfo          []PatProgramInfo
	crc32                uint32
}

// PatProgramInfo Program Info of mpeg.
type PatProgramInfo struct {
	programNumber uint16
	networkPid    uint16
	programMapPid uint16
}
