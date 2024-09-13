package model

// 位置信息汇报
type Msg0200 struct {
	Header       *MsgHeader    `json:"header"`
	LocationData *LocationData `json:"locationData"`
}

func (m *Msg0200) Decode(packet *PacketData) error {
	m.Header = packet.Header
	pkt, idx := packet.Body, 0

	m.LocationData = &LocationData{}

	m.LocationData.Decode(pkt, &idx)
	return nil
}

func (m *Msg0200) Encode() (pkt []byte, err error) {
	pkt, err = writeHeader(m, pkt)
	return pkt, err
}

func (m *Msg0200) GetHeader() *MsgHeader {
	return m.Header
}

func (m *Msg0200) GenOutgoing(_ JT808Msg) error {
	return nil
}
