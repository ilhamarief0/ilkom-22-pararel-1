package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/veritrans/go-midtrans"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"}, // Pastikan port sesuai dengan frontend
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.POST("/payment", func(c *gin.Context) {
		var req struct {
			OrderID     string `json:"order_id"`
			GrossAmount int64  `json:"gross_amount"`
		}

		// Bind request JSON ke struct dan validasi
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		snapGateway := midtrans.SnapGateway{
			Client: midtrans.Client{
				ServerKey:  "SB-Mid-server-Dh8ojLkmziWFkCqiJGiDAkq0", // Pastikan ini adalah Server Key Sandbox
				ClientKey:  "SB-Mid-client-etaFzjV97U45y5La",
				APIEnvType: midtrans.Sandbox, // Gunakan midtrans.Production untuk production
			},
		}

		snapReq := &midtrans.SnapReq{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  req.OrderID,
				GrossAmt: req.GrossAmount,
			},
		}

		snapResp, err := snapGateway.GetToken(snapReq)
		if err != nil {
			log.Printf("Failed to create transaction: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
			return
		}

		// Log untuk memastikan URL yang dikirim benar
		log.Printf("Redirect URL dari Midtrans: %s", snapResp.RedirectURL)

		// Kirim URL ke frontend
		c.JSON(http.StatusOK, gin.H{"redirect_url": snapResp.RedirectURL})

	})

	// Jalankan server di port 8080
	r.Run(":8080")
}
