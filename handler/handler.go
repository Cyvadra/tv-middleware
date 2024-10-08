package handler

import (
	"github.com/Cyvadra/tv-middleware/structures"
	_ "github.com/Cyvadra/tv-middleware/structures"
	"github.com/gin-gonic/gin"
)

func HandleTvcSignal(c *gin.Context) (err error) {
	f := structures.TvcSignal{}
	err = c.ShouldBind(&f)
	return
}
