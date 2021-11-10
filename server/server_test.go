// nolint:staticcheck
package server

import (
	"bufio"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/YasminTeles/CatMQ/exchange"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServerTestSuite struct {
	suite.Suite
	server *Server
	port   string
}

func (suite *ServerTestSuite) SetupSuite() {
	suite.port = "23001"
	suite.server = NewServer(suite.port)

	go suite.server.ListenAndServe()

	time.Sleep(1 * time.Second)
}

func (suite *ServerTestSuite) TearDownSuite() {
	time.Sleep(1 * time.Second)
	go suite.server.Close()
}

func (suite *ServerTestSuite) TestNewServer() {
	server := NewServer(suite.port)

	newServer := &Server{
		port:       suite.port,
		connection: nil,
		listener:   nil,
		exch:       exchange.NewExchange(),
	}

	assert.Exactly(suite.T(), newServer, server)
}

func (suite *ServerTestSuite) TestNewServerDefault() {
	server := NewServerDefault()

	newServer := &Server{
		port:       ServerPort,
		connection: nil,
		listener:   nil,
		exch:       exchange.NewExchange(),
	}

	assert.Exactly(suite.T(), newServer, server)
}

func (suite *ServerTestSuite) TestServerRun() {
	address := suite.server.getAddress()

	conn, err := net.Dial(ServerProtocol, address)
	defer conn.Close()

	assert.NoError(suite.T(), err)
}

func (suite *ServerTestSuite) TestSendMessage() {
	address := suite.server.getAddress()

	conn, err := net.Dial(ServerProtocol, address)
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
	server := NewServerDefault()

	address := server.getAddress()

	expected := "127.0.0.1:23023"

	assert.Equal(suite.T(), expected, address)
}

func TestServerTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(ServerTestSuite))
}
