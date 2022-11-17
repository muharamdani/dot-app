//go:build integration || pkg || comments || e2e || all

package tests

import (
	"dot-app/pkg"
	"net/http"
	"testing"

	"dot-app/db"
	"dot-app/db/migrations"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type CommentsTestSuite struct {
	suite.Suite
	DB        *gorm.DB
	Server    *gin.Engine
	Client    *httpexpect.Expect
	ArticleId string
	CommentId string
}

// Setup for the entire suite, for specific test setup use SetupTest() below
func (suite *CommentsTestSuite) SetupSuite() {
	db.Connect()
	suite.DB = db.DB

	gin.SetMode("test")
	suite.Server = gin.Default()
	pkg.Router(suite.Server.Group("/api"))

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
func (suite *CommentsTestSuite) TearDownSuite() {
	suite.DB.Exec("DELETE FROM comments")
	suite.DB.Exec("DELETE FROM articles")
}

func (suite *CommentsTestSuite) TearDownTest() {
	suite.DB.Exec("DELETE FROM comments")
	suite.DB.Exec("DELETE FROM articles")
}

// Executed before each test
func (suite *CommentsTestSuite) SetupTest() {
	// Create an article
	suite.Client.POST("/api/articles").
		WithJSON(map[string]interface{}{
			"title":   "Test Init Article",
			"content": "Test Init Content",
			"author":  "Test Init Author",
		}).
		Expect().Status(http.StatusCreated).JSON().Object().
		ValueEqual("status", true).
		ContainsKey("data").
		ContainsKey("message")

	// Get the article id
	suite.ArticleId = suite.Client.GET("/api/articles").Expect().Status(http.StatusOK).JSON().
		Object().Value("data").
		Object().Value("data").
		Array().First().
		Object().Value("id").String().Raw()

	// Create a comment
	suite.Client.POST("/api/comments").
		WithJSON(map[string]interface{}{
			"article_id": suite.ArticleId,
			"content":    "1",
		}).
		Expect().Status(http.StatusCreated).JSON().Object().
		ValueEqual("status", true).
		ContainsKey("data").
		ContainsKey("message")

	// Get the comment id
	suite.CommentId = suite.Client.POST("/api/comments").
		WithJSON(map[string]interface{}{
			"article_id": suite.ArticleId,
			"content":    "1",
		}).
		Expect().Status(http.StatusCreated).JSON().
		Object().Value("data").
		Object().Value("id").String().Raw()
}

// TestCreate tests the create comment
func (suite *CommentsTestSuite) TestCreateComment() {
	suite.Client.POST("/api/comments").
		WithJSON(map[string]interface{}{
			"article_id": suite.ArticleId,
			"content":    "Test Init Comment 2",
		}).
		Expect().Status(http.StatusCreated).JSON().Object().
		ValueEqual("status", true).
		ContainsKey("data").
		ContainsKey("message")

}

// TestCreateChildComment tests the create child comment
func (suite *CommentsTestSuite) TestCreateChildComment() {
	suite.Client.POST("/api/comments").
		WithJSON(map[string]interface{}{
			"parent_id": suite.CommentId,
			"content":   "Test Init Comment 2",
		}).
		Expect().Status(http.StatusCreated).JSON().Object().
		ContainsMap(map[string]interface{}{
			"status": true,
			"data": map[string]interface{}{
				"parent_id": suite.CommentId,
				"content":   "Test Init Comment 2",
			},
		}).ContainsKey("message")
}

// TestPatch tests the patch comment
func (suite *CommentsTestSuite) TestPatchComment() {
	suite.Client.PATCH("/api/comments/{id}", suite.CommentId).
		WithJSON(map[string]interface{}{
			"content": "2",
		}).
		Expect().Status(http.StatusOK).JSON().Object().
		ValueEqual("status", true).
		ContainsKey("data").
		ContainsKey("message")
}

// TestDelete tests the delete comment
func (suite *CommentsTestSuite) TestDeleteComment() {
	suite.Client.DELETE("/api/comments/{id}", suite.CommentId).
		Expect().Status(http.StatusOK).JSON().Object().
		ValueEqual("status", true).
		ContainsKey("data").
		ContainsKey("message")
}

func TestCommentsTestSuite(t *testing.T) {
	suite.Run(t, new(CommentsTestSuite))
}
