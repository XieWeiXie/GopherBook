package symbol

import (
	"GopherBook/chapter12/fina/models"
	"GopherBook/chapter12/fina/pkg/database"
)

type ControllerSymbol struct {
}

var Default = ControllerSymbol{}

func (C ControllerSymbol) GetSymbol() (models.SymbolSerializer, error) {
	var symbol models.Symbol
	var result models.SymbolSerializer
	if has, dbError := database.MySQL.Get(&symbol); dbError != nil || !has {
		return result, dbError
	}
	result = symbol.Serializer()
	return result, nil
}
