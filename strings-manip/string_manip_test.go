package stringsmanip

import (
	"fmt"
	"testing"
)

func TestCompare(t *testing.T) {
	returned := CompareString("rhyn", "Rhyn")
	expected := 0

	if returned != expected {
		t.Errorf("expected %d but got %d instead", expected, returned)
	}
}

func ExampleCompare() {
	fmt.Println(CompareString("test", "test"))
	// Output: 0
}
