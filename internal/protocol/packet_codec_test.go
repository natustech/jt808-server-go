package protocol

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/fakeyanss/jt808-server-go/internal/codec/hex"
	"github.com/fakeyanss/jt808-server-go/internal/protocol/model"
)

func TestMsg704_Decode(t *testing.T) {
	type args struct {
		payload []byte
	}
	tests := []struct {
		name    string
		pc      *JT808PacketCodec
		args    args
		want    *model.PacketData
		wantErr bool
	}{
		{
			name: "case704",
			pc:   &JT808PacketCodec{},
			args: args{
				payload: hex.Str2Byte("7e070402cd04545004522931e6000601007500000000000c000002713d3201bc1f4700000000000024091617581601040000038b30011f310100e4020135e50100e60100e7080000000600000000ec2348225402bc1ed54e225402bc1ed54e22540303cac94822540303cac2c8544b1560dcc0e10c011e000100170d002f4f2500f50101fb0100007500000000000c000002713d3201bc1f4700000000000024091618031601040000038b30011f310100e4020132e50100e60100e7080000000600000000ec2348225402bc1ed54e225402bc1ed54e22540303cac94822540303cac2c8544b1560dcc0e10c011e000100170d002f4f2500f50101fb0100007500000000000c000002713d3201bc1f4700000000000024091618081601040000038b30011f310100e4020132e50100e60100e7080000000600000000ec2348225402bc1ed54e225402bc1ed54e22540303cac94822540303cac2c8544b1560dcc0e10c011e000100170d002f4f2500f50101fb0100007500000000000c000002713d3201bc1f4700000000000024091618131601040000038b30011f310100e4020132e50100e60100e7080000000600000000ec2348225402bc1ed54e225402bc1ed54e22540303cac94822540303cac2c8544b1560dcc0e10c011e000100170d002f4f2500f50101fb0100007500000000000c000002713d3201bc1f4700000000000024091618181701040000038b30011f310100e4020132e50100e60100e7080000000600000000ec2348225402bc1ed54e225402bc1ed54e22540303cac94822540303cac2c8544b1560dcc0e10c011e000100170d002f4f2500f50101fb0100007500000000000c000002713d3201bc1f4700000000000024091618231601040000038b30011e310100e4020131e50100e60100e7080000000600000000ec2348225402bc1ed54e225402bc1ed54e22540303cac94822540303cac2c8544b1560dcc0e10c011e000100170d002f4f2500f50101fb0100137e"),
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &JT808PacketCodec{}
			packet, err := pc.Decode(tt.args.payload)
			msg := &model.Msg0704{}
			msg.Decode(packet)
			b, err := json.MarshalIndent(msg, "", "  ")
			t.Error(string(b))

			require.Equal(t, tt.wantErr, err != nil, err)
			if tt.want == nil {
				return
			}
			require.Equal(t, tt.want, msg)
		})
	}
}

func TestJT808PacketCodec_Decode(t *testing.T) {
	type args struct {
		payload []byte
	}
	tests := []struct {
		name    string
		pc      *JT808PacketCodec
		args    args
		want    *model.PacketData
		wantErr bool
	}{
		{
			name: "case704",
			pc:   &JT808PacketCodec{},
			args: args{
				payload: hex.Str2Byte("7e070402cd04545004522931e6000601007500000000000c000002713d3201bc1f4700000000000024091617581601040000038b30011f310100e4020135e50100e60100e7080000000600000000ec2348225402bc1ed54e225402bc1ed54e22540303cac94822540303cac2c8544b1560dcc0e10c011e000100170d002f4f2500f50101fb0100007500000000000c000002713d3201bc1f4700000000000024091618031601040000038b30011f310100e4020132e50100e60100e7080000000600000000ec2348225402bc1ed54e225402bc1ed54e22540303cac94822540303cac2c8544b1560dcc0e10c011e000100170d002f4f2500f50101fb0100007500000000000c000002713d3201bc1f4700000000000024091618081601040000038b30011f310100e4020132e50100e60100e7080000000600000000ec2348225402bc1ed54e225402bc1ed54e22540303cac94822540303cac2c8544b1560dcc0e10c011e000100170d002f4f2500f50101fb0100007500000000000c000002713d3201bc1f4700000000000024091618131601040000038b30011f310100e4020132e50100e60100e7080000000600000000ec2348225402bc1ed54e225402bc1ed54e22540303cac94822540303cac2c8544b1560dcc0e10c011e000100170d002f4f2500f50101fb0100007500000000000c000002713d3201bc1f4700000000000024091618181701040000038b30011f310100e4020132e50100e60100e7080000000600000000ec2348225402bc1ed54e225402bc1ed54e22540303cac94822540303cac2c8544b1560dcc0e10c011e000100170d002f4f2500f50101fb0100007500000000000c000002713d3201bc1f4700000000000024091618231601040000038b30011e310100e4020131e50100e60100e7080000000600000000ec2348225402bc1ed54e225402bc1ed54e22540303cac94822540303cac2c8544b1560dcc0e10c011e000100170d002f4f2500f50101fb0100137e"),
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "case1",
			pc:   &JT808PacketCodec{},
			args: args{
				payload: hex.Str2Byte("7E0200001c2234567890150000000000000002080301cd48b50728d22b003902bc008f2301251438137d027E"),
			},
			want:    &model.PacketData{},
			wantErr: false,
		},
		{
			name: "case2",
			pc:   &JT808PacketCodec{},
			args: args{
				payload: hex.Str2Byte("7E0200001C2234567890150000000000000002080301CD779E0728C032003C0000008F230125145158FB7E"),
			},
			want: &model.PacketData{
				Header: &model.MsgHeader{
					MsgID: 0x0200,
					Attr: &model.MsgBodyAttr{
						BodyLength:       28,
						Encryption:       0b000,
						PacketFragmented: 0,
						VersionSign:      0,
						Extra:            0,

						VersionDesc: model.Version2013,
					},
					ProtocolVersion: 0,
					PhoneNumber:     "223456789015",
					SerialNumber:    0,
					Frag:            nil,
				},
				Body:       hex.Str2Byte("000000000002080301CD779E0728C032003C0000008F230125145158"),
				VerifyCode: 126,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &JT808PacketCodec{}
			got, err := pc.Decode(tt.args.payload)
			require.Equal(t, tt.wantErr, err != nil, err)
			if tt.want == nil {
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestJT808PacketCodec_genVerifier(t *testing.T) {
	type args struct {
		pkt []byte
	}
	tests := []struct {
		name string
		pc   *JT808PacketCodec
		args args
		want []byte
	}{
		{
			name: "case1",
			pc:   &JT808PacketCodec{},
			args: args{pkt: hex.Str2Byte("000140050100000000017299841738ffff007b01c803")},
			want: append(hex.Str2Byte("000140050100000000017299841738ffff007b01c803"), 0xb5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &JT808PacketCodec{}
			if got := pc.genVerifier(tt.args.pkt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JT808PacketCodec.genVerifier() = %v, want %v", got, tt.want)
			}
		})
	}
}
