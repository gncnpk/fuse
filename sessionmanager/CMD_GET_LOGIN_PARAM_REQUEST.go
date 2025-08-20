package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdGetLoginParamRequest(ctx context.Context, message *message.Message, manager *SessionManager) error {
	var err error
	t := tppmessage.CmdGetLoginParamRequest{}
	err = json.Unmarshal(message.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	data := FromJSON(ctx, t.Msgid)
	if data != nil {
		message.MData = data
		message.Compress = true
		return nil
	}

	// TODO  t.ServerItemPlatformInfo.PlatformBaseRank matters?
	d := GetCmdGetLoginParamResponse(ctx, message, manager)

	message.MData, err = json.Marshal(d)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}
	// TODO not always needed?
	message.Compress = true

	return nil
}
