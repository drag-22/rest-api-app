package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"syscall"
	"time"
)

func main() {
	Books.Add("книга", "mama")
	Books.Add("книгерwefwef", "mawefwefma")
	Books.Add("кнfwefwefwfdигер", "mamwefwefwea")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		templ, err := template.ParseFiles("index.html")
		if err != nil {
			log.Println(err)
			c.Status(http.StatusBadRequest)
			return
		}
		sort.Slice(Books.list, func(i, j int) bool {
			return Books.list[i].Id < Books.list[j].Id
		})
		err = templ.Execute(c.Writer, Books.list)
		if err != nil {
			log.Println(err)
			c.Status(http.StatusBadRequest)
			return
		}

	})
	r.POST("/books/add", func(c *gin.Context) {
		err := c.Request.ParseForm()
		if err != nil {
			log.Println(err)
			c.Status(http.StatusBadRequest)
			return
		}
		Books.Add(c.PostForm("author"), c.PostForm("name"))
		c.Status(http.StatusNoContent)

	})
	r.POST("/books/remove", func(c *gin.Context) {
		err := c.Request.ParseForm()
		if err != nil {
			log.Println(err)
			c.Status(http.StatusBadRequest)
			return
		}
		Books.Rem(StringToInt(c.PostForm("id")))
		c.Status(http.StatusNoContent)

	})
	r.POST("/books/reset", func(c *gin.Context) {
		err := c.Request.ParseForm()
		if err != nil {
			log.Println(err)
			c.Status(http.StatusBadRequest)
			return
		}
		Books.Reset()
		c.Status(http.StatusNoContent)

	})

	srv := &http.Server{Addr: ":8080", Handler: r}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}

	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println(err)
		return
	}

	log.Println("exited")

}
