package model

import (
	"errors"

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

	// Append Location Reports
	for i := uint16(0); i < m.NumberOfDataItems; i++ {
		reportLength := hex.ReadWord(pkt, &idx)

		var l LocationData

		if err := l.Decode(pkt, &idx, int(reportLength)); err != nil {
			return err
		}

		m.LocationReports[i] = l
		idx += int(reportLength) - 2
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

func (l *LocationData) Decode(pkt []byte, idx *int, length int) error {
	if len(pkt) < *idx+length {
		return errors.New("packet too short for LocationData")
	}

	// Save starting index
	startIdx := *idx

	// Decode fields based on length
	l.AlarmSign = hex.ReadDoubleWord(pkt, idx)
	l.StatusSign = hex.ReadDoubleWord(pkt, idx)
	l.Latitude = hex.ReadDoubleWord(pkt, idx)
	l.Longitude = hex.ReadDoubleWord(pkt, idx)
	l.Altitude = hex.ReadWord(pkt, idx)
	l.Speed = hex.ReadWord(pkt, idx)
	l.Direction = hex.ReadWord(pkt, idx)
	l.Time = hex.ReadBCD(pkt, idx, 6)

	// Check that we have read the exact amount of data we expected
	if *idx-startIdx != length {
		return errors.New("unexpected length for LocationData")
	}

	return nil
}
