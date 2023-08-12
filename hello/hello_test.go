package hello

import "testing"

func TestSayHello(t *testing.T) {

	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items:  []string{"Kevin"},
			result: "Hello, Kevin!",
		},
		{
			items:  []string{"Kevin", "Ngenoh"},
			result: "Hello, Kevin, Ngenoh!",
		},
	}

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("wanted %s (%v), got %s", st.result, st.items, s)
		}
	}

}
