package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
	"time"
)

func GetCmdGetRankingResponse(ctx context.Context, msg *message.Message, manager *SessionManager, request *tppmessage.CmdGetRankingRequest) tppmessage.CmdGetRankingResponse {
	t := tppmessage.CmdGetRankingResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_RANKING.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	// TODO get neighbors from database

	t.UpdateDate = int(time.Now().Unix())
	player, err := manager.PlayerRepo.GetByID(ctx, msg.Platform, msg.PlayerID)
	if err != nil {
		slog.Error("get player", "error", err.Error(), "playerID", msg.PlayerID)
		t.Result = tppmessage.RESULT_ERR
		return t
	}

	t.RankingList = []tppmessage.RankingEntry{
		{
			DispRank:    player.FOBRank,
			FobGrade:    player.FOBGrade,
			IsGradeTop:  0,
			LeagueGrade: player.LeagueRank, // pf league
			PlayerInfo: tppmessage.FobPlayerInfo{
				PlayerID: msg.PlayerID,
				Npid: tppmessage.Npid{
					Handler: tppmessage.NpidHandler{
						Data:  "",
						Dummy: make([]int, 3),
						Term:  0,
					},
					Opt:      make([]int, 8),
					Reserved: make([]int, 8),
				},
				PlayerName: fmt.Sprintf("%d_player%02d", player.PlatformID, player.IDX),
				Ugc:        1,
				Xuid:       player.PlatformID,
			},
			Rank:  player.FOBRank,
			Score: player.FOBPoint,
		},
	}

	t.RankingNum = len(t.RankingList)

	return t
}

func HandleCmdGetRankingResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	t := tppmessage.CmdGetRankingResponse{}

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
