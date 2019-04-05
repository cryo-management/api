package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type SchemaTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *SchemaTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5

	//db connect
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *SchemaTestSuite) Test001Example() {
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}

func (suite *SchemaTestSuite) Test002Example() {
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSchemaTestSuite(t *testing.T) {
	suite.Run(t, new(SchemaTestSuite))
}
