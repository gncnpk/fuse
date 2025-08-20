package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/emblem"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdSyncEmblemResponse(ctx context.Context, msg *message.Message, manager *SessionManager, request *tppmessage.CmdSyncEmblemRequest) tppmessage.CmdSyncEmblemResponse {
	t := tppmessage.CmdSyncEmblemResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_SYNC_EMBLEM.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	for i, v := range request.Emblem.Parts {
		e := emblem.Emblem{
			PlayerID:   msg.PlayerID,
			IDX:        i,
			BaseColor:  v.BaseColor,
			FrameColor: v.FrameColor,
			PositionX:  v.PositionX,
			PositionY:  v.PositionY,
			Rotate:     v.Rotate,
			Scale:      v.Scale,
			TextureTag: v.TextureTag,
		}

		if err := manager.EmblemRepo.AddOrUpdate(ctx, &e); err != nil {
			slog.Error("sync emblem", "error", err.Error(), "msgID", t.Msgid)
			t.Result = tppmessage.RESULT_ERR
			return t
		}
	}

	return t
}

func HandleCmdSyncEmblemResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	//t := GetCmdSyncEmblemResponse()
	t := tppmessage.CmdSyncEmblemResponse{}

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
