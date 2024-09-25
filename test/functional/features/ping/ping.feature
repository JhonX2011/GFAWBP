Feature: Host health checks
  In order to know status's application
  I need to be able to ping the application

  Scenario: The server is up and running
    When i verify if server is up and running
    Then the server is up and running
    And the server status code is 200
    And the response is "pong"