package garantex

import (
	"Gl0ven/kata_projects/rates/internal/models"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func NewMockServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		marketId := r.URL.Query().Get("market")

		if marketId == "usdtrub" {
			w.WriteHeader(http.StatusOK)
			
			w.Write([]byte(`{
				"timestamp": 1648047486,
				"asks": [
					{
					"price": "156612.04",
					"volume": "0.23625259",
					"amount": "37000.0",
					"factor": "0.008",
					"type": "limit"     
					}
				],
				"bids": [
					{
					"price": "153590.1",
					"volume": "0.04143581",
					"amount": "6364.13",
					"factor": "-0.011",
					"type": "factor"
					}
				]
				}`))
		} else {
			http.Error(w, "Unknown market ID", http.StatusBadRequest)
		}
	}))

	return ts
}

func Test_garantexProvider_GetRates(t *testing.T) {
	ms := NewMockServer()
	getRatesUrl = ms.URL
	cl := &http.Client{}
	type fields struct {
		client   *http.Client
		marketId string
	}
	tests := []struct {
		name    string
		fields  fields
		want    models.Rates
		wantErr bool
	}{
		{
			name: "success", 
			fields: fields{client: cl, marketId: "usdtrub"}, 
			want: models.Rates{Timestamp: 1648047486, AskPrice: 156612.04, BidPrice: 153590.1},
			wantErr: false,
		},
		{
			name: "fail", 
			fields: fields{client: cl, marketId: "$rub"}, 
			want: models.Rates{},
			wantErr: true,
		},
			
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &garantexProvider{
				client:   tt.fields.client,
				marketId: tt.fields.marketId,
			}
			got, err := p.GetRates()
			if (err != nil) != tt.wantErr {
				t.Errorf("garantexProvider.GetRates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("garantexProvider.GetRates() = %v, want %v", got, tt.want)
			}
		})
	}
}

