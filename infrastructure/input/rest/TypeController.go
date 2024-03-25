package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/redooz/podekex-hexagonal-architecture/application/dto/request"
	"github.com/redooz/podekex-hexagonal-architecture/application/handler"
	"net/http"
	"strconv"
)

type TypeController struct {
	typeHandler handler.ITypeHandler
}

func NewTypeController(typeHandler handler.ITypeHandler) *TypeController {
	return &TypeController{typeHandler: typeHandler}
}

func (t *TypeController) SaveType(c *gin.Context) {
	var body request.Type

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	httpStatus, err := t.typeHandler.CreatePokemonType(&body)

	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
	}

	c.JSON(httpStatus, gin.H{"message": "Type saved successfully"})
}

func (t *TypeController) GetAllTypes(c *gin.Context) {
	response, httpStatus, err := t.typeHandler.GetAllPokemonTypes()

	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}

	c.JSON(httpStatus, response)
}

func (t *TypeController) GetTypeByID(c *gin.Context) {
	typeIDParam := c.Param("id")

	typeID, err := strconv.Atoi(typeIDParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, httpStatus, err := t.typeHandler.GetPokemonTypeByID(typeID)

	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
	}
	c.JSON(httpStatus, response)
}

func (t *TypeController) UpdateType(c *gin.Context) {
	typeIDParam := c.Param("id")

	typeID, err := strconv.Atoi(typeIDParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var body request.Type

	if err = c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	httpStatus, err := t.typeHandler.UpdatePokemonType(&body, typeID)

	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
	}

	c.JSON(httpStatus, nil)
}

func (t *TypeController) DeleteType(c *gin.Context) {
	typeIDParam := c.Param("id")

	typeID, err := strconv.Atoi(typeIDParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	httpStatus, err := t.typeHandler.DeletePokemonType(typeID)

	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
	}

	c.JSON(httpStatus, nil)
}
