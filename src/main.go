// main.go

package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

  // Set the router as the default one provided by Gin
  router = gin.Default()
  router.Static("/static", "./static")
  // Process the templates at the start so that they don't have to be loaded
  // from the disk again. This makes serving HTML pages very fast.
  router.LoadHTMLGlob("templates/*")

  // Define the route for the index page and display the index.html template
  // To start with, we'll use an inline route handler. Later on, we'll create
  // standalone functions that will be used as route handlers.
  router.GET("/", func(c *gin.Context) {

    // Call the HTML method of the Context to render a template
    c.HTML(
      // Set the HTTP status to 200 (OK)
      http.StatusOK,
      // Use the index.html template
      "index.html",
      // Pass the data that the page uses (in this case, 'title')
      gin.H{
        "title": "Home Page",
      },
    )

  })

  router.GET("/post", func(c *gin.Context) {

    // Call the HTML method of the Context to render a template
    c.HTML(
      // Set the HTTP status to 200 (OK)
      http.StatusOK,
      // Use the index.html template
      "post.html",
      // Pass the data that the page uses (in this case, 'title')
      gin.H{
        "title": "xblog建设过程",
        "subtitle": "本站点的建设故事很精彩，欢迎你的到来",
        "bg": "static/img/me.png",
        "p1": "static/img/face.JPG",
        "author": "Frank AK",
      },
    )

  })

  router.GET("/about", func(c *gin.Context) {

    // Call the HTML method of the Context to render a template
    c.HTML(
      // Set the HTTP status to 200 (OK)
      http.StatusOK,
      // Use the index.html template
      "about.html",
      // Pass the data that the page uses (in this case, 'title')
      gin.H{
        "title": "Home Page",
      },
    )

  })

  router.GET("/contact", func(c *gin.Context) {

    // Call the HTML method of the Context to render a template
    c.HTML(
      // Set the HTTP status to 200 (OK)
      http.StatusOK,
      // Use the index.html template
      "contact.html",
      // Pass the data that the page uses (in this case, 'title')
      gin.H{
        "title": "Home Page",
      },
    )

  })
  // Start serving the application
  router.Run()

}
