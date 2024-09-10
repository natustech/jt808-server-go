package model

import (
	"github.com/fakeyanss/jt808-server-go/internal/codec/hex"
)

// Request Sync Time Reply
type Msg8109 struct {
	Header         *MsgHeader `json:"header"`
	Year           uint16     `json:"year"`
	Month          uint8      `json:"month"`
	Day            uint8      `json:"day"`
	When           uint8      `json:"when"`
	Points         uint8      `json:"points"`
	Seconds        uint8      `json:"seconds"`
	ResponseResult uint8      `json:"responseResult"`
}

func (m *Msg8109) Decode(packet *PacketData) error {
	m.Header = packet.Header
	pkt, idx := packet.Body, 0

	m.Year = hex.ReadWord(pkt, &idx)
	m.Month = hex.ReadByte(pkt, &idx)
	m.Day = hex.ReadByte(pkt, &idx)
	m.When = hex.ReadByte(pkt, &idx)
	m.Points = hex.ReadByte(pkt, &idx)
	m.Seconds = hex.ReadByte(pkt, &idx)
	m.ResponseResult = hex.ReadByte(pkt, &idx)

	return nil
}

func (m *Msg8109) Encode() (pkt []byte, err error) {
	pkt = hex.WriteWord(pkt, m.Year)
	pkt = hex.WriteByte(pkt, m.Month)
	pkt = hex.WriteByte(pkt, m.Day)
	pkt = hex.WriteByte(pkt, m.When)
	pkt = hex.WriteByte(pkt, m.Points)
	pkt = hex.WriteByte(pkt, m.Seconds)
	pkt = hex.WriteByte(pkt, m.ResponseResult)

	pkt, err = writeHeader(m, pkt)
	return pkt, err
}

func (m *Msg8109) GetHeader() *MsgHeader {
	return m.Header
}

func (m *Msg8109) GenOutgoing(incoming JT808Msg) error {
	in, ok := incoming.(*Msg0109)
	if !ok {
		return ErrGenOutgoingMsg
	}

	m.Header = in.Header
	m.Header.MsgID = 0x8109

	return nil
}
