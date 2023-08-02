package testing

import (
	"strings"
	"testing"

	os2 "github.com/foomo/go/os"
	"github.com/foomo/go/testing/tag"
)

const envTestTags = "GO_TEST_TAGS"

// Tags defines the tags that the test should run under.
//
// For example:
//
//	func TestDemo(t *testing.T) {
//	  testing.Tags(t, tag.Integration, tag.Short)
//	}
//
// Results being run with:
//
//   - no tags
//   - `GO_TEST_TAGS=fast`
//   - `GO_TEST_TAGS=integration`
//   - `GO_TEST_TAGS=fast,integration`
//
// But would be skipped with:
//
//   - `GO_TEST_TAGS=-fast`
//   - `GO_TEST_TAGS=-integration`
//   - `GO_TEST_TAGS=fast,-integration`
func Tags(t *testing.T, tags ...tag.Tag) {
	t.Helper()

	// always skip if no tags are provided so it can be used as block tests
	if len(tags) == 0 {
		t.Skip("skipping untagged test")
		return
	}

	if skipTags(tags) {
		t.Skipf("skipping test with tag: %s", tags)
	}
}

func skipTags(tags []tag.Tag) bool {
	// always skip if no tags are provided so it can be used as block tests
	if len(tags) == 0 {
		return true
	}

	envTags := os2.GetenvStrings(envTestTags, nil)
	// always return false if there are non tags defined
	if envTags == nil {
		return false
	}

	// translate tags
	envTagsMap := make(map[tag.Tag]bool, len(tags))
	for _, s := range envTags {
		envTagsMap[tag.Tag(strings.TrimPrefix(s, "-"))] = !strings.HasPrefix(s, "-")
	}

	var (
		anyInclude bool
		anyExclude bool
	)
	for _, v := range tags {
		if v, ok := envTagsMap[v]; ok && v {
			anyInclude = true
		} else if ok && !v {
			anyExclude = true
		}
	}
	return !(anyInclude && !anyExclude)
}
