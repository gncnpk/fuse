package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdSendSneakResultResponse(ctx context.Context, msg *message.Message, manager *SessionManager, request *tppmessage.CmdSendSneakResultRequest) tppmessage.CmdSendSneakResultResponse {
	t := tppmessage.CmdSendSneakResultResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_SEND_SNEAK_RESULT.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	// TODO from database
	t.EventPoint = 0
	t.IsWormholeOpen = 0
	t.IsSecurityChallenge = 0
	t.ResultReward = tppmessage.SneakResultReward{
		AfricanPeach:   0,
		Biotic:         0,
		BlackCarrot:    0,
		Common:         0,
		DigitalisL:     0,
		DigitalisP:     0,
		Fuel:           0,
		Gmp:            0,
		GoldenCrescent: 0,
		Haoma:          0,
		IsBefore:       0,
		KeyItem:        0,
		MainType:       0,
		Minor:          0,
		ParamID:        0,
		Precious:       0,
		Rate:           0,
		Section:        0,
		StaffCount:     0,
		StaffRank:      make([]int, 64),
		StaffType:      0,
		Tarragon:       0,
		Wormwood:       0,
	}

	t.SneakPoint = request.SneakPoint

	return t
}

func HandleCmdSendSneakResultResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	//t := GetCmdSendSneakResultResponse()
	t := tppmessage.CmdSendSneakResultResponse{}

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
