# steps to get ready with implementing cocumber using golang godog library

## 1. install godog

```
go get github.com/cucumber/godog/cmd/godog
```
or, to install a specific version instead of master
```
go get github.com/cucumber/godog/cmd/godog@v0.12.0
```

 

## 2. define features
see `features/google.feature` file

## 3. generate the golang snippets 
From the project folder where the `features` folder exists, run the below command to genetrate golang code snippet for all the features
```
godog
```
to generate the code snippet for a single feature, provide the relative path of the corresponding feature file
```
godog features/google.feature
```

generated code snippet looks as below:
``` golang
func iSentRequestTo(arg1, arg2 string) error {
        return godog.ErrPending
}

func theResponseCodeShouldBe(arg1 int) error {
        return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
        ctx.Step(`^I sent "([^"]*)" request to "([^"]*)"$`, iSentRequestTo)
        ctx.Step(`^the response code should be (\d+)$`, theResponseCodeShouldBe)
}
```

copy the generated code snippet and save to a `_test.go` file

## 3. implement feature functions in go
- in the generated code snippets, implement the actual test code that tests the given feature.
- modify all the code as necessary. see `google_test.go`

## 4. run the tests 

From the project folder where the `features` folder exists, run the below command to test the all the features
```
godog
```
to test a single feature, provide the relative path of the corresponding feature file
```
godog features/google.feature
```

# references:

inspired by the example: https://github.com/cucumber/godog/tree/main/_examples/api

