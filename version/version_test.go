package version_test

import (
	"fmt"
	"testing"

	"github.com/lxygwqf9527/demo-api/version"
)

func TestVer(t *testing.T) {
	fmt.Println(version.FullVersion())
}
