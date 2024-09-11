package model

import (
	"github.com/fakeyanss/jt808-server-go/internal/codec/hex"
)

// Locate Data Bulk Upload
type Msg0704 struct {
	Header            *MsgHeader     `json:"header"`
	NumberOfDataItems uint16         `json:"numberOfDataItems"`
	LocationDataType  uint8          `json:"locationDataType"`
	LocationReports   []LocationData `json:"locationReports"`
}

func (m *Msg0704) Decode(packet *PacketData) error {
	m.Header = packet.Header
	pkt, idx := packet.Body, 0

	m.NumberOfDataItems = hex.ReadWord(pkt, &idx)
	m.LocationDataType = hex.ReadByte(pkt, &idx)
	m.LocationReports = make([]LocationData, m.NumberOfDataItems)

	// Append Location Reports
	for i := uint16(0); i < m.NumberOfDataItems; i++ {
		reportLength := hex.ReadWord(pkt, &idx)

		fakeIdx := idx

		l := &LocationData{}

		l.AlarmSign = hex.ReadDoubleWord(pkt, &fakeIdx)
		l.StatusSign = hex.ReadDoubleWord(pkt, &fakeIdx)
		l.Latitude = hex.ReadDoubleWord(pkt, &fakeIdx)
		l.Longitude = hex.ReadDoubleWord(pkt, &fakeIdx)
		l.Altitude = hex.ReadWord(pkt, &fakeIdx)
		l.Speed = hex.ReadWord(pkt, &fakeIdx)
		l.Direction = hex.ReadWord(pkt, &fakeIdx)
		l.Time = hex.ReadBCD(pkt, &fakeIdx, 6)

		m.LocationReports[i] = *l

		idx += int(reportLength)
	}

	return nil
}

func (m *Msg0704) Encode() (pkt []byte, err error) {
	// This message is just for receiving maybe TODO add
	pkt, err = writeHeader(m, pkt)
	return pkt, err
}

func (m *Msg0704) GetHeader() *MsgHeader {
	return m.Header
}

func (m *Msg0704) GenOutgoing(_ JT808Msg) error {
	return nil
}
