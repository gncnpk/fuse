package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdGetUrllistRequest(ctx context.Context, message *message.Message, repo *tppmessage.URLListEntryRepo) error {
	var err error
	t := tppmessage.CmdGetUrllistRequest{}
	err = json.Unmarshal(message.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	d, err := GetCmdGetUrllistResponse(ctx, repo)
	if err != nil {
		return err
	}
	
	message.MData, err = json.Marshal(d)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
