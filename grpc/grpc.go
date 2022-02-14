package grpc

import (
	"github.com/hashicorp/go-plugin"
	pg "github.com/ananrafs1/gomic/orchestrator/shared/plugin"
	"github.com/ananrafs1/gomic/model"
	"github.com/ananrafs1/gomic-pg-template/utils"
	"github.com/ananrafs1/gomic/orchestrator"
)
type Scrapper struct{}

func (Scrapper) Scrap(Title string, Page, Quantity int) model.Comic {
	return model.Comic{
		model.ComicFlat{
			Id:   100,
			Name: "Testing",
			Host: "Template",
		},
		utils.CreateDummyChapter(5),
	}
}

func (Scrapper) ScrapPerChapter(Title, Id string) model.Chapter {
	dummy := utils.CreateDummyChapter(1)
	return dummy[0]
}

func Serve() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: orchestrator.Handshake,
		Plugins: map[string]plugin.Plugin{
			"grpcscrapper": &pg.GRPCPlugin{Impl: &Scrapper{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}