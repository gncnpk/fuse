package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/clusterparam"
	"github.com/unknown321/fuse/clustersecurity"
	"github.com/unknown321/fuse/clustersecurityparam"
	"github.com/unknown321/fuse/equipflag"
	"github.com/unknown321/fuse/equipgrade"
	"github.com/unknown321/fuse/fobplaced"
	"github.com/unknown321/fuse/localbase"
	"github.com/unknown321/fuse/localbaseparam"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/motherbaseparam"
	"github.com/unknown321/fuse/pfskillstaff"
	"github.com/unknown321/fuse/securitylevel"
	"github.com/unknown321/fuse/tapeflag"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func HandleCmdSyncMotherBaseRequest(ctx context.Context, msg *message.Message, manager *SessionManager) error {
	var err error

	t := tppmessage.CmdSyncMotherBaseRequest{}
	err = json.Unmarshal(msg.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	if t.InvalidFob != 0 {
		slog.Info("invalid fob", "value", t.InvalidFob, "playerID", msg.PlayerID, "msgid", t.Msgid)
	}

	if t.Flag != "SYNC" {
		slog.Info("unexpected flag", "value", t.Flag, "playerID", msg.PlayerID, "msgid", t.Msgid)
	}

	data := FromJSON(ctx, t.Msgid)
	if data != nil {
		msg.MData = data
		return nil
	}

	d := GetCmdSyncMotherBaseResponse()

	// TODO figure out equip flag
	// equip flag
	{
		ef := equipflag.EquipFlag{PlayerID: msg.PlayerID}
		if err = ef.FromArray(t.EquipFlag); err != nil {
			slog.Error("flag", "error", err.Error())
			d.Result = tppmessage.RESULT_ERR
			goto marshalAndReturn
		}

		if err = manager.EquipFlagRepo.AddOrUpdate(ctx, &ef); err != nil {
			slog.Error("flag", "error", err.Error())
			d.Result = tppmessage.RESULT_ERR
			goto marshalAndReturn
		}
	}

	// equip grade
	{
		ef := equipgrade.EquipGrade{PlayerID: msg.PlayerID}
		if err = ef.FromArray(t.EquipGrade); err != nil {
			slog.Error("grade", "error", err.Error())
			d.Result = tppmessage.RESULT_ERR
			goto marshalAndReturn
		}

		if err = manager.EquipGradeRepo.AddOrUpdate(ctx, &ef); err != nil {
			slog.Error("grade", "error", err.Error())
			d.Result = tppmessage.RESULT_ERR
			goto marshalAndReturn
		}
	}

	// tape flag
	{
		ef := tapeflag.TapeFlag{PlayerID: msg.PlayerID}
		if err = ef.FromArray(t.TapeFlag); err != nil {
			slog.Error("tape flag", "error", err.Error())
			d.Result = tppmessage.RESULT_ERR
			goto marshalAndReturn
		}

		if err = manager.TapeFlagRepo.AddOrUpdate(ctx, &ef); err != nil {
			slog.Error("tape flag", "error", err.Error())
			d.Result = tppmessage.RESULT_ERR
			goto marshalAndReturn
		}
	}

	// security level
	{
		ef := securitylevel.SecurityLevel{PlayerID: msg.PlayerID}
		if err = ef.FromArray(t.SecurityLevel); err != nil {
			slog.Error("security level", "error", err.Error())
			d.Result = tppmessage.RESULT_ERR
			goto marshalAndReturn
		}

		if err = manager.SecurityLevelRepo.AddOrUpdate(ctx, &ef); err != nil {
			slog.Error("security level", "error", err.Error())
			d.Result = tppmessage.RESULT_ERR
			goto marshalAndReturn
		}
	}

	// local base
	{
		ef := localbase.LocalBase{PlayerID: msg.PlayerID}
		if err = ef.WithTime(t.LocalBaseTime); err != nil {
			slog.Error("local base", "error", err.Error())
			d.Result = tppmessage.RESULT_ERR
			goto marshalAndReturn
		}

		if err = ef.WithParam(t.LocalBaseParam); err != nil {
			slog.Error("local base", "error", err.Error())
			d.Result = tppmessage.RESULT_ERR
			goto marshalAndReturn
		}

		ef.MotherBaseNum = t.MotherBaseNum
		ef.NamePlateID = t.NamePlateID
		ef.PickupOpen = t.PickupOpen
		ef.SectionOpen = t.SectionOpen

		if err = manager.LocalBaseRepo.AddOrUpdate(ctx, &ef); err != nil {
			slog.Error("security level", "error", err.Error())
			d.Result = tppmessage.RESULT_ERR
			goto marshalAndReturn
		}
	}

	// pf skill staff
	{
		ef := pfskillstaff.PFSkillStaff{PlayerID: msg.PlayerID}
		ef.AllStaffNum = t.PfSkillStaff.AllStaffNum
		ef.Defender1Num = t.PfSkillStaff.Defender1Num
		ef.Defender2Num = t.PfSkillStaff.Defender2Num
		ef.Defender3Num = t.PfSkillStaff.Defender3Num
		ef.InterceptorMissile1Num = t.PfSkillStaff.InterceptorMissile1Num
		ef.InterceptorMissile2Num = t.PfSkillStaff.InterceptorMissile2Num
		ef.InterceptorMissile3Num = t.PfSkillStaff.InterceptorMissile3Num
		ef.LiquidCarbonMissile1Num = t.PfSkillStaff.LiquidCarbonMissile1Num
		ef.LiquidCarbonMissile2Num = t.PfSkillStaff.LiquidCarbonMissile2Num
		ef.LiquidCarbonMissile3Num = t.PfSkillStaff.LiquidCarbonMissile3Num
		ef.Medic1Num = t.PfSkillStaff.Medic1Num
		ef.Medic2Num = t.PfSkillStaff.Medic2Num
		ef.Medic3Num = t.PfSkillStaff.Medic3Num
		ef.Ranger1Num = t.PfSkillStaff.Ranger1Num
		ef.Ranger2Num = t.PfSkillStaff.Ranger2Num
		ef.Ranger3Num = t.PfSkillStaff.Ranger3Num
		ef.Sentry1Num = t.PfSkillStaff.Sentry1Num
		ef.Sentry2Num = t.PfSkillStaff.Sentry2Num
		ef.Sentry3Num = t.PfSkillStaff.Sentry3Num

		if err = manager.PFSkillStaffRepo.AddOrUpdate(ctx, &ef); err != nil {
			slog.Error("pf skill staff", "error", err.Error())
			d.Result = tppmessage.RESULT_ERR
			goto marshalAndReturn
		}
	}

	// mother base param
	{
		for i, v := range t.MotherBaseParam {
			mbp := motherbaseparam.MotherBaseParam{
				PlayerID:       msg.PlayerID,
				ConstructParam: v.ConstructParam,
				FobIndex:       i,
				PlatformCount:  v.PlatformCount,
				Price:          v.Price,
				SecurityRank:   v.SecurityRank,
			}

			/*
				fob index is always 0 for this request, this check is redundant

				if v.FobIndex != i {
					slog.Warn("fob index not equal", "index", v.FobIndex, "i", i)
				}
			*/

			params, err := manager.MotherBaseParamRepo.Get(ctx, msg.PlayerID, i)
			if err != nil {
				slog.Error("get mb param", "error", err.Error(), "msgid", msg.MsgID)
				d.Result = tppmessage.RESULT_ERR
				goto marshalAndReturn
			}

			var mbpID int
			if len(params) < 1 {
				mbp.FobIndex = i
				if mbpID, err = manager.MotherBaseParamRepo.AddOrUpdate(ctx, &mbp); err != nil {
					slog.Error("update mb param", "error", err.Error(), "msgid", msg.MsgID)
					d.Result = tppmessage.RESULT_ERR
					goto marshalAndReturn
				}
			} else {
				mbpID = params[0].ID
			}

			for platformID, cp := range v.ClusterParam { // 7 clusters per 1 mb
				cps, err := manager.ClusterParamRepo.GetByPlatformID(ctx, mbpID, platformID)
				if err != nil {
					slog.Error("get cluster param", "error", err.Error(), "msgid", msg.MsgID)
					d.Result = tppmessage.RESULT_ERR
					goto marshalAndReturn
				}

				var clusterParamID int
				if len(cps) < 1 {
					build := localbaseparam.LocalBaseParam{}
					if err = build.FromInt(cp.Build); err != nil {
						slog.Error("local base param parse", "error", err.Error())
						d.Result = tppmessage.RESULT_ERR
						goto marshalAndReturn
					}
					if build.ToInt() != cp.Build {
						slog.Error("build pack validation failed", "value", cp.Build)
						d.Result = tppmessage.RESULT_ERR
						goto marshalAndReturn
					}

					secParam := clustersecurityparam.ClusterSecurityParam{}
					if err = secParam.FromInt(cp.ClusterSecurity); err != nil {
						slog.Error("security param parse", "error", err.Error())
						d.Result = tppmessage.RESULT_ERR
						goto marshalAndReturn
					}
					if secParam.ToInt() != cp.ClusterSecurity {
						slog.Error("security param pack validation failed", "cluster_security", cp.ClusterSecurity)
						d.Result = tppmessage.RESULT_ERR
						goto marshalAndReturn
					}

					cluP := clusterparam.ClusterParam{
						ID:                   0,
						PlatformID:           platformID,
						MotherBaseParamID:    mbpID,
						Build:                build,
						ClusterSecurityParam: secParam,
						SoldierRank:          cp.SoldierRank,
					}

					if clusterParamID, err = manager.ClusterParamRepo.Add(ctx, &cluP); err != nil {
						slog.Error("cluster param", "error", err.Error(), "msgid", msg.MsgID)
						d.Result = tppmessage.RESULT_ERR
						goto marshalAndReturn
					}
					cluP.ID = clusterParamID

					cps = append(cps, cluP)
				}

				for _, cluP := range cps {
					clusterParamID = cluP.ID
					cluP.SoldierRank = cp.SoldierRank

					build := localbaseparam.LocalBaseParam{}
					if err := build.FromInt(cp.Build); err != nil {
						slog.Error("local base param parse", "error", err.Error())
						d.Result = tppmessage.RESULT_ERR
						goto marshalAndReturn
					}
					if build.ToInt() != cp.Build {
						slog.Error("build pack validation failed", "value", cp.Build)
						d.Result = tppmessage.RESULT_ERR
						goto marshalAndReturn
					}
					cluP.Build = build

					secParam := clustersecurityparam.ClusterSecurityParam{}
					if err = secParam.FromInt(cp.ClusterSecurity); err != nil {
						slog.Error("security param parse", "error", err.Error())
						d.Result = tppmessage.RESULT_ERR
						goto marshalAndReturn
					}
					if secParam.ToInt() != cp.ClusterSecurity {
						slog.Error("security param pack validation failed", "cluster_security", cp.ClusterSecurity)
						d.Result = tppmessage.RESULT_ERR
						goto marshalAndReturn
					}
					cluP.ClusterSecurityParam = secParam

					if _, err = manager.ClusterParamRepo.Update(ctx, &cluP); err != nil {
						slog.Error("cluster param", "error", err.Error(), "msgid", msg.MsgID)
						d.Result = tppmessage.RESULT_ERR
						goto marshalAndReturn
					}

					sec, err := manager.ClusterSecurityRepo.Get(ctx, clusterParamID, 0)
					if err != nil {
						slog.Error("cluster security get", "error", err.Error(), "msgid", msg.MsgID)
						d.Result = tppmessage.RESULT_ERR
						goto marshalAndReturn
					}
					found := true
					if len(sec) < 1 {
						cs0 := clustersecurity.ClusterSecurity{
							ClusterParamID: clusterParamID,
							IDX:            0,
							IsUnique:       0,
						}
						sec = append(sec, cs0)
						found = false
					}
					sec[0].Antitheft = cp.Common1Security.Antitheft
					sec[0].Camera = cp.Common1Security.Camera
					sec[0].CautionArea = cp.Common1Security.CautionArea
					sec[0].Decoy = cp.Common1Security.Decoy
					sec[0].IrSensor = cp.Common1Security.IrSensor
					sec[0].Mine = cp.Common1Security.Mine
					sec[0].Soldier = cp.Common1Security.Soldier
					sec[0].Uav = cp.Common1Security.Uav
					sec[0].VoluntaryCoordCameraCount = cp.Common1Security.VoluntaryCoordCameraCount
					sec[0].VoluntaryCoordMineCount = cp.Common1Security.VoluntaryCoordMineCount
					if !found {
						if err = manager.ClusterSecurityRepo.Add(ctx, &sec[0]); err != nil {
							slog.Error("cluster security add", "error", err.Error(), "msgid", msg.MsgID)
							d.Result = tppmessage.RESULT_ERR
							goto marshalAndReturn
						}
					} else {
						if err = manager.ClusterSecurityRepo.Update(ctx, &sec[0]); err != nil {
							slog.Error("cluster security update", "error", err.Error(), "msgid", msg.MsgID)
							d.Result = tppmessage.RESULT_ERR
							goto marshalAndReturn
						}
					}

					for _, placed := range cp.Common1Security.VoluntaryCoordCameraParams {
						pl := fobplaced.Placed{
							ClusterParamID: clusterParamID,
							Type:           fobplaced.CAMERA,
							PlacedIndex:    placed.PlacedIndex,
							PositionX:      placed.PositionX,
							PositionY:      placed.PositionY,
							PositionZ:      placed.PositionZ,
							RotationW:      placed.RotationW,
							RotationX:      placed.RotationX,
							RotationY:      placed.RotationY,
							RotationZ:      placed.RotationZ,
							SecurityIDX:    0,
						}
						if err = manager.FOBPlacedRepo.AddOrUpdate(ctx, &pl); err != nil {
							slog.Error("fob placed update", "error", err.Error(), "msgid", msg.MsgID)
							d.Result = tppmessage.RESULT_ERR
							goto marshalAndReturn
						}
					}

					for _, placed := range cp.Common1Security.VoluntaryCoordMineParams {
						pl := fobplaced.Placed{
							ClusterParamID: clusterParamID,
							Type:           fobplaced.MINE,
							PlacedIndex:    placed.PlacedIndex,
							PositionX:      placed.PositionX,
							PositionY:      placed.PositionY,
							PositionZ:      placed.PositionZ,
							RotationW:      placed.RotationW,
							RotationX:      placed.RotationX,
							RotationY:      placed.RotationY,
							RotationZ:      placed.RotationZ,
							SecurityIDX:    0,
						}
						if err = manager.FOBPlacedRepo.AddOrUpdate(ctx, &pl); err != nil {
							slog.Error("fob placed update", "error", err.Error(), "msgid", msg.MsgID)
							d.Result = tppmessage.RESULT_ERR
							goto marshalAndReturn
						}
					}

					// ====================

					sec, err = manager.ClusterSecurityRepo.Get(ctx, clusterParamID, 1)
					if err != nil {
						slog.Error("cluster security get", "error", err.Error(), "msgid", msg.MsgID)
						d.Result = tppmessage.RESULT_ERR
						goto marshalAndReturn
					}
					found = true
					if len(sec) < 1 {
						cs1 := clustersecurity.ClusterSecurity{
							ClusterParamID: clusterParamID,
							IDX:            1,
							IsUnique:       0,
						}
						sec = append(sec, cs1)
						found = false
					}

					sec[0].Antitheft = cp.Common2Security.Antitheft
					sec[0].Camera = cp.Common2Security.Camera
					sec[0].CautionArea = cp.Common2Security.CautionArea
					sec[0].Decoy = cp.Common2Security.Decoy
					sec[0].IrSensor = cp.Common2Security.IrSensor
					sec[0].Mine = cp.Common2Security.Mine
					sec[0].Soldier = cp.Common2Security.Soldier
					sec[0].Uav = cp.Common2Security.Uav
					sec[0].VoluntaryCoordCameraCount = cp.Common2Security.VoluntaryCoordCameraCount
					sec[0].VoluntaryCoordMineCount = cp.Common2Security.VoluntaryCoordMineCount
					if !found {
						if err = manager.ClusterSecurityRepo.Add(ctx, &sec[0]); err != nil {
							slog.Error("cluster security add", "error", err.Error(), "msgid", msg.MsgID)
							d.Result = tppmessage.RESULT_ERR
							goto marshalAndReturn
						}
					} else {
						if err = manager.ClusterSecurityRepo.Update(ctx, &sec[0]); err != nil {
							slog.Error("cluster security update", "error", err.Error(), "msgid", msg.MsgID)
							d.Result = tppmessage.RESULT_ERR
							goto marshalAndReturn
						}
					}

					for _, placed := range cp.Common2Security.VoluntaryCoordCameraParams {
						pl := fobplaced.Placed{
							ClusterParamID: clusterParamID,
							Type:           fobplaced.CAMERA,
							PlacedIndex:    placed.PlacedIndex,
							PositionX:      placed.PositionX,
							PositionY:      placed.PositionY,
							PositionZ:      placed.PositionZ,
							RotationW:      placed.RotationW,
							RotationX:      placed.RotationX,
							RotationY:      placed.RotationY,
							RotationZ:      placed.RotationZ,
							SecurityIDX:    1,
						}
						if err = manager.FOBPlacedRepo.AddOrUpdate(ctx, &pl); err != nil {
							slog.Error("fob placed update", "error", err.Error(), "msgid", msg.MsgID)
							d.Result = tppmessage.RESULT_ERR
							goto marshalAndReturn
						}
					}

					for _, placed := range cp.Common2Security.VoluntaryCoordMineParams {
						pl := fobplaced.Placed{
							ClusterParamID: clusterParamID,
							Type:           fobplaced.MINE,
							PlacedIndex:    placed.PlacedIndex,
							PositionX:      placed.PositionX,
							PositionY:      placed.PositionY,
							PositionZ:      placed.PositionZ,
							RotationW:      placed.RotationW,
							RotationX:      placed.RotationX,
							RotationY:      placed.RotationY,
							RotationZ:      placed.RotationZ,
							SecurityIDX:    1,
						}
						if err = manager.FOBPlacedRepo.AddOrUpdate(ctx, &pl); err != nil {
							slog.Error("fob placed update", "error", err.Error(), "msgid", msg.MsgID)
							d.Result = tppmessage.RESULT_ERR
							goto marshalAndReturn
						}
					}

					// ====================

					sec, err = manager.ClusterSecurityRepo.Get(ctx, clusterParamID, 2)
					if err != nil {
						slog.Error("cluster security get", "error", err.Error(), "msgid", msg.MsgID)
						d.Result = tppmessage.RESULT_ERR
						goto marshalAndReturn
					}
					found = true
					if len(sec) < 1 {
						cs2 := clustersecurity.ClusterSecurity{
							ClusterParamID: clusterParamID,
							IDX:            2,
							IsUnique:       0,
						}
						sec = append(sec, cs2)
						found = false
					}

					sec[0].Antitheft = cp.Common3Security.Antitheft
					sec[0].Camera = cp.Common3Security.Camera
					sec[0].CautionArea = cp.Common3Security.CautionArea
					sec[0].Decoy = cp.Common3Security.Decoy
					sec[0].IrSensor = cp.Common3Security.IrSensor
					sec[0].Mine = cp.Common3Security.Mine
					sec[0].Soldier = cp.Common3Security.Soldier
					sec[0].Uav = cp.Common3Security.Uav
					sec[0].VoluntaryCoordCameraCount = cp.Common3Security.VoluntaryCoordCameraCount
					sec[0].VoluntaryCoordMineCount = cp.Common3Security.VoluntaryCoordMineCount
					if !found {
						if err = manager.ClusterSecurityRepo.Add(ctx, &sec[0]); err != nil {
							slog.Error("cluster security add", "error", err.Error(), "msgid", msg.MsgID)
							d.Result = tppmessage.RESULT_ERR
							goto marshalAndReturn
						}
					} else {
						if err = manager.ClusterSecurityRepo.Update(ctx, &sec[0]); err != nil {
							slog.Error("cluster security update", "error", err.Error(), "msgid", msg.MsgID)
							d.Result = tppmessage.RESULT_ERR
							goto marshalAndReturn
						}
					}

					for _, placed := range cp.Common3Security.VoluntaryCoordCameraParams {
						pl := fobplaced.Placed{
							ClusterParamID: clusterParamID,
							Type:           fobplaced.CAMERA,
							PlacedIndex:    placed.PlacedIndex,
							PositionX:      placed.PositionX,
							PositionY:      placed.PositionY,
							PositionZ:      placed.PositionZ,
							RotationW:      placed.RotationW,
							RotationX:      placed.RotationX,
							RotationY:      placed.RotationY,
							RotationZ:      placed.RotationZ,
							SecurityIDX:    2,
						}
						if err = manager.FOBPlacedRepo.AddOrUpdate(ctx, &pl); err != nil {
							slog.Error("fob placed update", "error", err.Error(), "msgid", msg.MsgID)
							d.Result = tppmessage.RESULT_ERR
							goto marshalAndReturn
						}
					}

					for _, placed := range cp.Common3Security.VoluntaryCoordMineParams {
						pl := fobplaced.Placed{
							ClusterParamID: clusterParamID,
							Type:           fobplaced.MINE,
							PlacedIndex:    placed.PlacedIndex,
							PositionX:      placed.PositionX,
							PositionY:      placed.PositionY,
							PositionZ:      placed.PositionZ,
							RotationW:      placed.RotationW,
							RotationX:      placed.RotationX,
							RotationY:      placed.RotationY,
							RotationZ:      placed.RotationZ,
							SecurityIDX:    2,
						}
						if err = manager.FOBPlacedRepo.AddOrUpdate(ctx, &pl); err != nil {
							slog.Error("fob placed update", "error", err.Error(), "msgid", msg.MsgID)
							d.Result = tppmessage.RESULT_ERR
							goto marshalAndReturn
						}
					}

					// ===============

					sec, err = manager.ClusterSecurityRepo.Get(ctx, clusterParamID, 3)
					if err != nil {
						slog.Error("cluster security get", "error", err.Error(), "msgid", msg.MsgID)
						d.Result = tppmessage.RESULT_ERR
						goto marshalAndReturn
					}
					found = true
					if len(sec) < 1 {
						cs3 := clustersecurity.ClusterSecurity{
							ClusterParamID: clusterParamID,
							IDX:            3,
							IsUnique:       1,
						}
						sec = append(sec, cs3)
						found = false
					}

					sec[0].Antitheft = cp.UniqueSecurity.Antitheft
					sec[0].Camera = cp.UniqueSecurity.Camera
					sec[0].CautionArea = cp.UniqueSecurity.CautionArea
					sec[0].Decoy = cp.UniqueSecurity.Decoy
					sec[0].IrSensor = cp.UniqueSecurity.IrSensor
					sec[0].Mine = cp.UniqueSecurity.Mine
					sec[0].Soldier = cp.UniqueSecurity.Soldier
					sec[0].Uav = cp.UniqueSecurity.Uav
					sec[0].VoluntaryCoordCameraCount = cp.UniqueSecurity.VoluntaryCoordCameraCount
					sec[0].VoluntaryCoordMineCount = cp.UniqueSecurity.VoluntaryCoordMineCount
					if !found {
						if err = manager.ClusterSecurityRepo.Add(ctx, &sec[0]); err != nil {
							slog.Error("cluster security add", "error", err.Error(), "msgid", msg.MsgID)
							d.Result = tppmessage.RESULT_ERR
							goto marshalAndReturn
						}
					} else {
						if err = manager.ClusterSecurityRepo.Update(ctx, &sec[0]); err != nil {
							slog.Error("cluster security update", "error", err.Error(), "msgid", msg.MsgID)
							d.Result = tppmessage.RESULT_ERR
							goto marshalAndReturn
						}
					}

					for _, placed := range cp.UniqueSecurity.VoluntaryCoordCameraParams {
						pl := fobplaced.Placed{
							ClusterParamID: clusterParamID,
							Type:           fobplaced.CAMERA,
							PlacedIndex:    placed.PlacedIndex,
							PositionX:      placed.PositionX,
							PositionY:      placed.PositionY,
							PositionZ:      placed.PositionZ,
							RotationW:      placed.RotationW,
							RotationX:      placed.RotationX,
							RotationY:      placed.RotationY,
							RotationZ:      placed.RotationZ,
							SecurityIDX:    3,
						}
						if err = manager.FOBPlacedRepo.AddOrUpdate(ctx, &pl); err != nil {
							slog.Error("fob placed update", "error", err.Error(), "msgid", msg.MsgID)
							d.Result = tppmessage.RESULT_ERR
							goto marshalAndReturn
						}
					}

					for _, placed := range cp.UniqueSecurity.VoluntaryCoordMineParams {
						pl := fobplaced.Placed{
							ClusterParamID: clusterParamID,
							Type:           fobplaced.MINE,
							PlacedIndex:    placed.PlacedIndex,
							PositionX:      placed.PositionX,
							PositionY:      placed.PositionY,
							PositionZ:      placed.PositionZ,
							RotationW:      placed.RotationW,
							RotationX:      placed.RotationX,
							RotationY:      placed.RotationY,
							RotationZ:      placed.RotationZ,
							SecurityIDX:    3,
						}
						if err = manager.FOBPlacedRepo.AddOrUpdate(ctx, &pl); err != nil {
							slog.Error("fob placed update", "error", err.Error(), "msgid", msg.MsgID)
							d.Result = tppmessage.RESULT_ERR
							goto marshalAndReturn
						}
					}
				}
			}
		}
	}

marshalAndReturn:
	msg.MData, err = json.Marshal(d)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
