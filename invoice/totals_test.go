package invoice

import "testing"

func TestTotal(t *testing.T) {
	s, tax, tot := Total([]Line{{"widget", 2, 1000}}, 2300)
	if s != 2000 || tax != 460 || tot != 2460 {
		t.Fatalf("got %d %d %d", s, tax, tot)
	}
}
