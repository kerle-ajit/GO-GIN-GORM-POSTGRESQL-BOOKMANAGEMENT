package books

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ajit-kerle/go-gin-api-medium/pkg/common/models"

)


func (h handler) GetBooks (c *gin.Context) {
	var books []models.Book
    
	result := h.DB.Find(&books)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
        return
	}

	c.JSON(http.StatusOK, &books)
}