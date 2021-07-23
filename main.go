package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
	"github.com/unrolled/secure"

	// ginSwagger "github.com/swaggo/gin-swagger"
	"web-service-gin/api/album"
	"web-service-gin/api/user"
)

type authHeader struct {
	IDToken string `header:"Authorization"`
}

func main() {
	secureFunc := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			secureMiddleware := secure.New(secure.Options{
				SSLRedirect: false,
				SSLHost:     "localhost:8080",
			})
			err := secureMiddleware.Process(c.Writer, c.Request)

			// If there was an error, do not continue.
			if err != nil {
				return
			}

			c.Next()
		}
	}()
	validateToken := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			h := authHeader{}
			fmt.Println("validateToken --", c.Request.Header)
			if err := c.ShouldBindHeader(&h); err != nil {
				c.JSON(http.StatusUnauthorized, "Please provide valid details")
				return
			}
			idTokenHeader := strings.Split(h.IDToken, "Bearer ")
			fmt.Println("idTokenHeader --", idTokenHeader)
			// validate ID token here
			// user, err := user.ValidateToken(idTokenHeader[1])
			c.Next()
		}
	}()

	router := gin.Default()
	router.Use(secureFunc)
	router.Use(validateToken)

	v1 := router.Group("/api")
	v1.Use(SwaggerMiddleware(router))
	router.POST("authenticate", user.Authenticate)
	router.GET("/albums", album.GetAlbums)
	router.GET("/albums/:id", album.GetAlbumById)
	router.POST("/albums", album.PostAlbums)

	router.Run()
}

func SwaggerMiddleware(server *gin.Engine) gin.HandlerFunc {
	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yml"}
	sh := middleware.SwaggerUI(opts, nil)

	server.GET("/docs", gin.WrapH(sh))
	server.GET("/swagger.yml", gin.WrapH(http.FileServer(http.Dir("./docs"))))

	return func(c *gin.Context) {
		c.Next()
	}
}
