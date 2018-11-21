package link

type Request struct {
	TimeStamp string `json:"time_stamp"`
	AppKey    string `json:"app_key"`
	Sign      string `json:"sign"`
	UserId    string `json:"user_id"`
	Nick      string `json:"nick"`
	Avt       string `json:"avt"`
}

type Response struct {
	LinkKey string `json:"link_key"`
}
