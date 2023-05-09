package bootstrap

import (
	"github.com/gin-gonic/gin"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
	rkgin "github.com/rookie-ninja/rk-gin/v2/boot"
)

func newGinEntry(boot []byte) *rkgin.GinEntry {
	rkentry.BootstrapBuiltInEntryFromYAML(boot)
	rkentry.GlobalAppCtx.GetAppInfoEntry().AppName = "users"

	res := rkgin.RegisterGinEntryYAML(boot)

	return res["users"].(*rkgin.GinEntry)
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
