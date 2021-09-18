package mock

import (
	. "Moonus.Go/base/mock/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestMockFoo(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockFoo(ctrl)

	// Asserts that the first and only call to Bar() is passed 99.
	// Anything else will fail.
	m.
		EXPECT().
		Bar(gomock.Eq(99)).
		Return(101)

	m.
		EXPECT().
		Bar(gomock.Any()).
		Return(102)

	sut(m)
}
