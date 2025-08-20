package sessionmanager

import (
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func GetCmdAuthSteamticketResponse() tppmessage.CmdAuthSteamticketResponse {
	t := tppmessage.CmdAuthSteamticketResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMMON
	t.Msgid = tppmessage.CMD_AUTH_STEAMTICKET.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0
	t.SmartDeviceID = "QUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFB"
	t.LoginidPassword = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	t.AccountID = "11111111111111111"
	t.Currency = "NOK"

	return t
}

func HandleCmdAuthSteamticketResponse(message *message.Message) error {
	return nil
}
