package http

import (
	"errors"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

type EventStreamRequest struct {
	Message string `form:"message" json:"message" binding:"required,max=100"`
}

func HandleEventStreamPost(c *gin.Context, ch chan string) {
	fmt.Println("receiving post event: ", ch)
	var request EventStreamRequest
	if err := c.ShouldBind(&request); err != nil {
		errorMessage := fmt.Sprintf("request validation error: %s", err.Error())
		BadRequestResponse(c, errors.New(errorMessage))

		return
	}

	CreatedResponse(c, &request.Message)

	return
}

func HandleEventStreamGet(c *gin.Context, ch chan string) {
	fmt.Println("receiving get event: ", ch)
	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-ch; ok {
			c.SSEvent("message", msg)
			return true
		}
		fmt.Println("inside stream: ")
		return false
	})

	return
}
