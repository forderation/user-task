package db

import (
	"github.com/forderation/user-task/util/pagination"
	"gorm.io/gen"
)

func ApplyPaginationQuery(query gen.DO, input pagination.PaginationInput) gen.DO {
	if input.Page.Valid && input.PageSize.Valid {
		offset := pagination.GetOffsetValue((input.Page.Int64), input.PageSize.Int64)
		dao := query.
			Offset(int(offset)).
			Limit(int(input.PageSize.Int64))
		return *dao.(*gen.DO)
	}
	return query
}
