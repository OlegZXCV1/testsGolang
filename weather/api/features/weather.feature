Feature: Weather API
  Scenario: Get weather for a city
    Given I am a user
    When I request the weather for "London"
    Then the response should contain "London"

  Scenario Outline: Get weather for different cities
    Given I am a user
    When I request the weather for "<city>"
    Then the response should contain "<city>"

    Examples:
      | city   |
      | Paris  |
      | Berlin |
      | Tokyo  |

  Scenario: Get weather for a non-existent city
    Given I am a user
    When I request the weather for "NonExistentCity"
    Then the response should have status code 404

  Scenario: Check response header
    Given I am a user
    When I request the weather for "Berlin"
    Then the response header "Content-Type" should be "text/plain; charset=utf-8"
