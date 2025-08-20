package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdUpdateSessionResponse(ctx context.Context, msg *message.Message, manager *SessionManager) tppmessage.CmdUpdateSessionResponse {
	t := tppmessage.CmdUpdateSessionResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_UPDATE_SESSION.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	if err := manager.Update(ctx, *msg.SessionKey); err != nil {
		slog.Error("update session", "error", err.Error(), "msgid", t.Msgid, "key", *msg.SessionKey)
		t.Result = tppmessage.RESULT_ERR
		return t
	}

	// always -1?
	t.SneakMode = -1
	t.FobIndex = -1

	return t
}

func HandleCmdUpdateSessionResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	//t := GetCmdUpdateSessionResponse()
	t := tppmessage.CmdUpdateSessionResponse{}

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
