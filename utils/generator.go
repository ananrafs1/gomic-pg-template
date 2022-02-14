package utils
import (
	"fmt"
	"github.com/ananrafs1/gomic/model"
)

func CreateDummyChapter(n int) []model.Chapter {
	ret := make([]model.Chapter, 0)
	Images := []model.ImageProvider{
		model.ImageProvider{
			Episode: 1,
			Provider: "url1",
			Link:    "test",
		},
		model.ImageProvider{
			Episode: 1,
			Provider: "url2",
			Link:    "test2",
		},
		model.ImageProvider{
			Episode: 2,
			Provider: "url1",
			Link:    "test3",
		},
		model.ImageProvider{
			Episode: 2,
			Provider: "url2",
			Link:    "test4",
		},
	}

	for {
		if n < 1 {
			return ret
		}
		ret = append(ret, model.Chapter{
			model.ChapterFlat{
				Id: fmt.Sprintf("chapter %d", n),
			},
			Images,
		})
		n--
	}
}