package domain

// MaxAuditPageSize bounds the number of audit rows a single query may read,
// so an oversized page_size cannot drag the service down in one request.
const MaxAuditPageSize = 200

// DefaultAuditPageSize is applied when the caller omits or supplies a
// non-positive page size.
const DefaultAuditPageSize = 50

// NormalizeAuditPage is the single place that clamps audit pagination: it
// supplies a default size, enforces an upper limit, and guards against a
// negative offset.
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
