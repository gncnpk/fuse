package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"fuse/message"
	"fuse/tppmessage"
)

func HandleCmdGetWormholeListRequest(ctx context.Context, msg *message.Message, manager *SessionManager) error {
	var err error
	t := tppmessage.CmdGetWormholeListRequest{}
	err = json.Unmarshal(msg.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	data := FromJSON(ctx, t.Msgid)
	if data != nil {
		msg.MData = data
		return nil
	}

	d := GetCmdGetWormholeListResponse(ctx, msg, manager, &t)

	msg.MData, err = json.Marshal(d)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
