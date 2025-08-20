package sessionmanager

import (
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/steamid"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
	"strconv"
)

func HandleCmdAuthSteamticketRequest(message *message.Message) error {
	in := tppmessage.CmdAuthSteamticketRequest{}
	var err error
	err = json.Unmarshal(message.MData, &in)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	slog.Debug("steam auth request", "ticket", in.SteamTicket)

	t := GetCmdAuthSteamticketResponse()

	steamID, err := steamid.FromTicket(in.SteamTicket)
	if err != nil {
		return fmt.Errorf("cannot get steam id from ticket: %w", err)
	}

	t.AccountID = strconv.Itoa(int(steamID))

	slog.Info("CMD_AUTH_STEAMTICKET", "steamID", t.AccountID)

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
