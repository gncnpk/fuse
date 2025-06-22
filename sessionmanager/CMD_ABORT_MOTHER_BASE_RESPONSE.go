
package sessionmanager

import (
	"encoding/json"
	"fmt"
	"fuse/message"
	"fuse/tppmessage"
	"log/slog"
)
func GetCmdAbortMotherBaseResponse() tppmessage.CmdAbortMotherBaseResponse {
	t := tppmessage.CmdAbortMotherBaseResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_ABORT_MOTHER_BASE.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	// code

	return t
}

func HandleCmdAbortMotherBaseResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	t := GetCmdAbortMotherBaseResponse()
	
	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
