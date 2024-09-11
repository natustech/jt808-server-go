package model

import (
	"github.com/fakeyanss/jt808-server-go/internal/codec/hex"
)

// 位置信息汇报
type Msg0200 struct {
	Header       *MsgHeader    `json:"header"`
	LocationData *LocationData `json:"locationData"`
}

func (m *Msg0200) Decode(packet *PacketData) error {
	m.Header = packet.Header
	pkt, idx := packet.Body, 0
	m.LocationData.AlarmSign = hex.ReadDoubleWord(pkt, &idx)
	m.LocationData.StatusSign = hex.ReadDoubleWord(pkt, &idx)
	m.LocationData.Latitude = hex.ReadDoubleWord(pkt, &idx)
	m.LocationData.Longitude = hex.ReadDoubleWord(pkt, &idx)
	m.LocationData.Altitude = hex.ReadWord(pkt, &idx)
	m.LocationData.Speed = hex.ReadWord(pkt, &idx)
	m.LocationData.Direction = hex.ReadWord(pkt, &idx)
	m.LocationData.Time = hex.ReadBCD(pkt, &idx, 6)
	return nil
}

func (m *Msg0200) Encode() (pkt []byte, err error) {
	pkt = hex.WriteDoubleWord(pkt, m.LocationData.AlarmSign)
	pkt = hex.WriteDoubleWord(pkt, m.LocationData.StatusSign)
	pkt = hex.WriteDoubleWord(pkt, m.LocationData.Latitude)
	pkt = hex.WriteDoubleWord(pkt, m.LocationData.Longitude)
	pkt = hex.WriteWord(pkt, m.LocationData.Altitude)
	pkt = hex.WriteWord(pkt, m.LocationData.Speed)
	pkt = hex.WriteWord(pkt, m.LocationData.Direction)
	pkt = hex.WriteBCD(pkt, m.LocationData.Time)

	pkt, err = writeHeader(m, pkt)
	return pkt, err
}

func (m *Msg0200) GetHeader() *MsgHeader {
	return m.Header
}

func (m *Msg0200) GenOutgoing(_ JT808Msg) error {
	return nil
}
