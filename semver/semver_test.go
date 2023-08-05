package semver

import (
	"errors"
	"reflect"
	"testing"
)

func TestParseSemver(t *testing.T) {
	testName := "TestParseSemver"
	testCases := []struct {
		input     string
		want      Semver
		mustError bool
	}{
		{"0.0.4", Semver{"0.0.4", 0, 0, 4, []string{}, []string{}, nil}, false},
		{"1.2.3", Semver{"1.2.3", 1, 2, 3, []string{}, []string{}, nil}, false},
		{"10.20.30", Semver{"10.20.30", 10, 20, 30, []string{}, []string{}, nil}, false},
		{"1.1.2-preRelease+meta", Semver{"1.1.2-preRelease+meta", 1, 1, 2, []string{"preRelease"}, []string{"meta"}, nil}, false},
		{"1.1.2+meta", Semver{"1.1.2+meta", 1, 1, 2, []string{}, []string{"meta"}, nil}, false},
		{"1.1.2+meta-valid", Semver{"1.1.2+meta-valid", 1, 1, 2, []string{}, []string{"meta-valid"}, nil}, false},
		{"1.0.0-alpha", Semver{"1.0.0-alpha", 1, 0, 0, []string{"alpha"}, []string{}, nil}, false},
		{"1.0.0-beta", Semver{"1.0.0-beta", 1, 0, 0, []string{"beta"}, []string{}, nil}, false},
		{"1.0.0-alpha.beta", Semver{"1.0.0-alpha.beta", 1, 0, 0, []string{"alpha", "beta"}, []string{}, nil}, false},
		{"1.0.0-alpha.beta.1", Semver{"1.0.0-alpha.beta.1", 1, 0, 0, []string{"alpha", "beta", "1"}, []string{}, nil}, false},
		{"1.0.0-alpha.1", Semver{"1.0.0-alpha.1", 1, 0, 0, []string{"alpha", "1"}, []string{}, nil}, false},
		{"1.0.0-alpha0.valid", Semver{"1.0.0-alpha0.valid", 1, 0, 0, []string{"alpha0", "valid"}, []string{}, nil}, false},
		{"1.0.0-alpha.0valid", Semver{"1.0.0-alpha.0valid", 1, 0, 0, []string{"alpha", "0valid"}, []string{}, nil}, false},
		{"1.0.0-alpha-a.b-c-somethinglong+build.1-aef.1-its-okay", Semver{"1.0.0-alpha-a.b-c-somethinglong+build.1-aef.1-its-okay", 1, 0, 0, []string{"alpha-a", "b-c-somethinglong"}, []string{"build", "1-aef", "1-its-okay"}, nil}, false},
		{"1.0.0-rc.1+build.1", Semver{"1.0.0-rc.1+build.1", 1, 0, 0, []string{"rc", "1"}, []string{"build", "1"}, nil}, false},
		{"2.0.0-rc.1+build.123", Semver{"2.0.0-rc.1+build.123", 2, 0, 0, []string{"rc", "1"}, []string{"build", "123"}, nil}, false},
		{"1.2.3-beta", Semver{"1.2.3-beta", 1, 2, 3, []string{"beta"}, []string{}, nil}, false},
		{"10.2.3-DEV-SNAPSHOT", Semver{"10.2.3-DEV-SNAPSHOT", 10, 2, 3, []string{"DEV-SNAPSHOT"}, []string{}, nil}, false},
		{"1.2.3-SNAPSHOT-123", Semver{"1.2.3-SNAPSHOT-123", 1, 2, 3, []string{"SNAPSHOT-123"}, []string{}, nil}, false},
		{"1.0.0", Semver{"1.0.0", 1, 0, 0, []string{}, []string{}, nil}, false},
		{"2.0.0", Semver{"2.0.0", 2, 0, 0, []string{}, []string{}, nil}, false},
		{"1.1.7", Semver{"1.1.7", 1, 1, 7, []string{}, []string{}, nil}, false},
		{"2.0.0+build.1848", Semver{"2.0.0+build.1848", 2, 0, 0, []string{}, []string{"build", "1848"}, nil}, false},
		{"2.0.1-alpha.1227", Semver{"2.0.1-alpha.1227", 2, 0, 1, []string{"alpha", "1227"}, []string{}, nil}, false},
		{"1.0.0-alpha+beta", Semver{"1.0.0-alpha+beta", 1, 0, 0, []string{"alpha"}, []string{"beta"}, nil}, false},
		{"1.2.3----RC-SNAPSHOT.12.9.1--.12+788", Semver{"1.2.3----RC-SNAPSHOT.12.9.1--.12+788", 1, 2, 3, []string{"---RC-SNAPSHOT", "12", "9", "1--", "12"}, []string{"788"}, nil}, false},
		{"1.2.3----R-S.12.9.1--.12+meta", Semver{"1.2.3----R-S.12.9.1--.12+meta", 1, 2, 3, []string{"---R-S", "12", "9", "1--", "12"}, []string{"meta"}, nil}, false},
		{"1.2.3----RC-SNAPSHOT.12.9.1--.12", Semver{"1.2.3----RC-SNAPSHOT.12.9.1--.12", 1, 2, 3, []string{"---RC-SNAPSHOT", "12", "9", "1--", "12"}, []string{}, nil}, false},
		{"1.0.0+0.build.1-rc.10000aaa-kk-0.1", Semver{"1.0.0+0.build.1-rc.10000aaa-kk-0.1", 1, 0, 0, []string{}, []string{"0", "build", "1-rc", "10000aaa-kk-0", "1"}, nil}, false},
		{"1.0.0-0A.is.legal", Semver{"1.0.0-0A.is.legal", 1, 0, 0, []string{"0A", "is", "legal"}, []string{}, nil}, false},
		{input: "Invalid Semantic Versions", mustError: true},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got := ParseSemver(tc.input)
			if tc.mustError {
				if !errors.Is(got.Error(), ErrInvalidVersionString) {
					t.Fatalf(`%s("%s"): must error`, testName, tc.input)
				}
				return
			}
			if got.String() != tc.want.versionStr {
				t.Fatalf(`%s("%s"): versionStr = %s, want %s`, testName, tc.input, got.String(), tc.want.versionStr)
			}
			if got.Major() != tc.want.major {
				t.Fatalf(`%s("%s"): major = %d, want %d`, testName, tc.input, got.Major(), tc.want.major)
			}
			if got.Minor() != tc.want.minor {
				t.Fatalf(`%s("%s"): minor = %d, want %d`, testName, tc.input, got.Minor(), tc.want.minor)
			}
			if got.Patch() != tc.want.patch {
				t.Fatalf(`%s("%s"): patch = %d, want %d`, testName, tc.input, got.Patch(), tc.want.patch)
			}
			if !reflect.DeepEqual(got.PreRelease(), tc.want.preRelease) {
				t.Fatalf(`%s("%s"): preRelease = %#v, want %#v`, testName, tc.input, got.PreRelease(), tc.want.preRelease)
			}
			if !reflect.DeepEqual(got.Build(), tc.want.build) {
				t.Fatalf(`%s("%s"): build = %#v, want %#v`, testName, tc.input, got.Build(), tc.want.build)
			}
		})
	}
}

func TestSemver_Compare(t *testing.T) {
	testName := "TestSemver_Compare"
	testCases := []struct {
		name      string
		me, other Semver
		want      int
	}{
		{
			"Different majors",
			Semver{"", 1, 0, 0, []string{}, []string{}, nil},
			Semver{"", 2, 0, 0, []string{}, []string{}, nil},
			-1,
		},
		{
			"Different minors",
			Semver{"", 2, 0, 0, []string{}, []string{}, nil},
			Semver{"", 2, 1, 0, []string{}, []string{}, nil},
			-1,
		},
		{
			"Different patches",
			Semver{"", 2, 1, 0, []string{}, []string{}, nil},
			Semver{"", 2, 1, 1, []string{}, []string{}, nil},
			-1,
		},
		{
			"Pre-release < normal",
			Semver{"", 1, 0, 0, []string{}, []string{}, nil},
			Semver{"", 1, 0, 0, []string{"alpha"}, []string{}, nil},
			1,
		},
		{
			"1.0.0-alpha < 1.0.0-alpha.1",
			Semver{"", 1, 0, 0, []string{"alpha"}, []string{}, nil},
			Semver{"", 1, 0, 0, []string{"alpha", "1"}, []string{}, nil},
			-1,
		},
		{
			"1.0.0-alpha.1 < 1.0.0-alpha.beta",
			Semver{"", 1, 0, 0, []string{"alpha", "1"}, []string{}, nil},
			Semver{"", 1, 0, 0, []string{"alpha", "beta"}, []string{}, nil},
			-1,
		},
		{
			"1.0.0-alpha.beta < 1.0.0-beta",
			Semver{"", 1, 0, 0, []string{"alpha", "beta"}, []string{}, nil},
			Semver{"", 1, 0, 0, []string{"beta"}, []string{}, nil},
			-1,
		},
		{
			"1.0.0-beta < 1.0.0-beta.2",
			Semver{"", 1, 0, 0, []string{"beta"}, []string{}, nil},
			Semver{"", 1, 0, 0, []string{"beta", "2"}, []string{}, nil},
			-1,
		},
		{
			"1.0.0-beta.2 < 1.0.0-beta.11",
			Semver{"", 1, 0, 0, []string{"beta", "2"}, []string{}, nil},
			Semver{"", 1, 0, 0, []string{"beta", "11"}, []string{}, nil},
			-1,
		},
		{
			"1.0.0-beta.11 < 1.0.0-rc.1",
			Semver{"", 1, 0, 0, []string{"beta", "11"}, []string{}, nil},
			Semver{"", 1, 0, 0, []string{"rc", "1"}, []string{}, nil},
			-1,
		},
		{
			"1.0.0-rc.1 < 1.0.0",
			Semver{"", 1, 0, 0, []string{"rc", "1"}, []string{}, nil},
			Semver{"", 1, 0, 0, []string{}, []string{}, nil},
			-1,
		},
		{
			"Build-meta does not affect comparison",
			Semver{"", 1, 0, 0, []string{}, []string{"met1a"}, nil},
			Semver{"", 1, 0, 0, []string{}, []string{"meta2"}, nil},
			0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.me.Compare(tc.other)
			if got != tc.want {
				t.Fatalf(`%s/%s: got %d, want %d`, testName, tc.name, got, tc.want)
			}
			ngot := tc.other.Compare(tc.me)
			if ngot != 0-tc.want {
				t.Fatalf(`%s/%s: got %d, want %d`, testName, tc.name, ngot, 0-tc.want)
			}
		})
	}
}
