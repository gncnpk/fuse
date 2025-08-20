package sessionmanager

import (
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdGetPfPointExchangeParamsResponse() tppmessage.CmdGetPfPointExchangeParamsResponse {
	t := tppmessage.CmdGetPfPointExchangeParamsResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_PF_POINT_EXCHANGE_PARAMS.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	// TODO from database

	return t
}

func HandleCmdGetPfPointExchangeParamsResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	t := GetCmdGetPfPointExchangeParamsResponse()

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
