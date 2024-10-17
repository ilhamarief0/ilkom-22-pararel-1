package controllers
import(
	"product_service/backend-api/models"
	"github.com/gin-gonic/gin"
)

func Finpost(c *gin.Context){
	var posts []models.Post
	models.DB.Find(&posts)

	c.JSON(200, gin.H{
		"succes": true,
		"messaage": "list data posts",
		"data": posts,
	})
}