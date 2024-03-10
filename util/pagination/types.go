package pagination

import "github.com/forderation/null"

type PaginationInput struct {
	Page     null.Int
	PageSize null.Int
}

type PaginationOutput struct {
	Page      int64
	PageSize  int64
	PageCount int64
	TotalData int64
}
