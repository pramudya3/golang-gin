package main

import (
	"go-gin/domain"
	"go-gin/helper"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// type Person struct {
// 	Name     string    `form:"name"`
// 	Address  string    `form:"address"`
// 	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"+7"`
// }

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

type Booking struct {
	Name  string `json:"name" binding:"required,alpha,startswith=ok"`
	Total int    `json:"total" binding:"gte=1,lte=50"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	r := gin.Default()
	r.Use(Logger())

	r.GET("/healthz", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, helper.ResponseSuccess{Data: data, Meta: nil})
	})

	r.POST("/users", func(c *gin.Context) {
		user := &domain.User{}
		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusBadRequest, helper.ResponseFailed{Message: err.Error()})
			return
		}

		c.JSON(http.StatusCreated, helper.ResponseSuccess{Data: user, Meta: nil})
	})

	r.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			log.Printf("%+v", err)
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})

	r.GET("/logger", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// it would print: "12345"
		log.Println(example)
	})

	r.POST("/bookings", func(c *gin.Context) {
		var b Booking

		if err := c.ShouldBindJSON(&b); err != nil {
			out := helper.ValidationError(err)
			c.JSON(http.StatusBadRequest, helper.ResponseFailed{Message: out})
			return
		}

		c.JSON(http.StatusCreated, helper.ResponseSuccess{Data: b, Meta: nil})
	})

	r.Run(":8080")
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
