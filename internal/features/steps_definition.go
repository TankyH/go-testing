package features

import (
	. "github.com/gucumber/gucumber"
	"github.com/ru-rocker/go-testing/xrates"
	"github.com/stretchr/testify/assert"
	"fmt"
	"github.com/jarcoal/httpmock"
)

func init() {

	var rates xrates.Rates
	var resp xrates.RatesQuotedResponse

	Before("@xrates", func() {
		// runs before every feature or scenario tagged with @xrates
		rates = xrates.Rates{}
	})

	Given(`^I select base rates "(.+?)"$`, func(base string) {
		req := xrates.RatesQuotedRequest{Base: base}
		rates.Request = req
	})

	And(`^retrieve rates from REST endpoint$`, func() {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		responder := httpmock.NewStringResponder(200, `{
			"base": "USD",
			"date": "2016-06-23",
			"rates": {
				"AUD": 1.3214,
				"CAD": 1.3231,
				"IDR": 13305.0,
				"NZD": 1.3734,
				"SGD": 1.3869
			}
		}`)

		url := fmt.Sprintf("http://api.fixer.io/latest?base=%s", rates.Request.Base)
		httpmock.RegisterResponder("GET", url, responder)
		resp, _ = xrates.GetCurrentRates(rates)
	})

	Then(`^retrieve five rates "(.+?)", "(.+?)", "(.+?)", "(.+?)" and "(.+?)"`,
		func(nzd, aud, sgd, cad, idr string) {
			rates := resp.Rates
			assert.NotEmpty(T, rates[nzd], fmt.Sprintf("%s currency is %f", nzd, rates[nzd]))
			assert.NotEmpty(T, rates[aud], fmt.Sprintf("%s currency is %f", aud, rates[aud]))
			assert.NotEmpty(T, rates[sgd], fmt.Sprintf("%s currency is %f", sgd, rates[sgd]))
			assert.NotEmpty(T, rates[cad], fmt.Sprintf("%s currency is %f", cad, rates[cad]))
			assert.NotEmpty(T, rates[idr], fmt.Sprintf("%s currency is %f", idr, rates[idr]))
		})
}
