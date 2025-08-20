package clustersecurity

import "testing"

func TestCautionAreaToString(t *testing.T) {
	type args struct {
		c int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				c: 0x12345678,
			},
			want: "ABCDEFGH",
		},
		{
			name: "",
			args: args{
				c: 0x80000000,
			},
			want: "H-------",
		},
		{
			name: "",
			args: args{
				c: 0x87654321,
			},
			want: "HGFEDCBA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CautionAreaToString(tt.args.c); got != tt.want {
				t.Errorf("CautionAreaToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
