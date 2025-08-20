package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/playerresource"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdSyncResourceResponse(ctx context.Context, msg *message.Message, req *tppmessage.CmdSyncResourceRequest, manager *SessionManager) tppmessage.CmdSyncResourceResponse {
	t := tppmessage.CmdSyncResourceResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_SYNC_RESOURCE.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	// TODO: use request.CompensateResource

	// local resources, validate and save input
	{
		t.LocalGmp = req.Gmp
		t.DiffResource1 = req.DiffResource1
		t.DiffResource2 = req.DiffResource2

		d1 := playerresource.PlayerResource{PlayerID: msg.PlayerID}
		if err := d1.FromArray(req.DiffResource1, req.DiffResource2); err != nil {
			slog.Error("diff", "error", err.Error(), "msgid", t.Msgid)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		if err := manager.PlayerResourceRepo.AddOrUpdate(ctx, &d1); err != nil {
			slog.Error("diff save", "error", err.Error(), "msgid", t.Msgid)
			t.Result = tppmessage.RESULT_ERR
			return t
		}
	}

	// online resources
	{
		onlineResources, err := manager.PlayerResourceRepo.Get(ctx, msg.PlayerID, true)
		if err != nil {
			slog.Error("get online resources", "error", err.Error(), "msgid", t.Msgid, "playerID", msg.PlayerID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		if len(onlineResources) > 1 {
			slog.Error("too many online resources", "error", err.Error(), "msgid", t.Msgid, "playerID", msg.PlayerID, "count", len(onlineResources))
		}

		if len(onlineResources) == 0 {
			cc := playerresource.PlayerResource{
				PlayerID:    msg.PlayerID,
				IsOnline:    true,
				Raw:         manager.ManagerOpts.SignupBonus.Resources.Raw,
				Processed:   manager.ManagerOpts.SignupBonus.Resources.Processed,
				Plants:      manager.ManagerOpts.SignupBonus.Resources.Plants,
				Vehicles:    manager.ManagerOpts.SignupBonus.Resources.Vehicles,
				WalkerGears: manager.ManagerOpts.SignupBonus.Resources.WalkerGears,
				Nuclear:     manager.ManagerOpts.SignupBonus.Resources.Nuclear,
				Parasites:   manager.ManagerOpts.SignupBonus.Resources.Parasites,
				Mystery01:   0,
				Placed:      manager.ManagerOpts.SignupBonus.Resources.Placed,
				Mystery02:   [7]int{},
				Unused:      [13]int{},
			}
			if err = manager.PlayerResourceRepo.AddOrUpdate(ctx, &cc); err != nil {
				slog.Error("add new online resources", "error", err.Error(), "msgid", t.Msgid, "playerID", msg.PlayerID)
				t.Result = tppmessage.RESULT_ERR
				return t
			}
			onlineResources = append(onlineResources, cc)
		}

		fix1 := onlineResources[0].ToArray(false)
		fix2 := onlineResources[0].ToArray(true)
		t.FixResource1 = fix1
		t.FixResource2 = fix2
	}

	// player status
	{
		playerStatus, err := manager.PlayerStatusRepo.Get(ctx, msg.PlayerID)
		if err != nil {
			slog.Error("get player status", "error", err.Error(), "msgid", t.Msgid, "playerID", msg.PlayerID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		playerStatus.Gmp = req.Gmp
		playerStatus.Hero = req.Hero
		playerStatus.IsForceBalance = req.IsForceBalance
		playerStatus.IsWallet = req.IsWallet
		playerStatus.CumulativeGrade = req.CumulativeGrade

		if err = manager.PlayerStatusRepo.AddOrUpdate(ctx, &playerStatus); err != nil {
			slog.Error("add new player status", "error", err.Error(), "msgid", t.Msgid, "playerID", msg.PlayerID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		t.ServerGmp = playerStatus.ServerGmp
		t.InsuranceGmp = playerStatus.InsuranceGmp
		t.InjuryGmp = playerStatus.InjuryGmp
		t.LoadoutGmp = playerStatus.LoadoutGmp
	}

	// TODO track version?
	t.Version = req.Version + 1

	return t
}

func HandleCmdSyncResourceResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	//t := GetCmdSyncResourceResponse()
	t := tppmessage.CmdSyncResourceResponse{}

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
