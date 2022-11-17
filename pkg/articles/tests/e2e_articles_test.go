//go:build integration || pkg || articles || e2e || all

package tests

import (
	"net/http"
	"testing"

	"dot-app/db"
	"dot-app/db/migrations"
	"dot-app/pkg/articles/routers"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ArticlesTestSuite struct {
	suite.Suite
	DB        *gorm.DB
	Server    *gin.Engine
	Client    *httpexpect.Expect
	ArticleId string
}

// Setup for the entire suite, for specific test setup use SetupTest() below
func (suite *ArticlesTestSuite) SetupSuite() {
	db.Connect()
	suite.DB = db.DB

	gin.SetMode("test")
	suite.Server = gin.Default()
	routers.Router(suite.Server.Group(""))

	suite.Client = httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(suite.Server),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(suite.T()),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(suite.T(), true),
		},
	})

	migrations.Execute(suite.DB)
}

// Teardown for the entire suite, for specific test teardown use TearDownTest() below
func (suite *ArticlesTestSuite) TearDownSuite() {
	suite.DB.Exec("DELETE FROM articles")
}

func (suite *ArticlesTestSuite) TearDownTest() {
	suite.DB.Exec("DELETE FROM articles")
}

// Executed before each test
func (suite *ArticlesTestSuite) SetupTest() {
	// Create an article
	suite.Client.POST("/articles").
		WithJSON(map[string]interface{}{
			"title":   "Test Article",
			"content": "Test Content",
			"author":  "Test Author",
		}).
		Expect().Status(http.StatusCreated).JSON().Object().
		ValueEqual("status", true).
		ContainsKey("data").
		ContainsKey("message")

	// Get the article id
	suite.ArticleId = suite.Client.GET("/articles").Expect().Status(http.StatusOK).JSON().
		Object().Value("data").
		Object().Value("data").
		Array().First().
		Object().Value("id").String().Raw()
}

// TestIndex tests the index route
func (suite *ArticlesTestSuite) TestGetArticles() {
	suite.Client.GET("/articles").Expect().Status(http.StatusOK).JSON().Object().
		ValueEqual("status", true).
		ContainsKey("data").
		ContainsKey("message")
}

// TestShow tests the show route
func (suite *ArticlesTestSuite) TestShowArticle() {
	suite.Client.GET("/articles/{id}", suite.ArticleId).Expect().Status(http.StatusOK).JSON().
		Object().ContainsMap(map[string]interface{}{
		"status": true,
		"data": map[string]interface{}{
			"id":      suite.ArticleId,
			"title":   "Test Article",
			"content": "Test Content",
			"author":  "Test Author",
		},
	}).ContainsKey("message")
}

// TestCreate tests the create route
func (suite *ArticlesTestSuite) TestCreateArticle() {
	suite.Client.POST("/articles").
		WithJSON(map[string]interface{}{
			"title":   "Test Create Article",
			"content": "Test Create Content",
			"author":  "Test Create Author",
		}).
		Expect().Status(http.StatusCreated).JSON().
		Object().ContainsMap(map[string]interface{}{
		"status": true,
		"data": map[string]interface{}{
			"title":   "Test Create Article",
			"content": "Test Create Content",
			"author":  "Test Create Author",
		},
	}).ContainsKey("message")
}

// TestUpdate tests the update route
func (suite *ArticlesTestSuite) TestUpdateArticle() {
	suite.Client.PUT("/articles/{id}", suite.ArticleId).
		WithJSON(map[string]interface{}{
			"title":   "Test Update Article",
			"content": "Test Update Content",
			"author":  "Test Update Author",
		}).
		Expect().Status(http.StatusOK).JSON().
		Object().ContainsMap(map[string]interface{}{
		"status": true,
		"data": map[string]interface{}{
			"title":   "Test Update Article",
			"content": "Test Update Content",
			"author":  "Test Update Author",
		},
	}).ContainsKey("message")
}

// TestPatch tests the patch route
func (suite *ArticlesTestSuite) TestPatchArticle() {
	suite.Client.PATCH("/articles/{id}", suite.ArticleId).
		WithJSON(map[string]interface{}{
			"title": "Test Patch Article",
		}).
		Expect().Status(http.StatusOK).JSON().Object().
		ValueEqual("status", true).
		ContainsKey("data").
		ContainsKey("message")
}

// TestDelete tests the delete route
func (suite *ArticlesTestSuite) TestDeleteArticle() {
	suite.Client.DELETE("/articles/{id}", suite.ArticleId).Expect().Status(http.StatusOK).JSON().Object().
		ValueEqual("status", true).
		ContainsKey("data").
		ContainsKey("message")
}

func TestArticlesTestSuite(t *testing.T) {
	suite.Run(t, new(ArticlesTestSuite))
}
