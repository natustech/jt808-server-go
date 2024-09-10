package model

// Location Information Query
type Msg8201 struct {
	Header *MsgHeader `json:"header"`
}

func (m *Msg8201) Decode(packet *PacketData) error {
	m.Header = packet.Header

	return nil
}

func (m *Msg8201) Encode() (pkt []byte, err error) {

	pkt, err = writeHeader(m, pkt)
	return pkt, err
}

func (m *Msg8201) GetHeader() *MsgHeader {
	return m.Header
}

func (m *Msg8201) GenOutgoing(_ JT808Msg) error {
	return nil
}
