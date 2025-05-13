package stat

type GetStatsResponse struct {
	Period string `json:"period"`
	Sum    int    `json:"sum"`
}
