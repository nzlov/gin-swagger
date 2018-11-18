package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/basic/web"
)

//
// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} web.Pet	"ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /testapi/get-string-by-int/{some_id} [get]
func GetStringByInt(c *gin.Context) {
	id := c.Param("some_id")
	if id != "1" {
		c.JSON(http.StatusOK, web.APIError{
			ErrorCode:    10001,
			ErrorMessage: "not found",
		})
		return
	}
	c.JSON(http.StatusOK, web.Pet{
		ID: 1,
		Category: struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}{
			ID:   2,
			Name: "a",
		},
		Name: "dog",
	})

}

// @Description get struct array by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    string     true        "Some ID"
// @Param   offset     query    int     true        "Offset"
// @Param   limit      query    int     true        "Offset"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /testapi/get-struct-array-by-string/{some_id} [get]
func GetStructArrayByString(c *gin.Context) {

}

type Pet3 struct {
	ID int `json:"id"`
}
