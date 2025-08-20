package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdGetFobEventListResponse(ctx context.Context, msg *message.Message, manager *SessionManager) tppmessage.CmdGetFobEventListResponse {
	t := tppmessage.CmdGetFobEventListResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_FOB_EVENT_LIST.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	// TODO get from database

	/*
	   "attacker_id": playerID,
	   "cluster": 6,
	   "event_index": 19, -> descending, might have missing entries
	   "fob_index": 1,
	   "is_win": 0,
	   "layout_code": 73
	*/

	t.EventList = []tppmessage.CmdGetFobEventListResponseEvent{}
	t.EventNum = len(t.EventList) // 10 max

	return t
}

func HandleCmdGetFobEventListResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	t := tppmessage.CmdGetFobEventListResponse{}

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
