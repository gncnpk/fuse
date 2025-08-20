package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdGetAbolitionCountResponse(ctx context.Context, manager *SessionManager, msg *message.Message) tppmessage.CmdGetAbolitionCountResponse {
	t := tppmessage.CmdGetAbolitionCountResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_ABOLITION_COUNT.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	count, err := manager.AbolitionRepo.GetCount(ctx, msg.PlayerID)
	if err != nil {
		slog.Error("abolition count", "error", err.Error(), "playerID", msg.PlayerID)
		t.Result = tppmessage.RESULT_ERR
		return t
	}

	date, err := manager.AbolitionRepo.GetLatest(ctx, msg.PlayerID)
	if err != nil {
		slog.Error("abolition date", "error", err.Error(), "playerID", msg.PlayerID)
		t.Result = tppmessage.RESULT_ERR
		return t
	}

	t.Info.Count = count
	t.Info.Date = date

	// TODO get from database
	t.Info.Max = 2147483647
	t.Info.Num = 24958
	t.Info.Status = 0

	return t
}

func HandleCmdGetAbolitionCountResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	//t := GetCmdGetAbolitionCountResponse()
	t := tppmessage.CmdGetAbolitionCountResponse{}

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
