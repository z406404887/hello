package network

const (
	COMMON_HEADER_LENGTH = 10

)

const (
	MaxUint16 = 0xffff
)

type CommonHeader struct {
	MainType uint8
	SubType uint8
	Len uint16
	ClientId uint32
	Result uint16
}

func (header *CommonHeader) Encode() []byte{
	data := make([]byte, COMMON_HEADER_LENGTH)
	data[0] = header.MainType
	data[1] = header.SubType
	BinaryCoder.PutUint16(data[2:4],header.Len)
	BinaryCoder.PutUint32(data[4:8],header.ClientId)
	BinaryCoder.PutUint16(data[8:10],header.Result)
	return data
}

func (header *CommonHeader) Decode(data []byte) bool {
	if len(data) < COMMON_HEADER_LENGTH {
		return false
	}

	header.MainType = data[0]
	header.SubType = data[1]
	header.Len = BinaryCoder.Uint16(data[2:4])
	header.ClientId = BinaryCoder.Uint32(data[4:8])
	header.Result = BinaryCoder.Uint16(data[8:10])
	return true
}

type Message struct {
	SrcId uint32
	Head *CommonHeader
	Data []byte
}