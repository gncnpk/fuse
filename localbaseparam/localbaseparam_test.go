package localbaseparam

import (
	"reflect"
	"testing"
)

func TestLocalBaseParam_FromInt(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		fields  LocalBaseParam
		args    args
		wantErr bool
	}{
		{
			name: "",
			fields: LocalBaseParam{
				PlatformsBuilt: 4,
				Mystery0:       0,
				Mystery1:       1,
			},
			args: args{
				i: 16385,
			},
			wantErr: false,
		},
		{
			name: "",
			fields: LocalBaseParam{
				PlatformsBuilt: 3,
				Mystery0:       0,
				Mystery1:       1,
			},
			args: args{
				i: 12289,
			},
			wantErr: false,
		},
		{
			name: "",
			fields: LocalBaseParam{
				PlatformsBuilt: 2,
				Mystery0:       0,
				Mystery1:       1,
			},
			args: args{
				i: 8193,
			},
			wantErr: false,
		},
		{
			name: "",
			fields: LocalBaseParam{
				PlatformsBuilt: 1,
				Mystery0:       0,
				Mystery1:       1,
			},
			args: args{
				i: 4097,
			},
			wantErr: false,
		},
		{
			name: "",
			fields: LocalBaseParam{
				PlatformsBuilt: 0,
				Mystery0:       0,
				Mystery1:       0,
			},
			args: args{
				i: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := LocalBaseParam{}
			if err := c.FromInt(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("FromInt() error = %+v, wantErr %+v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(c, tt.fields) {
				t.Errorf("not equal %+v", c)
			}

		})
	}
}

func TestLocalBaseParam_ToInt(t *testing.T) {
	tests := []struct {
		name   string
		fields LocalBaseParam
		want   int
	}{
		{
			name: "",
			fields: LocalBaseParam{
				PlatformsBuilt: 4,
				Mystery0:       0,
				Mystery1:       1,
			},
			want: 16385,
		},
		{
			name: "",
			fields: LocalBaseParam{
				PlatformsBuilt: 1,
				Mystery0:       0,
				Mystery1:       1,
			},
			want: 4097,
		},
		{
			name: "",
			fields: LocalBaseParam{
				PlatformsBuilt: 2,
				Mystery0:       0,
				Mystery1:       1,
			},
			want: 8193,
		},
		{
			name: "",
			fields: LocalBaseParam{
				PlatformsBuilt: 3,
				Mystery0:       0,
				Mystery1:       1,
			},
			want: 12289,
		},
		{
			name: "",
			fields: LocalBaseParam{
				PlatformsBuilt: 0,
				Mystery0:       0,
				Mystery1:       0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LocalBaseParam{
				PlatformsBuilt: tt.fields.PlatformsBuilt,
				Mystery0:       tt.fields.Mystery0,
				Mystery1:       tt.fields.Mystery1,
			}
			if got := c.ToInt(); got != tt.want {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
