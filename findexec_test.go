package findexec

import (
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var osPathSep = string(os.PathListSeparator)

func TestFindAbsExisting(t *testing.T) {
	if runtime.GOOS == "windows" {
		assert.Equal(t, "C:\\Windows\\system32\\cmd.exe",
			Find("C:\\Windows\\system32\\cmd.exe", ""))
	} else {
		assert.Equal(t, "/bin/sh", Find("/bin/sh", ""))
	}
}

func TestFindExpectExist(t *testing.T) {
	if runtime.GOOS == "windows" {
		assert.Equal(t, "C:\\Windows\\system32\\cmd.exe", Find("cmd.exe", ""))
	} else {
		assert.Equal(t, "/bin/sh", Find("sh", ""))
	}
}

func TestFindExpectNotExisting(t *testing.T) {
	assert.Equal(t, "", Find("something_that_does_not_exist", ""))
}

func TestFindExpectExistWithSuppliedPath(t *testing.T) {
	if runtime.GOOS == "windows" {
		assert.Equal(t, "C:\\Windows\\system32\\cmd.exe",
			Find("cmd.exe", "C:\\Windows\\system32"))
	} else {
		assert.Equal(t, "/bin/sh", Find("sh", "/bin"))
	}
}

func TestFindExpectExistWithMultipleSuppliedPaths(t *testing.T) {
	if runtime.GOOS == "windows" {
		paths := strings.Join([]string{"C:\\invalidOne", "D:\\invalidTwo", "C:\\Windows\\system32"}, osPathSep)
		assert.Equal(t, "C:\\Windows\\system32\\cmd.exe", Find("cmd.exe", paths))
	} else {
		paths := strings.Join([]string{"/invalidOne", "/invalidTwo", "/bin"}, osPathSep)
		assert.Equal(t, "/bin/sh", Find("sh", paths))
	}
}

func TestFindExpectNotExistingWithSuppliedPath(t *testing.T) {
	if runtime.GOOS == "windows" {
		assert.Equal(t, "", Find("cmd.exe", "C:\\invalidOne"))
	} else {
		assert.Equal(t, "", Find("sh", "/invalid"))
	}
}

func TestFindExpectNotExistingWithMultipleSuppliedPaths(t *testing.T) {
	if runtime.GOOS == "windows" {
		paths := strings.Join([]string{"C:\\invalidOne", "C:\\invalidTwo"}, osPathSep)
		assert.Equal(t, "", Find("cmd.exe", paths))
	} else {
		paths := strings.Join([]string{"/invalidOne", "/invalidTwo", "/invalidThree"}, osPathSep)
		assert.Equal(t, "", Find("sh", paths))
	}
}

func TestFindWithMissingOSPATHs(t *testing.T) {
	_ = os.Setenv("PATH", "")
	if runtime.GOOS == "windows" {
		assert.Equal(t, "", Find("cmd.exe", ""))
	} else {
		assert.Equal(t, "", Find("sh", ""))
	}
}
