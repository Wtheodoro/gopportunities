package handler

import (
	"net/http"

	"github.com/Wtheodoro/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error on listing openings")
		return
	}

	sendSuccess(ctx, "list openings", openings)
}