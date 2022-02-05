package main
import (
	"github.com/ananrafs1/gomic/orchestrator/shared"
	"github.com/ananrafs1/gomic/model"
	"github.com/hashicorp/go-plugin"
	"fmt"
	
)
type Scrapper struct{}
func CreateDummyChapter(n int) []model.Chapter {
	ret := make([]model.Chapter,0)
	for  {
		if n < 1 {
			return ret
		}
		ret = append(ret, model.Chapter{
			Id: fmt.Sprintf("chapter %d", n),
			Images: []model.Image{
				model.Image{
					Episode: 1,
					Link: map[string]string{"url1": "test", "url2":"test2"},
				},
				model.Image{
					Episode: 2,
					Link: map[string]string{"url1": "test", "url2":"test2"},
				},
				model.Image{
					Episode: 3,
					Link: map[string]string{"url1": "test", "url2":"test2"},
				},
			},
		})
		n--
	}
}

func (Scrapper) ScrapAll(Title string) (model.Comic, error) {
	return model.Comic{
		Id: 100,
		Name: "Testing",
		Host: "Template",
		Chapters: CreateDummyChapter(5),
		}, nil
}


func (Scrapper) ScrapPerChapter(Title, Id string) (model.Chapter, error) {
	dummy := CreateDummyChapter(1)
	return dummy[0], nil
}

var pluginMap = map[string]plugin.Plugin{
	"scrapper": &shared.Scrapper{Impl: Scrapper},
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins:         pluginMap,
	})
}