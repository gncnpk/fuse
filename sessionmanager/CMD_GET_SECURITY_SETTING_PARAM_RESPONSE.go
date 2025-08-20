package sessionmanager

import (
	"context"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func GetCmdGetSecuritySettingParamResponse(ctx context.Context, msg *message.Message, manager *SessionManager) tppmessage.CmdGetSecuritySettingParamResponse {
	t := tppmessage.CmdGetSecuritySettingParamResponse{}
	t.Result = tppmessage.RESULT_NOERR
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_SECURITY_SETTING_PARAM.String()
	t.Rqid = 0

	t.SecuritySettings = []tppmessage.SecuritySettings{}

	// TODO HARDCODED, from/to database
	// 0 = infrared
	// 1 = ???
	// 2 = antitheft
	// 3 = ??
	// 4 = cameras
	// 5 = uav
	// 6 = decoys
	// 7 = mines
	// 8 = ???
	infra := []tppmessage.SecuritySettingsTypes{{LimitNums: []int{0, 1, 1, 2, 2, 3, 4, 5}}, {LimitNums: []int{0, 1, 1, 2, 2, 3, 4, 5}}}
	mystery := []tppmessage.SecuritySettingsTypes{{LimitNums: []int{0, 1, 2, 3, 4, 4, 4, 4}}, {LimitNums: []int{0, 1, 1, 2, 4, 6, 8, 10}}}
	antitheft := []tppmessage.SecuritySettingsTypes{{LimitNums: []int{0, 1, 2, 3, 4, 5, 6, 6}}, {LimitNums: []int{0, 1, 1, 2, 4, 6, 8, 10}}}
	mystery2 := []tppmessage.SecuritySettingsTypes{{LimitNums: []int{0, 1, 2, 6, 10, 14, 18, 22}}, {LimitNums: []int{0, 1, 1, 2, 4, 6, 8, 10}}}
	cameras := []tppmessage.SecuritySettingsTypes{{LimitNums: []int{0, 1, 2, 3, 4, 5, 6, 8}}, {LimitNums: []int{0, 1, 2, 3, 4, 5, 6, 7}}}
	uav := []tppmessage.SecuritySettingsTypes{{LimitNums: []int{0, 1, 1, 1, 1, 2, 3, 4}}, {LimitNums: []int{0, 1, 1, 1, 1, 2, 2, 2}}}
	decoy := []tppmessage.SecuritySettingsTypes{{LimitNums: []int{0, 1, 2, 4, 6, 8, 10, 12}}, {LimitNums: []int{0, 1, 2, 4, 6, 7, 8, 9}}}
	mines := []tppmessage.SecuritySettingsTypes{{LimitNums: []int{0, 1, 2, 4, 6, 8, 10, 12}}, {LimitNums: []int{0, 1, 2, 4, 6, 7, 8, 9}}}
	mystery3 := []tppmessage.SecuritySettingsTypes{{LimitNums: []int{0, 2, 3, 4, 5, 6, 7, 8}}, {LimitNums: []int{0, 2, 3, 4, 5, 6, 7, 8}}}

	t.SecuritySettings = []tppmessage.SecuritySettings{
		{Types: infra},
		{Types: mystery},
		{Types: antitheft},
		{Types: mystery2},
		{Types: cameras},
		{Types: uav},
		{Types: decoy},
		{Types: mines},
		{Types: mystery3},
	}

	return t
}
