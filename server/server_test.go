// nolint:staticcheck
package server

import (
	"bufio"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServerTestSuite struct {
	suite.Suite
}

func (suite *ServerTestSuite) SetupSuite() {
	go ListenAndServe()

	time.Sleep(1 * time.Second)
}

func (suite *ServerTestSuite) TearDownSuite() {
	go Close()

	time.Sleep(1 * time.Second)
}

func (suite *ServerTestSuite) TestServerRun() {
	address := GetAddress()

	conn, err := net.Dial(PROTOCOL, address)
	defer conn.Close()

	assert.NoError(suite.T(), err)
}

func (suite *ServerTestSuite) TestSendMessage() {
	address := GetAddress()

	conn, err := net.Dial(PROTOCOL, address)
	defer conn.Close()

	assert.NoError(suite.T(), err)

	request := `{"operation": "GET","data":""}
`
	fmt.Fprintln(conn, request)

	response, _ := bufio.NewReader(conn).ReadString('\n')

	assert.Equal(suite.T(), `{"operation":"EMP","data":""}
`, response)
}

func (suite *ServerTestSuite) TestGetAddress() {
	address := GetAddress()

	expected := "127.0.0.1:23023"

	assert.Equal(suite.T(), expected, address)
}

func TestServerTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(ServerTestSuite))
}
