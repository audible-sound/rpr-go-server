package db

import (
	"fmt"
	"reflect"
	"strings"

	modelManager "github.com/audible-sound/rpr-go-server/models"
	"gorm.io/gorm"
)

// Function to determine table dependencies
func getDependencies(models []interface{}) map[string][]string {
	dependencyMap := make(map[string][]string)
	for _, model := range models {
		modelType := reflect.TypeOf(model).Elem()
		tableName := modelType.Name()

		dependencyMap[tableName] = []string{}

		// Iterate through struct fields
		for i := 0; i < modelType.NumField(); i++ {
			field := modelType.Field(i)

			if field.Type.Kind() == reflect.Slice {
				continue
			}

			// If foreign key exists add dependency to dependencyMap

			tag, containsGormTag := field.Tag.Lookup("gorm")
			var isFk bool = strings.Contains(tag, "foreignKey")

			if containsGormTag && isFk {
				var parentTable string = field.Type.Name()
				dependencyMap[tableName] = append(dependencyMap[tableName], parentTable)
			}
		}
	}
	return dependencyMap
}

// Sort Models using Khan's Algorithm
func sortModels() []interface{} {
	models := modelManager.GetModelRegistry()
	dependencyMap := getDependencies(models)
	modelMap := modelManager.ModelNameMap()

	inDegree := make(map[string]int)
	adjList := make(map[string][]string)

	for node, deps := range dependencyMap {
		_, exists := inDegree[node]

		if !exists {
			inDegree[node] = 0
		}

		for _, dep := range deps {
			adjList[dep] = append(adjList[dep], node)
			inDegree[node]++
		}
	}

	var queue []string
	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	var sortedOrder []interface{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:] // Pop front
		model := modelMap[node]
		sortedOrder = append(sortedOrder, model)

		// Reduce in-degree of adjacent nodes
		for _, neighbor := range adjList[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return sortedOrder
}

func reverseModels(models []interface{}) {
	n := len(models)
	for i := 0; i < n/2; i++ {
		models[i], models[n-1-i] = models[n-1-i], models[i] // swap values
	}
}

func MigrateTables(db *gorm.DB) {
	fmt.Println("Migrating tables...")

	sortedModels := sortModels()
	err := db.Migrator().AutoMigrate(sortedModels...)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Tables sorted successfully!")
	}
}

func DropTables(db *gorm.DB) {
	fmt.Println("Dropping tables...")

	sortedModels := sortModels()
	reverseModels(sortedModels)
	err := db.Migrator().DropTable(sortedModels...)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Tables Dropped successfully!")
	}
}
