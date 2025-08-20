package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdGetSecurityInfoResponse(ctx context.Context, msg *message.Message, manager *SessionManager) tppmessage.CmdGetSecurityInfoResponse {
	t := tppmessage.CmdGetSecurityInfoResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_SECURITY_INFO.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	// todo from database? no good example, might be insurance?

	return t
}

func HandleCmdGetSecurityInfoResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	t := tppmessage.CmdGetSecurityInfoResponse{}

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
