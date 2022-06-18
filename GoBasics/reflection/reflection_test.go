package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q, but it didn't", haystack, needle)
	}
}

func TestWalk(t *testing.T) {
	t.Run("functions", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{23, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{23, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("separate for maps, cause of order", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}
		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
	t.Run("Other tests", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				"struct with one string field",
				struct {
					Name string
				}{"Chris"},
				[]string{"Chris"},
			},
			{
				"struct with two string fields",
				struct {
					Name string
					City string
				}{"Chris", "London"},
				[]string{"Chris", "London"},
			},
			{
				"struct with non string field",
				struct {
					Name string
					Age  int
				}{"Chris", 32},
				[]string{"Chris"},
			},
			{
				"nested fields in struct",
				Person{
					"Chris",
					Profile{33, "London"},
				},
				[]string{"Chris", "London"},
			},
			{
				"pointer to things",
				&Person{
					"Chris",
					Profile{33, "London"},
				},
				[]string{"Chris", "London"},
			},
			{
				"slices",
				[]Profile{
					{33, "London"},
					{53, "Lisbon"},
				},
				[]string{"London", "Lisbon"},
			},
			{
				"arrays",
				[2]Profile{
					{33, "London"},
					{23, "Zakopane"},
				},
				[]string{"London", "Zakopane"},
			},
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var got []string
				walk(test.Input, func(input string) {
					got = append(got, input)
				})

				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("got %v, wanted %v", got, test.ExpectedCalls)
				}
			})
		}
	})

}
