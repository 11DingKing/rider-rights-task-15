package domain

const (
	DefaultAuditPageSize = 50
	MaxAuditPageSize     = 200
)

func NormalizeAuditPage(size, offset int) (int, int) {
	if size <= 0 {
		size = DefaultAuditPageSize
	}
	if size > MaxAuditPageSize {
		size = MaxAuditPageSize
	}
	if offset < 0 {
		offset = 0
	}
	return size, offset
}
