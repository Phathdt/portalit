package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"portalit/shared/sctx/core"
)

func WriteErrorResponse(c *gin.Context, err error) {
	if errSt, ok := err.(core.StatusCodeCarrier); ok {
		c.JSON(errSt.StatusCode(), errSt)
		return
	}

	c.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithError(err.Error()))
}
