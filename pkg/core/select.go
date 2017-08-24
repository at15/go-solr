package core

import (
	"context"
	"fmt"
	"github.com/at15/go-solr/pkg/search"
	"github.com/pkg/errors"
	"net/http"
)

type SelectResponse struct {
	ResponseHeader struct {
		Status int `json:"status"`
		QTime  int `json:"QTime"`
		Params struct {
			Q      string `json:"q"`
			Df     string `json:"df"`
			Indent string `json:"indent"`
			Start  string `json:"start"`
			Sort   string `json:"sort"`
			Wt     string `json:"wt"`
			//NAMING_FAILED string `json:"_"`
		} `json:"params"`
	} `json:"responseHeader"`
	Response struct {
		NumFound int                      `json:"numFound"`
		Start    int                      `json:"start"`
		Docs     []map[string]interface{} `json:"docs"`
	} `json:"response"`
}

func (svc *Service) Select(ctx context.Context, query search.Query) (*SelectResponse, error) {
	res := &SelectResponse{}
	req, err := svc.client.NewRequest(http.MethodGet, svc.baseSelectURL, nil)
	if err != nil {
		return nil, err
	}
	q := query.Encode()
	// TODO: maybe we can pass *url.Values to encode so we can reuse existing parameters
	q.Set("wt", "json")
	req.URL.RawQuery = q.Encode()
	if _, err := svc.client.Do(ctx, req, res); err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("solr: can't select %s", req.URL.String()))
	}
	return res, nil
}
