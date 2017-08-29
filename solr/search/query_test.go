package search

import "testing"

func TestCommonQuery_Start(t *testing.T) {
	c := CommonQuery{}
	c.Start(10).Rows(10)
	t.Log(c.Encode().Encode())
}

func TestCommonQuery_IncludeField(t *testing.T) {
	c := CommonQuery{}
	c.IncludeField("name").IncludeField("created")
	t.Log(c.Encode().Encode())
}

func TestCommonQuery_SortBy(t *testing.T) {
	c := CommonQuery{}
	c.SortBy("inStock", SortOrderDesc).
		SortBy("price", SortOrderDesc)
	t.Log(c.Encode())
	t.Log(c.Encode().Encode()) // sort=inStock+desc%2Cprice+desc
}

func TestCommonQuery_Encode(t *testing.T) {
	c := CommonQuery{}
	c.IncludeField("price").
		SortBy("price", SortOrderDesc).
		SortBy("inStock", SortOrderAsc).
		Start(10).
		Rows(10)
	t.Log(c.Encode().Encode())
}

func TestStdQuery_Encode(t *testing.T) {
	s := StdQuery{}
	s.IncludeField("name").
		IncludeField("description").
		SortBy("rating", SortOrderDesc).
		Start(10).
		Rows(10)
	s.And("name", "docker*").Or("description", "container")
	s.FacetField("name").FacetField("rating")
	t.Log(s.Encode().Encode())
}
