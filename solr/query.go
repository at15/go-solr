package solr

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type SortOrder string
type DebugType string

const (
	SortOrderAsc        SortOrder = "asc"
	SortOrderDesc       SortOrder = "desc"
	StandardQueryParser           = "lucene"
	AllFields                     = "*"
	DebugQuery          DebugType = "query"
	DebugTiming         DebugType = "timing"
	DebugResults        DebugType = "results"
	DebugAll            DebugType = "all"
)

type CommonQuery struct {
	sortFields []string // NOTE: we use array because we have to keep the order
	sortOrders []string
	start      int
	rows       int
	fq         string
	fl         []string
	debug      string
}

// https://lucene.apache.org/solr/guide/6_6/common-query-parameters.html#CommonQueryParameters-ThesortParameter
func (q *CommonQuery) SortBy(field string, order SortOrder) *CommonQuery {
	// TODO: we should handle if one field is specified twice
	q.sortFields = append(q.sortFields, field)
	q.sortOrders = append(q.sortOrders, string(order))
	return q
}

func (q *CommonQuery) Start(i int) *CommonQuery {
	q.start = i
	return q
}

func (q *CommonQuery) Rows(i int) *CommonQuery {
	q.rows = i
	return q
}

// IncludeField set field that will be returned by the query
// https://lucene.apache.org/solr/guide/6_6/common-query-parameters.html#CommonQueryParameters-Thefl_FieldList_Parameter
// TODO: function can be pseudo-field https://lucene.apache.org/solr/guide/6_6/common-query-parameters.html#CommonQueryParameters-Thefl_FieldList_Parameter
func (q *CommonQuery) IncludeField(field string) *CommonQuery {
	q.fl = append(q.fl, field)
	return q
}

// Debug sets the debug parameter in query
// https://lucene.apache.org/solr/guide/6_6/common-query-parameters.html#CommonQueryParameters-ThedebugParameter
func (q *CommonQuery) Debug(t DebugType) *CommonQuery {
	q.debug = string(t)
	return q
}

func (q *CommonQuery) Encode() *url.Values {
	p := &url.Values{}
	if q.start != 0 {
		p.Set("start", strconv.Itoa(q.start))
	}
	if q.rows != 0 {
		p.Set("rows", strconv.Itoa(q.rows))
	}
	if len(q.fl) > 0 {
		p.Set("fl", strings.Join(q.fl, ","))
	}
	if len(q.sortFields) > 0 {
		sorts := []string{}
		for i := 0; i < len(q.sortFields); i++ {
			sorts = append(sorts, fmt.Sprintf("%s %s", q.sortFields[i], q.sortOrders[i]))
		}
		p.Set("sort", strings.Join(sorts, ","))
	}
	if q.debug != "" {
		p.Set("debug", q.debug)
	}
	return p
}

//TODO: set filter query https://lucene.apache.org/solr/guide/6_6/common-query-parameters.html#CommonQueryParameters-Thefq_FilterQuery_Parameter
// The fq parameter defines a query that can be used to restrict the superset of documents that can be returned, without influencing score.
// It can be very useful for speeding up complex queries, since the queries specified with fq are cached independently of the main query
//// no filter cache
//q=singer(bob marley) title:(redemption song) language:english genre:rock
//
//// one cache entry
//q=singer(bob marley) title:(redemption song)&fq=language:english AND genre:rock
//
//// two cache entry
//q=singer(bob marley) title:(redemption song)&fq=language:english&fq=genre:rock

type Query interface {
	DefType() string
	Encode() *url.Values // TODO: maybe change it to ToParams
}

var _ Query = (*StdQuery)(nil)

// TODO: https://github.paypal.com/piguo/tiller-server/issues/30
// StdQuery is for Standard Query Parser https://cwiki.apache.org/confluence/display/solr/The+Standard+Query+Parser
type StdQuery struct {
	CommonQuery
	q           string
	df          string
	facetFields []string // TODO: there are many more configuration for facet
}

func (q *StdQuery) DefType() string {
	return StandardQueryParser
}

func (q *StdQuery) DefaultField(field string) *StdQuery {
	q.df = field
	return q
}

func (q *StdQuery) FacetField(field string) *StdQuery {
	q.facetFields = append(q.facetFields, field)
	return q
}

// https://lucene.apache.org/solr/guide/6_6/the-standard-query-parser.html#TheStandardQueryParser-BooleanOperatorsSupportedbytheStandardQueryParser
// title:"The Right Way" AND text:go
// TODO: not, +, -
// TODO: grouping https://lucene.apache.org/solr/guide/6_6/the-standard-query-parser.html#TheStandardQueryParser-GroupingTermstoFormSub-Queries
func (q *StdQuery) And(field string, val string) *StdQuery {
	if q.q != "" {
		q.q += fmt.Sprintf(" AND %s:%s", field, val)
	} else {
		q.q = fmt.Sprintf("%s:%s", field, val)
	}
	return q
}

func (q *StdQuery) Or(field string, val string) *StdQuery {
	if q.q != "" {
		q.q += fmt.Sprintf(" OR %s:%s", field, val)
	} else {
		q.q = fmt.Sprintf("%s:%s", field, val)
	}
	return q
}

// Q sets the q parameter directly
func (q *StdQuery) Q(s string) *StdQuery {
	q.q = s
	return q
}

func (q *StdQuery) Encode() *url.Values {
	p := q.CommonQuery.Encode()
	p.Set("q", q.q)
	if q.df != "" {
		p.Set("df", q.df)
	}
	if len(q.facetFields) > 0 {
		p.Set("facet", "on")
		for _, field := range q.facetFields {
			p.Add("facet.field", field)
		}
	}
	return p
}
