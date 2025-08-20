package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdGetPurchasableAreaListResponse(ctx context.Context, manager *SessionManager, msg *message.Message) tppmessage.CmdGetPurchasableAreaListResponse {
	t := tppmessage.CmdGetPurchasableAreaListResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_PURCHASABLE_AREA_LIST.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	// TODO from database, location index depends on previously bought fobs, HARDCODED

	t.Area = []tppmessage.CmdGetPurchasableAreaListResponseArea{
		// location index - where is it on the map in area id
		{AreaId: 10, LocationIndex: 3, MiningResource: []int{10, 10, 0, 0}, Price: 1200},  // mid atlantic
		{AreaId: 20, LocationIndex: 3, MiningResource: []int{10, 10, 0, 0}, Price: 1100},  // hawaii
		{AreaId: 32, LocationIndex: 24, MiningResource: []int{10, 10, 0, 0}, Price: 1100}, // south atlantic
		{AreaId: 40, LocationIndex: 1, MiningResource: []int{10, 10, 0, 0}, Price: 1200},  // indian
		{AreaId: 53, LocationIndex: 16, MiningResource: []int{10, 10, 0, 0}, Price: 1000}, // north pacific
		{AreaId: 62, LocationIndex: 6, MiningResource: []int{10, 10, 0, 0}, Price: 1000},  // south pacific
		{AreaId: 71, LocationIndex: 12, MiningResource: []int{10, 10, 0, 0}, Price: 1000}, // north atlantic ocean
	}

	return t
}

func HandleCmdGetPurchasableAreaListResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	//t := GetCmdGetPurchasableAreaListResponse()
	t := tppmessage.CmdGetPurchasableAreaListResponse{}

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
