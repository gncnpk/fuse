package sessionmanager

import (
	"encoding/json"
	"fmt"
	"fuse/message"
	"fuse/tppmessage"
	"log/slog"
)

func GetCmdGetCombatDeployListResponse() tppmessage.CmdGetCombatDeployListResponse {
	t := tppmessage.CmdGetCombatDeployListResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_COMBAT_DEPLOY_LIST.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	// TODO from database

	return t
}

func HandleCmdGetCombatDeployListResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	t := GetCmdGetCombatDeployListResponse()

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
