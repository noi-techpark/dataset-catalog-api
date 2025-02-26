// © 2024 NOI Techpark <digital@noi.bz.it>
// SPDX-License-Identifier: AGPL-3.0-or-later

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthcheckHandler responds with a simple OK status.
func HealthcheckHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
