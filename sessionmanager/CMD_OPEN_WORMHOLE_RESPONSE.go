package sessionmanager

import (
	"encoding/json"
	"fmt"
	"fuse/message"
	"fuse/tppmessage"
	"log/slog"
)

func GetCmdOpenWormholeResponse() tppmessage.CmdOpenWormholeResponse {
	t := tppmessage.CmdOpenWormholeResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_OPEN_WORMHOLE.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	// TODO from database

	return t
}

func HandleCmdOpenWormholeResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	t := GetCmdOpenWormholeResponse()

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
