package testing

import (
	"strings"
	"testing"

	osx "github.com/foomo/go/os"
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

	envTags := osx.GetenvStringSlice(envTestTags, nil)
	// always return false if there are non tags defined
	if envTags == nil {
		return false
	}

	// translate tags
	allExclude := true
	envTagsMap := make(map[tag.Tag]bool, len(tags))
	for _, s := range envTags {
		include := !strings.HasPrefix(s, "-")
		envTagsMap[tag.Tag(strings.TrimPrefix(s, "-"))] = include
		if include {
			allExclude = false
		}
	}

	var (
		anyFound   bool
		anyInclude bool
		anyExclude bool
	)
	for _, v := range tags {
		include, ok := envTagsMap[v]
		if ok {
			anyFound = true
			if include {
				anyInclude = true
			} else {
				anyExclude = true
			}
		}
	}

	if allExclude && !anyFound {
		return false
	}

	return !(anyInclude && !anyExclude)
}
