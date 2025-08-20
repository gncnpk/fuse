package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/clusterparam"
	"github.com/unknown321/fuse/fobplaced"
	"github.com/unknown321/fuse/fobweaponplacement"
	"github.com/unknown321/fuse/intruder"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/playerresource"
	"github.com/unknown321/fuse/sectionstat"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdSneakMotherBaseResponse(ctx context.Context, msg *message.Message, manager *SessionManager, request *tppmessage.CmdSneakMotherBaseRequest) tppmessage.CmdSneakMotherBaseResponse {
	t := tppmessage.CmdSneakMotherBaseResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_SNEAK_MOTHER_BASE.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	t.DamageParam = []any{}           // always []
	t.EventFobParams = make([]int, 5) // always [0,0,0,0,0]
	t.FobDeployDamageParam = tppmessage.FobDeployDamageParam{
		ClusterIndex:   0,              // always 0
		DamageValues:   make([]int, 6), // always 0
		ExpirationDate: 0,              // always 0
		MotherbaseID:   0,              // always 0
	}

	t.IsEvent = 0            // always 0, even on events
	t.IsSecurityContract = 0 // TODO always 0?

	mbParams, err := manager.MotherBaseParamRepo.GetByMotherBaseID(ctx, request.MotherBaseID)
	if err != nil {
		slog.Error("mother base param", "error", err.Error(), "msgID", t.Msgid, "mbID", request.MotherBaseID)
		t.Result = tppmessage.RESULT_ERR
		return t
	}

	if len(mbParams) < 1 {
		slog.Error("mother base param", "error", "motherbase not found", "msgID", t.Msgid, "mbID", request.MotherBaseID)
		t.Result = tppmessage.RESULT_ERR
		return t
	}

	player, err := manager.PlayerRepo.GetByID(ctx, msg.Platform, mbParams[0].PlayerID)
	if err != nil {
		slog.Error("get owner", "error", err.Error(), "msgID", t.Msgid, "mbID", request.MotherBaseID, "playerID", mbParams[0].PlayerID, "platform", msg.Platform)
		t.Result = tppmessage.RESULT_ERR
		return t
	}

	playerStatus, err := manager.PlayerStatusRepo.Get(ctx, player.ID)
	if err != nil {
		slog.Error("get owner status", "error", err.Error(), "msgID", t.Msgid, "playerID", player.ID)
		t.Result = tppmessage.RESULT_ERR
		return t
	}

	t.OwnerGmp = playerStatus.Gmp // TODO sometimes it's -5000000, sometimes 0, sometimes positive, why?

	// TODO resources that would be taken in retaliation? there was an instance of values being > 0
	t.RecoverResource = tppmessage.CmdMiningResourceEntry{
		BioticResource: 0,
		CommonMetal:    0,
		FuelResource:   0,
		MinorMetal:     0,
		PreciousMetal:  0,
	}

	t.RecoverSoldier = []any{} // TODO always empty?
	t.RecoverSoldierNum = len(t.RecoverSoldier)
	t.RecoverSoldierCount = make([]int, 10) // TODO always 10 zeroes?

	t.RewardID = 8 // TODO  spotted ids: 8,13,14,16,17,24,1401,1405,1407,1411

	// TODO generate random solly?
	// this is rank E soldier
	t.RewardSoldier = []tppmessage.FobSoldier{
		{
			Header:       2311064631, // 0x89c00c37
			Seed:         1925251072, // 0x72c10000
			StatusNoSync: 4192,       // 0x1060
			StatusSync:   822083584,  // 0x31000000
		},
	}
	t.RewardSoldierNum = len(t.RewardSoldier) // TODO soldiers are given for specific rewards, spotted ids: 13, 14, 16, 17, 8. Where do reward ids come from?
	t.RewardSoldierRank = 0                   // always 0?
	t.RewardSoldierType = 0                   // always 0?

	// TODO solly, taken from security platform only
	t.SecuritySoldier = []tppmessage.FobSoldier{
		{
			Header:       2311064631,
			Seed:         1925251072,
			StatusNoSync: 4192,
			StatusSync:   822083584,
		},
	}
	t.SecuritySoldierNum = len(t.SecuritySoldier) // up to 500
	t.SecuritySoldierRank = 0                     // always 0

	t.WormholePlayerID = 0 // TODO set to attacker if there is an attacker?

	t.StageParam = tppmessage.FobSneakMotherBaseStageParam{
		ClusterParam: tppmessage.ClusterParam{
			Build:       0, // always 0
			SoldierRank: 0, // always 0
		},
	}

	// equip grade
	{
		eqg, err := manager.EquipGradeRepo.Get(ctx, mbParams[0].PlayerID)
		if err != nil {
			slog.Error("get equipgrade", "error", err.Error(), "msgID", t.Msgid, "mbID", request.MotherBaseID, "playerID", mbParams[0].PlayerID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		t.StageParam.EquipGrade = eqg.ToArray()
	}

	// security level
	{
		secL, err := manager.SecurityLevelRepo.Get(ctx, mbParams[0].PlayerID)
		if err != nil {
			slog.Error("get security level", "error", err.Error(), "msgID", t.Msgid, "mbID", request.MotherBaseID, "playerID", mbParams[0].PlayerID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		t.StageParam.SecurityLevel = secL.ToArray()
	}

	// section level, directly affects time remaining on fob
	{
		sl, err := manager.SectionStatRepo.GetByPlayerID(ctx, mbParams[0].PlayerID)
		if err != nil {
			slog.Error("get section stat", "error", err.Error(), "msgID", t.Msgid, "mbID", request.MotherBaseID, "playerID", mbParams[0].PlayerID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}
		for _, v := range sl {
			switch v.SectionID {
			case sectionstat.Base:
				t.StageParam.SectionLevel.Base = v.Level
			case sectionstat.Combat:
				t.StageParam.SectionLevel.Combat = v.Level
			case sectionstat.Develop:
				t.StageParam.SectionLevel.Develop = v.Level
			case sectionstat.Medical:
				t.StageParam.SectionLevel.Medical = v.Level
			case sectionstat.Security:
				t.StageParam.SectionLevel.Security = v.Level
			case sectionstat.Spy:
				t.StageParam.SectionLevel.Spy = v.Level
			case sectionstat.Support:
				t.StageParam.SectionLevel.Suport = v.Level
			default:
				slog.Error("invalid section stat id", "value", v.SectionID, "msgID", t.Msgid, "mbID", request.MotherBaseID, "playerID", mbParams[0].PlayerID)
				t.Result = tppmessage.RESULT_ERR
				return t
			}
		}
	}

	t.StageParam.ConstructParam = mbParams[0].ConstructParam
	t.StageParam.FobIndex = mbParams[0].FobIndex
	t.StageParam.MotherBaseID = mbParams[0].ID
	t.StageParam.OwnerPlayerID = mbParams[0].PlayerID
	t.StageParam.Platform = request.Platform // TODO might be different value in special cases? Don't know for sure

	var platform clusterparam.ClusterParam
	// cluster param
	{
		clusterParams, err := manager.ClusterParamRepo.Get(ctx, request.MotherBaseID)
		if err != nil {
			slog.Error("get cluster params", "error", err.Error(), "msgID", t.Msgid, "mbID", request.MotherBaseID, "platform", request.Platform)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		if len(clusterParams) < 1 {
			slog.Error("get cluster params", "error", "not found", "msgID", t.Msgid, "mbID", request.MotherBaseID, "platform", request.Platform)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		for _, v := range clusterParams {
			t.StageParam.Build = append(t.StageParam.Build, v.Build.ToInt())
			if request.Platform == v.PlatformID {
				platform = v
			}
		}
		t.StageParam.ClusterParam.ClusterSecurity = platform.ClusterSecurityParam.ToInt()
	}

	// cluster security 1
	{
		cs, err := manager.ClusterSecurityRepo.Get(ctx, platform.ID, 0)
		if err != nil {
			slog.Error("get cluster security1", "error", err.Error(), "msgID", t.Msgid, "mbID", request.MotherBaseID, "platform", platform.ID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}
		if len(cs) < 1 {
			slog.Error("get cluster security1", "error", "not found", "msgID", t.Msgid, "mbID", request.MotherBaseID, "platform", platform.ID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		t.StageParam.ClusterParam.Common1Security.Antitheft = cs[0].Antitheft
		t.StageParam.ClusterParam.Common1Security.Camera = cs[0].Camera
		t.StageParam.ClusterParam.Common1Security.CautionArea = cs[0].CautionArea
		t.StageParam.ClusterParam.Common1Security.Decoy = cs[0].Decoy
		t.StageParam.ClusterParam.Common1Security.IrSensor = cs[0].IrSensor
		t.StageParam.ClusterParam.Common1Security.Mine = cs[0].Mine
		t.StageParam.ClusterParam.Common1Security.Soldier = cs[0].Soldier
		t.StageParam.ClusterParam.Common1Security.Uav = cs[0].Uav

		placed, err := manager.FOBPlacedRepo.Get(ctx, platform.ID, 0)
		if err != nil {
			slog.Error("get cluster placed", "error", err.Error(), "cluster param id", platform.ID)
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
				t.StageParam.ClusterParam.Common1Security.VoluntaryCoordMineCount++
				t.StageParam.ClusterParam.Common1Security.VoluntaryCoordMineParams = append(t.StageParam.ClusterParam.Common1Security.VoluntaryCoordMineParams, vc)
			case fobplaced.CAMERA:
				t.StageParam.ClusterParam.Common1Security.VoluntaryCoordCameraCount++
				t.StageParam.ClusterParam.Common1Security.VoluntaryCoordCameraParams = append(t.StageParam.ClusterParam.Common1Security.VoluntaryCoordCameraParams, vc)
			default:
				slog.Error("placed invalid type", "cluster param id", platform.ID)
				t.Result = tppmessage.RESULT_ERR
				return t
			}
		}
	}

	// cluster security 2
	{
		cs, err := manager.ClusterSecurityRepo.Get(ctx, platform.ID, 1)
		if err != nil {
			slog.Error("get cluster security2", "error", err.Error(), "msgID", t.Msgid, "mbID", request.MotherBaseID, "platform", platform.ID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}
		if len(cs) < 1 {
			slog.Error("get cluster security2", "error", "not found", "msgID", t.Msgid, "mbID", request.MotherBaseID, "platform", platform.ID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		t.StageParam.ClusterParam.Common2Security.Antitheft = cs[0].Antitheft
		t.StageParam.ClusterParam.Common2Security.Camera = cs[0].Camera
		t.StageParam.ClusterParam.Common2Security.CautionArea = cs[0].CautionArea
		t.StageParam.ClusterParam.Common2Security.Decoy = cs[0].Decoy
		t.StageParam.ClusterParam.Common2Security.IrSensor = cs[0].IrSensor
		t.StageParam.ClusterParam.Common2Security.Mine = cs[0].Mine
		t.StageParam.ClusterParam.Common2Security.Soldier = cs[0].Soldier
		t.StageParam.ClusterParam.Common2Security.Uav = cs[0].Uav

		placed, err := manager.FOBPlacedRepo.Get(ctx, platform.ID, 1)
		if err != nil {
			slog.Error("get cluster placed", "error", err.Error(), "cluster param id", platform.ID)
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
				t.StageParam.ClusterParam.Common2Security.VoluntaryCoordMineCount++
				t.StageParam.ClusterParam.Common2Security.VoluntaryCoordMineParams = append(t.StageParam.ClusterParam.Common2Security.VoluntaryCoordMineParams, vc)
			case fobplaced.CAMERA:
				t.StageParam.ClusterParam.Common2Security.VoluntaryCoordCameraCount++
				t.StageParam.ClusterParam.Common2Security.VoluntaryCoordCameraParams = append(t.StageParam.ClusterParam.Common2Security.VoluntaryCoordCameraParams, vc)
			default:
				slog.Error("placed invalid type", "cluster param id", platform.ID)
				t.Result = tppmessage.RESULT_ERR
				return t
			}
		}
	}

	// cluster security 3
	{
		cs, err := manager.ClusterSecurityRepo.Get(ctx, platform.ID, 2)
		if err != nil {
			slog.Error("get cluster security3", "error", err.Error(), "msgID", t.Msgid, "mbID", request.MotherBaseID, "platform", platform.ID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}
		if len(cs) < 1 {
			slog.Error("get cluster security3", "error", "not found", "msgID", t.Msgid, "mbID", request.MotherBaseID, "platform", platform.ID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		t.StageParam.ClusterParam.Common3Security.Antitheft = cs[0].Antitheft
		t.StageParam.ClusterParam.Common3Security.Camera = cs[0].Camera
		t.StageParam.ClusterParam.Common3Security.CautionArea = cs[0].CautionArea
		t.StageParam.ClusterParam.Common3Security.Decoy = cs[0].Decoy
		t.StageParam.ClusterParam.Common3Security.IrSensor = cs[0].IrSensor
		t.StageParam.ClusterParam.Common3Security.Mine = cs[0].Mine
		t.StageParam.ClusterParam.Common3Security.Soldier = cs[0].Soldier
		t.StageParam.ClusterParam.Common3Security.Uav = cs[0].Uav

		placed, err := manager.FOBPlacedRepo.Get(ctx, platform.ID, 2)
		if err != nil {
			slog.Error("get cluster placed", "error", err.Error(), "cluster param id", platform.ID)
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
				t.StageParam.ClusterParam.Common3Security.VoluntaryCoordMineCount++
				t.StageParam.ClusterParam.Common3Security.VoluntaryCoordMineParams = append(t.StageParam.ClusterParam.Common3Security.VoluntaryCoordMineParams, vc)
			case fobplaced.CAMERA:
				t.StageParam.ClusterParam.Common3Security.VoluntaryCoordCameraCount++
				t.StageParam.ClusterParam.Common3Security.VoluntaryCoordCameraParams = append(t.StageParam.ClusterParam.Common3Security.VoluntaryCoordCameraParams, vc)
			default:
				slog.Error("placed invalid type", "cluster param id", platform.ID)
				t.Result = tppmessage.RESULT_ERR
				return t
			}
		}
	}

	// cluster security unique
	{
		cs, err := manager.ClusterSecurityRepo.Get(ctx, platform.ID, 3)
		if err != nil {
			slog.Error("get cluster security3", "error", err.Error(), "msgID", t.Msgid, "mbID", request.MotherBaseID, "platform", platform.ID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}
		if len(cs) < 1 {
			slog.Error("get cluster security3", "error", "not found", "msgID", t.Msgid, "mbID", request.MotherBaseID, "platform", platform.ID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		t.StageParam.ClusterParam.UniqueSecurity.Antitheft = cs[0].Antitheft
		t.StageParam.ClusterParam.UniqueSecurity.Camera = cs[0].Camera
		t.StageParam.ClusterParam.UniqueSecurity.CautionArea = cs[0].CautionArea
		t.StageParam.ClusterParam.UniqueSecurity.Decoy = cs[0].Decoy
		t.StageParam.ClusterParam.UniqueSecurity.IrSensor = cs[0].IrSensor
		t.StageParam.ClusterParam.UniqueSecurity.Mine = cs[0].Mine
		t.StageParam.ClusterParam.UniqueSecurity.Soldier = cs[0].Soldier
		t.StageParam.ClusterParam.UniqueSecurity.Uav = cs[0].Uav

		placed, err := manager.FOBPlacedRepo.Get(ctx, platform.ID, 3)
		if err != nil {
			slog.Error("get cluster placed", "error", err.Error(), "cluster param id", platform.ID)
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
				t.StageParam.ClusterParam.UniqueSecurity.VoluntaryCoordMineCount++
				t.StageParam.ClusterParam.UniqueSecurity.VoluntaryCoordMineParams = append(t.StageParam.ClusterParam.UniqueSecurity.VoluntaryCoordMineParams, vc)
			case fobplaced.CAMERA:
				t.StageParam.ClusterParam.UniqueSecurity.VoluntaryCoordCameraCount++
				t.StageParam.ClusterParam.UniqueSecurity.VoluntaryCoordCameraParams = append(t.StageParam.ClusterParam.UniqueSecurity.VoluntaryCoordCameraParams, vc)
			default:
				slog.Error("placed invalid type", "cluster param id", platform.ID)
				t.Result = tppmessage.RESULT_ERR
				return t
			}
		}
	}

	// weapon emplacement
	{
		placement, err := manager.FOBWeaponPlacementRepo.Get(ctx, request.MotherBaseID)
		if err != nil {
			slog.Error("get weapon placement", "error", err.Error(), "msgID", t.Msgid, "mbID", request.MotherBaseID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}
		if len(placement) < 1 {
			slog.Warn("weapon placement not found, using dummy", "msgID", t.Msgid, "mbID", request.MotherBaseID)
			placement = append(placement, fobweaponplacement.WeaponPlacement{})
		}

		t.StageParam.Placement.MortarNormal = placement[0].MortarNormal
		t.StageParam.Placement.GatlingGunWest = placement[0].GatlingGunWest
		t.StageParam.Placement.GatlingGunEast = placement[0].GatlingGunEast
		t.StageParam.Placement.EmplacementGunWest = placement[0].EmplacementGunWest
		t.StageParam.Placement.EmplacementGunEast = placement[0].EmplacementGunEast
		t.StageParam.Placement.GatlingGun = placement[0].GatlingGun
	}

	// resources
	{
		resources, err := manager.PlayerResourceRepo.Get(ctx, mbParams[0].PlayerID, true)
		if err != nil {
			slog.Error("get resources", "error", err.Error(), "msgID", t.Msgid, "mbID", request.MotherBaseID, "playerID", mbParams[0].PlayerID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}
		if len(resources) < 1 {
			slog.Error("weapon placement not found, using dummy", "msgID", t.Msgid, "mbID", request.MotherBaseID, "playerID", mbParams[0].PlayerID)
			resources = append(resources, playerresource.PlayerResource{})
		}

		resource := resources[0]
		t.StageParam.ProcessingResource.FuelResource = resource.Raw.Fuel
		t.StageParam.ProcessingResource.BioticResource = resource.Raw.Bio
		t.StageParam.ProcessingResource.MinorMetal = resource.Raw.MinorMetal
		t.StageParam.ProcessingResource.CommonMetal = resource.Raw.CommonMetal
		t.StageParam.ProcessingResource.PreciousMetal = resource.Raw.PreciousMetal
		t.StageParam.UsableResource.FuelResource = resource.Processed.Fuel
		t.StageParam.UsableResource.BioticResource = resource.Processed.Bio
		t.StageParam.UsableResource.MinorMetal = resource.Processed.MinorMetal
		t.StageParam.UsableResource.CommonMetal = resource.Processed.CommonMetal
		t.StageParam.UsableResource.PreciousMetal = resource.Processed.PreciousMetal
	}

	// TODO nuclear
	t.StageParam.Nuclear = 1

	if mbParams[0].PlayerID != msg.PlayerID {
		intr := intruder.Intruder{
			PlayerID:             msg.PlayerID,
			OwnerID:              mbParams[0].PlayerID,
			MotherBaseID:         request.MotherBaseID,
			MotherBasePlatformID: request.Platform,
			Mode:                 0,
			IsSneak:              1, // TODO from database?
		}

		if err = manager.IntruderRepo.AddOrUpdate(ctx, &intr); err != nil {
			slog.Error("add intruder", "error", err.Error(), "playerID", msg.PlayerID, "mbID", request.MotherBaseID, "platform", request.Platform)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		if err = manager.FobStatusRepo.SetRescue(ctx, mbParams[0].PlayerID); err != nil {
			slog.Error("set rescue", "error", err.Error(), "playerID", msg.PlayerID, "mbID", request.MotherBaseID, "platform", request.Platform)
			t.Result = tppmessage.RESULT_ERR
			return t
		}
	}

	return t
}

func HandleCmdSneakMotherBaseResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	//t := GetCmdSneakMotherBaseResponse()
	t := tppmessage.CmdActiveSneakMotherBaseResponse{}

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
