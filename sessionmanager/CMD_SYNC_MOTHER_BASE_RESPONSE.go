package sessionmanager

import (
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdSyncMotherBaseResponse() tppmessage.CmdSyncMotherBaseResponse {
	t := tppmessage.CmdSyncMotherBaseResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_SYNC_MOTHER_BASE.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	t.Version = 0 // always 0

	return t
}

func HandleCmdSyncMotherBaseResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	t := GetCmdSyncMotherBaseResponse()

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
