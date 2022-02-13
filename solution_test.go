package solution

import "testing"

func TestGetMessage(t *testing.T) {
	t.Log("Check GetMessage function")
	expect := "Hello 🗺️ !"

	if expect != GetMessage() {
		t.Error("Wrong result")
	}
}
