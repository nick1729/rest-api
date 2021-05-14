package main

import (
	"testing"
)

// Tests UUID validation
func TestIsUUID(t *testing.T) {

	type tTest struct {
		u    string
		expd bool
	}

	var (
		tests = []tTest{
			{"58657d7b-f442-411a-a630-d79d46150692", true},
			{"25281c42-832c-4982-8bc9-de9d660c956c", true},
			{"ea20bcf8-fb71-41a3-8c31-c68b8d4361d9", true},
			{"133f8fe0-b67b-468e-97c5-72194d6070cf", true},
			{"c21efa44-a56f-42e3-9de6-b512c63f3756", true},

			{"314a6aa4-59e3-4fa5-bee5-d6ec1995370a", true},
			{"fdd91101-d301-40b0-89ce-aba8d3ae1b03", true},
			{"bb7d4430-3467-48c4-9eca-20ae298c794b", true},
			{"fec45ac3-2ca6-49d6-9200-e327e6dd88d1", true},
			{"90a268c9-245e-46d8-af24-a612449ee58e", true},

			{"b53171f16_9099-460e-a051-2a6fae25021", false},
			{"d64f2da-6e829-4f4d-9b08-88a8631db5ce", false},
			{"*00a07e8-8fc-42ec-9b6f--9d885c9af1be", false},
			{"d0wc4679b-84f7-4938-9f05-f54a0fbc71f", false},
			{"01edff5-0660-4bfe-a9a5-02396954670a6", false},

			{"90a268c9-245e5-46d8-af24-a612449ee58", false},
			{"cc70ee09-63d9-40e2-b5b0-034c5%8c9528", false},
			{"12$1c48c-3af5-4296-b4aa-a21e92657319", false},
			{"fec45ac3-2ca6-49d6-92-0-e327e6dd88d1", false},
			{"bb7d<430-3467-48c4-9eca-20ae298c794b", false},
		}
		got bool
	)

	for _, pair := range tests {
		got = isUUID(pair.u)
		if got != pair.expd {
			t.Error("UUID:", pair.u, "Expected:", pair.expd, "got:", got)
		}
	}
}

// Tests name validation
func TestIsName(t *testing.T) {

	type tTest struct {
		n    string
		expd bool
	}

	var (
		tests = []tTest{
			{"Gilbert", true},
			{"Jerome", true},
			{"Theron", true},
			{"Rafael", true},
			{"Carolina", true},

			{"Jessie", true},
			{"Veronica", true},
			{"Cristina", true},
			{"Sylvia", true},
			{"Kari", true},

			{"Charles9", false},
			{"_Zachary", false},
			{"Darius.J", false},
			{"Reginald!", false},
			{"F.Marvin", false},

			{"CarlY", false},
			{"Kitty<3", false},
			{"Maur33n", false},
			{"Carmella-22", false},
			{"@Sheree", false},
		}
		got bool
	)

	for _, pair := range tests {
		got = isName(pair.n)
		if got != pair.expd {
			t.Error("UUID:", pair.n, "Expected:", pair.expd, "got:", got)
		}
	}
}

// Tests email validation
func TestIsEmail(t *testing.T) {

	type tTest struct {
		e    string
		expd bool
	}

	var (
		tests = []tTest{
			{"qwerty@gmail.com", true},
			{"q2w7erty@ma-il.com", true},
			{"qwerty@mail.io", true},
			{"qw_er-ty@gmail.ru", true},
			{"qwerty@yahoo.com", true},

			{"cuevas-victoria@hotmail.com", true},
			{"william_09@gmail.com", true},
			{"fbutler@murillo-sawyer.com.ru", true},
			{"wardryan@anderson.net", true},
			{"kayla.buckley@yahoo.com", true},

			{"gsullivan@gmail..com", false},
			{"qwe&rty@ya.ru", false},
			{"qwertygmail.com", false},
			{"@qwerty@gmail.commm", false},
			{"qwerty@mail,com", false},

			{"renee(shaw)@ryan-baker.com", false},
			{"tclay@gmail$com", false},
			{"gardner%alvin@miller.net", false},
			{"=grant79@george.com", false},
			{"ant*honymiller@yahoo.com", false},
		}
		got bool
	)

	for _, pair := range tests {
		got = isEmail(pair.e)
		if got != pair.expd {
			t.Error("Email:", pair.e, "Expected:", pair.expd, "got:", got)
		}
	}
}
