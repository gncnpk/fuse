package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func HandleCmdGetInformationlist2Request(ctx context.Context, message *message.Message, manager *SessionManager) error {
	var err error
	t := tppmessage.CmdGetInformationlist2Request{}
	err = json.Unmarshal(message.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	slog.Info("CMD_GET_INFORMATIONLIST2", "region", t.Region)

	data := FromJSON(ctx, t.Msgid)
	if data != nil {
		message.MData = data
		return nil
	}

	d := GetCmdGetInformationlist2Response(ctx, t.Region, t.Lang, manager)

	message.MData, err = json.Marshal(d)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
