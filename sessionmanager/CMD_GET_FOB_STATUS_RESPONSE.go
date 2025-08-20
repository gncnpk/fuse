package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/fobrecord"
	"github.com/unknown321/fuse/fobstatus"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

// TODO check what params mean

func GetCmdGetFobStatusResponse(ctx context.Context, msg *message.Message, manager *SessionManager) tppmessage.CmdGetFobStatusResponse {
	t := tppmessage.CmdGetFobStatusResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_FOB_STATUS.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	t.FobIndex = -1  // always -1
	t.IsRescue = 0   // emergency
	t.IsReward = 0   // always 0?
	t.KillCount = 0  // always 0?
	t.SneakMode = -1 // always -1?

	// record
	{
		record, err := manager.FobRecordRepo.GetByPlayerID(ctx, msg.PlayerID)
		if err != nil {
			slog.Error("fob record get", "error", err.Error(), "msgid", msg.MsgID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		var rr fobrecord.FobRecord
		if len(record) == 0 {
			rr.PlayerID = msg.PlayerID
			if err = manager.FobRecordRepo.Add(ctx, &rr); err != nil {
				slog.Error("fob record add", "error", err.Error(), "msgid", msg.MsgID)
				t.Result = tppmessage.RESULT_ERR
				return t
			}
		} else {
			rr = record[0]
		}

		t.Record.Defense.Win = rr.DefenseWin
		t.Record.Defense.Lose = rr.DefenseWin
		t.Record.Insurance = rr.Insurance
		t.Record.Score = rr.Score
		t.Record.ShieldDate = rr.ShieldDate
		t.Record.Sneak.Win = rr.SneakWin
		t.Record.Sneak.Lose = rr.SneakLose
	}

	// status
	{
		status, err := manager.FobStatusRepo.GetByPlayerID(ctx, msg.PlayerID)
		if err != nil {
			slog.Error("fob status get", "error", err.Error(), "msgid", msg.MsgID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		var ss fobstatus.FobStatus
		if len(status) == 0 {
			ss.PlayerID = msg.PlayerID
			ss.SneakMode = -1
			if err = manager.FobStatusRepo.Add(ctx, &ss); err != nil {
				slog.Error("fob status add", "error", err.Error(), "msgid", msg.MsgID)
				t.Result = tppmessage.RESULT_ERR
				return t
			}
		} else {
			ss = status[0]
		}

		// TODO check if intruder present and set IsRescue based on that

		t.IsRescue = ss.IsRescue
		t.IsReward = ss.IsReward
		t.SneakMode = ss.SneakMode
	}

	return t
}

func HandleCmdGetFobStatusResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	//t := GetCmdGetFobStatusResponse()
	t := tppmessage.CmdGetFobStatusResponse{}

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
