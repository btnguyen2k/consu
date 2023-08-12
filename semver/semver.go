package semver

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	reSemver = regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

	// ErrInvalidVersionString is returned when the input string is not a valid semantic version.
	ErrInvalidVersionString = errors.New("invalid semantic version string")
)

// ParseSemver parses a string into a semantic version.
func ParseSemver(input string) Semver {
	input = strings.TrimSpace(input)
	result := Semver{versionStr: input, preRelease: make(PreRelease, 0), build: make(BuildMeta, 0)}
	matches := reSemver.FindAllStringSubmatch(input, -1)
	if len(matches) == 0 {
		result.err = ErrInvalidVersionString
	}
	if result.err == nil {
		result.major, result.err = strconv.Atoi(matches[0][1])
	}
	if result.err == nil {
		result.minor, result.err = strconv.Atoi(matches[0][2])
	}
	if result.err == nil {
		result.patch, result.err = strconv.Atoi(matches[0][3])
	}
	if result.err == nil {
		if matches[0][4] != "" {
			result.preRelease = strings.Split(matches[0][4], ".")
		}
		if matches[0][5] != "" {
			result.build = strings.Split(matches[0][5], ".")
		}
	}
	return result
}

// PreRelease represents the "preRelease" extension of a semantic version.
type PreRelease []string

// BuildMeta represents the "build metadata" extension of a semantic version.
type BuildMeta []string

// Semver represents a semantic version number (see https://semver.org/).
type Semver struct {
	versionStr string
	major      int
	minor      int
	patch      int
	preRelease PreRelease
	build      BuildMeta
	err        error
}

func normalize(value int) int {
	if value < 0 {
		return -1
	}
	if value > 0 {
		return 1
	}
	return 0
}

func toNumberIfAllDigits(input string) int {
	for _, c := range input {
		if c < '0' || c > '9' {
			return -1
		}
	}
	result, _ := strconv.Atoi(input)
	return result
}

// Compare performs semantic version comparison against another Semver.
// This function returns -1 if the receiver is less than the other Semver, 0 if they are equal, and 1 if the receiver
// is greater than the other Semver.
func (s Semver) Compare(other Semver) int {
	if v := s.major - other.major; v != 0 {
		return normalize(v)
	}
	if v := s.minor - other.minor; v != 0 {
		return normalize(v)
	}
	if v := s.patch - other.patch; v != 0 {
		return normalize(v)
	}
	if len(s.preRelease) == 0 && len(other.preRelease) > 0 {
		return 1
	}
	if len(s.preRelease) > 0 && len(other.preRelease) == 0 {
		return -1
	}
	for i := 0; i < len(s.preRelease) && i < len(other.preRelease); i++ {
		v1, v2 := toNumberIfAllDigits(s.preRelease[i]), toNumberIfAllDigits(other.preRelease[i])
		if v1 >= 0 && v2 >= 0 {
			return normalize(v1 - v2)
		}
		if v1 < 0 && v2 >= 0 {
			return 1
		}
		if v1 >= 0 && v2 < 0 {
			return -1
		}
		if v := strings.Compare(s.preRelease[i], other.preRelease[i]); v != 0 {
			return v
		}
	}
	if v := len(s.preRelease) - len(other.preRelease); v != 0 {
		return normalize(v)
	}
	return 0
}

// Error returns error if any.
func (s Semver) Error() error {
	return s.err
}

// String returns the string representation of the semantic version.
//
// @Available since v0.1.1
func (s Semver) String() string {
	return s.versionStr
}

// Major returns the major version number.
func (s Semver) Major() int {
	return s.major
}

// Minor returns the minor version number.
func (s Semver) Minor() int {
	return s.minor
}

// Patch returns the patch version number.
func (s Semver) Patch() int {
	return s.patch
}

// PreRelease returns the preRelease version number.
func (s Semver) PreRelease() PreRelease {
	result := make(PreRelease, len(s.preRelease))
	copy(result, s.preRelease)
	return result
}

// Build returns the build version number.
func (s Semver) Build() BuildMeta {
	result := make(BuildMeta, len(s.build))
	copy(result, s.build)
	return result
}

// IncMajor increments the major version number and return new Semver instance.
//
// @available since <<VERSION>>
func (s Semver) IncMajor() Semver {
	return Semver{
		versionStr: fmt.Sprintf("%d.%d.%d", s.major+1, 0, 0),
		major:      s.major + 1, minor: 0, patch: 0,
		preRelease: make(PreRelease, 0),
		build:      make(BuildMeta, 0),
	}
}

// IncMinor increments the minor version number and return new Semver instance.
//
// @available since <<VERSION>>
func (s Semver) IncMinor() Semver {
	return Semver{
		versionStr: fmt.Sprintf("%d.%d.%d", s.major, s.minor+1, 0),
		major:      s.major, minor: s.minor + 1, patch: 0,
		preRelease: make(PreRelease, 0),
		build:      make(BuildMeta, 0),
	}
}

// IncPatch increments the patch version number and return new Semver instance.
//
// @available since <<VERSION>>
func (s Semver) IncPatch() Semver {
	return Semver{
		versionStr: fmt.Sprintf("%d.%d.%d", s.major, s.minor, s.patch+1),
		major:      s.major, minor: s.minor, patch: s.patch + 1,
		preRelease: make(PreRelease, 0),
		build:      make(BuildMeta, 0),
	}
}

// MakeRelease removes the preRelease and return new Semver instance.
//
// @available since <<VERSION>>
func (s Semver) MakeRelease() Semver {
	return Semver{
		versionStr: fmt.Sprintf("%d.%d.%d", s.major, s.minor, s.patch),
		major:      s.major, minor: s.minor, patch: s.patch,
		preRelease: make(PreRelease, 0),
		build:      make(BuildMeta, 0),
	}
}
