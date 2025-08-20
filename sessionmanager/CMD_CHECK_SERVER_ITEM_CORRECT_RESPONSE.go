package sessionmanager

import (
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdCheckServerItemCorrectResponse() tppmessage.CmdCheckServerItemCorrectResponse {
	t := tppmessage.CmdCheckServerItemCorrectResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_CHECK_SERVER_ITEM_CORRECT.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	// TODO check?
	t.CheckResult = 0

	return t
}

func HandleCmdCheckServerItemCorrectResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	t := GetCmdCheckServerItemCorrectResponse()

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
