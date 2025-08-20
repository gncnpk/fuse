
package sessionmanager

import (
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)
func GetCmdActiveSneakMotherBaseResponse() tppmessage.CmdActiveSneakMotherBaseResponse {
	t := tppmessage.CmdActiveSneakMotherBaseResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_ACTIVE_SNEAK_MOTHER_BASE.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	// code

	return t
}

func HandleCmdActiveSneakMotherBaseResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	t := GetCmdActiveSneakMotherBaseResponse()
	
	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
