package adsinsights

import (
	"encoding/json"
	"github.com/mazti/facebook-go-business-sdk/sdk"
	"net/http"
)

const (
	endpoint = "/insights"
)

type AdsInsights struct {
	Context         *sdk.APIContext
	AccountCurrency string `json:"account_currency"`
	AccountID       string `json:"account_id"`
	AccountName     string `json:"account_name"`

	ActionsPerImpression string `json:"actions_per_impression"`
	ActivityRecency      string `json:"activity_recency"`

	AdBidType  string `json:"ad_bid_type"`
	AdBidValue string `json:"ad_bid_value"`

	// TODO: adding more field here
}

func CreateAPIRequestGet(id string, context *sdk.APIContext) *sdk.APIRequest {
	return sdk.NewAPIRequest(
		context,
		id,
		endpoint,
		http.MethodGet,
		sdk.Parser(parserResponse),
		sdk.ParamNames(params),
	)
}

func parserResponse(data json.RawMessage) (sdk.APIResponse, error) {
	ent := &AdsInsights{}
	if err := json.Unmarshal(data, ent); err != nil {
		return ent, err
	}
	return ent, nil
}