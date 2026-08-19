package domain

func NormalizeAuditPage(size, offset int) (int, int) {
	if size <= 0 {
		size = 50
	}
	if offset < 0 {
		offset = 0
	}
	return size, offset
}
