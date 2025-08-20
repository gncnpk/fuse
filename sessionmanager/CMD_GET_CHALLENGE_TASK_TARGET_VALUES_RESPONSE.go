package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdGetChallengeTaskTargetValuesResponse(ctx context.Context, msg *message.Message, manager *SessionManager) tppmessage.CmdGetChallengeTaskTargetValuesResponse {
	t := tppmessage.CmdGetChallengeTaskTargetValuesResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_CHALLENGE_TASK_TARGET_VALUES.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	{
		ps, err := manager.PlayerStatusRepo.Get(ctx, msg.PlayerID)
		if err != nil {
			slog.Error("player status", "error", err.Error(), "playerID", msg.PlayerID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		t.EspionageRatingGrade = ps.EspionageRatingGrade
		t.FobDeployToSupportersEmergencyCount = ps.FobDeployToSupportersEmergencyCount
		t.FobSupportingUserCount = ps.FobSupportingUserCount
		t.PfRatingDefenseForce = ps.PfRatingDefenseForce
		t.PfRatingDefenseLife = ps.PfRatingDefenseLife
		t.PfRatingOffenceForce = ps.PfRatingOffenceForce
		t.PfRatingOffenceLife = ps.PfRatingOffenceLife
		t.PfRatingRank = ps.PfRatingRank
		t.TotalDevelopmentGrade = ps.TotalDevelopmentGrade
		t.TotalFobSecurityLevel = ps.TotalFobSecurityLevel
	}

	{
		recs, err := manager.FobRecordRepo.GetByPlayerID(ctx, msg.PlayerID)
		if err != nil {
			slog.Error("fob record", "error", err.Error(), "playerID", msg.PlayerID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}
		if len(recs) < 1 {
			slog.Error("fob record not found", "playerID", msg.PlayerID)
			t.Result = tppmessage.RESULT_ERR
			return t
		}

		t.FobDefenseSuccessCount = recs[0].DefenseWin
		t.FobSneakCount = recs[0].SneakLose + recs[0].SneakWin // might be wrong?
		t.FobSneakSuccessCount = recs[0].SneakWin
	}

	return t
}

func HandleCmdGetChallengeTaskTargetValuesResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	//t := GetCmdGetChallengeTaskTargetValuesResponse()
	t := tppmessage.CmdGetChallengeTaskTargetValuesResponse{}

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
