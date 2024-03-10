package pagination

func GetOffsetValue(page int64, pageSize int64) int64 {
	offset := int64(0)
	if page > 0 {
		offset = (page - 1) * pageSize
	}
	return offset
}

func GetPageCount(pageSize int64, totalData int64) int64 {
	pageCount := int64(1)
	if pageSize > 0 {
		totalData := totalData
		if pageSize >= totalData {
			return pageCount
		}
		if totalData%pageSize == 0 {
			pageCount = totalData / pageSize
		} else {
			pageCount = (totalData / pageSize) + 1
		}
	}
	return pageCount
}

func CreatePaginationOutput(input PaginationInput, totalData int64) PaginationOutput {
	pageCount := GetPageCount(input.PageSize.Int64, totalData)
	return PaginationOutput{
		Page:      input.Page.Int64,
		PageSize:  input.PageSize.Int64,
		TotalData: totalData,
		PageCount: pageCount,
	}
}
