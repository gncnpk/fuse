package tppmessage

type CmdAuthSteamticketRequest struct {
	Country         string `json:"country"`
	Lang            string `json:"lang"`
	Msgid           string `json:"msgid"`
	Region          int    `json:"region"`
	Rqid            int    `json:"rqid"`
	SteamTicket     string `json:"steam_ticket"`
	SteamTicketSize int    `json:"steam_ticket_size"`
}
