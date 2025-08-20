package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

// TODO not working, player is created on REQHTTPS

func GetCmdCreatePlayerResponse(ctx context.Context, msg *message.Message, manager *SessionManager) (tppmessage.CmdCreatePlayerResponse, error) {
	t := tppmessage.CmdCreatePlayerResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_CREATE_PLAYER.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	pid, err := manager.PlayerRepo.Add(ctx, msg.Platform, msg.PlatformID)
	if err != nil {
		return tppmessage.CmdCreatePlayerResponse{}, err
	}
	slog.Info("created player", "id", pid, "platformID", msg.PlatformID, "msgid", msg.MsgID)

	return t, nil
}

func HandleCmdCreatePlayerRequest(ctx context.Context, message *message.Message, manager *SessionManager) error {
	var err error
	t := tppmessage.CmdCreatePlayerRequest{}
	err = json.Unmarshal(message.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	slog.Info("CMD_CREATE_PLAYER", "player_name", t.PlayerName)

	d := t

	message.MData, err = json.Marshal(d)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
