package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdGetOwnFobListResponse(ctx context.Context, msg *message.Message, manager *SessionManager) tppmessage.CmdGetOwnFobListResponse {
	t := tppmessage.CmdGetOwnFobListResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_OWN_FOB_LIST.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	// TODO save to database where?
	t.EnableSecurityChallenge = 0

	fobs, err := manager.MotherBaseParamRepo.GetByPlayerID(ctx, msg.PlayerID)
	if err != nil {
		slog.Error("get fobs", "error", err.Error(), "msgid", msg.MsgID, "playerID", msg.PlayerID)
		t.Result = tppmessage.RESULT_ERR
		return t
	}

	for _, f := range fobs {
		t.Fob = append(t.Fob, tppmessage.MotherBaseParam{
			AreaID:         f.AreaID,
			ClusterParam:   []tppmessage.ClusterParam{}, // always empty
			ConstructParam: f.ConstructParam,
			FobIndex:       f.FobIndex,
			MotherBaseID:   f.ID,
			PlatformCount:  0,    // always 0
			Price:          1000, // always 1000
			SecurityRank:   0,    // always 0
		})
	}

	return t
}

func HandleCmdGetOwnFobListResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	//t := GetCmdGetOwnFobListResponse()
	t := tppmessage.CmdGetOwnFobListResponse{}

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
