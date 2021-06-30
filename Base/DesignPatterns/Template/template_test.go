package Template

import "testing"

func TestTemplate(t *testing.T) {
	cook := ChineseCook{}
	DoCook(cook)

	foreignCook := ForeignCook{}
	DoCook(foreignCook)
}
