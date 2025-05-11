package main

import (
	"fmt"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"

	models "card-master/model"
)

var modelRegistry []interface{} = []interface{}{
	&models.CardEntity{},
	&models.CardSeriesEntity{},
	&models.CardBrandEntity{},
}

func main() {
	// Export the schema
	for _, model := range modelRegistry {
		stmts, err := gormschema.New("sqlite").Load(model)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading model %T: %v\n", model, err)
			os.Exit(1)
		}
		fmt.Println(stmts)
	}
}