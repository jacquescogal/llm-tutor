package repository

func pageOffset(page, pageSize uint32) uint32 {
	if page == 0 {
		page = 1
	}
	return (page - 1) * pageSize
}