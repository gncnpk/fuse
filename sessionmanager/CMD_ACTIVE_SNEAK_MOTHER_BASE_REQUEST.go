package sessionmanager

import (
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func HandleCmdActiveSneakMotherBaseRequest(message *message.Message) error {
	var err error
	t := tppmessage.CmdActiveSneakMotherBaseRequest{}
	err = json.Unmarshal(message.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	slog.Debug("CMD_ACTIVE_SNEAK_MOTHER_BASE", "mother_base_id", t.MotherBaseID)

	d := GetCmdActiveSneakMotherBaseResponse()

	message.MData, err = json.Marshal(d)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
