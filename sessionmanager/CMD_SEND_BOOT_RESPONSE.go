package sessionmanager

import (
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdSendBootResponse() tppmessage.CmdSendBootResponse {
	t := tppmessage.CmdSendBootResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_SEND_BOOT.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	t.Flag = 0

	// TODO save req to database

	return t
}

func HandleCmdSendBootResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	t := GetCmdSendBootResponse()

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
