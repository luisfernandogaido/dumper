package fs

import (
	"testing"
)

func TestZipFile(t *testing.T) {
	file := "../static/ect-20180413-075051.sql"
	if err := ZipFile(file, ""); err != nil {
		t.Fatal(err)
	}
}
