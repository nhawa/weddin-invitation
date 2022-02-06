package controller

import (
	//"html/template"
	"net/http"
	_ "strconv"

	//"time"

	"github.com/labstack/echo/v4"
)

type (
	ArticleController interface {
		GetArticleList(echo echo.Context) error
		GetArticleDetail(echo echo.Context) error
	}

	articleController struct {
	}

	Welcome struct {
		Name string
		Time string
	}
)

func NewArticleController() ArticleController {

	articleController := articleController{}

	return articleController
}

func (articleController articleController) GetArticleList(c echo.Context) error {
	return c.Render(http.StatusOK, "weddingIndex", map[string]interface{}{
		"banner1": "http://localhost:9010/static/1.jpeg",
		"banner2": "http://localhost:9010/static/2.jpeg",
		"banner3": "http://localhost:9010/static/3.jpeg",
		"groom":   "http://localhost:9010/static/a1.jpeg",
		"bride":   "http://localhost:9010/static/a2.jpeg",
		"galery1": "http://localhost:9010/static/e1.jpeg",
		"galery2": "http://localhost:9010/static/e2.jpeg",
		"galery3": "http://localhost:9010/static/g1.jpeg",
		"galery4": "http://localhost:9010/static/g22.jpeg",
		"galery5": "http://localhost:9010/static/g3.jpeg",
		"galery6": "http://localhost:9010/static/g4.jpeg",
		"galery7": "http://localhost:9010/static/g5.jpeg",
		"galery8": "http://localhost:9010/static/g6.jpeg",
		"galery9": "http://localhost:9010/static/g7.jpeg",
	})
}

func (articleController articleController) GetArticleDetail(echo echo.Context) error {
	//var res presenter.BasePresenterModel
	//
	//articleId, err := strconv.Atoi(echo.Param("articleId"))
	//if err != nil {
	//    res = presenter.CreateFailResponse(presenter.InvalidArticleDetailCode, presenter.InvalidArticleDetailReason)
	//
	//    return echo.JSON(http.StatusBadRequest, res)
	//}
	//
	//article, err := articleController.articleService.GetArticleDetail(articleId)
	//if err != nil {
	//    res = presenter.CreateFailResponse(presenter.InvalidArticleDetailCode, presenter.InvalidArticleDetailReason)
	//
	//    return echo.JSON(http.StatusBadRequest, res)
	//}
	//responseData := presenter.CreateArticleDetailSuccessResponse(article)
	//
	//return echo.JSON(http.StatusOK, responseData)
	var err error
	return err
}
