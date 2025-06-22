package sessionmanager

import (
	"encoding/json"
	"fmt"
	"fuse/message"
	"fuse/tppmessage"
	"log/slog"
)

func GetCmdSendMissionResultResponse() tppmessage.CmdSendMissionResultResponse {
	t := tppmessage.CmdSendMissionResultResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_SEND_MISSION_RESULT.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	return t
}

func HandleCmdSendMissionResultResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	t := GetCmdSendMissionResultResponse()

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
