package dtomapper

import (
	"github.com/redooz/podekex-hexagonal-architecture/application/dto/request"
	"github.com/redooz/podekex-hexagonal-architecture/application/dto/response"
	"github.com/redooz/podekex-hexagonal-architecture/domain/model"
)

func TypeRequestToModel(request *request.Type) *model.Type {
	return &model.Type{
		FirstType:  request.FirstType,
		SecondType: request.SecondType,
	}
}

func TypeModelToResponse(model *model.Type) *response.Type {
	return &response.Type{
		ID:         model.ID,
		FirstType:  model.FirstType,
		SecondType: model.SecondType,
	}
}

func SliceTypeModelToSliceResponse(types []*model.Type) []*response.Type {
	var typeResponse []*response.Type
	for _, typeModel := range types {
		typeResponse = append(typeResponse, TypeModelToResponse(typeModel))
	}
	return typeResponse
}
