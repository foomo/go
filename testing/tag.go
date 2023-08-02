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
	testing.Short()
	if len(tags) == 0 {
		t.Skip("skipping untagged test")
		return
	}
	ok := true
	for _, v := range tags {
		if !getTag(v) {
			ok = false
		}
	}
	if !ok {
		t.Skipf("skipping test with tag: %s", tags)
	}
}

func getTag(v tag.Tag) bool {
	tags := os2.GetenvStrings(envTestTags, nil)

	// always return true if there are non tags defined
	if tags == nil {
		return true
	}

	// translate tags
	tagsMap := make(map[tag.Tag]bool, len(tags))
	for _, s := range tags {
		tagsMap[tag.Tag(strings.TrimPrefix(s, "-"))] = !strings.HasPrefix(s, "-")
	}

	if v, ok := tagsMap[v]; ok {
		return v
	}

	return false
}
