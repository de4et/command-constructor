package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/de4et/command-constructor/db"
	"github.com/de4et/command-constructor/db/fixtures"
	"github.com/de4et/command-constructor/types"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testdburi = "mongodb://localhost:27017"
	dbname    = "command-constructor-test"
)

type testapp struct {
	app   *fiber.App
	store db.Store
	l     logger
}

func (ta *testapp) teardown(t *testing.T) {
	if err := ta.store.User.Drop(context.Background()); err != nil {
		t.Fatal(err)
	}
	// if err := ta.store.Command.Drop(context.Background()); err != nil {
	// 	t.Fatal(err)
	// }
}

func appsetup() *testapp {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testdburi))
	if err != nil {
		log.Fatal(err)
	}

	var config = fiber.Config{
		ErrorHandler: ErrorHandler,
	}
	return &testapp{
		store: db.Store{
			User:    db.NewMongoUserStore(client, dbname),
			Command: db.NewMongoCommandStore(client, dbname),
		},
		app: fiber.New(config),
		l: logger{
			Debug: false,
		},
	}
}

type logger struct {
	Debug bool
}

func (l logger) log(s string, args ...any) {
	if l.Debug {
		fmt.Printf(s, args...)
	}
}

func (ta *testapp) testRoute(method string, path string, token string, reqParams, respParams any) (int, map[string]any) { // statusCode + raw json
	ta.l.log("Requesting %s %s\n", method, path)
	b, _ := json.Marshal(reqParams)
	ta.l.log("reqParams -- %s\n", string(b))

	req := httptest.NewRequest(
		method,
		path,
		bytes.NewReader(b),
	)
	req.Header.Add("Content-Type", "application/json") // always json
	req.Header.Add("X-Api-Token", token)               // always json

	resp, _ := ta.app.Test(req)
	m := make(map[string]any)
	json.NewDecoder(resp.Body).Decode(&m)

	ta.l.log("Status --- %s\n", resp.Status)
	if resp.StatusCode == fiber.StatusOK && respParams != nil {
		json.NewDecoder(resp.Body).Decode(respParams)
		ta.l.log("Response --- %+v\n", respParams)
	} else {
		ta.l.log("Response --- %+v\n", m)
	}
	return resp.StatusCode, m
}

// CreateUser
func TestCreateUserSuccess(t *testing.T) {
	// test /api/reg
	ta := appsetup()
	defer ta.teardown(t)
	SetupRoutes(ta.app, &ta.store)
	ta.l.Debug = false

	params := types.CreateUserParams{
		Name:     "Timur",
		Email:    "timur@foo.com",
		Password: "timurchik",
	}
	scode, resp := ta.testRoute(
		"POST",
		"/api/reg",
		"",
		params,
		nil,
	)

	if scode != http.StatusOK {
		t.Fail()
	}
	if resp["name"] != "Timur" {
		t.Fatalf("expected name '%s' but got '%s'", "Timur", resp["name"])
	}
	if resp["email"] != "timur@foo.com" {
		t.Fatalf("expected email '%s' but got '%s'", "timur@foo.com", resp["email"])
	}
	if len(resp["id"].(string)) == 0 {
		t.Fatalf("Length of id is 0")
	}
}

func TestCreateUserWrongCredentials(t *testing.T) {
	ta := appsetup()
	defer ta.teardown(t)
	SetupRoutes(ta.app, &ta.store)
	ta.l.Debug = false

	params := types.CreateUserParams{
		Name:     "T",
		Email:    "timur@foo.ckasdjfkasjdkfom-",
		Password: "timurchik",
	}
	scode, _ := ta.testRoute(
		"POST",
		"/api/reg",
		"",
		params,
		nil,
	)
	if scode == http.StatusOK {
		t.FailNow()
	}
}

func TestCreateUserAlreadyExists(t *testing.T) {
	ta := appsetup()
	defer ta.teardown(t)
	SetupRoutes(ta.app, &ta.store)
	ta.l.Debug = true

	fixtures.AddUser(&ta.store, "Timur")

	params := types.CreateUserParams{
		Name:     "Timur",
		Email:    "timur@foo.ckasdjfkasjdkfom",
		Password: "timurchik",
	}
	scode, _ := ta.testRoute(
		"POST",
		"/api/reg",
		"",
		params,
		nil,
	)

	if scode == http.StatusOK {
		t.FailNow()
	}
}

func TestAuth(t *testing.T) {
	ta := appsetup()
	defer ta.teardown(t)
	SetupRoutes(ta.app, &ta.store)
	ta.l.Debug = true
	fixtures.AddUser(&ta.store, "Timur")

	params := AuthParams{
		Name:     "Timur",
		Password: "Timur_123!",
	}

	scode, _ := ta.testRoute(
		"POST",
		"/api/auth",
		"",
		params,
		nil,
	)

	assert.Equal(t, scode, http.StatusOK)

}

// DeleteUser
