package component

import "testing"

func TestComponent(t *testing.T) {
	c := NewComponentType(testComponentData{}, nil)
	if c.String() != "testComponentData" {
		t.Errorf("expected name testComponentData, got %s", c.String())
	}
}

type testComponentData struct{}
