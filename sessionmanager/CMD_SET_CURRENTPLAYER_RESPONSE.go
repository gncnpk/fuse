package sessionmanager

import (
	"context"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/platform"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdSetCurrentplayerResponse(ctx context.Context, manager *SessionManager, sessionKey string, plat platform.Platform, platformID uint64, index int, isReset bool) tppmessage.CmdSetCurrentplayerResponse {
	t := tppmessage.CmdSetCurrentplayerResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_SET_CURRENTPLAYER.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	player, err := manager.PlayerRepo.Get(ctx, plat, platformID, index)
	if err != nil {
		slog.Error("player not found", "msgID", t.Msgid, "platform", plat, "platformID", platformID, "index", index, "error", err.Error())
		t.Result = tppmessage.RESULT_ERR
		return t
	}

	t.PlayerID = player.ID
	if err = manager.SetPlayerID(ctx, sessionKey, player.ID); err != nil {
		slog.Error("cannot set player id", "error", err.Error())
		t.Result = tppmessage.RESULT_ERR
		return t
	}

	return t
}

func HandleCmdSetCurrentplayerResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	//var err error
	//t := GetCmdSetCurrentplayerResponse()
	//
	//message.MData, err = json.Marshal(t)
	//if err != nil {
	//	return fmt.Errorf("cannot marshal: %w", err)
	//}

	return nil
}
