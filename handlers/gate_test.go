package handlers

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/unknown321/fuse/coder"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/sessionmanager"
	"github.com/unknown321/fuse/tppmessage"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

// this test emulates client-server communication focused on switching from COMMON encryption type to COMPOUND

func TestGateHandler(t *testing.T) {
	var err error

	gh := GateHandler{}
	sm := sessionmanager.SessionManager{}
	gh.WithManager(&sm)
	c := coder.Coder{}
	if err = c.WithKey(nil); err != nil {
		t.Fatalf("%s", err.Error())
	}
	gh.WithCoder(&c)

	dsnURI := "./testdb.dat"
	_ = os.Remove(dsnURI)
	if err = gh.DBConnect(dsnURI); err != nil {
		t.Fatalf("%s", err.Error())
	}

	ctx := context.Background()
	if err = gh.InitDB(ctx, "http://127.0.0.1/", "tppstm"); err != nil {
		t.Fatalf("initdb %s", err.Error())
	}

	reqauth := tppmessage.CMDReqAuthHTTPSRequest{
		Hash:     "t1AfMMWEWHb2lmv+P2lOyA==",
		IsTpp:    1,
		Msgid:    "CMD_REQAUTH_HTTPS",
		Platform: "Steam",
		Rqid:     0,
		Ugc:      1,
		UserName: "76561197960287930",
		Ver:      "NotImplement",
	}

	clientMsg := message.Message{
		MsgID:     tppmessage.CMD_REQAUTH_HTTPS,
		IsRequest: true,
	}
	if clientMsg.MData, err = json.Marshal(reqauth); err != nil {
		t.Fatalf("%s", err.Error())
	}
	clientMsg.OriginalSize = len(clientMsg.MData)
	clientCoder := coder.Coder{}
	_ = clientCoder.WithKey(nil)
	clientMsg.WithCoder(&clientCoder)
	var res []byte
	if res, err = clientMsg.Encode(); err != nil {
		t.Fatalf("%s", err.Error())
	}

	buf := bytes.NewReader([]byte(url.QueryEscape(string(res))))
	req, err := http.NewRequest("test", "url", buf)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	recorder := httptest.NewRecorder()
	gh.Handle(recorder, req)
	if recorder.Result().StatusCode != http.StatusOK {
		t.Fatalf("code %d", recorder.Result().StatusCode)
	}

	body := make([]byte, recorder.Result().ContentLength)
	if _, err = recorder.Result().Body.Read(body); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	reqauthResp := message.Message{}
	reqauthResp.WithCoder(&clientCoder)
	if err = reqauthResp.Decode(bytes.ReplaceAll(body, []byte("\\r\\n"), nil)); err != nil {
		t.Fatalf("%s", err.Error())
	}
	reqauthRespMsg := tppmessage.CMDReqAuthHTTPSResponse{}
	if err = json.Unmarshal([]byte(reqauthResp.Data), &reqauthRespMsg); err != nil {
		t.Fatalf("%s", err.Error())
	}
	if reqauthRespMsg.Result != tppmessage.RESULT_NOERR {
		t.Fatalf("%s", reqauthRespMsg.Result)
	}
	if reqauthRespMsg.Msgid != tppmessage.CMD_REQAUTH_HTTPS.String() {
		t.Fatalf("%s", reqauthRespMsg.Msgid)
	}

	// ===========

	key, err := base64.StdEncoding.DecodeString(reqauthRespMsg.CryptoKey)
	if err != nil {
		t.Fatalf("key %s", err.Error())
	}
	sessionCoder := coder.Coder{}
	if err = sessionCoder.WithKey(key); err != nil {
		t.Fatalf("bad key %s", err.Error())
	}

	setCP := tppmessage.CmdSetCurrentplayerRequest{Index: 0, IsReset: 0, Msgid: tppmessage.CMD_SET_CURRENTPLAYER.String(), Rqid: 0}
	setCPmsg := message.Message{
		Compress:      false,
		SessionCrypto: true,
		SessionKey:    &reqauthRespMsg.Session,
		MsgID:         tppmessage.CMD_SET_CURRENTPLAYER,
		IsRequest:     true,
	}
	setCPmsg.WithCoder(&clientCoder)
	if setCPmsg.MData, err = json.Marshal(setCP); err != nil {
		t.Fatalf("%s", err.Error())
	}
	setCPmsg.OriginalSize = len(setCPmsg.MData)
	setCPmsg.MData = sessionCoder.Encode(setCPmsg.MData)
	if res, err = setCPmsg.Encode(); err != nil {
		t.Fatalf("%s", err.Error())
	}

	buf = bytes.NewReader([]byte(url.QueryEscape(string(res))))
	req2, err := http.NewRequest("test", "url", buf)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	recorder = httptest.NewRecorder()
	gh.Handle(recorder, req2)
	if recorder.Result().StatusCode != http.StatusOK {
		t.Fatalf("code %d", recorder.Result().StatusCode)
	}

	body = make([]byte, recorder.Result().ContentLength)
	if _, err = recorder.Result().Body.Read(body); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	setCPResp := message.Message{}

	setCPResp.WithCoder(&clientCoder)
	if err = setCPResp.Decode(bytes.ReplaceAll(body, []byte("\\r\\n"), nil)); err != nil {
		t.Fatalf("%s", err.Error())
	}
	setCPResp.Data = string(sessionCoder.Decode([]byte(strings.ReplaceAll(setCPResp.Data, `\r\n`, ""))))
	setCPRespMsg := tppmessage.CmdSetCurrentplayerResponse{}
	if err = json.Unmarshal([]byte(setCPResp.Data), &setCPRespMsg); err != nil {
		t.Fatalf("%s", err.Error())
	}
	// no rows in result set, expected
	if setCPRespMsg.Result != tppmessage.RESULT_ERR {
		t.Fatalf("result %s", setCPRespMsg.Result)
	}
	if setCPRespMsg.Msgid != tppmessage.CMD_SET_CURRENTPLAYER.String() {
		t.Fatalf("msgid %s", setCPRespMsg.Msgid)
	}
	if setCPRespMsg.CryptoType != tppmessage.CRYPTO_TYPE_COMPOUND {
		t.Fatalf("crypto type %s", setCPRespMsg.CryptoType)
	}
}
