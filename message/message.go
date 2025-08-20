package message

import (
	"bytes"
	"compress/flate"
	"compress/zlib"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/coder"
	"github.com/unknown321/fuse/platform"
	"github.com/unknown321/fuse/tppmessage"
	"github.com/unknown321/fuse/util"
	"io"
	"log/slog"
	"os"
	"strconv"
	"time"
)

type Message struct {
	Compress      bool                     `json:"compress"`
	Data          string                   `json:"data"`
	OriginalSize  int                      `json:"original_size"`
	SessionCrypto bool                     `json:"session_crypto"`
	SessionKey    *string                  `json:"session_key"`
	MData         []byte                   `json:"-"` // message data
	MsgID         tppmessage.ETppMessageID `json:"-"`

	IsRequest  bool `json:"-"`
	coder      *coder.Coder
	UserID     int               `json:"-"`
	PlayerID   int               `json:"-"`
	PlatformID uint64            `json:"-"` // steamID64, psn, xboneID
	Platform   platform.Platform `json:"-"` // steam, ps4, xbox etc
}

func (m *Message) ToFile(dir string) error {
	reqName := "request"
	if !m.IsRequest {
		reqName = "response"
	}
	if dir == "" {
		dir = "."
	}

	name := fmt.Sprintf("%s/%s_%s_%s", dir, strconv.Itoa(int(time.Now().UnixMilli())), m.MsgID, reqName)
	file, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("cannot create file %s: %w", name, err)
	}
	_, err = file.Write(m.MData)
	if err != nil {
		return fmt.Errorf("cannot write to %s: %w", name, err)
	}
	_ = file.Close()

	return nil
}

func (m *Message) WithCoder(c *coder.Coder) {
	m.coder = c
}

func (m *Message) Decode(src []byte) error {
	if len(src) < 8 {
		return fmt.Errorf("not enough data")
	}

	res := m.coder.Decode(src)
	err := json.Unmarshal(res, m)
	if err != nil {
		return fmt.Errorf("cannot unmarshal message: %w", err)
	}

	//if m.Compress && !m.SessionCrypto {
	//	err = m.Decompress()
	//	if err != nil {
	//		return fmt.Errorf("cannot decompress: %w", err)
	//	}
	//} else {
	m.MData = []byte(m.Data)
	//}

	slog.Debug("decoded message", "len", len(res), "raw", fmt.Sprintf("%s", res))
	//if !m.SessionCrypto {
	//	slog.Info("message contents", "data", m.MData)
	//}
	//
	//m.Data = ""

	return nil
}

func (m *Message) Encode() ([]byte, error) {
	//m.OriginalSize = len(m.MData)
	//if m.Compress {
	//	if err := m.DoCompress(); err != nil {
	//		return nil, fmt.Errorf("cannot encode %w", err)
	//	}
	//}

	m.Data = string(m.MData)
	m.MData = nil

	marshal, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal: %w", err)
	}
	slog.Debug("marsh", "d", marshal, "len", len(marshal))

	res := m.coder.Encode(marshal)

	return res, nil
}

func (m *Message) Decompress() error {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(m.Data)))
	n, err1 := base64.StdEncoding.Decode(dst, []byte(m.Data))
	if err1 != nil {
		return fmt.Errorf("cannot decode data in message: %w", err1)
	}

	slog.Debug("decompress", "header", hex.EncodeToString(dst[0:4]), "len", len(m.Data))

	reader, err2 := zlib.NewReader(bytes.NewReader(dst[:n]))
	if err2 != nil {
		return fmt.Errorf("cannot create zlib reader: %w", err2)
	}
	defer reader.Close()

	var b []byte
	buf := bytes.NewBuffer(b)

	_, err := io.Copy(buf, reader)
	if err != nil {
		return fmt.Errorf("cannot copy compressed data back: %w", err)
	}

	m.MData = buf.Bytes()

	slog.Debug("decompressed", "len", len(m.MData), "data", string(m.MData))

	return nil
}

func (m *Message) DoCompress() error {
	var b bytes.Buffer

	slog.Debug("compressing", "len", len(m.MData), "data", string(m.MData))
	writer, err := zlib.NewWriterLevel(&b, flate.BestCompression)
	if err != nil {
		return fmt.Errorf("cannot create new zlib writer: %w", err)
	}
	defer writer.Close()
	_, err = writer.Write(m.MData)
	if err != nil {
		return fmt.Errorf("cannot compress: %w", err)
	}

	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("cannot flush compressed data: %w", err)
	}

	writer.Close()

	slog.Debug("compressed length", "len", b.Len())

	//dst := make([]byte, base64.StdEncoding.EncodedLen(b.Len()+))
	res := base64.StdEncoding.EncodeToString(b.Bytes())
	m.MData = []byte(res)

	lines := util.SplitByteString(m.MData, 76)
	m.MData = bytes.Join(lines, []byte("\r\n"))
	m.MData = append(m.MData, []byte("\r\n")...)

	return nil
}

func (m *Message) GetDataType() error {
	var iData interface{}
	err := json.Unmarshal(m.MData, &iData)
	if err != nil {
		return fmt.Errorf("cannot unmarshal message data: %w", err)
	}

	mm := iData.(map[string]interface{})
	msgid, ok := mm["msgid"]
	if !ok {
		return fmt.Errorf("msgid not found")
	}

	m.MsgID = tppmessage.CMD_INVALID

	switch msgid {
	case "CMD_GET_URLLIST":
		//m.TppMessage = tppmessage.CMDGetURLListResponse{}
		m.MsgID = tppmessage.CMD_GET_URLLIST
	case "CMD_AUTH_STEAMTICKET":
		m.MsgID = tppmessage.CMD_AUTH_STEAMTICKET
	case "CMD_REQAUTH_HTTPS":
		//if isRequest {
		//	m.TppMessage = tppmessage.CMDReqAuthHTTPSRequest{}
		//} else {
		//	m.TppMessage = tppmessage.CMDReqAuthHTTPSResponse{}
		//}
		m.MsgID = tppmessage.CMD_REQAUTH_HTTPS
	case "CMD_SEND_IPANDPORT":
		m.MsgID = tppmessage.CMD_SEND_IPANDPORT
	case "CMD_GET_PLAYERLIST":
		m.MsgID = tppmessage.CMD_GET_PLAYERLIST
	case "CMD_SET_CURRENTPLAYER":
		m.MsgID = tppmessage.CMD_SET_CURRENTPLAYER
	case "CMD_CREATE_PLAYER":
		m.MsgID = tppmessage.CMD_CREATE_PLAYER
	case "CMD_GDPR_CHECK":
		m.MsgID = tppmessage.CMD_GDPR_CHECK
	case "CMD_REQAUTH_SESSIONSVR":
		m.MsgID = tppmessage.CMD_REQAUTH_SESSIONSVR
	case "CMD_SEND_HEARTBEAT":
		m.MsgID = tppmessage.CMD_SEND_HEARTBEAT
	case "CMD_NOTICE_SNEAK_MOTHER_BASE":
		m.MsgID = tppmessage.CMD_NOTICE_SNEAK_MOTHER_BASE
	case "CMD_GET_INFORMATIONLIST2":
		m.MsgID = tppmessage.CMD_GET_INFORMATIONLIST2
	case "CMD_GET_SVRLIST":
		m.MsgID = tppmessage.CMD_GET_SVRLIST
	case "CMD_GET_SVRTIME":
		m.MsgID = tppmessage.CMD_GET_SVRTIME
	case "CMD_SYNC_MOTHER_BASE":
		m.MsgID = tppmessage.CMD_SYNC_MOTHER_BASE
	case "CMD_SYNC_RESOURCE":
		m.MsgID = tppmessage.CMD_SYNC_RESOURCE
	case "CMD_SYNC_SOLDIER_BIN":
		m.MsgID = tppmessage.CMD_SYNC_SOLDIER_BIN
	case "CMD_SYNC_SOLDIER_DIFF":
		m.MsgID = tppmessage.CMD_SYNC_SOLDIER_DIFF
	case "CMD_SYNC_EMBLEM":
		m.MsgID = tppmessage.CMD_SYNC_EMBLEM
	case "CMD_CREATE_NUCLEAR":
		m.MsgID = tppmessage.CMD_CREATE_NUCLEAR
	case "CMD_DESTRUCT_NUCLEAR":
		m.MsgID = tppmessage.CMD_DESTRUCT_NUCLEAR
	case "CMD_SEND_NUCLEAR":
		m.MsgID = tppmessage.CMD_SEND_NUCLEAR
	case "CMD_MINING_RESOURCE":
		m.MsgID = tppmessage.CMD_MINING_RESOURCE
	case "CMD_SYNC_RESET":
		m.MsgID = tppmessage.CMD_SYNC_RESET
	case "CMD_RESET_MOTHER_BASE":
		m.MsgID = tppmessage.CMD_RESET_MOTHER_BASE
	case "CMD_GET_FOB_STATUS":
		m.MsgID = tppmessage.CMD_GET_FOB_STATUS
	case "CMD_GET_ONLINE_PRISON_LIST":
		m.MsgID = tppmessage.CMD_GET_ONLINE_PRISON_LIST
	case "CMD_GET_FOB_DAMAGE":
		m.MsgID = tppmessage.CMD_GET_FOB_DAMAGE
	case "CMD_GET_FOB_TARGET_LIST":
		m.MsgID = tppmessage.CMD_GET_FOB_TARGET_LIST
	case "CMD_GET_FOB_TARGET_DETAIL":
		m.MsgID = tppmessage.CMD_GET_FOB_TARGET_DETAIL
	case "CMD_ABORT_MOTHER_BASE":
		m.MsgID = tppmessage.CMD_ABORT_MOTHER_BASE
	case "CMD_SNEAK_MOTHER_BASE":
		m.MsgID = tppmessage.CMD_SNEAK_MOTHER_BASE
	case "CMD_ACTIVE_SNEAK_MOTHER_BASE":
		m.MsgID = tppmessage.CMD_ACTIVE_SNEAK_MOTHER_BASE
	case "CMD_REQUEST_RELIEF":
		m.MsgID = tppmessage.CMD_REQUEST_RELIEF
	case "CMD_GET_FOB_REWARD_LIST":
		m.MsgID = tppmessage.CMD_GET_FOB_REWARD_LIST
	case "CMD_GET_FOB_EVENT_LIST":
		m.MsgID = tppmessage.CMD_GET_FOB_EVENT_LIST
	case "CMD_GET_FOB_EVENT_DETAIL":
		m.MsgID = tppmessage.CMD_GET_FOB_EVENT_DETAIL
	case "CMD_SEND_SNEAK_RESULT":
		m.MsgID = tppmessage.CMD_SEND_SNEAK_RESULT
	case "CMD_OPEN_WORMHOLE":
		m.MsgID = tppmessage.CMD_OPEN_WORMHOLE
	case "CMD_GET_WORMHOLE_LIST":
		m.MsgID = tppmessage.CMD_GET_WORMHOLE_LIST
	case "CMD_GET_ABOLITION_COUNT":
		m.MsgID = tppmessage.CMD_GET_ABOLITION_COUNT
	case "CMD_GET_CONTRIBUTE_PLAYER_LIST":
		m.MsgID = tppmessage.CMD_GET_CONTRIBUTE_PLAYER_LIST
	case "CMD_GET_RENTAL_LOADOUT_LIST":
		m.MsgID = tppmessage.CMD_GET_RENTAL_LOADOUT_LIST
	case "CMD_SYNC_LOADOUT":
		m.MsgID = tppmessage.CMD_SYNC_LOADOUT
	case "CMD_RENTAL_LOADOUT":
		m.MsgID = tppmessage.CMD_RENTAL_LOADOUT
	case "CMD_GET_COMBAT_DEPLOY_LIST":
		m.MsgID = tppmessage.CMD_GET_COMBAT_DEPLOY_LIST
	case "CMD_DEPLOY_MISSION":
		m.MsgID = tppmessage.CMD_DEPLOY_MISSION
	case "CMD_CANCEL_COMBAT_DEPLOY":
		m.MsgID = tppmessage.CMD_CANCEL_COMBAT_DEPLOY
	case "CMD_CANCEL_COMBAT_DEPLOY_SINGLE":
		m.MsgID = tppmessage.CMD_CANCEL_COMBAT_DEPLOY_SINGLE
	case "CMD_ELAPSE_COMBAT_DEPLOY":
		m.MsgID = tppmessage.CMD_ELAPSE_COMBAT_DEPLOY
	case "CMD_GET_COMBAT_DEPLOY_RESULT":
		m.MsgID = tppmessage.CMD_GET_COMBAT_DEPLOY_RESULT
	case "CMD_SEND_DEPLOY_INJURE":
		m.MsgID = tppmessage.CMD_SEND_DEPLOY_INJURE
	case "CMD_GET_FOB_DEPLOY_LIST":
		m.MsgID = tppmessage.CMD_GET_FOB_DEPLOY_LIST
	case "CMD_DEPLOY_FOB_ASSIST":
		m.MsgID = tppmessage.CMD_DEPLOY_FOB_ASSIST
	case "CMD_CALC_COST_FOB_DEPLOY_REPLACE":
		m.MsgID = tppmessage.CMD_CALC_COST_FOB_DEPLOY_REPLACE
	case "CMD_GET_SERVER_ITEM_LIST":
		m.MsgID = tppmessage.CMD_GET_SERVER_ITEM_LIST
	case "CMD_GET_SERVER_ITEM":
		m.MsgID = tppmessage.CMD_GET_SERVER_ITEM
	case "CMD_DEVELOP_SERVER_ITEM":
		m.MsgID = tppmessage.CMD_DEVELOP_SERVER_ITEM
	case "CMD_SPEND_SERVER_WALLET":
		m.MsgID = tppmessage.CMD_SPEND_SERVER_WALLET
	case "CMD_CHECK_SERVER_ITEM_CORRECT":
		m.MsgID = tppmessage.CMD_CHECK_SERVER_ITEM_CORRECT
	case "CMD_DESTRUCT_ONLINE_NUCLEAR":
		m.MsgID = tppmessage.CMD_DESTRUCT_ONLINE_NUCLEAR
	case "CMD_GET_CHALLENGE_TASK_REWARDS":
		m.MsgID = tppmessage.CMD_GET_CHALLENGE_TASK_REWARDS
	case "CMD_GET_CHALLENGE_TASK_TARGET_VALUES":
		m.MsgID = tppmessage.CMD_GET_CHALLENGE_TASK_TARGET_VALUES
	case "CMD_SEND_ONLINE_CHALLENGE_TASK_STATUS":
		m.MsgID = tppmessage.CMD_SEND_ONLINE_CHALLENGE_TASK_STATUS
	case "CMD_GET_PF_DETAIL_PARAMS":
		m.MsgID = tppmessage.CMD_GET_PF_DETAIL_PARAMS
	case "CMD_GET_LEAGUE_RESULT":
		m.MsgID = tppmessage.CMD_GET_LEAGUE_RESULT
	case "CMD_EXCHANGE_LEAGUE_POINT":
		m.MsgID = tppmessage.CMD_EXCHANGE_LEAGUE_POINT
	case "CMD_GET_PF_POINT_EXCHANGE_PARAMS":
		m.MsgID = tppmessage.CMD_GET_PF_POINT_EXCHANGE_PARAMS
	case "CMD_EXCHANGE_LEAGUE_POINT2":
		m.MsgID = tppmessage.CMD_EXCHANGE_LEAGUE_POINT2
	case "CMD_USE_PF_ITEM":
		m.MsgID = tppmessage.CMD_USE_PF_ITEM
	case "CMD_GET_RANKING":
		m.MsgID = tppmessage.CMD_GET_RANKING
	case "CMD_GET_NEXT_MAINTENANCE":
		m.MsgID = tppmessage.CMD_GET_NEXT_MAINTENANCE
	case "CMD_CHECK_SHORT_PFLEAGUE_ENTERABLE":
		m.MsgID = tppmessage.CMD_CHECK_SHORT_PFLEAGUE_ENTERABLE
	case "CMD_ENTER_SHORT_PFLEAGUE":
		m.MsgID = tppmessage.CMD_ENTER_SHORT_PFLEAGUE
	case "CMD_CANCEL_SHORT_PFLEAGUE":
		m.MsgID = tppmessage.CMD_CANCEL_SHORT_PFLEAGUE
	case "CMD_GET_SHORT_PFLEAGUE_RESULT":
		m.MsgID = tppmessage.CMD_GET_SHORT_PFLEAGUE_RESULT
	case "CMD_GET_PREVIOUS_SHORT_PFLEAGUE_RESULT":
		m.MsgID = tppmessage.CMD_GET_PREVIOUS_SHORT_PFLEAGUE_RESULT
	case "CMD_USE_SHORT_PF_ITEM":
		m.MsgID = tppmessage.CMD_USE_SHORT_PF_ITEM
	case "CMD_GET_MBCOIN_REMAINDER":
		m.MsgID = tppmessage.CMD_GET_MBCOIN_REMAINDER
	case "CMD_GET_OWN_FOB_LIST":
		m.MsgID = tppmessage.CMD_GET_OWN_FOB_LIST
	case "CMD_GET_PURCHASABLE_AREA_LIST":
		m.MsgID = tppmessage.CMD_GET_PURCHASABLE_AREA_LIST
	case "CMD_PURCHASE_FOB":
		m.MsgID = tppmessage.CMD_PURCHASE_FOB
	case "CMD_PURCHASE_FIRST_FOB":
		m.MsgID = tppmessage.CMD_PURCHASE_FIRST_FOB
	case "CMD_RELOCATE_FOB":
		m.MsgID = tppmessage.CMD_RELOCATE_FOB
	case "CMD_GET_PAY_ITEM_LIST":
		m.MsgID = tppmessage.CMD_GET_PAY_ITEM_LIST
	case "CMD_PURCHASE_RESOURCES_PROCESSING":
		m.MsgID = tppmessage.CMD_PURCHASE_RESOURCES_PROCESSING
	case "CMD_PURCHASE_PLATFORM_CONSTRUCTION":
		m.MsgID = tppmessage.CMD_PURCHASE_PLATFORM_CONSTRUCTION
	case "CMD_PURCHASE_SECURITY_SERVICE":
		m.MsgID = tppmessage.CMD_PURCHASE_SECURITY_SERVICE
	case "CMD_PURCHASE_SEND_TROOPS_COMPLETION":
		m.MsgID = tppmessage.CMD_PURCHASE_SEND_TROOPS_COMPLETION
	case "CMD_PURCHASE_WEPON_DEVELOPMENT_COMPLETION":
		m.MsgID = tppmessage.CMD_PURCHASE_WEPON_DEVELOPMENT_COMPLETION
	case "CMD_PURCHASE_NUCLEAR_COMPLETION":
		m.MsgID = tppmessage.CMD_PURCHASE_NUCLEAR_COMPLETION
	case "CMD_PURCHASE_ONLINE_DEPLOYMENT_COMPLETION":
		m.MsgID = tppmessage.CMD_PURCHASE_ONLINE_DEPLOYMENT_COMPLETION
	case "CMD_PURCHASE_ONLINE_DEVELOPMENT_COMPLETION":
		m.MsgID = tppmessage.CMD_PURCHASE_ONLINE_DEVELOPMENT_COMPLETION
	case "CMD_GET_SECURITY_INFO":
		m.MsgID = tppmessage.CMD_GET_SECURITY_INFO
	case "CMD_GET_SECURITY_PRODUCT_LIST":
		m.MsgID = tppmessage.CMD_GET_SECURITY_PRODUCT_LIST
	case "CMD_GET_DEVELOPMENT_PROGRESS":
		m.MsgID = tppmessage.CMD_GET_DEVELOPMENT_PROGRESS
	case "CMD_GET_ONLINE_DEVELOPMENT_PROGRESS":
		m.MsgID = tppmessage.CMD_GET_ONLINE_DEVELOPMENT_PROGRESS
	case "CMD_GET_PLATFORM_CONSTRUCTION_PROGRESS":
		m.MsgID = tppmessage.CMD_GET_PLATFORM_CONSTRUCTION_PROGRESS
	case "CMD_SEND_TROOPS":
		m.MsgID = tppmessage.CMD_SEND_TROOPS
	case "CMD_GET_TROOPS_LIST":
		m.MsgID = tppmessage.CMD_GET_TROOPS_LIST
	case "CMD_DELETE_TROOPS_LIST":
		m.MsgID = tppmessage.CMD_DELETE_TROOPS_LIST
	case "CMD_EXTEND_PLATFORM":
		m.MsgID = tppmessage.CMD_EXTEND_PLATFORM
	case "CMD_DEVELOP_WEPON":
		m.MsgID = tppmessage.CMD_DEVELOP_WEPON
	case "CMD_CALC_COST_TIME_REDUCTION":
		m.MsgID = tppmessage.CMD_CALC_COST_TIME_REDUCTION
	case "CMD_CONSUME_RESERVE":
		m.MsgID = tppmessage.CMD_CONSUME_RESERVE
	case "CMD_CHECK_CONSUME_TRANSACTION":
		m.MsgID = tppmessage.CMD_CHECK_CONSUME_TRANSACTION
	case "CMD_START_CONSUME_TRANSACTION":
		m.MsgID = tppmessage.CMD_START_CONSUME_TRANSACTION
	case "CMD_COMMIT_CONSUME_TRANSACTION":
		m.MsgID = tppmessage.CMD_COMMIT_CONSUME_TRANSACTION
	case "CMD_GET_PURCHASE_HISTORY_NUM":
		m.MsgID = tppmessage.CMD_GET_PURCHASE_HISTORY_NUM
	case "CMD_GET_PURCHASE_HISTORY":
		m.MsgID = tppmessage.CMD_GET_PURCHASE_HISTORY
	case "CMD_GET_ENTITLEMENT_ID_LIST":
		m.MsgID = tppmessage.CMD_GET_ENTITLEMENT_ID_LIST
	case "CMD_GET_SHOP_ITEM_NAME_LIST":
		m.MsgID = tppmessage.CMD_GET_SHOP_ITEM_NAME_LIST
	case "CMD_OPEN_STEAM_SHOP":
		m.MsgID = tppmessage.CMD_OPEN_STEAM_SHOP
	case "CMD_APPROVE_STEAM_SHOP":
		m.MsgID = tppmessage.CMD_APPROVE_STEAM_SHOP
	case "CMD_GET_STEAM_SHOP_ITEM_LIST":
		m.MsgID = tppmessage.CMD_GET_STEAM_SHOP_ITEM_LIST
	case "CMD_ADD_FOLLOW":
		m.MsgID = tppmessage.CMD_ADD_FOLLOW
	case "CMD_DELETE_FOLLOW":
		m.MsgID = tppmessage.CMD_DELETE_FOLLOW
	case "CMD_UPDATE_SESSION":
		m.MsgID = tppmessage.CMD_UPDATE_SESSION
	case "CMD_SEND_MISSION_RESULT":
		m.MsgID = tppmessage.CMD_SEND_MISSION_RESULT
	case "CMD_SEND_BOOT":
		m.MsgID = tppmessage.CMD_SEND_BOOT
	case "CMD_GET_FOB_NOTICE":
		m.MsgID = tppmessage.CMD_GET_FOB_NOTICE
	case "CMD_GET_DAILY_REWARD":
		m.MsgID = tppmessage.CMD_GET_DAILY_REWARD
	case "CMD_GET_CAMPAIGN_DIALOG_LIST":
		m.MsgID = tppmessage.CMD_GET_CAMPAIGN_DIALOG_LIST
	case "CMD_GET_PLAYER_PLATFORM_LIST":
		m.MsgID = tppmessage.CMD_GET_PLAYER_PLATFORM_LIST
	case "CMD_GET_FOB_PARAM":
		m.MsgID = tppmessage.CMD_GET_FOB_PARAM
	case "CMD_GET_LOGIN_PARAM":
		m.MsgID = tppmessage.CMD_GET_LOGIN_PARAM
	case "CMD_GET_SECURITY_SETTING_PARAM":
		m.MsgID = tppmessage.CMD_GET_SECURITY_SETTING_PARAM
	case "CMD_GET_RESOURCE_PARAM":
		m.MsgID = tppmessage.CMD_GET_RESOURCE_PARAM
	case "CMD_SALE_RESOURCE":
		m.MsgID = tppmessage.CMD_SALE_RESOURCE
	case "CMD_CHECK_DEFENCE_MOTHERBASE":
		m.MsgID = tppmessage.CMD_CHECK_DEFENCE_MOTHERBASE
	case "CMD_SEND_SUSPICION_PLAY_DATA":
		m.MsgID = tppmessage.CMD_SEND_SUSPICION_PLAY_DATA
	case "CMD_SET_SECURITY_CHALLENGE":
		m.MsgID = tppmessage.CMD_SET_SECURITY_CHALLENGE
	case "CMD_GET_FOB_EVENT_POINT_EXCHANGE_PARAMS":
		m.MsgID = tppmessage.CMD_GET_FOB_EVENT_POINT_EXCHANGE_PARAMS
	case "CMD_EXCHANGE_FOB_EVENT_POINT":
		m.MsgID = tppmessage.CMD_EXCHANGE_FOB_EVENT_POINT
	case "CMD_GET_SNEAK_TARGET_LIST":
		m.MsgID = tppmessage.CMD_GET_SNEAK_TARGET_LIST
	default:
		return fmt.Errorf("unknown data type")
	}

	slog.Info("message", "type", m.MsgID)

	return nil
}
