package handler

import (
	"errors"

	"github.com/Cyvadra/tv-middleware/structures"
	"github.com/gin-gonic/gin"
)

func ReceiveTvcSignal(c *gin.Context) (f *structures.TvFormatedSignal, err error) {
	f = &structures.TvFormatedSignal{}
	err = c.ShouldBind(f)
	if err != nil {
		return
	} else if f.Ticker == "" || len(f.Action) == 0 {
		err = errors.New("ticker or action is empty")
		return
	} else {
		return
	}
}
