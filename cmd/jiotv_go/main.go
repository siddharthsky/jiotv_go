package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"

	"github.com/gin-gonic/gin"
	"github.com/rabilrbl/jiotv_go/internals/handlers"
	"github.com/rabilrbl/jiotv_go/internals/middleware"
	"github.com/rabilrbl/jiotv_go/internals/utils"
)

func main() {
	// Create a WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	// Start the Gin web server
	wg.Add(1)
	go func() {
		defer wg.Done()

		r := gin.Default()
		r.Use(middleware.CORS())

		utils.Log = utils.GetLogger()

		r.StaticFS("/static", http.FS(staticEmbed))
		tmpl := template.Must(template.ParseFS(tmplEmbed, "templates/*"))
		r.SetHTMLTemplate(tmpl)

		// Initialize the television client
		handlers.Init()

		r.GET("/", handlers.IndexHandler)
		r.GET("/login", handlers.LoginHandler)
		r.POST("/login", handlers.LoginHandler)
		r.GET("/live/:id", handlers.LiveHandler)
		r.GET("/render.m3u8", handlers.RenderHandler)
		r.GET("/render.key", handlers.RenderKeyHandler)
		r.GET("/channels", handlers.ChannelsHandler)
		r.GET("/playlist.m3u", handlers.PlaylistHandler)
		r.GET("/play/:id", handlers.PlayHandler)
		r.GET("/player/:id", handlers.PlayerHandler)
		r.GET("/clappr/:id", handlers.ClapprHandler)
		r.GET("/favicon.ico", handlers.FaviconHandler)

		// Start server on local IP address
		log.Fatal(startServer(r, "localhost:7778"))
	}()

	// Start the Fyne GUI application in the main goroutine
	a := app.New()
	w := a.NewWindow("Hello World")
	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()

	// Wait for all goroutines to finish
	wg.Wait()
}

// startServer starts the HTTP server on the specified address.
func startServer(r *gin.Engine, addr string) error {
	return r.Run(addr)
}
