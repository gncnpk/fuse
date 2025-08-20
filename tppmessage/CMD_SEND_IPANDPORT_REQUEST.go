package tppmessage

type CmdSendIpandportRequest struct {
	ExIp                string `json:"ex_ip"`
	ExPort              int    `json:"ex_port"`
	InIp                string `json:"in_ip"`
	InPort              int    `json:"in_port"`
	Msgid               string `json:"msgid"`
	Nat                 string `json:"nat"`
	Rqid                int    `json:"rqid"`
	SecureDeviceAddress string `json:"secure_device_address"`
	Xnaddr              string `json:"xnaddr"`
}
