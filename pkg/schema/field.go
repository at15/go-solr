package schema

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/at15/go-solr/pkg/common"
	"github.com/pkg/errors"
)

const (
	opAddField    = "add-field"
	opDeleteField = "delete-field"
)

func (svc *Service) AddField(ctx context.Context, field common.Field) error {
	p := make(map[string]common.Field)
	p[opAddField] = field
	if _, err := svc.client.Post(ctx, svc.baseURL, p, ioutil.Discard); err != nil {
		return errors.WithMessage(err, fmt.Sprintf("can't create field %s", field.Name))
	}
	// FIXME: 200 is returned when error happens, extra check is needed
	return nil
}

func (svc *Service) DeleteField(ctx context.Context, name string) error {
	p := make(map[string]map[string]string)
	p[opDeleteField] = map[string]string{"name": name}
	if _, err := svc.client.Post(ctx, svc.baseURL, p, ioutil.Discard); err != nil {
		return errors.WithMessage(err, fmt.Sprintf("can't delete field %s", name))
	}
	// FIXME: 200 is returned when error happens, extra check is needed
	return nil
}
