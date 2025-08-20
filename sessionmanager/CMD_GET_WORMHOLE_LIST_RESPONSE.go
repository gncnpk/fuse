package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdGetWormholeListResponse(ctx context.Context, msg *message.Message, manager *SessionManager, request *tppmessage.CmdGetWormholeListRequest) tppmessage.CmdGetWormholeListResponse {
	t := tppmessage.CmdGetWormholeListResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_WORMHOLE_LIST.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	// TODO from database

	t.WormholeList = make([]tppmessage.WormholeList, 0)
	t.WormholeNum = len(t.WormholeList)

	return t
}

func HandleCmdGetWormholeListResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	t := tppmessage.CmdGetWormholeListResponse{}

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
