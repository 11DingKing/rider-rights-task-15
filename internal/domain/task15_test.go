package domain

import "testing"

func TestAuditPageSizeIsCapped(t *testing.T) {
	size, offset := NormalizeAuditPage(10000, -5)
	if size != 200 {
		t.Fatalf("page size was not capped: %d", size)
	}
	if offset != 0 {
		t.Fatalf("negative offset was not normalized: %d", offset)
	}
}
