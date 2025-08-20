package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
	"strings"
)

func GetCmdGetUrllistResponse(ctx context.Context, repo *tppmessage.URLListEntryRepo) (tppmessage.CMDGetURLListResponse, error) {
	t := tppmessage.CMDGetURLListResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMMON
	t.Msgid = tppmessage.CMD_GET_URLLIST.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	all, err := repo.GetAll(ctx)
	if err != nil {
		return t, fmt.Errorf("cannot get all: %w", err)
	}

	t.UrlList = all
	t.UrlNum = len(all)

	return t, nil
}

func HandleCmdGetUrllistResponse(ctx context.Context, message *message.Message, override bool, repo *tppmessage.URLListEntryRepo) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	t, err := GetCmdGetUrllistResponse(ctx, repo)

	// TODO replace with real server ip?
	for n := range t.UrlList {
		t.UrlList[n].Url = strings.ReplaceAll(t.UrlList[n].Url, "https", "http")
		t.UrlList[n].Url = strings.ReplaceAll(t.UrlList[n].Url, "mgstpp-game.konamionline.com", "127.0.0.1")
	}

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
