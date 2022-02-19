package rpc

import (
	"github.com/ananrafs1/gomic-pg-template/utils"
	"github.com/ananrafs1/gomic/model"
	"github.com/ananrafs1/gomic/orchestrator"
	pg "github.com/ananrafs1/gomic/orchestrator/shared/plugin"
	"github.com/hashicorp/go-plugin"
)

type Scrapper struct{}

func (Scrapper) Scrap(Title string, Page, Quantity int) model.Comic {
	return model.Comic{
		ComicFlat: model.ComicFlat{
			Id:   100,
			Name: "Testing",
			Host: "Template",
		},
		Chapters: utils.CreateDummyChapter(5),
	}
}

func (Scrapper) ScrapPerChapter(Title, Id string) model.Chapter {
	dummy := utils.CreateDummyChapter(1)
	return dummy[0]
}

var pluginMap = map[string]plugin.Plugin{
	"scrapper": &pg.ScrapperPlugin{Impl: Scrapper{}},
}

func Serve() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: orchestrator.Handshake,
		Plugins:         pluginMap,
		GRPCServer:      plugin.DefaultGRPCServer,
	})
}
