package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func HandleCmdSetCurrentplayerRequest(ctx context.Context, message *message.Message, manager *SessionManager) error {
	var err error
	t := tppmessage.CmdSetCurrentplayerRequest{}
	err = json.Unmarshal(message.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	slog.Info("CMD_SET_CURRENTPLAYER", "index", t.Index, "playerID", message.PlayerID)

	data := FromJSON(ctx, t.Msgid)
	if data != nil {
		message.MData = data
		return nil
	}

	reset := false
	if t.IsReset > 0 {
		reset = true
		slog.Warn("reset ON, why?", "msgID", t.Msgid)
	}
	d := GetCmdSetCurrentplayerResponse(ctx, manager, *message.SessionKey, message.Platform, message.PlatformID, t.Index, reset)

	message.MData, err = json.Marshal(d)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
