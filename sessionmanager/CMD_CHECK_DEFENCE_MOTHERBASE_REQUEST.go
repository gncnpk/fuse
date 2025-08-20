package sessionmanager

import (
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdCheckDefenceMotherbaseRequest(message *message.Message) error {
	var err error
	t := tppmessage.CmdCheckDefenceMotherbaseRequest{}
	err = json.Unmarshal(message.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	d := tppmessage.GetCmdCheckDefenceMotherbaseResponse()

	message.MData, err = json.Marshal(d)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
