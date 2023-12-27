package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Rest) home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
	})
}

func (r *Rest) pong(c *gin.Context) {
	c.HTML(http.StatusOK, "bar.tmpl", gin.H{
		"title": "Foo website",
	})
}

//func (r *Rest) icon(c *gin.Context) {
//	file, _ := r.Files.Open("images/example.png")
//	c.Data(
//		http.StatusOK,
//		"image/x-icon",
//		file,
//	)
//
//}
