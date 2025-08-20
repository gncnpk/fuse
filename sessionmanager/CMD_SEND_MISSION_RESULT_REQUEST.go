package sessionmanager

import (
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func HandleCmdSendMissionResultRequest(message *message.Message) error {
	var err error
	t := tppmessage.CmdSendMissionResultRequest{}
	err = json.Unmarshal(message.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	slog.Debug("CMD_SEND_MISSION_RESULT", "mission_id", t.MissionID)

	// TODO save to database

	d := GetCmdSendMissionResultResponse()

	message.MData, err = json.Marshal(d)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
