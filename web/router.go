package web

import (
	"io"
	"log"
	"net/http/httptest"
	"os"

	"github.com/airdb/mina-api/mocks"
	"github.com/airdb/sailor/config"
	"github.com/airdb/sailor/gin/middlewares"
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

	router := gin.New()

	router.GET("/", Status)

	APIs := router.Group("/apis")
	APIs.GET("/bbs/v0/robot/query", QueryBBS)

	// v1API := router.Group("/apis/mina/v1")
	v2API := router.Group("/apis/mina/v1")
	v2API.Use(
		middlewares.Jsonifier(),
	)

	v2API.GET("/robot/query", QueryBBS)
	v2API.GET("/category/list", ListCategory)
	v2API.GET("/lost/list", ListLost)

	// v1API := router.Group("/apis/mina/v1")
	v1API := router.Group("/v1")
	v1API.Use(
		middlewares.Jsonifier(),
	)

	// v1API.GET("/db/initdb", ListUser)
	// v1API.GET("/api/user", ListUser)
	// v1API.GET("/api/users", ListUser)
	v1API.GET("/api/category/list", ListCategory)
	v1API.GET("/api/lost/list", ListLost)
	// v1API.GET("/api/topics", ListUser)
	// v1API.GET("/api/weapp/authorizations", ListUser)
	// v1API.GET("/api/weapp/users", ListUser)
	// v1API.GET("/api/wechat", ListUser)

	// router.Use(Logger(), Recovery())

	return router
}

func APIRequest(uri, method string, param io.Reader) *httptest.ResponseRecorder {
	// Change to the root directory for handler test case.
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	defer func() {
		err = os.Chdir(wd)
		if err != nil {
			panic(err)
		}
	}()

	err = os.Chdir("../")
	if err != nil {
		panic(err)
	}

	db, err := mocks.SetUpMockDatabases()
	if err != nil {
		panic(err)
	}

	defer mocks.DestroyMockDatabases(db)

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
