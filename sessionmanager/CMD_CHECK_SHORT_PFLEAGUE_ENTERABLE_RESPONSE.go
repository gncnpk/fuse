package sessionmanager

import (
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdCheckShortPfleagueEnterableResponse() tppmessage.CmdCheckShortPfleagueEnterableResponse {
	t := tppmessage.CmdCheckShortPfleagueEnterableResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_CHECK_SHORT_PFLEAGUE_ENTERABLE.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	t.PfleagueDate = 0      // timestamp
	t.ResultAlreadyRead = 1 // ?, 0
	t.Status = "ACCEPTING"  // ALREADY_ENTERED, ALREADY_HELD

	if t.Status == "ALREADY_HELD" || t.Status == "ACCEPTING" {
		t.PfleagueDate = 0
	}

	if t.Status == "ALREADY_HELD" {
		t.ResultAlreadyRead = 0
	}

	// TODO from database

	return t
}

func HandleCmdCheckShortPfleagueEnterableResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	t := GetCmdCheckShortPfleagueEnterableResponse()

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
