package clustersecurityparam

import (
	"fmt"
	"github.com/unknown321/fuse/guardrank"
	"github.com/unknown321/fuse/weaponrange"
	"reflect"
	"testing"
)

func TestClusterSecurityParam_FromInt(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		fields  ClusterSecurityParam
		args    args
		wantErr bool
	}{
		{
			name: "",
			fields: ClusterSecurityParam{
				DefenseLevel:   78,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.SPlusPlus,
				EquipmentGrade: 11,
				WeaponRange:    weaponrange.Long,
				HasGuards:      1,
				Bit25_31:       0,
			},
			args: args{
				i: 322158,
			},
			wantErr: false,
		},
		{
			name: "",
			fields: ClusterSecurityParam{
				DefenseLevel:   25,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.SPlusPlus,
				EquipmentGrade: 5,
				WeaponRange:    weaponrange.Mid,
			},
			args: args{
				i: 102997,
			},
			wantErr: false,
		},
		{
			name: "",
			fields: ClusterSecurityParam{
				DefenseLevel:   35,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.SPlusPlus,
				EquipmentGrade: 5,
				WeaponRange:    weaponrange.Mid,
			},
			args: args{
				i: 143957,
			},
			wantErr: false,
		},
		{
			name: "prep level 0",
			fields: ClusterSecurityParam{
				DefenseLevel:   1,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.E,
				EquipmentGrade: 1,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      0,
				Bit25_31:       0,
			},
			args: args{
				i: 4101,
			},
			wantErr: false,
		},
		{
			name: "prep level 0, guards 1",
			fields: ClusterSecurityParam{
				DefenseLevel:   1,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.E,
				EquipmentGrade: 1,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      1,
				Bit25_31:       0,
			},
			args: args{
				i: 6149,
			},
			wantErr: false,
		},
		{
			name: "prep level 0, guards 12",
			fields: ClusterSecurityParam{
				DefenseLevel:   10,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.E,
				EquipmentGrade: 1,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      1,
				Bit25_31:       0,
			},
			args: args{
				i: 43013,
			},
			wantErr: false,
		},
		{
			name: "prep level 0, guard rank E",
			fields: ClusterSecurityParam{
				DefenseLevel:   10,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.E,
				EquipmentGrade: 1,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      1,
				Bit25_31:       0,
			},
			args: args{
				i: 43013,
			},
			wantErr: false,
		},
		{
			name: "prep level 0, guard rank S++",
			fields: ClusterSecurityParam{
				DefenseLevel:   35,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.SPlusPlus,
				EquipmentGrade: 1,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      1,
				Bit25_31:       0,
			},
			args: args{
				i: 145989,
			},
			wantErr: false,
		},
		{
			name: "prep level 0, guard rank A+",
			fields: ClusterSecurityParam{
				DefenseLevel:   30,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.APlus,
				EquipmentGrade: 1,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      1,
				Bit25_31:       0,
			},
			args: args{
				i: 125253,
			},
			wantErr: false,
		},
		{
			name: "key zone disabled/enabled",
			fields: ClusterSecurityParam{
				DefenseLevel:   10,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.E,
				EquipmentGrade: 6,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      1,
			},
			args: args{
				i: 43033,
			},
		},
		{
			name: "equipment grade 3",
			fields: ClusterSecurityParam{
				DefenseLevel:   9,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.E,
				EquipmentGrade: 3,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      1,
			},
			args: args{
				i: 38925,
			},
		},
		{
			name: "equipment grade 6",
			fields: ClusterSecurityParam{
				DefenseLevel:   9,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.E,
				EquipmentGrade: 6,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      1,
			},
			args: args{
				i: 38937,
			},
		},
		{
			name: "swimsuit olive drab",
			fields: ClusterSecurityParam{
				DefenseLevel:   10,
				NonLethal:      0,
				SwimsuitID:     1,
				GuardRank:      guardrank.C,
				EquipmentGrade: 1,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      1,
			},
			args: args{
				i: 567429,
			},
			wantErr: false,
		},
		{
			name: "swimsuit tiger stripe",
			fields: ClusterSecurityParam{
				DefenseLevel:   10,
				NonLethal:      0,
				SwimsuitID:     2,
				GuardRank:      guardrank.C,
				EquipmentGrade: 1,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      1,
			},
			args: args{
				i: 1091717,
			},
			wantErr: false,
		},
		{
			name: "swimsuit meg(mud) (max option)",
			fields: ClusterSecurityParam{
				DefenseLevel:   10,
				NonLethal:      0,
				SwimsuitID:     36,
				GuardRank:      guardrank.C,
				EquipmentGrade: 1,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      1,
				Bit25_31:       0,
			},
			args: args{
				i: 18917509,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ClusterSecurityParam{}
			if err := c.FromInt(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("FromInt() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(c, tt.fields) {
				t.Errorf("not equal %+v", c)
			}
		})
	}
}

func TestClusterSecurityParam_ToInt(t *testing.T) {
	type want struct {
		i int
	}
	tests := []struct {
		name    string
		fields  ClusterSecurityParam
		want    want
		wantErr bool
	}{
		{
			name: "",
			fields: ClusterSecurityParam{
				DefenseLevel:   78,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.SPlusPlus,
				EquipmentGrade: 11,
				WeaponRange:    weaponrange.Long,
				HasGuards:      1,
				Bit25_31:       0,
			},
			want: want{
				i: 322158,
			},
			wantErr: false,
		},
		{
			name: "",
			fields: ClusterSecurityParam{
				DefenseLevel:   25,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.SPlusPlus,
				EquipmentGrade: 5,
				WeaponRange:    weaponrange.Mid,
			},
			want: want{
				i: 102997,
			},
			wantErr: false,
		},
		{
			name: "",
			fields: ClusterSecurityParam{
				DefenseLevel:   35,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.SPlusPlus,
				EquipmentGrade: 5,
				WeaponRange:    weaponrange.Mid,
			},
			want: want{
				i: 143957,
			},
			wantErr: false,
		},
		{
			name: "prep level 0",
			fields: ClusterSecurityParam{
				DefenseLevel:   1,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.E,
				EquipmentGrade: 1,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      0,
				Bit25_31:       0,
			},
			want: want{
				i: 4101,
			},
			wantErr: false,
		},
		{
			name: "prep level 0, guards 1",
			fields: ClusterSecurityParam{
				DefenseLevel:   1,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.E,
				EquipmentGrade: 1,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      1,
				Bit25_31:       0,
			},
			want: want{
				i: 6149,
			},
			wantErr: false,
		},
		{
			name: "prep level 0, guards 12",
			fields: ClusterSecurityParam{
				DefenseLevel:   10,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.E,
				EquipmentGrade: 1,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      1,
				Bit25_31:       0,
			},
			want: want{
				i: 43013,
			},
			wantErr: false,
		},
		{
			name: "prep level 0, guard rank E",
			fields: ClusterSecurityParam{
				DefenseLevel:   10,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.E,
				EquipmentGrade: 1,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      1,
				Bit25_31:       0,
			},
			want: want{
				i: 43013,
			},
			wantErr: false,
		},
		{
			name: "prep level 0, guard rank S++",
			fields: ClusterSecurityParam{
				DefenseLevel:   35,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.SPlusPlus,
				EquipmentGrade: 1,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      1,
				Bit25_31:       0,
			},
			want: want{
				i: 145989,
			},
			wantErr: false,
		},
		{
			name: "prep level 0, guard rank A+",
			fields: ClusterSecurityParam{
				DefenseLevel:   30,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.APlus,
				EquipmentGrade: 1,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      1,
				Bit25_31:       0,
			},
			want: want{
				i: 125253,
			},
			wantErr: false,
		},
		{
			name: "key zone disabled/enabled",
			fields: ClusterSecurityParam{
				DefenseLevel:   10,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.E,
				EquipmentGrade: 6,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      1,
			},
			want: want{
				i: 43033,
			},
		},
		{
			name: "equipment grade 3",
			fields: ClusterSecurityParam{
				DefenseLevel:   9,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.E,
				EquipmentGrade: 3,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      1,
			},
			want: want{
				i: 38925,
			},
		},
		{
			name: "equipment grade 6",
			fields: ClusterSecurityParam{
				DefenseLevel:   9,
				NonLethal:      0,
				SwimsuitID:     0,
				GuardRank:      guardrank.E,
				EquipmentGrade: 6,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      1,
			},
			want: want{
				i: 38937,
			},
		},
		{
			name: "swimsuit olive drab",
			fields: ClusterSecurityParam{
				DefenseLevel:   10,
				NonLethal:      0,
				SwimsuitID:     1,
				GuardRank:      guardrank.C,
				EquipmentGrade: 1,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      1,
			},
			want: want{
				i: 567429,
			},
			wantErr: false,
		},
		{
			name: "swimsuit tiger stripe",
			fields: ClusterSecurityParam{
				DefenseLevel:   10,
				NonLethal:      0,
				SwimsuitID:     2,
				GuardRank:      2,
				EquipmentGrade: 1,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      1,
			},
			want: want{
				i: 1091717,
			},
			wantErr: false,
		},
		{
			name: "swimsuit meg(mud) (max option)",
			fields: ClusterSecurityParam{
				DefenseLevel:   10,
				NonLethal:      0,
				SwimsuitID:     36,
				GuardRank:      guardrank.C,
				EquipmentGrade: 1,
				WeaponRange:    weaponrange.Mid,
				HasGuards:      1,
			},
			want: want{
				i: 18917509,
			},
			wantErr: false,
		},

		{
			name: "for testing only",
			fields: ClusterSecurityParam{
				DefenseLevel:   127,
				NonLethal:      1,
				SwimsuitID:     0,
				GuardRank:      guardrank.E,
				EquipmentGrade: 0,
				WeaponRange:    weaponrange.Close,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ClusterSecurityParam{
				DefenseLevel:   tt.fields.DefenseLevel,
				NonLethal:      tt.fields.NonLethal,
				SwimsuitID:     tt.fields.SwimsuitID,
				GuardRank:      tt.fields.GuardRank,
				EquipmentGrade: tt.fields.EquipmentGrade,
				WeaponRange:    tt.fields.WeaponRange,
				Bit25_31:       tt.fields.Bit25_31,
				HasGuards:      tt.fields.HasGuards,
			}
			if got := c.ToInt(); got != tt.want.i {
				if tt.wantErr {
					fmt.Printf("ToInt() = %d, want %d, diff %d (it's ok)\n", got, tt.want.i, got-tt.want.i)
				} else {
					t.Errorf("ToInt() = %d, want %d, diff %d", got, tt.want.i, got-tt.want.i)
				}
			}
		})
	}
}

func TestBulk(t *testing.T) {
	var securities = []int{
		102997,
		143957,
		162402,
		166498,
		204398,
		231910,
		233058,
		242285,
		272874,
		276970,
		4101,
		51244,
		52718,
		56814,
		6193, // eq grade 12?
	}

	for _, v := range securities {
		ss := ClusterSecurityParam{}
		if err := ss.FromInt(v); err != nil {
			t.Fatalf("%s", err.Error())
		}
		//fmt.Println(v, ss.EquipmentGrade)
	}
}
