package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

func GetCmdGetInformationlist2Response(ctx context.Context, region string, lang string, manager *SessionManager) tppmessage.CmdGetInformationlist2Response {
	t := tppmessage.CmdGetInformationlist2Response{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_INFORMATIONLIST2.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	msg, err := manager.InformationMessageRepo.GetByRegionLang(ctx, region, lang)
	if err != nil {
		slog.Error("info message", "error", err.Error(), "region", region, "lang", lang, "msgid", t.Msgid)
		t.Result = tppmessage.RESULT_ERR
		return t
	}

	if len(msg) == 0 {
		slog.Info("no messages found for", "lang", lang, "region", region)
	}

	for _, m := range msg {
		imp := "FALSE"
		if m.Important {
			imp = "TRUE"
		}
		v := tppmessage.CmdGetInfoListEntry{
			Date:       m.Date,
			Important:  imp,
			InfoID:     m.InfoID,
			MesBody:    m.MesBody,
			MesSubject: m.MesSubject,
		}
		t.InfoList = append(t.InfoList, v)
	}
	t.InfoNum = len(t.InfoList)

	return t
}

func HandleCmdGetInformationlist2Response(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	//t := GetCmdGetInformationlist2Response()
	t := tppmessage.CmdGetInformationlist2Response{}

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
