package sessionmanager

import (
	"encoding/json"
	"fmt"
	"fuse/message"
	"fuse/tppmessage"
	"log/slog"
)

func GetCmdSyncLoadoutResponse() tppmessage.CmdSyncLoadoutResponse {
	t := tppmessage.CmdSyncLoadoutResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_SYNC_LOADOUT.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	return t
}

func HandleCmdSyncLoadoutResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	t := GetCmdSyncLoadoutResponse()

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
