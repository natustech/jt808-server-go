package model

// Request synchronization time
type Msg0109 struct {
	Header *MsgHeader `json:"header"`
}

func (m *Msg0109) Decode(packet *PacketData) error {
	m.Header = packet.Header
	return nil
}

func (m *Msg0109) Encode() (pkt []byte, err error) {
	pkt, err = writeHeader(m, pkt)
	return pkt, err
}

func (m *Msg0109) GetHeader() *MsgHeader {
	return m.Header
}

func (m *Msg0109) GenOutgoing(_ JT808Msg) error {
	return nil
}
