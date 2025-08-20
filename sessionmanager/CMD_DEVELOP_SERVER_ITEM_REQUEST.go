package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/serveritem"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

// HandleCmdDevelopServerItemRequest processes a request to develop a server item and persists it
func HandleCmdDevelopServerItemRequest(ctx context.Context, msg *message.Message, manager *SessionManager) error {
	// parse request
	t := tppmessage.CmdDevelopServerItemRequest{}
	if err := json.Unmarshal(msg.MData, &t); err != nil {
		return fmt.Errorf("cannot unmarshal develop server item request: %w", err)
	}
	// allow JSON override for testing
	if data := FromJSON(ctx, t.Msgid); data != nil {
		msg.MData = data
		return nil
	}
	// record development entry
	now := int(time.Now().Unix())
   si := serveritem.ServerItem{
	   ProductID:  t.ID,
	   PlayerID:   msg.PlayerID,
	   CreateDate: now,
	   Develop:    1,
	   MbCoin:     0,
	   Open:       1, // mark as developed/open
   }
	if err := manager.ServerItemRepo.Add(ctx, &si); err != nil {
		slog.Error("develop server item failed", "error", err.Error(), "playerID", msg.PlayerID, "itemID", t.ID)
		// respond with error
		resp := tppmessage.CmdDevelopServerItemResponse{
			CryptoType: tppmessage.CRYPTO_TYPE_COMPOUND,
			Flowid:     nil,
			Msgid:      tppmessage.CMD_DEVELOP_SERVER_ITEM.String(),
			Result:     tppmessage.RESULT_ERR,
			Rqid:       t.Rqid,
			Xuid: 		nil,
		}
		data, _ := json.Marshal(resp)
		msg.MData = data
		return nil
	}
	// respond with success
	resp := tppmessage.CmdDevelopServerItemResponse{
		CryptoType: tppmessage.CRYPTO_TYPE_COMPOUND,
		Flowid:     nil,
		Msgid:      tppmessage.CMD_DEVELOP_SERVER_ITEM.String(),
		Result:     tppmessage.RESULT_NOERR,
		Rqid:       t.Rqid,
		Xuid: 		nil,
	}
	data, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("cannot marshal develop server item response: %w", err)
	}
	msg.MData = data
	return nil
}
