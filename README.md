# go-testing
TDD and BDD using go.

# USE CASE
#### Story
As a user I want to see latest currency rates in order to exchange my current currency

#### Acceptance Criteria
Given input USD then user will see rates from CAD, AUD, SGD, NZD and IDR

# Running

    $ cd $GOPATH/src/github.com/ru-rocker/go-testing
    $ gucumber

# Notes
* Currency API [fixer.io](http://fixer.io/)
* Http Mock [httpmock](https://github.com/jarcoal/httpmock)
* Unit test API [testify](https://github.com/stretchr/testify)
* BDD API [gucumber](https://github.com/gucumber/gucumber)

# References
* https://medium.com/@PurdonKyle/building-a-unit-testable-api-in-golang-b42ff1fcbae7