package e2e

import (
	"fmt"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
)

type inputData struct {
	Id   int    // Note: the column names and the field names in the data table in feature file should have exact match. not sure if there are any cucumber struct flags that can be added to make this work without exact match
	Name string 
}

func belowData(data *godog.Table) error {
	assist := assistdog.NewDefault()
	input, err := assist.CreateSlice(new(inputData), data)
	if err != nil {
		return err
	}
	for _, r := range input.([]*inputData) {
		fmt.Println(r.Id, r.Name)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^below data$`, belowData)
}
