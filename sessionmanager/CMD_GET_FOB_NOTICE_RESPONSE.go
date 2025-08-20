package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	fobevent "github.com/unknown321/fuse/fobevent/event"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
	"time"
)

func GetCmdGetFobNoticeResponse(ctx context.Context, msg *message.Message, manager *SessionManager) tppmessage.CmdGetFobNoticeResponse {
	t := tppmessage.CmdGetFobNoticeResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_FOB_NOTICE.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	t.CommonServerText = "NotImplement"
	t.CommonServerTextTitle = "NotImplement"
	t.PointExchangeEventServerText = "NotImplement"
	t.CampaignParamList = []any{}                 // always 0?
	t.Daily = 0                                   // always 0?
	t.LeagueUpdate = tppmessage.FobNoticeUpdate{} // always 0?
	t.SneakUpdate = tppmessage.FobNoticeUpdate{}  // always 0?
	t.ExistsEventPointCombatDeploy = 0            // always 0?

	{
		event, err := manager.FobEventRepo.GetActive(ctx)
		if err != nil {
			slog.Error("get active fob event", "error", err.Error())
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		if len(event) == 0 {
			slog.Warn("no active fob event found, using dummy")
			event = append(event, fobevent.Event{
				ID:         1,
				Active:     true,
				EndDate:    int(time.Now().Unix()) + 24*60*60,
				DeleteDate: 0,
				Flag:       1488, // 1489, 465, 1528, 1520, 1496
			})
		}

		t.ActiveEventServerText = fmt.Sprintf("mb_fob_event_name_%02d", event[0].ID)
		t.EventEndDate = event[0].EndDate
		t.Flag = event[0].Flag
		t.EventDeleteDate = event[0].DeleteDate // delete date is bigger than end date by 3 hours, but not always, usually it's 0
	}

	// mb coins
	{
		ps, err := manager.PlayerStatusRepo.Get(ctx, msg.PlayerID)
		if err != nil {
			slog.Error("get player status", "error", err.Error(), "msgid", t.Msgid, "playerID", msg.PlayerID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		t.MbCoin = ps.MbCoin
	}

	{
		record, err := manager.FobRecordRepo.GetByPlayerID(ctx, msg.PlayerID)
		if err != nil {
			slog.Error("get fob record", "error", err.Error(), "msgid", t.Msgid)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		if len(record) != 1 {
			slog.Error("get fob record", "len", len(record), "want", 1, "msgid", t.Msgid)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		// other fields are 0?
		t.Record = tppmessage.FobRecord{
			Defense: tppmessage.FobRecordWinLose{Lose: record[0].DefenseLose, Win: record[0].DefenseWin},
		}
	}

	// pf league
	{
		seasons, err := manager.PFSeasonRepo.GetActive(ctx)
		if err != nil {
			slog.Error("get pf seasons", "error", err.Error(), "msgid", t.Msgid)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		if len(seasons) != 2 {
			slog.Error("get pf seasons", "len", len(seasons), "want", 2, "msgid", t.Msgid)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		for _, s := range seasons {
			if s.IsShort == 0 {
				t.PfCurrentSeason = s.ID
				t.PfFinishNumMax = s.PlayerNum
			} else {
				t.ShortPfCurrentSeason = s.ID
				t.ShortPfFinishNumMax = s.PlayerNum
			}
		}

		ranks, err := manager.PFRankingRepo.GetByPlayerID(ctx, msg.PlayerID, t.PfCurrentSeason, t.ShortPfCurrentSeason)
		if err != nil {
			slog.Error("get pf ranks", "error", err.Error(), "msgid", t.Msgid, "playerID", msg.PlayerID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		for _, r := range ranks {
			if r.IsShort == 1 {
				t.ShortPfFinishNum = r.Finish
			} else {
				t.PfFinishNum = r.Finish
			}
		}
	}

	return t
}

func HandleCmdGetFobNoticeResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	//t := GetCmdGetFobNoticeResponse()
	t := tppmessage.CmdGetFobNoticeResponse{}

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
