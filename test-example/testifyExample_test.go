package test_example

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MyMockedObject struct {
	mock.Mock
}

func (m *MyMockedObject) DoSomething(number int) (bool, error) {

	args := m.Called(number)
	return args.Bool(0), args.Error(1)

}

func TestSomething(t *testing.T) {

	assertions := assert.New(t)

	// assertions equality
	assertions.Equal(123, 123, "they should be equal")

	// assertions inequality
	assertions.NotEqual(123, 456, "they should not be equal")

	var nilObject interface{}
	// assertions for nil (good for errors)
	assertions.Nil(nilObject)

	noNilObject := ""
	// assertions for not nil (good when you expect something)
	if assertions.NotNil(noNilObject) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assertions.Equal("", noNilObject)
	}

}

func TestMock(t *testing.T) {

	// create an instance of our test object
	testObj := new(MyMockedObject)

	// setup expectations
	testObj.On("DoSomething", mock.Anything).Return(true, nil)

	// call the code we are testing

	b, e := testObj.DoSomething(10)
	assert.Equal(t, true, b)
	assert.Equal(t, nil, e)

	// assert that the expectations were met
	testObj.AssertExpectations(t)

}

type ExampleTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *ExampleTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *ExampleTestSuite) TestExample() {
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
