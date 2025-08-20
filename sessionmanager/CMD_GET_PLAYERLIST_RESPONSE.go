package sessionmanager

import (
	"context"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/platform"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdGetPlayerlistResponse(ctx context.Context, manager *SessionManager, plat platform.Platform, userID int) (tppmessage.CmdGetPlayerListResponse, error) {
	t := tppmessage.CmdGetPlayerListResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_PLAYERLIST.String()
	t.Result = tppmessage.RESULT_ERR
	t.Rqid = 0

	entries, err := manager.PlayerRepo.GetAllByUserID(ctx, plat, userID)
	if err != nil {
		return tppmessage.CmdGetPlayerListResponse{}, fmt.Errorf("player list by userID %d: %w", userID, err)
	}

	for _, e := range entries {
		c := tppmessage.CmdGetPlayerListEntry{
			EspionageLose: e.EspionageLose,
			EspionageWin:  e.EspionageWin,
			FobGrade:      e.FOBGrade,
			FobPoint:      e.FOBPoint,
			FobRank:       e.FOBRank,
			Index:         e.IDX,
			IsInsurance:   e.IsInsurance,
			LeagueGrade:   e.LeagueGrade,
			LeagueRank:    e.LeagueRank,
			Name:          fmt.Sprintf("%d_player%02d", e.PlatformID, e.IDX),
			Playtime:      e.Playtime,
			Point:         e.Point,
		}
		t.PlayerList = append(t.PlayerList, c)
	}
	t.PlayerNum = len(t.PlayerList)
	t.Result = tppmessage.RESULT_NOERR

	return t, nil
}

func HandleCmdGetPlayerlistResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	//var err error
	//t := GetCmdGetPlayerlistResponse()

	//message.MData, err = json.Marshal(t)
	//if err != nil {
	//	return fmt.Errorf("cannot marshal: %w", err)
	//}

	return nil
}
