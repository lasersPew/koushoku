package server

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"path/filepath"
	"strings"
	"time"

	. "koushoku/config"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type Handler func(c *Context)
type Handlers []Handler

var server *gin.Engine
var secretHandler func()

func Init() {
	if strings.EqualFold(Config.Mode, "production") {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	server = gin.Default()
	LoadTemplates()

	server.ForwardedByClientIP = true
	server.RedirectTrailingSlash = true
	server.RemoveExtraSlash = true

	server.Use(gzip.Gzip(gzip.DefaultCompression))
	if secretHandler != nil {
		secretHandler()
	}
	server.NoRoute(Handler(noRoute).wrap())

	assets := server.Group("/")
	assets.Use(func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=300")
	})

	assets.Static("/js", filepath.Join(Config.Directories.Root, "assets/js"))
	assets.Static("/css", filepath.Join(Config.Directories.Root, "assets/css"))
	assets.Static("/fonts", filepath.Join(Config.Directories.Root, "assets/fonts"))

	assets.StaticFile("/serviceWorker.js", filepath.Join(Config.Directories.Root, "assets/js/serviceWorker.js"))
	assets.StaticFile("/favicon.ico", filepath.Join(Config.Directories.Root, "assets/favicon.ico"))
	assets.StaticFile("/robots.txt", filepath.Join(Config.Directories.Root, "assets/robots.txt"))

	if gin.Mode() == gin.DebugMode {
		group := server.Group("/debug/pprof")
		{
			group.GET("/", gin.WrapH(http.HandlerFunc(pprof.Index)))
			group.GET("/cmdline", gin.WrapH(http.HandlerFunc(pprof.Cmdline)))
			group.GET("/profile", gin.WrapH(http.HandlerFunc(pprof.Profile)))
			group.POST("/symbol", gin.WrapH(http.HandlerFunc(pprof.Symbol)))
			group.GET("/symbol", gin.WrapH(http.HandlerFunc(pprof.Symbol)))
			group.GET("/trace", gin.WrapH(http.HandlerFunc(pprof.Trace)))
			group.GET("/allocs", gin.WrapH(http.HandlerFunc(pprof.Handler("allocs").ServeHTTP)))
			group.GET("/block", gin.WrapH(http.HandlerFunc(pprof.Handler("block").ServeHTTP)))
			group.GET("/goroutine", gin.WrapH(http.HandlerFunc(pprof.Handler("goroutine").ServeHTTP)))
			group.GET("/heap", gin.WrapH(http.HandlerFunc(pprof.Handler("heap").ServeHTTP)))
			group.GET("/mutex", gin.WrapH(http.HandlerFunc(pprof.Handler("mutex").ServeHTTP)))
			group.GET("/threadcreate", gin.WrapH(http.HandlerFunc(pprof.Handler("threadcreate").ServeHTTP)))
		}
	}
}

func Start() {
	if gin.Mode() != gin.DebugMode {
		log.Println("Listening and serving HTTP on :", Config.Server.Port)
	}

	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", Config.Server.Port),
		Handler:        server,
		ReadTimeout:    time.Minute,
		WriteTimeout:   time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func noRoute(c *Context) {
	c.HTML(http.StatusNotFound, "error.html")
}

func (h Handler) wrap() gin.HandlerFunc {
	return func(c *gin.Context) {
		var context *Context
		if v, exists := c.Get("context"); exists {
			context = v.(*Context)
		} else {
			context = &Context{Context: c}
			c.Set("context", context)
		}
		h(context)
	}
}

func (h Handlers) wrap() []gin.HandlerFunc {
	ginHandlers := make([]gin.HandlerFunc, len(h))
	for i := range h {
		ginHandlers[i] = h[i].wrap()
	}
	return ginHandlers
}

func Handle(method string, relativePath string, handlers ...Handler) {
	server.Handle(method, relativePath, Handlers(handlers).wrap()...)
}

func GET(relativePath string, handlers ...Handler) {
	server.GET(relativePath, Handlers(handlers).wrap()...)
}

func POST(relativePath string, handlers ...Handler) {
	server.POST(relativePath, Handlers(handlers).wrap()...)
}

func PATCH(relativePath string, handlers ...Handler) {
	server.PATCH(relativePath, Handlers(handlers).wrap()...)
}

func PUT(relativePath string, handlers ...Handler) {
	server.PUT(relativePath, Handlers(handlers).wrap()...)
}

func DELETE(relativePath string, handlers ...Handler) {
	server.DELETE(relativePath, Handlers(handlers).wrap()...)
}
