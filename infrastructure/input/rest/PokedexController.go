package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/redooz/podekex-hexagonal-architecture/application/dto/request"
	"github.com/redooz/podekex-hexagonal-architecture/application/handler"
	"net/http"
	"strconv"
)

type PokedexController struct {
	pokedexHandler handler.IPokedexHandler
}

func NewPokedexController(pokedexHandler handler.IPokedexHandler) *PokedexController {
	return &PokedexController{pokedexHandler: pokedexHandler}
}

func (p *PokedexController) SavePokemonToPokedex(c *gin.Context) {
	var body request.Pokedex

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	httpStatus, err := p.pokedexHandler.SavePokemonToPokedex(&body)

	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
	}

	c.JSON(httpStatus, gin.H{"message": "Pokemon saved successfully"})
}

func (p *PokedexController) GetAllPokemonFromPokedex(c *gin.Context) {
	response, httpStatus, err := p.pokedexHandler.GetAllPokemonFromPokedex()

	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}

	c.JSON(httpStatus, response)
}

func (p *PokedexController) GetPokemonFromPokedexByNumber(c *gin.Context) {
	pokedexNumberParam := c.Param("number")

	pokedexNumber, err := strconv.Atoi(pokedexNumberParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, httpStatus, err := p.pokedexHandler.GetPokemonFromPokedexByNumber(pokedexNumber)

	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}

	c.JSON(httpStatus, response)
}

func (p *PokedexController) UpdatePokemonInPokedex(c *gin.Context) {
	var body request.Pokedex
	pokedexNumberParam := c.Param("number")

	pokedexNumber, err := strconv.Atoi(pokedexNumberParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	if err = c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	httpStatus, err := p.pokedexHandler.UpdatePokemonInPokedex(&body, pokedexNumber)

	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}

	c.JSON(httpStatus, nil)
}

func (p *PokedexController) DeletePokemonFromPokedex(c *gin.Context) {
	pokedexNumberParam := c.Param("number")

	pokedexNumber, err := strconv.Atoi(pokedexNumberParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	httpStatus, err := p.pokedexHandler.DeletePokemonFromPokedex(pokedexNumber)

	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}

	c.JSON(httpStatus, nil)
}
