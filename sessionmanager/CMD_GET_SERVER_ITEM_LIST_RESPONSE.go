package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
	"time"
)

func GetCmdGetServerItemListResponse(ctx context.Context, msg *message.Message, manager *SessionManager) tppmessage.CmdGetServerItemListResponse {
	t := tppmessage.CmdGetServerItemListResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_SERVER_ITEM_LIST.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	// TODO from database HARDCODED
	t.DevelopLimit = 4

	items, err := manager.ServerItemRepo.GetByPlayerID(ctx, msg.PlayerID)
	if err != nil {
		slog.Error("get server items", "error", err.Error(), "playerID", msg.PlayerID)
		t.Result = tppmessage.RESULT_ERR
		return t
	}

	now := int(time.Now().Unix())
	for _, v := range items {
		   // determine develop status: 1 in-progress, 2 finished
		   developVal := v.Develop
		   finishTime := v.CreateDate + v.MaxSecond
		   if v.Develop != 0 && now >= finishTime {
			   developVal = 2
			   // persist finished status
			   if v.Develop != 2 {
				   if err := manager.ServerItemRepo.UpdateDevelop(ctx, msg.PlayerID, v.ProductID, 2); err != nil {
					   slog.Error("failed to persist develop status", "error", err.Error(), "playerID", msg.PlayerID, "productID", v.ProductID)
				   }
			   }
		   }
		q := tppmessage.ServerItemListEntry{
			CreateDate: v.CreateDate,
			Develop:    developVal,
			Gmp:        v.Gmp,
			ID:         v.ProductID,
			LeftSecond: 0,
			MaxSecond:  v.MaxSecond,
			MbCoin:     v.MbCoin,
			Open:       v.Open,
		}

		if v.Develop != 0 {
			if (v.MaxSecond + v.CreateDate) > now {
				q.LeftSecond = (v.CreateDate + v.MaxSecond) - now
			}
		}

		t.ItemList = append(t.ItemList, q)
	}

	t.ItemNum = len(t.ItemList)

	return t
}

func HandleCmdGetServerItemListResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	//t := GetCmdGetServerItemListResponse()
	t := tppmessage.CmdGetServerItemListResponse{}

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
