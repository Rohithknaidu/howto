package e2e

import (
	"fmt"
	"net/http"

	"github.com/cucumber/godog"
)

type apiFeature struct {
	resp *http.Response
}

func (a *apiFeature) resetResponse(*godog.Scenario) {
	a.resp = &http.Response{}
}

func (a *apiFeature) iSentRequestTo(method, endpoint string) error {
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return err
	}

	c := http.Client{}

	a.resp, err = c.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func (a *apiFeature) theResponseCodeShouldBe(statusCode int) error {
	if statusCode != a.resp.StatusCode {
		return fmt.Errorf("expected status code %d, but received %d", statusCode, a.resp.StatusCode)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	a := &apiFeature{}

	ctx.BeforeScenario(a.resetResponse)

	ctx.Step(`^I sent "([^"]*)" request to "([^"]*)"$`, a.iSentRequestTo)
	ctx.Step(`^the response code should be (\d+)$`, a.theResponseCodeShouldBe)
}
