package sessionmanager

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/platform"
	"github.com/unknown321/fuse/playerstatus"
	"github.com/unknown321/fuse/session"
	"github.com/unknown321/fuse/steamid"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
	"strconv"
	"time"
)

// HandleCmdReqAuthHTTPSResponse handles kojipro response
func HandleCmdReqAuthHTTPSResponse(ctx context.Context, message *message.Message, manager *SessionManager) error {
	tppm := &tppmessage.CMDReqAuthHTTPSResponse{}

	err := json.Unmarshal(message.MData, &tppm)
	if err != nil {
		return fmt.Errorf("cannot unmarshal tpp message: %w", err)
	}

	cryptoKey, err := base64.StdEncoding.DecodeString(tppm.CryptoKey)
	if err != nil {
		return fmt.Errorf("cannot decode crypt key: %w", err)
	}

	if err = manager.Add(ctx, tppm.Session, cryptoKey, tppm.UserID, strconv.Itoa(steamid.InvalidSteamID)); err != nil {
		return fmt.Errorf("cannot add session: %w", err)
	}

	return nil
}

func GetAuthResponse(ctx context.Context, steamID string, manager *SessionManager, pltf platform.Platform) (tppmessage.CMDReqAuthHTTPSResponse, error) {
	t := tppmessage.CMDReqAuthHTTPSResponse{}
	t.Msgid = tppmessage.CMD_REQAUTH_HTTPS.String()
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMMON
	t.Result = tppmessage.RESULT_ERR
	// TODO to database
	t.HeartbeatSec = 60
	t.IsUseApr = 0
	t.InquiryId = 1 // TODO ??
	t.TimeoutSec = session.Timeout

	sid, err := steamid.ValidateString(steamID)
	if err != nil {
		return t, err
	}

	sess, err := manager.SessionRepo.GetByPlatformID(ctx, steamID)
	if err == nil && sess != nil {
		if (time.Now().Unix() - sess.Timestamp) >= int64(session.Timeout) {
			if err = manager.SessionRepo.Remove(ctx, sess); err != nil {
				return t, fmt.Errorf("remove outdated session: %w", err)
			}
			slog.Info("removed outdated session", "steamID", steamID)
			sess = nil
		} else {
			slog.Info("found existing session", "steamID", steamID)
			t.CryptoKey = base64.StdEncoding.EncodeToString(sess.CryptoKey)
			t.Session = sess.ID
			t.UserID = sess.UserID
			t.SmartDeviceID = sess.SmartDeviceID
			t.Result = tppmessage.RESULT_NOERR
			return t, nil
		}
	}

	if sess != nil {
		return t, fmt.Errorf("session was found but not reused, steamID %s", steamID)
	}

	sess, err = session.New()
	if err != nil {
		return t, err
	}

	sess.PlatformID = sid
	sess.Platform = pltf

	t.CryptoKey = base64.StdEncoding.EncodeToString(sess.CryptoKey)
	t.Session = sess.ID
	t.SmartDeviceID = sess.SmartDeviceID

	uid := 0
	res, err := manager.UserRepo.Get(ctx, sid)
	if err != nil {
		slog.Warn("user not found", "error", err.Error(), "steamID", sid)
		if uid, err = manager.UserRepo.Add(ctx, sid); err != nil {
			return t, fmt.Errorf("cannot add new user %d: %w", sid, err)
		}
		slog.Info("created user", "id", uid, "steamID", sid)
		pid := 0
		if pid, err = manager.PlayerRepo.Add(ctx, platform.Steam, sid); err != nil {
			return t, err
		}
		slog.Info("created player", "id", pid, "platform", platform.Steam)

		if err = manager.PlayerStatusRepo.AddOrUpdate(ctx, &playerstatus.PlayerStatus{PlayerID: pid, ServerGmp: manager.ManagerOpts.SignupBonus.GMP}); err != nil {
			slog.Error("add player status", "error", err.Error(), "playerID", pid)
			t.Result = tppmessage.RESULT_ERR
			return t, err
		}
	} else {
		uid = res.ID
		slog.Info("found user", "id", uid, "steamID", sid)
	}

	sess.UserID = uid
	manager.sessions[sess.ID] = sess
	if err = manager.SessionRepo.Add(ctx, sess); err != nil {
		slog.Error("cannot add session to database", "error", err.Error())
		t.Result = tppmessage.RESULT_ERR
		return t, err
	}

	t.UserID = uid
	t.Result = tppmessage.RESULT_NOERR

	return t, nil
}

func HandleCmdReqAuthHTTPSRequest(ctx context.Context, message *message.Message, manager *SessionManager, pltf platform.Platform) error {
	tppm := &tppmessage.CMDReqAuthHTTPSRequest{}

	err := json.Unmarshal(message.MData, &tppm)
	if err != nil {
		return fmt.Errorf("cannot unmarshal tpp message: %w", err)
	}

	resp, err := GetAuthResponse(ctx, tppm.UserName, manager, pltf)
	if err != nil {
		return fmt.Errorf("cannot get auth response: %w", err)
	}

	if message.MData, err = json.Marshal(resp); err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
