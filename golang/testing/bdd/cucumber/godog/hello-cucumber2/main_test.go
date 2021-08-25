package main

import (
	"testing"

	"github.com/cucumber/godog"


	"github.com/muly/howto/testing/bdd/cucumber/godog/hello-cucumber2/e2e" 
)

func TestMain(m *testing.M) {
	var t *testing.T // Comes from your test function, e.g. func TestFeatures(t *testing.T).

	suite := godog.TestSuite{
		ScenarioInitializer: func(s *godog.ScenarioContext) {
			e2e.InitializeScenario(s)
		},
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"e2e/features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
