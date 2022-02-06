package main

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type (
	// Application Interface
	Application interface {
		ListenAndServe()
	}

	application struct {
		echo *echo.Echo
		DB   *sqlx.DB
	}

	TemplateRenderer struct {
		templates *template.Template
	}
)

var funcMap = template.FuncMap{
	"date": func(t time.Time) string {
		return fmt.Sprintf("%02d-%02d-%02d %02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())
	},
	"unescape": func(s string) template.HTML {
		return template.HTML(s)
	},
}

// NewApplication :nodoc:
func NewApplication() Application {

	app := &application{}

	return app
}

func (a *application) connectAppDB() {
	dsn := viper.GetString("application.db.url")
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Info("Connected to DB")

	a.DB = db
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	template := template.New("templates")

	template.Funcs(funcMap)

	filepath.Walk("templates", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Walker: ", err)
			return nil
		}

		if strings.HasSuffix(path, ".html") {
			template.ParseFiles(path)
		}

		return nil
	})

	t.templates = template

	return t.templates.ExecuteTemplate(w, name, data)
}

// StartServe :nodoc:
func (a *application) ListenAndServe() {
	e := echo.New()

	a.connectAppDB()

	defer a.DB.Close()

	e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	//e.Use(middleware.Secure())
	//e.Use(middleware.CORS())
	e.Use(middleware.Static(""))

	e.HideBanner = true

	// a.ConnectDB()
	e.Renderer = &TemplateRenderer{}

	a.echo = e

	a.registerRoutes()

	port := fmt.Sprintf(":%d", viper.GetInt("application.port"))
	a.echo.Logger.Fatal(e.Start(port))
}
