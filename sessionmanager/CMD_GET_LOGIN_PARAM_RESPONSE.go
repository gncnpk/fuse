package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/clusterbuildcost"
	fobeventreward "github.com/unknown321/fuse/fobevent/reward"
	fobeventtimebonus "github.com/unknown321/fuse/fobevent/timebonus"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
	"time"
)

func GetCmdGetLoginParamResponse(ctx context.Context, msg *message.Message, manager *SessionManager) tppmessage.CmdGetLoginParamResponse {
	t := tppmessage.CmdGetLoginParamResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_LOGIN_PARAM.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	// TODO from database
	t.EventPoint = 0
	t.HeroThresholdPoint = 150000
	t.IsAbleToBuyFob4 = 1
	t.IsStuckRescue = 0
	t.NotHeroThresholdPoint = 100000

	// fob event task result, normal defense
	{
		nd, err := manager.FOBEventTimeBonusRepo.GetByType(ctx, fobeventtimebonus.NormalDefense)
		if err != nil {
			slog.Error("fob event time bonus", "error", err.Error(), "msgid", t.Msgid, "type", fobeventtimebonus.NormalDefense)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		if len(nd) != 1 {
			slog.Error("fob event time bonus, too many values", "error", err.Error(), "msgid", t.Msgid, "type", fobeventtimebonus.NormalDefense)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		for _, v := range nd[0].SameTimeBonus {
			t.FobEventTaskResultParam.NormalDefenseSameTimeBonus = append(t.FobEventTaskResultParam.NormalDefenseSameTimeBonus, v)
		}
	}

	// fob event task result, normal sneak
	{
		nd, err := manager.FOBEventTimeBonusRepo.GetByType(ctx, fobeventtimebonus.NormalSneak)
		if err != nil {
			slog.Error("fob event time bonus", "error", err.Error(), "msgid", t.Msgid, "type", fobeventtimebonus.NormalSneak)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		if len(nd) != 1 {
			slog.Error("fob event time bonus, too many values", "error", err.Error(), "msgid", t.Msgid, "type", fobeventtimebonus.NormalSneak)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		for _, v := range nd[0].SameTimeBonus {
			t.FobEventTaskResultParam.NormalSneakSameTimeBonus = append(t.FobEventTaskResultParam.NormalSneakSameTimeBonus, v)
		}
	}

	// fob event task result, one event param
	{
		nd, err := manager.FOBEventTimeBonusRepo.GetByType(ctx, fobeventtimebonus.OneEventTaskSneak)
		if err != nil {
			slog.Error("fob event time bonus", "error", err.Error(), "msgid", t.Msgid, "type", fobeventtimebonus.OneEventTaskSneak)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		for _, v := range nd {
			p := tppmessage.OneEventParam{
				EventID:                 v.EventID,
				EventSneakClearPointMax: v.BonusMax,
				EventSneakClearPointMin: v.BonusMin,
				EventSneakSameTimeBonus: v.SameTimeBonus[:],
			}
			t.FobEventTaskResultParam.OneEventParam = append(t.FobEventTaskResultParam.OneEventParam, p)
		}
	}

	// online challenge tasks
	{
		tasks, err := manager.OnlineChallengeTaskRepo.GetByPlayerID(ctx, msg.PlayerID)
		if err != nil {
			slog.Error("online challenge tasks", "error", err.Error(), "msgid", t.Msgid, "playerID", msg.PlayerID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		for _, v := range tasks {
			e := tppmessage.OnlineChallengeTaskEntry{
				MissionID: v.MissionID,
				Reward: tppmessage.OnlineChallengeTaskReward{
					BottomType: v.RewardBottomType,
					Rate:       v.RewardRate,
					Section:    v.RewardSection,
					Type:       v.RewardType,
					Value:      v.RewardValue,
				},
				Status:     0,
				TaskTypeID: v.TaskTypeID,
				Threshold:  v.Threshold,
			}
			t.OnlineChallengeTask.TaskList = append(t.OnlineChallengeTask.TaskList, e)
		}

		if len(tasks) > 0 {
			t.OnlineChallengeTask.Version = tasks[0].Version
			t.OnlineChallengeTask.EndDate = tasks[0].EndDate
		} else {
			// TODO remove HARDCODED
			t.OnlineChallengeTask.TaskList = []tppmessage.OnlineChallengeTaskEntry{{
				MissionID: 10033,
				Reward: tppmessage.OnlineChallengeTaskReward{
					BottomType: 12,
					Rate:       1000000,
					Section:    0,
					Type:       12,
					Value:      1000,
				},
				Status:     0,
				TaskTypeID: 31,
				Threshold:  10,
			}}

			t.OnlineChallengeTask.Version = 20250502
			t.OnlineChallengeTask.EndDate = int(time.Now().Unix() + 24*60*60)
		}

	}

	// server product params
	{
		products, err := manager.ServerProductParamRepo.GetByPlayerID(ctx, msg.PlayerID)
		if err != nil {
			slog.Error("server product params", "error", err.Error(), "msgid", t.Msgid, "playerID", msg.PlayerID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		// TODO `open` depends on player
		for _, product := range products {
			t.ServerProductParams = append(t.ServerProductParams, tppmessage.ServerProductParam{
				DevCoin:            product.DevCoin,
				DevGmp:             product.DevGmp,
				DevItem1:           product.DevItem1,
				DevItem2:           product.DevItem2,
				DevPlatlv01:        product.DevPlatlv01,
				DevPlatlv02:        product.DevPlatlv02,
				DevPlatlv03:        product.DevPlatlv03,
				DevPlatlv04:        product.DevPlatlv04,
				DevPlatlv05:        product.DevPlatlv05,
				DevPlatlv06:        product.DevPlatlv06,
				DevPlatlv07:        product.DevPlatlv07,
				DevRescount01Value: product.DevRescount01Value,
				DevRescount02Value: product.DevRescount02Value,
				DevResource01ID:    product.DevResource01Id,
				DevResource02ID:    product.DevResource02Id,
				DevSkil:            product.DevSkil,
				DevSpecial:         product.DevSpecial,
				DevTime:            product.DevTime,
				ID:                 product.ID,
				Open:               product.Open,
				Type:               product.Type,
				UseGmp:             product.UseGmp,
				UseRescount01Value: product.UseRescount01Value,
				UseRescount02Value: product.UseRescount02Value,
				UseResource01ID:    product.UseResource01Id,
				UseResource02ID:    product.UseResource02Id,
			})
		}
	}

	// fob event task list, one event task
	{
		nd, err := manager.FOBEventRewardRepo.GetByType(ctx, fobeventreward.OneEventTaskSneak)
		if err != nil {
			slog.Error("fob reward", "error", err.Error(), "msgid", t.Msgid, "type", fobeventreward.OneEventTaskSneak)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		oet := make(map[int][]tppmessage.FobEventTaskListRewardInfo)
		for _, v := range nd {
			oet[v.EventID] = append(oet[v.EventID], tppmessage.FobEventTaskListRewardInfo{
				Reward:     v.Reward,
				TaskTypeID: v.TaskTypeID,
				Threshold:  v.Threshold,
			})
		}

		for k, v := range oet {
			t.FobEventTaskList.OneEventTask = append(t.FobEventTaskList.OneEventTask, tppmessage.OneEventTask{
				EventID:    k,
				EventSneak: v,
			})
		}
	}

	// fob event task list, normal defence
	{
		nd, err := manager.FOBEventRewardRepo.GetByType(ctx, fobeventreward.NormalDefense)
		if err != nil {
			slog.Error("fob reward", "error", err.Error(), "msgid", t.Msgid, "type", fobeventreward.NormalDefense)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		for _, v := range nd {
			t.FobEventTaskList.NormalDefense = append(t.FobEventTaskList.NormalDefense, tppmessage.FobEventTaskListRewardInfo{
				Reward:     v.Reward,
				TaskTypeID: v.TaskTypeID,
				Threshold:  v.Threshold,
			})
		}
	}

	// fob event task list, normal sneak
	{
		nd, err := manager.FOBEventRewardRepo.GetByType(ctx, fobeventreward.NormalSneak)
		if err != nil {
			slog.Error("fob reward", "error", err.Error(), "msgid", t.Msgid, "type", fobeventreward.NormalSneak)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		for _, v := range nd {
			t.FobEventTaskList.NormalSneak = append(t.FobEventTaskList.NormalSneak, tppmessage.FobEventTaskListRewardInfo{
				Reward:     v.Reward,
				TaskTypeID: v.TaskTypeID,
				Threshold:  v.Threshold,
			})
		}
	}

	// cluster build costs
	{
		t.ClusterBuildCostsPerCluster = make([]tppmessage.ClusterBuildCostsPerCluster, 2)
		for fobs := 0; fobs < 2; fobs++ {
			t.ClusterBuildCostsPerCluster[fobs].ClusterBuildCostsPerGrade = make([]tppmessage.ClusterBuildCostsPerGrade, 7)
			for grades := 0; grades < 7; grades++ {
				t.ClusterBuildCostsPerCluster[fobs].ClusterBuildCostsPerGrade[grades].ClusterBuildCosts = make([]clusterbuildcost.ClusterBuildCost, 4)
			}
		}

		costs, err := manager.ClusterBuildCostRepo.GetAll(ctx)
		if err != nil {
			slog.Error("cluster build cost", "error", err.Error(), "msgid", t.Msgid)
			t.Result = tppmessage.RESULT_ERR
			return t
		}
		for _, v := range costs {
			t.ClusterBuildCostsPerCluster[v.FOBNumber].ClusterBuildCostsPerGrade[v.Grade].ClusterBuildCosts[v.IDX] = v
		}
	}

	// espionage events
	{
		espIDs, err := manager.EspionageEventRepo.GetAll(ctx)
		if err != nil {
			slog.Error("espionage events", "error", err.Error(), "msgid", t.Msgid)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		for _, id := range espIDs {
			t.RankingEspiEventIDs = append(t.RankingEspiEventIDs, id.EventID)
		}
	}

	// pf events
	{
		pfIDS, err := manager.PFEventRepo.GetAll(ctx)
		if err != nil {
			slog.Error("pf events", "error", err.Error(), "msgid", t.Msgid)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		for _, id := range pfIDS {
			t.RankingPfEventIDs = append(t.RankingPfEventIDs, id.EventID)
		}
	}

	// server texts
	{
		texts, err := manager.ServerTextRepo.GetAll(ctx)
		if err != nil {
			slog.Error("server texts", "error", err.Error(), "msgid", t.Msgid)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		for _, v := range texts {
			t.ServerTexts = append(t.ServerTexts, tppmessage.ServerText{
				Identifier: v.Identifier,
				Language:   v.Language,
				Text:       v.Text,
			})
		}
	}

	// staff rank bonus rates
	{
		t.StaffRankBonusRates = []tppmessage.StaffRankBonusRate{}
		rates, err := manager.StaffRankBonusRateRepo.GetAll(ctx)
		if err != nil {
			slog.Error("staff rank bonus rate", "error", err.Error(), "msgid", t.Msgid)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		for _, r := range rates {
			nr := tppmessage.StaffRankBonusRate{Rates: []int{r.Negative, r.Positive}}
			t.StaffRankBonusRates = append(t.StaffRankBonusRates, nr)
		}
	}

	return t
}

func HandleCmdGetLoginParamResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	var err error
	//t := GetCmdGetLoginParamResponse()
	t := tppmessage.CmdGetLoginParamResponse{}

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
