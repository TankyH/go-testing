package xrates

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"fmt"
	"github.com/jarcoal/httpmock"
)

// Private Function
// HTTP Mock
func TestGetRatesReturn200(t *testing.T) {
	assert := assert.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responder := httpmock.NewStringResponder(200, `{
		"base":"USD",
		"date":"2016-06-23",
		"rates":{"AUD":1.3214,
			 "CAD":1.3231,
			 "IDR":13305.0,
			 "NZD":1.3734,
			 "SGD":1.3869
			 }
		}`)

	rates := Rates{
		Request: RatesQuotedRequest{Base: "USD"},
	}

	url := fmt.Sprintf("http://api.fixer.io/latest?base=%s", rates.Request.Base)
	httpmock.RegisterResponder("GET", url, responder)
	resp, _ := rates.getRates()

	assert.Equal("USD", resp.Base, "Base Currency")
	assert.Equal("2016-06-23", resp.Date, "Exchange Date")
	assert.Equal(float32(1.3214), resp.Rates["AUD"],"Australian X-Rate")
	assert.Equal(float32(13305.0), resp.Rates["IDR"],"Indonesian X-Rate")

}

func TestGetRatesErrorJSONFormat(t *testing.T) {
	//assert := assert.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responder := httpmock.NewStringResponder(200, `{
		"base":"USD",
		}`)

	rates := Rates{
		Request: RatesQuotedRequest{Base: "USD"},
	}

	url := fmt.Sprintf("http://api.fixer.io/latest?base=%s", rates.Request.Base)
	httpmock.RegisterResponder("GET", url, responder)
	_, err := rates.getRates()

	assert.NotNil(t, err, "Error while parsing json")

}

func TestGetRatesEmptyBody(t *testing.T) {
	//assert := assert.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responder := httpmock.NewStringResponder(200, "")

	rates := Rates{
		Request: RatesQuotedRequest{Base: "USD"},
	}

	url := fmt.Sprintf("http://api.fixer.io/latest?base=%s", rates.Request.Base)
	httpmock.RegisterResponder("GET", url, responder)
	_, err := rates.getRates()

	assert.NotNil(t, err, "EOF")

}

// Public Function
type ratesMock struct {
	mock.Mock
}

func (o *ratesMock) getRates() (RatesQuotedResponse, error) {
	args := o.Called()
	return args.Get(0).(RatesQuotedResponse), args.Error(1)
}

func TestFunctionGetCurrentRates(t *testing.T) {
	assert := assert.New(t)

	ex_rates := map[string]float32{
		"AUD": 1.0000,
		"NZD": 1.0000,
		"IDR": 13000.0000,
		"CAD": 1.0000,
		"SGD": 1.0000,
	}

	ar := RatesQuotedResponse{
		Base:  "USD",
		Rates: ex_rates,
	}

	rmock := new(ratesMock)
	rmock.On("GetRates").Return(ar, nil)
	resp, err := GetCurrentRates(rmock)

	assert.Nil(err, "no error")
	assert.Equal("USD", resp.Base, "The currency should be USD")
	assert.Equal(5, len(resp.Rates), "Total exchange currency")

	rmock.AssertExpectations(t)
}
