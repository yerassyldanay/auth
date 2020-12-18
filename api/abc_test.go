package api

import (
	"auth/model"
	"auth/utils/config"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"log"
	"net/http/httptest"
	"os"
	"testing"
)

var testServer *Server
var mq *model.MailerQueue
var testConfiguration config.Config

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	dbName := "postgres"
	dbSource := "postgres://simple:simple@localhost:8101/simple?sslmode=disable"

	fmt.Println("starting test main ...")
	connection, err := sql.Open(dbName, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
		return
	}
	defer connection.Close()

	if err := connection.Ping(); err != nil {
		log.Fatal("cannot ping the db:", err)
		return
	}

	// new test connection
	testServer = NewServer(connection)

	// start mailer
	// create one
	mq = model.InitiateNewMailerQueue()
	ctx, cancel := context.WithCancel(context.Background())

	// handle messages
	go mq.Handle(ctx)
	defer func() {
		fmt.Println("[TEST] Canceling handler...")
		cancel()
	}()

	// upload env variables
	testConfiguration, err = config.LoadConfig("../environment/")
	if err != nil {
		panic(err)
	}

	fmt.Println("[TEST] Running tests...")
	n := m.Run()
	fmt.Println("[TEST] Finished tests...")

	os.Exit(n)
}

type testErrorStruct struct {
	Error			string				`json:"error"`
}

func testParseResponseBody(t *testing.T, recorder *httptest.ResponseRecorder) {
	// parse response body
	terr := testErrorStruct{}
	err := json.NewDecoder(recorder.Body).Decode(&terr)
	require.NoError(t, err)

	if terr.Error != "" {
		fmt.Println("[ERROR]", terr)
	}
}
