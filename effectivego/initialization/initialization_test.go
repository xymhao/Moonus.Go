package initialization_test

import (
	"Moonus.Go/effectivego/initialization"
	"testing"
)

func TestPrint(t *testing.T) {
	initialization.User = "1"
	initialization.Print()
}
