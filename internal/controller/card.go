package controller

import (
	"net/http"

	"github.com/LinkinStars/dc/internal/base/config"
	"github.com/LinkinStars/dc/internal/base/handler"
	"github.com/LinkinStars/dc/internal/base/middleware"
	"github.com/LinkinStars/dc/internal/dao"
	"github.com/LinkinStars/dc/internal/val"
	"github.com/LinkinStars/go-scaffold/logger"
	"github.com/LinkinStars/go-scaffold/mistake"
	"github.com/gin-gonic/gin"
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
	req := &val.GetCardPageReq{}
	err := ctx.ShouldBind(req)
	if err != nil {
		return
	}

	cards, count, err := dao.GetCards(req.Page, 20)
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
		resp.Cards = append(resp.Cards, &val.CardResp{
			ID:        c.ID,
			CreatedAt: c.CreatedAt.Format("2006-01-02"),
			Content:   c.ParsedText,
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

	c, err := dao.GetCardDetail(req.ID)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}

	if e := dao.AddCardPv(req.ID, 1); e != nil {
		logger.Error(e)
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
