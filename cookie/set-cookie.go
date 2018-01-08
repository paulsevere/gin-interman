package cookie

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func determineCookieDomain(origin string) string {
	host := origin
	if strings.Contains(host, "localhost") {
		host = "localhost"
	} else if strings.Contains(host, "libertymutual") {
		host = ".libertymutual.com"
	}
	return host
}

func setSuiteCookie(suiteName string, c *gin.Context) {
	host := determineCookieDomain(c.GetHeader("origin"))
	c.SetCookie("SUITE", suiteName, 1000000, "/", host, false, false)
	c.String(200, "SET THE COOKIE:  %s", suiteName)

}

func deleteCookie(c *gin.Context) {
	host := determineCookieDomain(c.GetHeader("origin"))
	c.SetCookie("SUITE", "", -1, "/", host, false, false)
	c.String(200, "REMOVED THE COOKIE!!!")
}

// SetSuiteCookie sets or unsets suite cookie
func SetSuiteCookie(c *gin.Context) {
	suiteName := c.Query("suite")
	if suiteName != "" {
		setSuiteCookie(suiteName, c)
	} else {
		deleteCookie(c)
	}
}
