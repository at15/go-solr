package solr

import (
	"context"
	"net/http"

	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
)

const (
	adminBaseURL       = "/solr/admin"
	adminCoreBaseURL   = "/solr/admin/cores"
	pAction            = "action"
	actionCreate       = "CREATE"
	actionRename       = "RENAME"
	actionStatus       = "STATUS"
	actionUnload       = "UNLOAD"
	pInstanceDir       = "instanceDir"
	pConfigSet         = "configSet"
	pCore              = "core"
	pName              = "name"
	pIndexInfo         = "indexInfo"
	pDeleteInstanceDir = "deleteInstanceDir"
)

// TODO: it seems all the http methods work for core admin API ... you can use GET for everything that does not have request body

// CreateCore create cores directly, you will get error if the core already exists. NOTE: when using configSet, Solr does NOT
// copy configSet to instance folder, if you have multiple cores using same configSet, you have to make sure they have same schema,
// otherwise you will end up having a mixed schema when you ingest different document. It is different from the bin/solr command
// see https://github.com/at15/go-solr/issues/15 for detail
// http://localhost:8983/solr/admin/cores?action=CREATE&name=films&instanceDir=films&configSet=data_driven_schema_configs
func (c *Client) CreateCore(ctx context.Context, core Core) error {
	req, err := c.client.NewRequest(http.MethodPost, adminCoreBaseURL, nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Set(pAction, actionCreate)
	q.Set(pName, core.Name)
	if core.InstanceDir != "" {
		q.Set(pInstanceDir, core.InstanceDir)
	}
	if core.ConfigSet != "" {
		q.Set(pConfigSet, core.ConfigSet)
	}
	req.URL.RawQuery = q.Encode()
	if _, err := c.client.Do(ctx, req, ioutil.Discard); err != nil {
		return errors.WithMessage(err, fmt.Sprintf("solr: can't create core %s", req.URL.String()))
	}
	return nil
}

// CreateCoreIfNotExists gets all cores using status API and create core if the provided core is not found,
// the first return value (exists) tells if the core was already there.
// In fact we can call create directly and check if the error message contains 'already exists' to save a RTT, which is our old way
func (c *Client) CreateCoreIfNotExists(ctx context.Context, core Core) (exists bool, err error) {
	cores, err := c.CoresStatus(ctx, false, "")
	if _, ok := cores[core.Name]; ok {
		exists = true
		err = nil
		return
	}
	return false, c.CreateCore(ctx, core)
}

// CoresStatus returns status of all cores if core is set to "", it would also return index information (num of docs etc.) if includeIndexInfo is true
func (c *Client) CoresStatus(ctx context.Context, includeIndexInfo bool, core string) (map[string]CoreStatus, error) {
	req, err := c.client.NewRequest(http.MethodGet, adminCoreBaseURL, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Set(pAction, actionStatus)
	if core != "" {
		q.Set(pCore, core)
	}
	if includeIndexInfo {
		q.Set(pIndexInfo, "true")
	} else {
		q.Set(pIndexInfo, "false")
	}
	req.URL.RawQuery = q.Encode()
	res := &CoreStatusResponse{}
	if _, err := c.client.Do(ctx, req, res); err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("solr: can't get core status %s", req.URL.String()))
	}
	return res.Status, nil
}

// http://localhost:8983/solr/admin/cores?action=RENAME&core=old_name&other=new
// TODO: implementation
func (c *Client) RenameCore() error {
	return nil
}

// http://localhost:8983/solr/admin/cores?action=UNLOAD&core=my_solr_core
// TODO: test it
func (c *Client) UnloadCore(ctx context.Context, core string) error {
	req, err := c.client.NewRequest(http.MethodPost, adminCoreBaseURL, nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Set(pAction, actionUnload)
	q.Set(pCore, core)
	q.Set(pDeleteInstanceDir, "false") // the only different between delete and unload
	req.URL.RawQuery = q.Encode()
	if _, err := c.client.Do(ctx, req, ioutil.Discard); err != nil {
		return errors.WithMessage(err, fmt.Sprintf("solr: can't unload core %s", req.URL.String()))
	}
	return nil
}

// http:/localhost:8983/solr/admin/cores?acton=UNLOAD&core=mysolrcore&deleteInstanceDir=true
func (c *Client) DeleteCore(ctx context.Context, core string) error {
	req, err := c.client.NewRequest(http.MethodPost, adminCoreBaseURL, nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Set(pAction, actionUnload)
	q.Set(pCore, core)
	q.Set(pDeleteInstanceDir, "true")
	req.URL.RawQuery = q.Encode()
	if _, err := c.client.Do(ctx, req, ioutil.Discard); err != nil {
		return errors.WithMessage(err, fmt.Sprintf("solr: can't delete core %s", req.URL.String()))
	}
	return nil
}

func (c *Client) SystemInfo(ctx context.Context) (*SystemInfoResponse, error) {
	res := &SystemInfoResponse{}
	if _, err := c.client.Get(ctx, adminBaseURL+"/info/system", res); err != nil {
		return res, err
	}
	return res, nil
}
