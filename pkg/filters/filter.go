// This package is used to describe a standard search filters for
// any search functionality.
package filters

type Filter interface {
	SortColumn() string
	SortDirection() string
	Limit() int
	Offset() int
}
