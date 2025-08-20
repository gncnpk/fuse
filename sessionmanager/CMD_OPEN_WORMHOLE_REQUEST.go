package sessionmanager

import (
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdOpenWormholeRequest(message *message.Message) error {
	var err error
	t := tppmessage.CmdOpenWormholeRequest{}
	err = json.Unmarshal(message.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	// TODO use request parameters

	d := GetCmdOpenWormholeResponse()

	d.PlayerID = t.PlayerID
	d.ToPlayerID = t.ToPlayerID
	d.IsNewOpen = 1

	message.MData, err = json.Marshal(d)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
