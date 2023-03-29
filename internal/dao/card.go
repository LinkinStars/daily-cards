package dao

import (
	"bytes"
	"fmt"
	"time"

	"github.com/LinkinStars/dc/internal/base/db"
	"github.com/LinkinStars/dc/internal/model"
	"github.com/LinkinStars/dc/internal/val"
	"github.com/LinkinStars/go-scaffold/logger"
	"github.com/LinkinStars/go-scaffold/mistake"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func AddCard(content string) error {
	card := &model.Card{CreatedAt: time.Now(), OriginalText: content, ParsedText: Markdown2HTML(content)}
	_, err := db.Engine.InsertOne(card)
	if err != nil {
		return mistake.InternalServer("500", err.Error())
	}
	return nil
}

func UpdateCard(req *val.UpdateCardReq) error {
	card := &model.Card{
		OriginalText: req.Content,
		ParsedText:   Markdown2HTML(req.Content),
	}

	if len(req.CreatedAt) > 0 {
		date, err := time.Parse("2006-01-02", req.CreatedAt)
		if err == nil {
			card.CreatedAt = date
		} else {
			logger.Error(fmt.Sprintf("parse date error: %s", err.Error()))
		}
	}

	_, err := db.Engine.ID(req.ID).Update(card)
	if err != nil {
		return mistake.InternalServer("500", err.Error())
	}
	return nil
}

func AddCardPv(id int, pv int) error {
	card := &model.Card{}
	_, err := db.Engine.ID(id).Incr("pv", pv).Update(card)
	if err != nil {
		return mistake.InternalServer("500", err.Error())
	}
	return nil
}

func DeleteCard(id int) error {
	_, err := db.Engine.ID(id).Delete(&model.Card{})
	if err != nil {
		return mistake.InternalServer("500", err.Error())
	}
	return nil
}

func GetCardDetail(id int) (card *model.Card, err error) {
	card = &model.Card{}
	session := db.Engine.Desc("id")
	if id > 0 {
		session.ID(id)
	}
	_, err = session.Get(card)
	if err != nil {
		return nil, mistake.InternalServer("500", err.Error())
	}
	return card, nil
}

func GetCards(page, pageSize int) (cards []*model.Card, count int64, err error) {
	cards = make([]*model.Card, 0)
	startNum := (page - 1) * pageSize
	count, err = db.Engine.Desc("created_at").Limit(pageSize, startNum).FindAndCount(&cards)
	if err != nil {
		return nil, 0, mistake.InternalServer("500", err.Error())
	}
	return cards, count, nil
}

func Markdown2HTML(source string) string {
	mdConverter := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithUnsafe(),
		),
	)
	var buf bytes.Buffer
	if err := mdConverter.Convert([]byte(source), &buf); err != nil {
		logger.Error(err)
		return source
	}
	return buf.String()
}
