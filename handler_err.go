package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlerErr(ctx *gin.Context) {
	respondWithError(ctx.Writer, http.StatusInternalServerError, "Internal Server Error.")
}
