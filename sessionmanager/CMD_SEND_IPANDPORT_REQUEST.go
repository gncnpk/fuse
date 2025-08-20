package sessionmanager

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
	"net"
)

func HandleCmdSendIPAndPortRequest(ctx context.Context, message *message.Message, manager *SessionManager) error {
	var err error
	t := tppmessage.CmdSendIpandportRequest{}
	err = json.Unmarshal(message.MData, &t)
	if err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	slog.Debug("CMD_SEND_IPANDPORT", "in_ip", t.InIp, "in_port", t.InPort, "ex_ip", t.ExIp, "ex_port", t.ExPort, "nat", t.Nat, "account_id", message.PlayerID)

	data := FromJSON(ctx, t.Msgid)
	if data != nil {
		message.MData = data
		return nil
	}

	tt := GetCmdSendIpandportResponse()

	inIp := net.ParseIP(t.InIp)
	in := int(binary.BigEndian.Uint32(inIp.To4()))
	exIp := net.ParseIP(t.InIp)
	ex := int(binary.BigEndian.Uint32(exIp.To4()))

	if err = manager.SetIP(ctx, *message.SessionKey, in, t.InPort, ex, t.ExPort); err != nil {
		slog.Error("cannot set ip", "error", err.Error())
		tt.Result = tppmessage.RESULT_ERR
	}

	message.MData, err = json.Marshal(tt)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
