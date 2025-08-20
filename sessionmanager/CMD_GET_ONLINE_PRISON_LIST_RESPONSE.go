package sessionmanager

import (
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

/*
{
	"crypto_type": "COMPOUND",
	"flowid": null,
	"msgid": "CMD_GET_ONLINE_PRISON_LIST",
	"prison_soldier_param": [],
	"rescue_list": [],
	"rescue_num": 0,
	"result": "NOERR",
	"rqid": 0,
	"soldier_num": 0,
	"total_num": 53,     <-----------
	"xuid": null
}
*/

func GetCmdGetOnlinePrisonListResponse() tppmessage.CmdGetOnlinePrisonListResponse {
	t := tppmessage.CmdGetOnlinePrisonListResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_ONLINE_PRISON_LIST.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	// TODO from db
	t.PrisonSoldierParam = []any{}
	t.SoldierNum = len(t.PrisonSoldierParam)

	t.RescueList = []any{}
	t.RescueNum = len(t.RescueList)

	t.TotalNum = 0

	return t
}

func HandleCmdGetOnlinePrisonListResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	t := GetCmdGetOnlinePrisonListResponse()

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
