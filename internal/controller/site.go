package controller

import (
	"github.com/LinkinStars/dc/internal/base/config"
	"github.com/LinkinStars/dc/internal/base/handler"
	"github.com/LinkinStars/dc/internal/val"
	"github.com/gin-gonic/gin"
)

func GetSiteInfo(ctx *gin.Context) {
	resp := &val.GetSiteInfoResp{
		SiteName:   config.GlobalConf.SiteName,
		ShowQrcode: config.GlobalConf.ShowQrcode,
		Avatar:     config.GlobalConf.Avatar,
		Nickname:   config.GlobalConf.Nickname,
	}
	handler.HandleResponse(ctx, nil, resp)
}
