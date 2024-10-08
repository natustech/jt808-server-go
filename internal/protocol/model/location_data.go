package model

import (
	"github.com/fakeyanss/jt808-server-go/internal/codec/hex"
	"github.com/rs/zerolog/log"
)

type BaseStation struct {
	Lac            string `json:"lac"`
	CellId         uint32 `json:"cellId"`
	SignalStrength uint8  `json:"signalStrength"`
}

type LocationData struct {
	AlarmSign    uint32        `json:"alarmSign"`   // 报警标志位
	StatusSign   uint32        `json:"statusSign"`  // 状态标志位
	Latitude     uint32        `json:"latitude"`    // 纬度，以度为单位的纬度值乘以10的6次方，精确到百万分之一度
	Longitude    uint32        `json:"longitude"`   // 精度，以度为单位的经度值乘以10的6次方，精确到百万分之一度
	Altitude     uint16        `json:"altitude"`    // 高程，海拔高度，单位为米(m)
	Speed        uint16        `json:"speed"`       // 速度，单位为0.1公里每小时(1/10km/h)
	Direction    uint16        `json:"direction"`   // 方向，0-359，正北为 0，顺时针
	Time         string        `json:"time"`        // YY-MM-DD-hh-mm-ss(GMT+8 时间)
	Electricity  uint8         `json:"electricity"` // Additional Informations
	MCC          uint16        `json:"mcc"`
	MNC          uint16        `json:"mnc"`
	BaseStations []BaseStation `json:"baseStations"`
}

func (m *LocationData) Decode(pkt []byte, idx *int) error {
	m.AlarmSign = hex.ReadDoubleWord(pkt, idx)
	m.StatusSign = hex.ReadDoubleWord(pkt, idx)
	m.Latitude = hex.ReadDoubleWord(pkt, idx)
	m.Longitude = hex.ReadDoubleWord(pkt, idx)
	m.Altitude = hex.ReadWord(pkt, idx)
	m.Speed = hex.ReadWord(pkt, idx)
	m.Direction = hex.ReadWord(pkt, idx)
	m.Time = hex.ReadBCD(pkt, idx, 6)

	// Additional Informations Decode:
	for !(*idx+1 >= len(pkt)) {
		addInfoId := hex.ReadByte(pkt, idx)
		addInfoLength := hex.ReadByte(pkt, idx)
		addInfoBody := hex.ReadBytes(pkt, idx, int(addInfoLength))

		log.Debug().Str("AddInfoId : ", string(addInfoId))
		log.Debug().Str("AddInfoLength : ", string(addInfoLength))

		addInfoIdx := 0
		switch addInfoId {
		case 0xE4:
			addInfoIdx++
			m.Electricity = hex.ReadByte(addInfoBody, &addInfoIdx)
		case 0xE1:
			/** 2+2+(3+4+1)*N First 4 byte is static */
			m.MCC = hex.ReadWord(pkt, &addInfoIdx)
			m.MNC = hex.ReadWord(pkt, &addInfoIdx)

			totalDataCount := (addInfoLength - 4) / 8
			m.BaseStations = make([]BaseStation, totalDataCount)

			for i := 0; i < int(totalDataCount); i++ {
				baseStation := &BaseStation{}
				baseStation.Lac = hex.ReadBCD(addInfoBody, &addInfoIdx, 3)
				baseStation.CellId = hex.ReadDoubleWord(addInfoBody, &addInfoIdx)
				baseStation.SignalStrength = hex.ReadByte(addInfoBody, &addInfoIdx)

				m.BaseStations[i] = *baseStation

			}
		}
	}

	return nil
}
