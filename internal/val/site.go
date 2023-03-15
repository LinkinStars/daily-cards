package val

type GetSiteInfoResp struct {
	SiteName   string `json:"site_name"`
	ShowQrcode bool   `json:"show_qrcode"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
}
