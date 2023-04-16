package val

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PostCardReq struct {
	Content string `json:"content"`
}

type DeleteCardReq struct {
	ID int `json:"id"`
}

type UpdateCardReq struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	Content   string `json:"content"`
}

type GetCardDetailReq struct {
	ID int `form:"id"`
}

type GetCardPageReq struct {
	Page int `form:"page"`
}

type GetCardPageResp struct {
	Nickname string      `json:"nickname"`
	Avatar   string      `json:"avatar"`
	Cards    []*CardResp `json:"cards"`
	Count    int64       `json:"count"`
}

type CardResp struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	Content   string `json:"content"`
	PV        int    `json:"pv"`
}

type GetCardDetailResp struct {
	ID           int    `json:"id"`
	CreatedAt    string `json:"created_at"`
	Content      string `json:"content"`
	Nickname     string `json:"nickname"`
	Avatar       string `json:"avatar"`
	IsLogin      bool   `json:"is_login"`
	OriginalText string `json:"original_text"`
	PV           int    `json:"pv"`
}

type GetCardsStatResp struct {
	Day     int  `json:"day"`
	Checked bool `json:"checked"`
}
