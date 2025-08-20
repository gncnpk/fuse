package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdGetMBCoinRemainderResponse(ctx context.Context, manager *SessionManager, msg *message.Message) tppmessage.CmdGetMBCoinRemainderResponse {
	t := tppmessage.CmdGetMBCoinRemainderResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_MBCOIN_REMAINDER.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	pl, err := manager.PlayerStatusRepo.Get(ctx, msg.PlayerID)
	if err != nil {
		slog.Error("get player status", "error", err.Error(), "msgID", t.Msgid, "playerID", msg.PlayerID)
	}

	t.Remainder = pl.MbCoin
	return t
}

func HandleCmdGetMBCoinRemainderResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	//t := GetCmdGetMBCoinRemainderResponse()
	t := tppmessage.CmdGetMBCoinRemainderResponse{}

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
