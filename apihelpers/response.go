package apihelpers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// ResponseData struct
type ResponseData struct {
	Status int
	Meta   interface{}
	Data   interface{}
}

// RespondJSON for response to users in json
func RespondJSON(w *gin.Context, status int, payload interface{}) {
	fmt.Println("status:", status)
	var res ResponseData
	res.Status = status
	//res.Meta=utils.ResponseMessage(status)
	res.Data = payload

	w.JSON(200, res)
}
