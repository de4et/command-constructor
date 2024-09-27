package db

import (
	"fmt"
	"testing"

	"github.com/de4et/command-constructor/types"
)

func TestConvertStructToBson(t *testing.T) {
	s := types.CreateCommandTemplateParams{
		Name:        "scp send files",
		Description: "send file or directories of files via ssh",
		TemplateParams: []types.CommandParam{
			{
				Name:        "-r",
				Description: "for directories",
			},
		},
		ConstantParams: []types.CommandParam{
			{
				Name:        "-v",
				Description: "for debug",
			},
		},
	}
	b, err := ConvertStructToBson(s)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(b)

}
