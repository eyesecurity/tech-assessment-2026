package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHealth reports service liveness. The frontend calls this endpoint to
// prove end-to-end connectivity between the React app and the Go backend.
func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
