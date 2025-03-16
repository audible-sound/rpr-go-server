package models

import "reflect"

var modelRegistry []interface{}

func Register(model interface{}) {
	modelRegistry = append(modelRegistry, model)
}

func GetModelRegistry() []interface{} {
	return modelRegistry
}

// Map the model name to the interface
func ModelNameMap() map[string]interface{} {
	modelMap := make(map[string]interface{})
	for _, model := range modelRegistry {
		modelType := reflect.TypeOf(model).Elem()
		var modelName string = modelType.Name()

		modelMap[modelName] = model
	}
	return modelMap
}
