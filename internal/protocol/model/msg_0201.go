package model

import (
	"github.com/fakeyanss/jt808-server-go/internal/codec/hex"
)

// Location information query answer
type Msg0201 struct {
	Header               *MsgHeader `json:"header"`
	ResponseSerialNumber uint16     `json:"responseSerialNumber"`
	AlarmSign            uint32     `json:"alarmSign"`
	StatusSign           uint32     `json:"statusSign"`
	Latitude             uint32     `json:"latitude"`
	Longitude            uint32     `json:"longitude"`
	Altitude             uint16     `json:"altitude"`
	Speed                uint16     `json:"speed"`
	Direction            uint16     `json:"direction"`
	Time                 string     `json:"time"`
}

func (m *Msg0201) Decode(packet *PacketData) error {
	m.Header = packet.Header
	pkt, idx := packet.Body, 0

	m.ResponseSerialNumber = hex.ReadWord(pkt, &idx)
	m.AlarmSign = hex.ReadDoubleWord(pkt, &idx)
	m.StatusSign = hex.ReadDoubleWord(pkt, &idx)
	m.Latitude = hex.ReadDoubleWord(pkt, &idx)
	m.Longitude = hex.ReadDoubleWord(pkt, &idx)
	m.Altitude = hex.ReadWord(pkt, &idx)
	m.Speed = hex.ReadWord(pkt, &idx)
	m.Direction = hex.ReadWord(pkt, &idx)
	m.Time = hex.ReadBCD(pkt, &idx, 6)

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
