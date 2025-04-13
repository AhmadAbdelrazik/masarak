package filters

import (
	"errors"
	"strings"
)

type SQLFilters struct {
	page         int
	pageSize     int
	sort         string
	sortSafeList []string
}

type Opt func(f *SQLFilters) error

func NewSQLFilter(opts ...Opt) (*SQLFilters, error) {
	f := &SQLFilters{
		page:         1,
		pageSize:     20,
		sort:         "id",
		sortSafeList: []string{"id"},
	}

	for _, opt := range opts {
		if err := opt(f); err != nil {
			return nil, err
		}
	}

	return f, nil
}

func WithPage(page int) Opt {
	return func(f *SQLFilters) error {
		if page < 0 {
			return errors.New("page: must be greater than zero")
		}
		if page > 10_000_000 {
			return errors.New("page: must be a maximum of 10 million")
		}
		f.page = page
		return nil
	}
}

func WithPageSize(pageSize int) Opt {
	return func(f *SQLFilters) error {
		if pageSize < 0 {
			return errors.New("page_size: must be greater than zero")
		}
		if pageSize >= 100 {
			return errors.New("page_size: must be a maximum of 100")
		}
		f.pageSize = pageSize
		return nil
	}
}

func WithSort(sort, defaultValue string, sortSafeList []string) Opt {
	return func(f *SQLFilters) error {
		if !valueIn(defaultValue, sortSafeList) {
			panic("default value is not found in the sort safe list")
		}
		if sort == "" {
			sort = defaultValue
		}

		if !valueIn(strings.TrimPrefix(sort, "-"), sortSafeList) {
			return errors.New("sort: invalid sort value")
		}

		f.sort = sort
		f.sortSafeList = sortSafeList

		return nil
	}
}

func (f SQLFilters) SortColumn() string {
	return strings.TrimPrefix(f.sort, "-")
}

func (f SQLFilters) SortDirection() string {
	if strings.HasPrefix(f.sort, "-") {
		return "DESC"
	}

	return "ASC"
}

func (f SQLFilters) Limit() int {
	return f.pageSize
}

func (f SQLFilters) Offset() int {
	return (f.page - 1) * f.pageSize
}

func (f SQLFilters) Page() int {
	return f.page
}

func (f SQLFilters) PageSize() int {
	return f.pageSize
}

func valueIn(val string, values []string) bool {
	for _, v := range values {
		if v == val {
			return true
		}
	}

	return false
}
