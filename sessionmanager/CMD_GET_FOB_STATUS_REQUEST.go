package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdGetFobStatusRequest(ctx context.Context, msg *message.Message, manager *SessionManager) error {
	var err error
	t := tppmessage.CmdGetFobStatusRequest{}
	err = json.Unmarshal(msg.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	data := FromJSON(ctx, t.Msgid)
	if data != nil {
		msg.MData = data
		return nil
	}

	d := GetCmdGetFobStatusResponse(ctx, msg, manager)

	msg.MData, err = json.Marshal(d)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
