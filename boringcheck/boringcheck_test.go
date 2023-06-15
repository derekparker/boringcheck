package boringcheck

import "testing"

const testdatapath = "../testdata"

func TestBoringCheck(t *testing.T) {
	fns := BoringCheck(testdatapath)
	if len(fns) != 3 {
		t.Errorf("Expected 3 functions, got %d", len(fns))
	}
	for _, fn := range fns {
		if fn != "main.HasNoBoringCheck" && fn != "main.HasWrongBoringCheck" && fn != "main.HasWrongBoringCheck2" {
			t.Errorf("Expected function name to be HasNoBoringCheck or HasWrongBoringCheck, got %s", fn)
		}
	}
}
