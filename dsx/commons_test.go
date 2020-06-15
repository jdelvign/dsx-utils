package dsx

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExtractCharset(t *testing.T) {

	cs, err := extractCharset("../test/test.dsx")
	if err != nil {
		t.Error(err.Error())
	}

	actual := cs

	expected := "CP1252"

	// Using go-cmp instead of reflect.DeepEqual
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Errorf("%s:\tmismatch (-expected +actual):\n%s", "extractCharset", diff)
	}

}
