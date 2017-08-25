package core

import (
	"context"
	"encoding/json"
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
			Q          string      `json:"q"`
			Df         string      `json:"df"`
			FacetField interface{} `json:"facet.field"` // NOTE: this could be string/[]string
			Facet      string      `json:"facet"`
			Indent     string      `json:"indent"`
			Start      string      `json:"start"`
			Sort       string      `json:"sort"`
			Wt         string      `json:"wt"`
			//NAMING_FAILED string `json:"_"`
		} `json:"params"`
	} `json:"responseHeader"`
	Response struct {
		NumFound int                      `json:"numFound"`
		Start    int                      `json:"start"`
		Docs     []map[string]interface{} `json:"docs"`
	} `json:"response"`
	FacetCounts struct {
		FacetQueries   interface{}           `json:"facet_queries"`
		FacetFields    map[string]FacetField `json:"facet_fields"` // NOTE: facet fields mix string and number in array https://github.com/at15/go-solr/issues/17
		FacetRanges    interface{}           `json:"facet_ranges"`
		FacetIntervals interface{}           `json:"facet_intervals"`
		FacetHeatmaps  interface{}           `json:"facet_heatmaps"`
	} `json:"facet_counts"`
}

type FacetField map[string]int

// https://groups.google.com/forum/#!topic/golang-nuts/IxPipKwI-zQ
// https://play.golang.org/p/YgUIFxT7hA
func (f *FacetField) UnmarshalJSON(data []byte) error {
	//log.Info("facet field json unmarshaler called")
	var mixed []json.Number
	if err := json.Unmarshal(data, &mixed); err != nil {
		return err
	}
	m := make(map[string]int, len(mixed)/2)
	for i := 0; i < len(mixed); i += 2 {
		c, err := mixed[i+1].Int64()
		if err != nil {
			return err
		}
		m[mixed[i].String()] = int(c)
	}
	*f = m
	return nil
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
	log.Info(req.URL)
	if _, err := svc.client.Do(ctx, req, res); err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("solr: can't select %s", req.URL.String()))
	}
	return res, nil
}
