package types

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestCreateCommandFromParams(t *testing.T) {
	fmt.Printf("%+v\n\n", UpdateCommandTemplateParams{})
	var params = CreateCommandTemplateParams{
		Name:        "send files via ssh",
		Description: "send files by pscp(Putty)",
		CommandName: "pscp",
		TemplateParams: []CommandParam{
			{
				// pscp -i "%USERPROFILE%/Documents/prin.ppk" -r ./bin root@77.232.42.104:/root/
				Name:         "",
				Description:  "Path: example -- root@127.0.0.1:/root/",
				Type:         TypeNameless,
				Value:        []string{},
				DefaultValue: "",
			},
			{
				Name:         "-i",
				Description:  "private key to send without authentication\nexample -- %USERPROFILE%/Documents/prin.ppk",
				Type:         TypeString,
				Value:        []string{},
				DefaultValue: "",
			},
			{
				Name:         "-r",
				Description:  "for sending directory\nexample -- ./bin",
				Type:         TypeString,
				Value:        []string{},
				DefaultValue: "",
			},
		},
		ConstantParams: []CommandParam{},
	}
	c, err := NewCommandTemplateFromParams("66fae8ad5f4221f1bd7399eb", params)
	spew.Dump(c)
	if err != nil {
		t.Fatal(err)
	}
	_ = c
}
