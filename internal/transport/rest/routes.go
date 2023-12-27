package rest

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"io/fs"
	"net/http"
	"os"
)

type Rest struct {
	Router *gin.Engine
	Files  fs.FS
}

func New() *Rest {
	f := os.DirFS("assets")
	return &Rest{
		Router: gin.New(),
		Files:  f,
	}
}

const (
	templeDir = "html/templates/*.tmpl"
	templeFoo = "html/templates/foo/*.tmpl"
)

func (r *Rest) Routes() {
	temple := template.Must(template.New("").ParseFS(r.Files, templeDir, templeFoo))
	r.Router.SetHTMLTemplate(temple)
	r.Router.StaticFS("/public", http.FS(r.Files))

	r.Router.GET("/", r.home)
	r.Router.GET("/ping", r.pong)
	//r.Router.GET("favicon.ico", r.icon)
}
