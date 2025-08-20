package sessionmanager

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/clustersecurity"
	"github.com/unknown321/fuse/fobplaced"
	"github.com/unknown321/fuse/fobweaponplacement"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/sectionstat"
	"github.com/unknown321/fuse/steamid"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
	"net"
)

func GetCmdGetFobTargetDetailResponse(ctx context.Context, msg *message.Message, manager *SessionManager, request *tppmessage.CmdGetFobTargetDetailRequest) tppmessage.CmdGetFobTargetDetailResponse {
	t := tppmessage.CmdGetFobTargetDetailResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_FOB_TARGET_DETAIL.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	// TODO what do these flags do?
	t.EventClearBit = 0
	t.IsRestrict = 0

	t.Detail = tppmessage.FobDetail{
		CapturedRankBottom: 0,   // always 0
		CapturedRankTop:    0,   // always 0
		CapturedStaff:      0,   // always 0
		RewardRate:         0,   // always 0
		Platform:           255, // 0 is nuclear only, 4 for emergency (invaded platform id?), 255 for everyone else

		// TODO from database? should be calculated based upon materials and ranks on target fob
		// rewards per fob platform type, for reaching each platform?
		PrimaryReward: []tppmessage.FobPrimaryReward{
			{
				RewardInfo: []tppmessage.FobRewardInfo{
					{
						BottomType: 0,
						Rate:       900000, // always 100% in ui
						Section:    100,    // 0-6, 100
						Type:       1,
						Value:      907500, // online gmp
					},
					{
						BottomType: 2,      // minimal volunteer rank, 2 = S++, 11 = E
						Rate:       805100, // acq probability, 80.51%
						Section:    0,
						Type:       2,  // 0,1=not showing, 2 = S++, 4 = S, type 5 = A++ -B
						Value:      40, // 40 volunteers
					},
					{
						BottomType: 0,
						Rate:       1000000,
						Section:    0,
						Type:       24, // 12 = common metal, 16 = biomaterial, 17 = golden crescent, 24 = haoma (max)
						Value:      3000,
					},
					{
						BottomType: 0,
						Rate:       1000000,
						Section:    0,
						Type:       23,
						Value:      400, // not showing on rewards screen
					},
				},
			},
			{
				RewardInfo: []tppmessage.FobRewardInfo{
					{
						BottomType: 0,
						Rate:       0,
						Section:    100,
						Type:       1,
						Value:      907500,
					},
					{
						BottomType: 8,
						Rate:       1000000,
						Section:    1,
						Type:       4,
						Value:      30,
					},
					{
						BottomType: 0,
						Rate:       1000000,
						Section:    0,
						Type:       12,
						Value:      6000,
					},
					{
						BottomType: 0,
						Rate:       1000000,
						Section:    0,
						Type:       24,
						Value:      200,
					},
				},
			},
			{
				RewardInfo: []tppmessage.FobRewardInfo{
					{
						BottomType: 0,
						Rate:       0,
						Section:    100,
						Type:       1,
						Value:      860000,
					},
					{
						BottomType: 8,
						Rate:       1000000,
						Section:    2,
						Type:       4,
						Value:      30,
					},
					{
						BottomType: 0,
						Rate:       1000000,
						Section:    0,
						Type:       14,
						Value:      2200,
					},
					{
						BottomType: 0,
						Rate:       1000000,
						Section:    0,
						Type:       21,
						Value:      1000,
					},
				},
			},
			{
				RewardInfo: []tppmessage.FobRewardInfo{
					{
						BottomType: 0,
						Rate:       0,
						Section:    100,
						Type:       1,
						Value:      860000,
					},
					{
						BottomType: 8,
						Rate:       1000000,
						Section:    3,
						Type:       4,
						Value:      20,
					},
					{
						BottomType: 0,
						Rate:       1000000,
						Section:    0,
						Type:       13,
						Value:      7000,
					},
					{
						BottomType: 0,
						Rate:       1000000,
						Section:    0,
						Type:       15,
						Value:      3000,
					},
				},
			},
			{
				RewardInfo: []tppmessage.FobRewardInfo{
					{
						BottomType: 0,
						Rate:       0,
						Section:    100,
						Type:       1,
						Value:      907500,
					},
					{
						BottomType: 8,
						Rate:       1000000,
						Section:    4,
						Type:       4,
						Value:      20,
					},
					{
						BottomType: 0,
						Rate:       1000000,
						Section:    0,
						Type:       15,
						Value:      8100,
					},
					{
						BottomType: 0,
						Rate:       1000000,
						Section:    0,
						Type:       22,
						Value:      300,
					},
				},
			},
			{
				RewardInfo: []tppmessage.FobRewardInfo{
					{
						BottomType: 0,
						Rate:       0,
						Section:    100,
						Type:       1,
						Value:      860000,
					},
					{
						BottomType: 8,
						Rate:       1000000,
						Section:    5,
						Type:       4,
						Value:      20,
					},
					{
						BottomType: 0,
						Rate:       1000000,
						Section:    0,
						Type:       17,
						Value:      1200,
					},
					{
						BottomType: 0,
						Rate:       1000000,
						Section:    0,
						Type:       18,
						Value:      600,
					},
				},
			},
			{
				RewardInfo: []tppmessage.FobRewardInfo{
					{
						BottomType: 0,
						Rate:       0,
						Section:    100,
						Type:       1,
						Value:      860000,
					},
					{
						BottomType: 8,
						Rate:       1000000,
						Section:    6,
						Type:       4,
						Value:      20,
					},
					{
						BottomType: 0,
						Rate:       1000000,
						Section:    0,
						Type:       19,
						Value:      900,
					},
					{
						BottomType: 0,
						Rate:       1000000,
						Section:    0,
						Type:       20,
						Value:      400,
					},
				},
			},
		},
		// TODO soldiers per rank, must be taken from SOLDIERS_BIN, HARDCODED fix me please
		SectionStaff: []tppmessage.FobSectionStaff{
			{
				Base:     0,
				Combat:   0,
				Develop:  0,
				Medical:  0,
				Security: 0,
				Spy:      0,
				Suport:   0,
			},
			{
				Base:     0,
				Combat:   0,
				Develop:  0,
				Medical:  0,
				Security: 0,
				Spy:      0,
				Suport:   0,
			},
			{
				Base:     0,
				Combat:   0,
				Develop:  0,
				Medical:  0,
				Security: 0,
				Spy:      0,
				Suport:   0,
			},
			{
				Base:     0,
				Combat:   0,
				Develop:  0,
				Medical:  0,
				Security: 0,
				Spy:      0,
				Suport:   0,
			},
			{
				Base:     0,
				Combat:   0,
				Develop:  0,
				Medical:  0,
				Security: 90,
				Spy:      0,
				Suport:   0,
			},
			{
				Base:     70,
				Combat:   40,
				Develop:  30,
				Medical:  30,
				Security: 30,
				Spy:      70,
				Suport:   70,
			},
			{
				Base:     70,
				Combat:   60,
				Develop:  100,
				Medical:  70,
				Security: 30,
				Spy:      50,
				Suport:   70,
			},
			{
				Base:     30,
				Combat:   20,
				Develop:  30,
				Medical:  30,
				Security: 20,
				Spy:      20,
				Suport:   35,
			},
			{
				Base:     1,
				Combat:   1,
				Develop:  1,
				Medical:  1,
				Security: 1,
				Spy:      1,
				Suport:   1,
			},
			{
				Base:     1,
				Combat:   1,
				Develop:  1,
				Medical:  1,
				Security: 1,
				Spy:      1,
				Suport:   1,
			},
		},
	}

	// placement
	{
		pl, err := manager.FOBWeaponPlacementRepo.Get(ctx, request.MotherBaseID)
		if err != nil {
			slog.Error("weapon placement", "error", err.Error(), "mother base ID", request.MotherBaseID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		if len(pl) < 1 {
			pl = append(pl, fobweaponplacement.WeaponPlacement{})
		}

		t.Detail.Placement = tppmessage.FobWeaponPlacement{
			EmplacementGunEast: pl[0].EmplacementGunEast,
			EmplacementGunWest: pl[0].EmplacementGunWest,
			GatlingGun:         pl[0].GatlingGun,
			GatlingGunEast:     pl[0].GatlingGunEast,
			GatlingGunWest:     pl[0].GatlingGunWest,
			MortarNormal:       pl[0].MortarNormal,
		}
	}

	var motherBaseParamID int
	// mbparam
	{
		param, err := manager.MotherBaseParamRepo.GetByMotherBaseID(ctx, request.MotherBaseID)
		if err != nil {
			slog.Error("get mother base param", "error", err.Error(), "motherBaseID", request.MotherBaseID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		if len(param) < 1 {
			slog.Error("mother base param", "error", "not found", "motherBaseID", request.MotherBaseID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		t.Detail.MotherBaseParam = tppmessage.MotherBaseParam{
			AreaID:         0, // always 0
			ClusterParam:   []tppmessage.ClusterParam{},
			ConstructParam: param[0].ConstructParam,
			FobIndex:       param[0].FobIndex,
			MotherBaseID:   param[0].ID,
			PlatformCount:  0, // always 0
			Price:          0, // always 0
			SecurityRank:   0, // always 0
		}

		motherBaseParamID = param[0].ID
		t.Detail.OwnerPlayerID = param[0].PlayerID
	}

	// cluster param
	{
		params, err := manager.ClusterParamRepo.Get(ctx, motherBaseParamID)
		if err != nil {
			slog.Error("get cluster param", "error", err.Error(), "motherBaseParamID", motherBaseParamID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		for _, v := range params {
			qq := tppmessage.ClusterParam{
				Build:           v.Build.ToInt(),
				ClusterSecurity: v.ClusterSecurityParam.ToInt(),
				SoldierRank:     v.SoldierRank,
				Common1Security: tppmessage.ClusterParamSecurity{},
				Common2Security: tppmessage.ClusterParamSecurity{},
				Common3Security: tppmessage.ClusterParamSecurity{},
				UniqueSecurity:  tppmessage.ClusterParamSecurity{},
			}

			// TODO a nice function to fill these
			// common1
			{
				s1, err := manager.ClusterSecurityRepo.Get(ctx, v.ID, 0)
				if err != nil {
					slog.Error("get cluster security", "error", err.Error(), "cluster param id", v.ID)
					t.Result = tppmessage.RESULT_ERR
					return t
				}
				if len(s1) == 0 {
					sec := clustersecurity.ClusterSecurity{}
					s1 = append(s1, sec)
				}
				qq.Common1Security = tppmessage.ClusterParamSecurity{
					Antitheft:                  s1[0].Antitheft,
					Camera:                     s1[0].Camera,
					CautionArea:                s1[0].CautionArea,
					Decoy:                      s1[0].Decoy,
					IrSensor:                   s1[0].IrSensor,
					Mine:                       s1[0].Mine,
					Soldier:                    s1[0].Soldier,
					Uav:                        s1[0].Uav,
					VoluntaryCoordCameraCount:  0,
					VoluntaryCoordCameraParams: []tppmessage.VoluntaryCoordParams{},
					VoluntaryCoordMineCount:    0,
					VoluntaryCoordMineParams:   []tppmessage.VoluntaryCoordParams{},
				}

				placed, err := manager.FOBPlacedRepo.Get(ctx, v.ID, 0)
				if err != nil {
					slog.Error("get cluster placed", "error", err.Error(), "cluster param id", v.ID)
					t.Result = tppmessage.RESULT_ERR
					return t
				}
				for _, p := range placed {
					vc := tppmessage.VoluntaryCoordParams{
						PlacedIndex: p.PlacedIndex,
						PositionX:   p.PositionX,
						PositionY:   p.PositionY,
						PositionZ:   p.PositionZ,
						RotationW:   p.RotationW,
						RotationX:   p.RotationX,
						RotationY:   p.RotationY,
						RotationZ:   p.RotationZ,
					}

					switch p.Type {
					case fobplaced.MINE:
						qq.Common1Security.VoluntaryCoordMineCount++
						qq.Common1Security.VoluntaryCoordMineParams = append(qq.Common1Security.VoluntaryCoordMineParams, vc)
					case fobplaced.CAMERA:
						qq.Common1Security.VoluntaryCoordCameraCount++
						qq.Common1Security.VoluntaryCoordCameraParams = append(qq.Common1Security.VoluntaryCoordCameraParams, vc)
					default:
						slog.Error("placed invalid type", "cluster param id", v.ID)
						t.Result = tppmessage.RESULT_ERR
						return t
					}
				}
			}

			// common2
			{
				s2, err := manager.ClusterSecurityRepo.Get(ctx, v.ID, 1)
				if err != nil {
					slog.Error("get cluster security", "error", err.Error(), "cluster param id", v.ID)
					t.Result = tppmessage.RESULT_ERR
					return t
				}

				if len(s2) == 0 {
					sec := clustersecurity.ClusterSecurity{}
					s2 = append(s2, sec)
				}

				qq.Common2Security = tppmessage.ClusterParamSecurity{
					Antitheft:                  s2[0].Antitheft,
					Camera:                     s2[0].Camera,
					CautionArea:                s2[0].CautionArea,
					Decoy:                      s2[0].Decoy,
					IrSensor:                   s2[0].IrSensor,
					Mine:                       s2[0].Mine,
					Soldier:                    s2[0].Soldier,
					Uav:                        s2[0].Uav,
					VoluntaryCoordCameraCount:  0,
					VoluntaryCoordCameraParams: []tppmessage.VoluntaryCoordParams{},
					VoluntaryCoordMineCount:    0,
					VoluntaryCoordMineParams:   []tppmessage.VoluntaryCoordParams{},
				}

				placed, err := manager.FOBPlacedRepo.Get(ctx, v.ID, 1)
				if err != nil {
					slog.Error("get cluster placed", "error", err.Error(), "cluster param id", v.ID)
					t.Result = tppmessage.RESULT_ERR
					return t
				}
				for _, p := range placed {
					vc := tppmessage.VoluntaryCoordParams{
						PlacedIndex: p.PlacedIndex,
						PositionX:   p.PositionX,
						PositionY:   p.PositionY,
						PositionZ:   p.PositionZ,
						RotationW:   p.RotationW,
						RotationX:   p.RotationX,
						RotationY:   p.RotationY,
						RotationZ:   p.RotationZ,
					}

					switch p.Type {
					case fobplaced.MINE:
						qq.Common2Security.VoluntaryCoordMineCount++
						qq.Common2Security.VoluntaryCoordMineParams = append(qq.Common2Security.VoluntaryCoordMineParams, vc)
					case fobplaced.CAMERA:
						qq.Common2Security.VoluntaryCoordCameraCount++
						qq.Common2Security.VoluntaryCoordCameraParams = append(qq.Common2Security.VoluntaryCoordCameraParams, vc)
					default:
						slog.Error("placed invalid type", "cluster param id", v.ID)
						t.Result = tppmessage.RESULT_ERR
						return t
					}
				}
			}

			// common3
			{
				s3, err := manager.ClusterSecurityRepo.Get(ctx, v.ID, 2)
				if err != nil {
					slog.Error("get cluster security", "error", err.Error(), "cluster param id", v.ID)
					t.Result = tppmessage.RESULT_ERR
					return t
				}
				if len(s3) == 0 {
					s3 = append(s3, clustersecurity.ClusterSecurity{})
				}
				qq.Common3Security = tppmessage.ClusterParamSecurity{
					Antitheft:                  s3[0].Antitheft,
					Camera:                     s3[0].Camera,
					CautionArea:                s3[0].CautionArea,
					Decoy:                      s3[0].Decoy,
					IrSensor:                   s3[0].IrSensor,
					Mine:                       s3[0].Mine,
					Soldier:                    s3[0].Soldier,
					Uav:                        s3[0].Uav,
					VoluntaryCoordCameraCount:  0,
					VoluntaryCoordCameraParams: []tppmessage.VoluntaryCoordParams{},
					VoluntaryCoordMineCount:    0,
					VoluntaryCoordMineParams:   []tppmessage.VoluntaryCoordParams{},
				}

				placed, err := manager.FOBPlacedRepo.Get(ctx, v.ID, 2)
				if err != nil {
					slog.Error("get cluster placed", "error", err.Error(), "cluster param id", v.ID)
					t.Result = tppmessage.RESULT_ERR
					return t
				}
				for _, p := range placed {
					vc := tppmessage.VoluntaryCoordParams{
						PlacedIndex: p.PlacedIndex,
						PositionX:   p.PositionX,
						PositionY:   p.PositionY,
						PositionZ:   p.PositionZ,
						RotationW:   p.RotationW,
						RotationX:   p.RotationX,
						RotationY:   p.RotationY,
						RotationZ:   p.RotationZ,
					}

					switch p.Type {
					case fobplaced.MINE:
						qq.Common3Security.VoluntaryCoordMineCount++
						qq.Common3Security.VoluntaryCoordMineParams = append(qq.Common3Security.VoluntaryCoordMineParams, vc)
					case fobplaced.CAMERA:
						qq.Common3Security.VoluntaryCoordCameraCount++
						qq.Common3Security.VoluntaryCoordCameraParams = append(qq.Common3Security.VoluntaryCoordCameraParams, vc)
					default:
						slog.Error("placed invalid type", "cluster param id", v.ID)
						t.Result = tppmessage.RESULT_ERR
						return t
					}
				}
			}

			// unique
			{
				s4, err := manager.ClusterSecurityRepo.Get(ctx, v.ID, 3)
				if err != nil {
					slog.Error("get cluster security", "error", err.Error(), "cluster param id", v.ID)
					t.Result = tppmessage.RESULT_ERR
					return t
				}
				if len(s4) == 0 {
					s4 = append(s4, clustersecurity.ClusterSecurity{})
				}
				qq.UniqueSecurity = tppmessage.ClusterParamSecurity{
					Antitheft:                  s4[0].Antitheft,
					Camera:                     s4[0].Camera,
					CautionArea:                s4[0].CautionArea,
					Decoy:                      s4[0].Decoy,
					IrSensor:                   s4[0].IrSensor,
					Mine:                       s4[0].Mine,
					Soldier:                    s4[0].Soldier,
					Uav:                        s4[0].Uav,
					VoluntaryCoordCameraCount:  0,
					VoluntaryCoordCameraParams: []tppmessage.VoluntaryCoordParams{},
					VoluntaryCoordMineCount:    0,
					VoluntaryCoordMineParams:   []tppmessage.VoluntaryCoordParams{},
				}

				placed, err := manager.FOBPlacedRepo.Get(ctx, v.ID, 3)
				if err != nil {
					slog.Error("get cluster placed", "error", err.Error(), "cluster param id", v.ID)
					t.Result = tppmessage.RESULT_ERR
					return t
				}
				for _, p := range placed {
					vc := tppmessage.VoluntaryCoordParams{
						PlacedIndex: p.PlacedIndex,
						PositionX:   p.PositionX,
						PositionY:   p.PositionY,
						PositionZ:   p.PositionZ,
						RotationW:   p.RotationW,
						RotationX:   p.RotationX,
						RotationY:   p.RotationY,
						RotationZ:   p.RotationZ,
					}

					switch p.Type {
					case fobplaced.MINE:
						qq.UniqueSecurity.VoluntaryCoordMineCount++
						qq.UniqueSecurity.VoluntaryCoordMineParams = append(qq.UniqueSecurity.VoluntaryCoordMineParams, vc)
					case fobplaced.CAMERA:
						qq.UniqueSecurity.VoluntaryCoordCameraCount++
						qq.UniqueSecurity.VoluntaryCoordCameraParams = append(qq.UniqueSecurity.VoluntaryCoordCameraParams, vc)
					default:
						slog.Error("placed invalid type", "cluster param id", v.ID)
						t.Result = tppmessage.RESULT_ERR
						return t
					}
				}
			}

			t.Detail.MotherBaseParam.ClusterParam = append(t.Detail.MotherBaseParam.ClusterParam, qq)
		}
	}

	// section rank
	{
		ss, err := manager.SectionStatRepo.GetBySectionID(ctx, msg.PlayerID, sectionstat.Security)
		if err != nil {
			slog.Error("section stat", "error", err.Error(), "msgID", t.Msgid, "playerID", msg.PlayerID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		if len(ss) < 1 {
			slog.Error("section stat", "error", "not found", "msgID", t.Msgid, "playerID", msg.PlayerID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		t.Detail.SecuritySectionRank = ss[0].Level
	}

	t.Session = tppmessage.FobSession{
		Ip:        "FILL_ME",
		IsInvalid: 0,
		Npid: tppmessage.Npid{ // always empty
			Handler: tppmessage.NpidHandler{
				Data:  "",
				Dummy: make([]int, 3),
				Term:  0,
			},
			Opt:      make([]int, 8),
			Reserved: make([]int, 8),
		},
		Port:                5733,
		SecureDeviceAddress: "NotImplement",
		Steamid:             steamid.InvalidSteamID,
		Xnaddr:              "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
		Xnkey:               "",
		Xnkid:               "",
		Xuid:                steamid.InvalidSteamID,
	}

	// intruder
	{
		intr, err := manager.IntruderRepo.GetByOwnerID(ctx, msg.PlayerID)
		if err != nil {
			slog.Error("get intruder", "error", err.Error(), "mother base ID", request.MotherBaseID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		var playerID int

		if len(intr) > 0 {
			if intr[0].PlayerID == msg.PlayerID {
				slog.Warn("attempting to intrude yourself?", "playerID", msg.PlayerID)
			}
			// set only on emergency
			t.Session.Xnkey = "AAAAAAAAAAAAAAAAAAAAAA=="
			t.Session.Xnkid = "AAAAAAAAAAA="
			playerID = intr[0].PlayerID
			t.Detail.Platform = intr[0].MotherBasePlatformID
		} else {
			slog.Info("intruder not found", "mbID", request.MotherBaseID)
			playerID = msg.PlayerID
		}

		sess, err := manager.GetByPlayerID(playerID)
		if err != nil {
			slog.Error("intruder session not found", "error", err.Error(), "playerID", playerID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		//slog.Warn("player", "id", fmt.Sprintf("%+v", sess))

		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, uint32(sess.ExIp))
		t.Session.Ip = ip.String()
		t.Session.Xuid = sess.PlatformID
		t.Session.Steamid = sess.PlatformID
	}

	return t
}

func HandleCmdGetFobTargetDetailResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	//t := GetCmdGetFobTargetDetailResponse()
	t := tppmessage.CmdGetFobTargetDetailResponse{}

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
