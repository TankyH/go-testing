package features

import (
	. "github.com/gucumber/gucumber"
	"github.com/ru-rocker/go-testing/xrates"
	"github.com/stretchr/testify/assert"
	"fmt"
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
		resp, _ = xrates.GetCurrentRates(rates)
	})

	Then(`^retrieve five rates "(.+?)", "(.+?)", "(.+?)", "(.+?)" and "(.+?)"`,
		func(nzd, aud, sgd, cad, idr string) {
			rates := resp.Rates
			assert.NotNil(T, rates[nzd], fmt.Sprintf("%s currency is %f", nzd, rates[nzd]))
			assert.NotNil(T, rates[aud], fmt.Sprintf("%s currency is %f", aud, rates[aud]))
			assert.NotNil(T, rates[sgd], fmt.Sprintf("%s currency is %f", sgd, rates[sgd]))
			assert.NotNil(T, rates[cad], fmt.Sprintf("%s currency is %f", cad, rates[cad]))
			assert.NotNil(T, rates[idr], fmt.Sprintf("%s currency is %f", idr, rates[idr]))
		})
}
