Feature: Weather API
  Scenario: Get weather for a city
    Given I am a user
    When I request the weather for "London"
    Then the response should contain "London"
