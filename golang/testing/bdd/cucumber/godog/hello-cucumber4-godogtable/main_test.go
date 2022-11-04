package e2e

import (
	"testing"

	"github.com/cucumber/godog"

	"github.com/muly/howto/golang/testing/bdd/cucumber/godog/hello-cucumber4-godogtable/e2e"
)

func TestMain(m *testing.M) {
	var t *testing.T

	suite := godog.TestSuite{
		ScenarioInitializer: func(s *godog.ScenarioContext) {
			e2e.InitializeScenario(s)
		},
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"e2e/features"},
			TestingT: t, 
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
