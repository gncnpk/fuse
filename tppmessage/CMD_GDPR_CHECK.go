package tppmessage

type CmdGDPRCheckRequest struct {
	Authcode        string `json:"authcode"`
	Hash            string `json:"hash"`
	IssuerId        int    `json:"issuer_id"`
	Lang            int    `json:"lang"`
	Msgid           string `json:"msgid"`
	Platform        string `json:"platform"`
	Rqid            int    `json:"rqid"`
	SteamTicket     string `json:"steam_ticket"`
	SteamTicketSize int    `json:"steam_ticket_size"`
	UserName        string `json:"user_name"`
}

type GDPRAddendumMessage struct {
	Confirm string `json:"confirm"`
	Lang    int    `json:"lang"`
	Text    string `json:"text"`
	Title   string `json:"title"`
}

type GDPRAddendumList struct {
	Index       int                   `json:"index"`
	MessageList []GDPRAddendumMessage `json:"message_list"`
}

type CmdGDPRCheckResponse struct {
	AddendumList     []GDPRAddendumList `json:"addendum_list"`
	CcpaStateIndex   int                `json:"ccpa_state_index"`
	CryptoType       string             `json:"crypto_type"`
	Flowid           interface{}        `json:"flowid"`
	GdprCountryIndex int                `json:"gdpr_country_index"`
	GeoIpCountry     string             `json:"geo_ip_country"`
	GeoIpState       string             `json:"geo_ip_state"`
	Msgid            string             `json:"msgid"`
	Result           string             `json:"result"`
	Rqid             int                `json:"rqid"`
	Xuid             interface{}        `json:"xuid"`
}
