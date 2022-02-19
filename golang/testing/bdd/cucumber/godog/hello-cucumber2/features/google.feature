
Feature: get google query


    Scenario: should return google search results
        When I sent "GET" request to "https://www.google.com/search?q=golang"
        Then the response code should be 200
