package sessionmanager

import (
	"context"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdGetResourceParamResponse(ctx context.Context, manager *SessionManager) (tppmessage.CmdGetResourceParamResponse, error) {
	t := tppmessage.CmdGetResourceParamResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_RESOURCE_PARAM.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	// TODO from database?
	t.NuclearDevelopCosts = []int{
		750000,
		50000,
		75000,
	}

	t.Result = tppmessage.RESULT_NOERR

	return t, nil
}

func HandleCmdGetResourceParamResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	//var err error
	//t := GetCmdGetResourceParamResponse()

	//message.MData, err = json.Marshal(t)
	//if err != nil {
	//	return fmt.Errorf("cannot marshal: %w", err)
	//}

	return nil
}
