package rest

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"io/fs"
	"net/http"
	"os"
)

const (
	assets     = "assets"
	templeDir  = "html/templates/*.tmpl"
	templeFoo  = "html/templates/foo/*.tmpl"
	templeHome = "html/templates/home/*.tmpl"
	templePage = "html/templates/page/*.tmpl"
)

type Rest struct {
	Router *gin.Engine
	Files  fs.FS
}

func New() *Rest {
	f := os.DirFS("assets")
	return &Rest{
		Router: setupRouter(),
		Files:  f,
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func (r *Rest) Routes() {
	temple := template.Must(template.New("").ParseFS(r.Files, templeDir, templeFoo, templeHome, templePage))
	r.Router.SetHTMLTemplate(temple)
	r.Router.StaticFS("/public", http.Dir(assets))

	r.Router.GET("/", r.home)
	r.Router.GET("/ping", r.ping)
	r.Router.GET("/user/:name", r.getUser)
	r.Router.GET("favicon.ico", r.icon)
	r.Router.GET("/prologue", r.prologue)
	r.Router.GET("/test", r.test)
	r.Router.GET("/list", r.page)

	r.authorizedGroup()
}
