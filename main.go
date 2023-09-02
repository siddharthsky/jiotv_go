package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sync"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/gin-gonic/gin"
	"github.com/rabilrbl/jiotv_go/internals/handlers"
	"github.com/rabilrbl/jiotv_go/internals/utils"
)

func main() {
	// Create a WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	// Start server on local IP address
	wg.Add(1)
	go func() {
		defer wg.Done()
		//err := startServer("localhost:5001")
		if len(os.Args) > 1 {
			err := startServer(os.Args[1])
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := startServer("localhost:5001")
			if err != nil {
				log.Fatal(err)
			}
		}

	}()

	fmt.Println("")
	fmt.Println("Local Web Player at ⬇️")
	fmt.Printf("\x1b[33m%s\x1b[0m\n", "Web Player: http://localhost:7777")
	fmt.Println("")

	// Start GUI within the main goroutine
	startGUI()

	// Wait for all goroutines to finish
	wg.Wait()
}

// startServer starts the HTTP server on the specified address.
func startServer(addr string) error {
	r := gin.Default()

	if os.Getenv("GO_ENV") != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	utils.Log = utils.GetLogger()

	// Define your Gin routes and handlers here
	r.StaticFS("/static", http.FS(staticEmbed))
	tmpl := template.Must(template.ParseFS(tmplEmbed, "templates/*"))
	r.SetHTMLTemplate(tmpl)

	// Initialize the television client
	handlers.Init()

	r.GET("/", handlers.IndexHandler)
	r.GET("/login", handlers.LoginHandler)
	r.POST("/login", handlers.LoginHandler)
	r.GET("/live/:id", handlers.LiveHandler)
	r.GET("/render", handlers.RenderHandler)
	r.GET("/renderKey", handlers.RenderKeyHandler)
	r.GET("/channels", handlers.ChannelsHandler)
	r.GET("/play/:id", handlers.PlayHandler)
	r.GET("/player/:id", handlers.PlayerHandler)
	r.GET("/clappr/:id", handlers.ClapprHandler)
	r.POST("/blank", handlers.BlankHandler)
	r.GET("/favicon.ico", handlers.FaviconHandler)

	return r.Run(addr)
}

// startGUI starts the GUI window using fyne.
func startGUI() {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
