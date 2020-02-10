package web

import (
	"io"
	"log"
	"net/http/httptest"
	"os"

	"github.com/airdb/mina-api/mocks"
	"github.com/airdb/sailor/config"
	"github.com/gin-gonic/gin"
)

func Run() {
	log.Printf("Env: %s, bind: %s\n", config.GetEnv(), config.GetDefaultBindAddress())
	// err := NewRouter().Run("0.0.0.0:" + config.GetPort())
	err := NewRouter().Run(config.GetDefaultBindAddress())

	if err != nil {
		log.Println("error: ", err)
	}
}

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.GET("/", Status)

	v1API := router.Group("/apis/mina/v1")
	// For QQ robot.
	v1API.GET("/robot/query", QueryBBS)

	v1API.GET("/category/list", ListCategory)
	v1API.GET("/lost/list", ListLost)
	v1API.GET("/lost/query/:id", QueryLost)
	v1API.GET("/lost/query", QueryLost)
	v1API.GET("/user/login", Login)

	// router.Use(Logger(), Recovery())

	return router
}

func APIRequest(uri, method string, param io.Reader) *httptest.ResponseRecorder {
	if os.Getenv("TESTDB") == "sqlite" {
		db, err := mocks.SetUpMockDatabases()
		if err != nil {
			panic(err)
		}

		defer mocks.DestroyMockDatabases(db)
	}

	req := httptest.NewRequest(method, uri, param)

	if method == "GET" {
		req.Header.Set("Content-Type", "application/json")
	} else if method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	w := httptest.NewRecorder()
	NewRouter().ServeHTTP(w, req)

	return w
}
