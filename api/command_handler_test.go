package api

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/de4et/command-constructor/db/fixtures"
	. "github.com/de4et/command-constructor/types"
	"github.com/stretchr/testify/assert"
)

func TestCommandUpdate(t *testing.T) {
	// test POST /api/v1/command/:id
	ta := appsetup()
	defer ta.teardown(t)
	SetupRoutes(ta.app, &ta.store)
	ta.l.Debug = false

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
	user := fixtures.AddUser(&ta.store, "Timur")
	token := makeTokenFromUser(user)

	command := fixtures.AddCommand(&ta.store, user, params)

	cs, err := ta.store.Command.SearchCommands(context.TODO(), command.UserID, "send")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(cs)

	params.Name = "blalblalblab"
	scode, resp := ta.testRoute(
		"PUT",
		"/api/v1/command/"+command.ID,
		token,
		params,
		nil,
	)
	_ = resp
	cs, err = ta.store.Command.SearchCommands(context.TODO(), command.UserID, "send")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(cs)

	assert.Equal(t, scode, http.StatusOK)
	// assert.Equal(
}
