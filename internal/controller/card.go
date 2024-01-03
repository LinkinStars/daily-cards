package controller

import (
	"github.com/LinkinStars/dc/internal/model"
	"net/http"
	"time"

	"github.com/LinkinStars/dc/internal/base/config"
	"github.com/LinkinStars/dc/internal/base/handler"
	"github.com/LinkinStars/dc/internal/base/middleware"
	"github.com/LinkinStars/dc/internal/dao"
	"github.com/LinkinStars/dc/internal/val"
	"github.com/LinkinStars/go-scaffold/logger"
	"github.com/LinkinStars/go-scaffold/mistake"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/now"
)

func Login(ctx *gin.Context) {
	req := &val.LoginReq{}
	err := ctx.ShouldBind(req)
	if err != nil {
		return
	}

	if req.Username != config.GlobalConf.Username || req.Password != config.GlobalConf.Password {
		handler.HandleResponse(ctx, mistake.BadRequest("", "用户名密码错误"), nil)
		return
	}

	token, err := middleware.GenerateToken(req.Username)
	if err != nil {
		logger.Error(err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, handler.ErrRespBody("Unauthorized"))
		return
	}
	handler.HandleResponse(ctx, err, map[string]string{"token": token})
}

func PostCard(ctx *gin.Context) {
	req := &val.PostCardReq{}
	err := ctx.ShouldBind(req)
	if err != nil {
		return
	}

	err = dao.AddCard(req.Content)
	handler.HandleResponse(ctx, err, nil)
}

func UpdateCard(ctx *gin.Context) {
	req := &val.UpdateCardReq{}
	err := ctx.ShouldBind(req)
	if err != nil {
		return
	}

	err = dao.UpdateCard(req)
	handler.HandleResponse(ctx, err, nil)
}

func DeleteCard(ctx *gin.Context) {
	req := &val.DeleteCardReq{}
	err := ctx.ShouldBind(req)
	if err != nil {
		return
	}

	err = dao.DeleteCard(req.ID)
	handler.HandleResponse(ctx, err, nil)
}

func GetCardsPage(ctx *gin.Context) {
	_, isLogin := ctx.Get("username")

	req := &val.GetCardPageReq{}
	err := ctx.ShouldBind(req)
	if err != nil {
		return
	}

	cards, count, err := dao.GetCards(req.Page, 20, req.SearchKeyword, req.SearchDate)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}

	resp := &val.GetCardPageResp{
		Count:    count,
		Avatar:   config.GlobalConf.Avatar,
		Nickname: config.GlobalConf.Nickname,
		Cards:    make([]*val.CardResp, 0),
	}
	for _, c := range cards {
		if !isLogin {
			c.PV = 0
		}
		resp.Cards = append(resp.Cards, &val.CardResp{
			ID:        c.ID,
			CreatedAt: c.CreatedAt.Format("2006-01-02"),
			Content:   c.ParsedText,
			PV:        c.PV,
		})
	}
	handler.HandleResponse(ctx, err, resp)
}

func GetCardDetail(ctx *gin.Context) {
	_, isLogin := ctx.Get("username")

	req := &val.GetCardDetailReq{}
	err := ctx.ShouldBind(req)
	if err != nil {
		return
	}

	var c *model.Card
	if req.Offset == 0 {
		c, err = dao.GetCardDetail(req.ID)
		if e := dao.AddCardPv(req.ID, 1); e != nil {
			logger.Error(e)
		}
	} else {
		c, err = dao.GetCardDetailByOffset(req.ID, req.Offset)
	}
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	if c == nil {
		handler.HandleResponse(ctx, mistake.BadRequest("", "没有更多了"), nil)
		return
	}

	resp := &val.GetCardDetailResp{
		ID:        c.ID,
		CreatedAt: c.CreatedAt.Format("2006-01-02"),
		Content:   c.ParsedText,
		Avatar:    config.GlobalConf.Avatar,
		Nickname:  config.GlobalConf.Nickname,
		IsLogin:   isLogin,
	}
	if isLogin {
		resp.OriginalText = c.OriginalText
		resp.PV = c.PV
	}
	handler.HandleResponse(ctx, err, resp)
}

func GetCardsStat(ctx *gin.Context) {
	req := &val.GetCardsStatReq{}
	if err := ctx.ShouldBind(req); err != nil {
		return
	}

	endTime := now.EndOfDay()
	var startTime time.Time
	if len(req.StartTime) > 0 {
		startTime, _ = time.Parse("2006-01-02", req.StartTime)
	}
	if startTime.IsZero() {
		startTime = endTime.AddDate(-1, 0, 0)
	}
	records, err := dao.GetCardsRecord(startTime, endTime)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}

	resp := &val.GetCardsStatResp{CheckedDays: make([]string, 0)}
	for _, r := range records {
		parsed, _ := time.Parse(time.RFC3339, r)
		resp.CheckedDays = append(resp.CheckedDays, parsed.Format("2006-01-02"))
	}
	handler.HandleResponse(ctx, err, resp)
}
