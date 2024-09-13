package model

// Location information query answer
type Msg0201 struct {
	Header               *MsgHeader    `json:"header"`
	ResponseSerialNumber uint16        `json:"responseSerialNumber"`
	LocationData         *LocationData `json:"locationData"`
}

func (m *Msg0201) Decode(packet *PacketData) error {
	m.Header = packet.Header
	pkt, idx := packet.Body, 0

	m.LocationData = &LocationData{}

	m.LocationData.Decode(pkt, &idx)

	return nil
}

func (m *Msg0201) Encode() (pkt []byte, err error) {
	// This message is just for receiving maybe TODO add
	pkt, err = writeHeader(m, pkt)
	return pkt, err
}

func (m *Msg0201) GetHeader() *MsgHeader {
	return m.Header
}

func (m *Msg0201) GenOutgoing(_ JT808Msg) error {
	return nil
}
