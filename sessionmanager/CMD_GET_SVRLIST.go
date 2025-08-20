package sessionmanager

import (
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func GetSvrList() tppmessage.CmdGetSvrListResponse {
	t := tppmessage.CmdGetSvrListResponse{}
	t.Msgid = tppmessage.CMD_GET_SVRLIST.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMMON

	// TODO always 0, check the exe for details?

	t.Svrlist = []string{}
	t.ServerNum = 0
	return t
}

func HandleCmdGetSvrListRequest(message *message.Message) error {
	var err error

	t := GetSvrList()

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal tppmessage.CmdGetSvrListResponse: %w", err)
	}

	return nil
}
