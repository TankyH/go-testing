@xrates
Feature: Exchange Currency

    Scenario: User successfully logs in
        Given I select base rates "USD"
        And retrieve rates from REST endpoint
        Then retrieve five rates "NZD", "AUD", "SGD", "CAD" and "IDR"

