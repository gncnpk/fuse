package sessionmanager

import (
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"time"
)

func GetSvrTime() tppmessage.CmdGetSvrTimeResponse {
	t := tppmessage.CmdGetSvrTimeResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMMON
	t.Msgid = tppmessage.CMD_GET_SVRTIME.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	t.Date = int(time.Now().Unix())
	return t
}

func HandleCmdGetSvrTimeRequest(message *message.Message) error {
	t := GetSvrTime()

	var err error
	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal tppmessage.CmdGetSvrTimeResponse: %w", err)
	}

	return nil
}
