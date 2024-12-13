package hello

import (
	"testing"
)

// func TestSayHello(t *testing.T) {

// 	want := "Hello, test!"

// 	got := Say([]string{"test"})

// 	if want != got {
// 		t.Errorf("wanted %s got %s", want, got)
// 	}
// }

func TestSayHello(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, World!",
		},
		{
			items:  []string{"Aman"},
			result: "Hello, Aman!",
		},
		{
			items:  []string{"Aman", "Vishal", "Rahul"},
			result: "Hello, Aman, Vishal, Rahul!",
		},
	}

	for _, st := range subtests {

		if s := Say(st.items); s != st.result {
			t.Errorf("wanted %s (%v) got %s", st.result, st.items, s)
		}
	}

}
