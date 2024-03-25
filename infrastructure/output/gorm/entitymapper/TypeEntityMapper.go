package entitymapper

import (
	"github.com/redooz/podekex-hexagonal-architecture/domain/model"
	"github.com/redooz/podekex-hexagonal-architecture/infrastructure/output/gorm/entity"
)

func TypeEntityToModel(entity *entity.Type) *model.Type {
	return &model.Type{
		ID:         entity.ID,
		FirstType:  entity.FirstType,
		SecondType: entity.SecondType,
	}
}

func TypeModelToEntity(model *model.Type) *entity.Type {
	return &entity.Type{
		FirstType:  model.FirstType,
		SecondType: model.SecondType,
	}
}

func SliceTypeEntityToSliceModel(entityTypes []*entity.Type) []*model.Type {
	var modelTypes []*model.Type
	for _, entityType := range entityTypes {
		modelTypes = append(modelTypes, TypeEntityToModel(entityType))
	}
	return modelTypes
}

func SliceTypeModelToSliceEntity(modelTypes []*model.Type) []*entity.Type {
	var entityTypes []*entity.Type
	for _, modelType := range modelTypes {
		entityTypes = append(entityTypes, TypeModelToEntity(modelType))
	}
	return entityTypes
}
