package sessionmanager

import (
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdSendIpandportResponse() tppmessage.CmdSendIpandportResponse {
	t := tppmessage.CmdSendIpandportResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_SEND_IPANDPORT.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	return t
}

func HandleCmdSendIpandportResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	t := GetCmdSendIpandportResponse()

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
