package core

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func (svc *Service) Update(ctx context.Context, body interface{}) error {
	req, err := svc.client.NewRequest(http.MethodPost, svc.baseUpdateURL, body)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Set("commit", "true")
	req.URL.RawQuery = q.Encode()
	// FIXME: get error response it might be a 200 response, the ok response is {"responseHeader":{"status":0,"QTime":238}} we can ignore
	if _, err := svc.client.Do(ctx, req, ioutil.Discard); err != nil {
		return errors.WithMessage(err, fmt.Sprintf("solr: can't update document %s", req.URL.String()))
	}
	return nil
}
