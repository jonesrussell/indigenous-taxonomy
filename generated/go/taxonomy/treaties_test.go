package taxonomy_test

import (
	"slices"
	"testing"

	taxonomy "github.com/jonesrussell/indigenous-taxonomy/generated/go/taxonomy"
)

func TestTreatyCount(t *testing.T) {
	t.Helper()
	if got := len(taxonomy.AllTreaties); got != 11 {
		t.Errorf("AllTreaties length = %d, want 11", got)
	}
}

func TestTreatySlugsPresent(t *testing.T) {
	t.Helper()
	want := []string{
		"treaty:1", "treaty:2", "treaty:3", "treaty:4", "treaty:5",
		"treaty:6", "treaty:7", "treaty:8", "treaty:9", "treaty:10", "treaty:11",
	}
	// Convert AllTreaties to []string for slices.Contains usage.
	got := make([]string, len(taxonomy.AllTreaties))
	for i, tr := range taxonomy.AllTreaties {
		got[i] = string(tr)
	}
	for _, slug := range want {
		if !slices.Contains(got, slug) {
			t.Errorf("AllTreaties missing expected slug %q", slug)
		}
	}
}

func TestIsValidTreaty(t *testing.T) {
	t.Helper()
	validCases := []string{
		"treaty:1", "treaty:2", "treaty:3", "treaty:4", "treaty:5",
		"treaty:6", "treaty:7", "treaty:8", "treaty:9", "treaty:10", "treaty:11",
	}
	for _, slug := range validCases {
		if !taxonomy.IsValidTreaty(slug) {
			t.Errorf("IsValidTreaty(%q) = false, want true", slug)
		}
	}

	invalidCases := []string{"treaty:0", "treaty:12", "not-a-treaty", "", "Treaty:1", "treaty:"}
	for _, slug := range invalidCases {
		if taxonomy.IsValidTreaty(slug) {
			t.Errorf("IsValidTreaty(%q) = true, want false", slug)
		}
	}
}

func TestTreatyConstantValues(t *testing.T) {
	t.Helper()
	cases := []struct {
		constant taxonomy.Treaty
		slug     string
	}{
		{taxonomy.TreatyArea1, "treaty:1"},
		{taxonomy.TreatyArea2, "treaty:2"},
		{taxonomy.TreatyArea3, "treaty:3"},
		{taxonomy.TreatyArea4, "treaty:4"},
		{taxonomy.TreatyArea5, "treaty:5"},
		{taxonomy.TreatyArea6, "treaty:6"},
		{taxonomy.TreatyArea7, "treaty:7"},
		{taxonomy.TreatyArea8, "treaty:8"},
		{taxonomy.TreatyArea9, "treaty:9"},
		{taxonomy.TreatyArea10, "treaty:10"},
		{taxonomy.TreatyArea11, "treaty:11"},
	}
	for _, tc := range cases {
		if got := string(tc.constant); got != tc.slug {
			t.Errorf("constant value = %q, want %q", got, tc.slug)
		}
	}
}
