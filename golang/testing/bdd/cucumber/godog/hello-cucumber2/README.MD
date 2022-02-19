



```
s@mulys-new-mbp hello-cucumber2 % go test
Feature: get google query

  Scenario: should return google search results                           # features/google.feature:5
    When I sent "GET" request to "https://www.google.com/search?q=golang" # google_test.go:18 -> *apiFeature
    Then the response code should be 200                                  # google_test.go:33 -> *apiFeature

1 scenarios (1 passed)
2 steps (2 passed)
284.190888ms
ok      github.com/muly/howto/golang/testing/bdd/cucumber/godog/hello-cucumber2 0.464s
s@mulys-new-mbp hello-cucumber2 % godog
Feature: get google query

  Scenario: should return google search results                           # features/google.feature:5
    When I sent "GET" request to "https://www.google.com/search?q=golang" # google_test.go:18 -> *apiFeature
    Then the response code should be 200                                  # google_test.go:33 -> *apiFeature

1 scenarios (1 passed)
2 steps (2 passed)
308.104161ms
s@mulys-new-mbp hello-cucumber2 % 
```