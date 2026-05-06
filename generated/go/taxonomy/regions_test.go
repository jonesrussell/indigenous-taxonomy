package taxonomy_test

import (
	"slices"
	"testing"

	taxonomy "github.com/jonesrussell/indigenous-taxonomy/generated/go/taxonomy"
)

func regionSlugs() []string {
	got := make([]string, len(taxonomy.AllRegions))
	for i, r := range taxonomy.AllRegions {
		got[i] = string(r)
	}
	return got
}

func TestAllRegionsCount(t *testing.T) {
	t.Helper()
	// 16 original + 15 new (cities + community samples) = 31
	if got := len(taxonomy.AllRegions); got != 31 {
		t.Errorf("AllRegions length = %d, want 31", got)
	}
}

func TestRegionSlugsPresent(t *testing.T) {
	t.Helper()
	want := []string{
		// original provinces/regions
		"canada",
		"canada:british-columbia",
		"canada:alberta",
		"canada:saskatchewan",
		"canada:manitoba",
		"canada:manitoba:southern",
		"canada:ontario",
		"canada:ontario:northern",
		"canada:ontario:north-shore-huron",
		"canada:ontario:southern",
		"canada:quebec",
		"canada:atlantic",
		"canada:north",
		"canada:north:yukon",
		"canada:north:nwt",
		"canada:north:nunavut",
		// T005: cities
		"canada:manitoba:winnipeg",
		"canada:ontario:toronto",
		"canada:ontario:ottawa",
		"canada:british-columbia:vancouver",
		"canada:alberta:calgary",
		"canada:saskatchewan:saskatoon",
		// T006: First Nations community samples
		"canada:manitoba:sagkeeng-fn",
		"canada:british-columbia:musqueam-fn",
		"canada:alberta:siksika-fn",
		"canada:saskatchewan:beardy-okimasis-fn",
		"canada:ontario:mississaugas-of-the-credit-fn",
		"canada:ontario:fort-william-fn",
		"canada:quebec:kahnawake-fn",
		"canada:atlantic:membertou-fn",
		"canada:north:kluane-fn",
	}
	got := regionSlugs()
	for _, slug := range want {
		if !slices.Contains(got, slug) {
			t.Errorf("AllRegions missing expected slug %q", slug)
		}
	}
}

func TestIsValidRegion(t *testing.T) {
	t.Helper()
	validCases := []string{
		"canada",
		"canada:manitoba",
		"canada:manitoba:winnipeg",
		"canada:manitoba:sagkeeng-fn",
		"canada:ontario:toronto",
		"canada:ontario:ottawa",
		"canada:british-columbia:vancouver",
		"canada:british-columbia:musqueam-fn",
		"canada:alberta:calgary",
		"canada:alberta:siksika-fn",
		"canada:saskatchewan:saskatoon",
		"canada:north:nunavut",
	}
	for _, slug := range validCases {
		if !taxonomy.IsValidRegion(slug) {
			t.Errorf("IsValidRegion(%q) = false, want true", slug)
		}
	}

	invalidCases := []string{
		"",
		"Canada",
		"canada:unknown",
		"canada:ontario:london",
		"not-a-region",
		"canada:",
		":ontario",
	}
	for _, slug := range invalidCases {
		if taxonomy.IsValidRegion(slug) {
			t.Errorf("IsValidRegion(%q) = true, want false", slug)
		}
	}
}

func TestRegionConstantValues(t *testing.T) {
	t.Helper()
	cases := []struct {
		constant taxonomy.Region
		slug     string
	}{
		// cities
		{taxonomy.RegionCanadaManitobaWinnipeg, "canada:manitoba:winnipeg"},
		{taxonomy.RegionCanadaOntarioToronto, "canada:ontario:toronto"},
		{taxonomy.RegionCanadaOntarioOttawa, "canada:ontario:ottawa"},
		{taxonomy.RegionCanadaBritishColumbiaVancouver, "canada:british-columbia:vancouver"},
		{taxonomy.RegionCanadaAlbertaCalgary, "canada:alberta:calgary"},
		{taxonomy.RegionCanadaSaskatchewanSaskatoon, "canada:saskatchewan:saskatoon"},
		// community samples
		{taxonomy.RegionCanadaManitobaSagkeengFn, "canada:manitoba:sagkeeng-fn"},
		{taxonomy.RegionCanadaBritishColumbiaMusqueamFn, "canada:british-columbia:musqueam-fn"},
		{taxonomy.RegionCanadaAlbertaSiksikaFn, "canada:alberta:siksika-fn"},
		{taxonomy.RegionCanadaSaskatchewanBeardyOkimasisFn, "canada:saskatchewan:beardy-okimasis-fn"},
		{taxonomy.RegionCanadaOntarioMississaugasOfTheCreditFn, "canada:ontario:mississaugas-of-the-credit-fn"},
		{taxonomy.RegionCanadaOntarioFortWilliamFn, "canada:ontario:fort-william-fn"},
		{taxonomy.RegionCanadaQuebecKahnawakeFn, "canada:quebec:kahnawake-fn"},
		{taxonomy.RegionCanadaAtlanticMembertouFn, "canada:atlantic:membertou-fn"},
		{taxonomy.RegionCanadaNorthKluaneFn, "canada:north:kluane-fn"},
	}
	for _, tc := range cases {
		if got := string(tc.constant); got != tc.slug {
			t.Errorf("constant value = %q, want %q", got, tc.slug)
		}
	}
}
