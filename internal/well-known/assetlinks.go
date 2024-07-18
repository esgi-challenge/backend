package wellknown

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AssetLinksTarget struct {
	Namespace   string   `json:"namespace" binding:"required"`
	PackageName string   `json:"package_name" binding:"required"`
	Sha256      []string `json:"sha256_cert_fingerprints" binding:"required"`
}

type AssetLinks struct {
	Relation []string         `json:"relation" binding:"required"`
	Target   AssetLinksTarget `json:"target" binding:"required"`
}

func GetAssetLinks() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := []AssetLinks{
			{
				Relation: []string{
					"delegate_permission/common.handle_all_urls",
				},
				Target: AssetLinksTarget{
					Namespace:   "android_app",
					PackageName: "com.example.mobile",
					Sha256: []string{
						"40:6D:2A:42:4A:19:D0:E4:D2:FD:9A:39:7D:5E:38:70:89:71:4E:F9:E8:C1:08:6C:8D:D8:64:EA:88:BC:BF:09",
						"7B:5D:4B:11:F3:13:FD:22:2F:45:F9:73:B3:A5:32:68:D0:33:73:2E:24:21:A1:C1:30:6A:FB:F4:5D:6F:70:BB",
						"53:E0:01:14:35:D0:7D:CD:A3:A8:EB:B8:36:A5:72:7F:15:A4:5A:B0:D2:F2:3E:4B:7F:55:9D:F9:72:39:8C:0A",
					},
				},
			},
		}

		ctx.JSON(http.StatusOK, body)
	}
}
