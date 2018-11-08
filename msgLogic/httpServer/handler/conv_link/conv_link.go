package conv_link

type JoinConvRequest struct {
	ConvId int    `json:"ConvId"`
	UserToken      string `json:"userToken"`
}
