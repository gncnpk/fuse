package sessionmanager

import (
	"reflect"
	"testing"
)

func TestReadSoldierData(t *testing.T) {
	type args struct {
		soldierParams string
	}
	tests := []struct {
		name    string
		args    args
		want    []Solly
		wantErr bool
	}{
		{
			name: "",
			args: args{
				soldierParams: "AAEAAQAEAAQq0B+A8OcAAAnACYAAAB1QAAEABQAFAAUqoB+A9DDwACpAAIAAAB1Q",
			},
			want: []Solly{{
				DataClient: SollyDataClient{
					DirectContract:  1,
					MaybeGene:       1,
					MaybeBaseStats:  4,
					MaybeBaseStats2: 4,
				},
				DataServer: SollyDataServer{Header: [4]byte{0x2a, 0xd0, 0x1f, 0x80},
					Data: [12]byte{0xf0, 0xe7, 0x00, 0x00, 0x09, 0xc0, 0x09, 0x80, 0x00, 0x00, 0x1d, 0x50}},
			}, {
				DataClient: SollyDataClient{
					DirectContract:  1,
					MaybeGene:       5,
					MaybeBaseStats:  5,
					MaybeBaseStats2: 5,
				},
				DataServer: SollyDataServer{Header: [4]byte{0x2a, 0xa0, 0x1f, 0x80},
					Data: [12]byte{0xf4, 0x30, 0xf0, 0x00, 0x2a, 0x40, 0x00, 0x80, 0x00, 0x00, 0x1d, 0x50}},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadSoldierData(tt.args.soldierParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadSoldierData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadSoldierData() got = %v, want %v", got, tt.want)
			}
		})
	}
}
