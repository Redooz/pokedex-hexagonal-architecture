package infrastructure

import (
	"github.com/redooz/podekex-hexagonal-architecture/application/handler"
	"github.com/redooz/podekex-hexagonal-architecture/domain/api/usecase"
	"github.com/redooz/podekex-hexagonal-architecture/infrastructure/input/rest"
	"github.com/redooz/podekex-hexagonal-architecture/infrastructure/output/gorm"
	"github.com/redooz/podekex-hexagonal-architecture/infrastructure/output/gorm/adapter"
	"github.com/redooz/podekex-hexagonal-architecture/infrastructure/output/gorm/entity"
	"github.com/redooz/podekex-hexagonal-architecture/infrastructure/output/gorm/repository"
)

func initializeDatabase() *gorm.Database {
	connectionValues := &gorm.ConnectionValues{
		Host:         "localhost",
		Port:         "3306",
		User:         "root",
		Password:     "",
		DatabaseName: "pokedex",
	}

	entities := []interface{}{
		&entity.Pokemon{},
		&entity.Type{},
	}

	return gorm.NewDatabase(connectionValues, entities)
}

func InitializePokedexController() *rest.PokedexController {
	pokemonRepository := repository.NewPokemonRepository(initializeDatabase())
	typeRepository := repository.NewTypeRepository(initializeDatabase())
	typePersistenceAdapter := adapter.NewTypePersistenceAdapter(typeRepository)
	pokemonPersistenceAdapter := adapter.NewPokemonPersistenceAdapter(pokemonRepository)
	typeUsecase := usecase.NewTypeUseCase(typePersistenceAdapter)
	pokemonUseCase := usecase.NewPokemonUseCase(pokemonPersistenceAdapter, typeUsecase)
	pokedexHandler := handler.NewPokedexHandler(pokemonUseCase, typeUsecase)
	pokedexController := rest.NewPokedexController(pokedexHandler)

	return pokedexController
}

func InitializeTypeController() *rest.TypeController {
	typeRepository := repository.NewTypeRepository(initializeDatabase())
	typePersistenceAdapter := adapter.NewTypePersistenceAdapter(typeRepository)
	typeUsecase := usecase.NewTypeUseCase(typePersistenceAdapter)
	typeHandler := handler.NewTypeHandler(typeUsecase)
	typeController := rest.NewTypeController(typeHandler)

	return typeController
}
