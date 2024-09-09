package model

import (
	"github.com/fakeyanss/jt808-server-go/internal/codec/hex"
)

type Msg1007 struct {
	Header   *MsgHeader `json:"header"`
	TimeZone byte       `json:"timeZone"`
}

func (m *Msg1007) Decode(packet *PacketData) error {
	m.Header = packet.Header
	pkt, idx := packet.Body, 0
	m.TimeZone = hex.ReadByte(pkt, &idx)
	return nil
}

func (m *Msg1007) Encode() (pkt []byte, err error) {
	/** TODO */
	pkt, err = writeHeader(m, pkt)
	return pkt, err
}

func (m *Msg1007) GetHeader() *MsgHeader {
	return m.Header
}

func (m *Msg1007) GenOutgoing(_ JT808Msg) error {
	return nil
}
