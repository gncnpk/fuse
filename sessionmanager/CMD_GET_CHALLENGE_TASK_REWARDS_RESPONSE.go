package sessionmanager

import (
	"context"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdGetChallengeTaskRewardsResponse(ctx context.Context, manager *SessionManager, userID int) tppmessage.CmdGetChallengeTaskRewardsResponse {
	t := tppmessage.CmdGetChallengeTaskRewardsResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_CHALLENGE_TASK_REWARDS.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	res, err := manager.TaskRewardRepo.GetByUser(ctx, userID)
	if err != nil {
		slog.Error("get challenge task reward list", "user", userID, "error", err.Error())
		return t
	}

	for _, v := range res {
		q := tppmessage.CmdGetChallengeTaskRewardsTask{
			Reward: tppmessage.CmdGetChallengeTaskRewardsReward{
				BottomType: v.Reward.BottomType,
				Rate:       v.Reward.Rate,
				Section:    v.Reward.Section,
				Type:       v.Reward.Type,
				Value:      v.Reward.Value,
			},
			TaskID: v.ID,
		}

		t.TaskList = append(t.TaskList, q)
	}

	t.TaskCount = len(t.TaskList)
	slog.Debug("got tasks", "count", t.TaskCount, "userID", userID)

	return t
}
