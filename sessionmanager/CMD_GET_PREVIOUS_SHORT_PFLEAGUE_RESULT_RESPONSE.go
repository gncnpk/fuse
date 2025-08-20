package sessionmanager

import (
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdGetPreviousShortPfleagueResultResponse() tppmessage.CmdGetPreviousShortPfleagueResultResponse {
	t := tppmessage.CmdGetPreviousShortPfleagueResultResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_PREVIOUS_SHORT_PFLEAGUE_RESULT.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	// TODO from database

	return t
}

func HandleCmdGetPreviousShortPfleagueResultResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	t := GetCmdGetPreviousShortPfleagueResultResponse()

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
