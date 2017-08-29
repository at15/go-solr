package schema

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/at15/go-solr/pkg/common"
	"github.com/pkg/errors"
)

type FieldRequest struct {
	AddField    []common.Field      `json:"add-field,omitempty"`
	DeleteField []map[string]string `json:"delete-field,omitempty"`
}

func (svc *Service) AddField(ctx context.Context, field common.Field) error {
	p := FieldRequest{}
	p.AddField = append(p.AddField, field)
	if _, err := svc.client.Post(ctx, svc.baseURL, p, ioutil.Discard); err != nil {
		return errors.WithMessage(err, fmt.Sprintf("can't create field %s", field.Name))
	}
	// FIXME: 200 is returned when error happens, extra check is needed
	return nil
}

func (svc *Service) AddFields(ctx context.Context, fields ...common.Field) error {
	if len(fields) == 0 {
		return errors.New("no field specified")
	}
	p := FieldRequest{}
	p.AddField = fields
	if _, err := svc.client.Post(ctx, svc.baseURL, p, ioutil.Discard); err != nil {
		return errors.WithMessage(err, fmt.Sprintf("can't create %d fileds", len(fields)))
	}
	// FIXME: 200 is returned when error happens, extra check is needed, and the error response could be different
	return nil
}

func (svc *Service) DeleteField(ctx context.Context, name string) error {
	p := FieldRequest{}
	p.DeleteField = append(p.DeleteField, map[string]string{"name": name})
	if _, err := svc.client.Post(ctx, svc.baseURL, p, ioutil.Discard); err != nil {
		return errors.WithMessage(err, fmt.Sprintf("can't delete field %s", name))
	}
	// FIXME: 200 is returned when error happens, extra check is needed
	return nil
}
