package db

import (
	"fmt"
	"testing"

	"github.com/de4et/command-constructor/types"
	"go.mongodb.org/mongo-driver/bson"
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
	correct := bson.M{
		"name":        "scp send files",
		"description": "send file or directories of files via ssh",
		"templateparams": []bson.M{
			{
				"name":        "-r",
				"description": "for directories",
				"value":       "",
			},
		},
		"constantparams": []bson.M{
			{
				"name":        "-v",
				"description": "for debug",
				"value":       "",
			},
		},
	}

	// this is bs
	if fmt.Sprint(b) != fmt.Sprint(correct) {
		t.Errorf("expected to be equal")
	}

}
