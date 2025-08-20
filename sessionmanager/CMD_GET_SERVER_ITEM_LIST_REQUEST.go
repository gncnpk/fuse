package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdGetServerItemListRequest(ctx context.Context, message *message.Message, manager *SessionManager) error {
	var err error
	t := tppmessage.CmdGetServerItemListRequest{}
	err = json.Unmarshal(message.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	data := FromJSON(ctx, t.Msgid)
	if data != nil {
		message.MData = data
		return nil
	}

	// TODO use request params
	d := GetCmdGetServerItemListResponse(ctx, message, manager)

	message.MData, err = json.Marshal(d)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
