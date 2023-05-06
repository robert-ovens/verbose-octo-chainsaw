Feature: Check if the service is up
    So I can know that the service is operational
    As a cloud engineer
    I want to check that the service is running

Scenario: The service is running
    Given the status of the service is "up"
    When the status is requested
    Then the status code will be "200"
    
Scenario: The service is not running
    Given the status of the service is "down"
    When the status is requested
    Then the status code will be "500"
    