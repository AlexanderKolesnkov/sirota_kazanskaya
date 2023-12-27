package rest

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var db = make(map[string]string)

func (r *Rest) home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.page.tmpl", gin.H{
		"title": "Main website",
	})
}

func (r *Rest) ping(c *gin.Context) {
	c.HTML(http.StatusOK, "bar.tmpl", gin.H{
		"title": "Foo website",
	})
}

func (r *Rest) test(c *gin.Context) {
	c.HTML(http.StatusOK, "test", gin.H{})
}

func (r *Rest) page(c *gin.Context) {
	c.HTML(http.StatusOK, "list.page.tmpl", gin.H{})
}

func (r *Rest) prologue(c *gin.Context) {
	c.File("assets/scenario/drafts/prologue.txt")
}

func (r *Rest) icon(c *gin.Context) {
	file, err := r.Files.Open("favicon.ico")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	data := make([]byte, fileInfo.Size())

	_, err = file.Read(data)
	if err != nil {
		log.Fatalln(err)
	}
	c.Data(
		http.StatusOK,
		"image/x-icon",
		data,
	)

}

func (r *Rest) getUser(c *gin.Context) {
	user := c.Params.ByName("name")
	value, ok := db[user]
	if ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	}
}

func (r *Rest) authorizedGroup() {
	authorized := r.Router.Group("/group", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})
}
