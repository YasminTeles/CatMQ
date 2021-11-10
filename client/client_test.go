// nolint:wsl,errcheck
package client

import (
	"testing"
	"time"

	"github.com/YasminTeles/CatMQ/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ClientTestSuite struct {
	suite.Suite
	Server *server.Server
	port   string
}

func (suite *ClientTestSuite) SetupSuite() {
	suite.port = "23002"
	suite.Server = server.NewServer(suite.port)

	go suite.Server.ListenAndServe()

	time.Sleep(1 * time.Second)
}

func (suite *ClientTestSuite) TearDownSuite() {
	time.Sleep(1 * time.Second)
	go suite.Server.Close()
}

func (suite *ClientTestSuite) TestNewClient() {
	client := NewClient(suite.port)

	newClient := &Client{
		port:       suite.port,
		connection: nil,
	}
	assert.Exactly(suite.T(), newClient, client)
}

func (suite *ClientTestSuite) TestNewClientDefault() {
	client := NewClientDefault()

	newClient := &Client{
		port:       server.ServerPort,
		connection: nil,
	}
	assert.Exactly(suite.T(), newClient, client)
}

func (suite *ClientTestSuite) TestConnect() {
	client := NewClient(suite.port)

	client.Connect()
	defer client.Disconnect()

	assert.NotNil(suite.T(), client.connection)
}

func (suite *ClientTestSuite) TestDisconnect() {
	client := NewClient(suite.port)
	client.Connect()

	err := client.Disconnect()

	assert.NoError(suite.T(), err)
}

func (suite *ClientTestSuite) TestPublish() {
	client := NewClient(suite.port)
	client.Connect()
	defer client.Disconnect()

	message := "Na feira de ontem não tinha tangerina."
	result := client.Publish(message)

	assert.True(suite.T(), result)
}

func (suite *ClientTestSuite) TestSend() {
	client := NewClient(suite.port)
	client.Connect()
	defer client.Disconnect()

	request := `{"operation": "GET","data":""}
`

	result, err := client.send(request)

	assert.NoError(suite.T(), err)

	response := `{"operation":"EMP","data":""}
`
	assert.Equal(suite.T(), response, result)
}

func (suite *ClientTestSuite) TestGet() {
	client := NewClient(suite.port)
	client.Connect()
	defer client.Disconnect()

	client.Producer()
	client.Publish("Me gusta oír el mar.")
	client.Publish("Na feira de ontem não tinha tangerina.")
	client.Consumer()

	result := client.Get()

	assert.Equal(suite.T(), "Me gusta oír el mar.", result)
}

func (suite *ClientTestSuite) TestConsumer() {
	client := NewClient(suite.port)
	client.Connect()
	defer client.Disconnect()

	result := client.Consumer()

	assert.True(suite.T(), result)
}

func (suite *ClientTestSuite) TestProducer() {
	client := NewClient(suite.port)
	client.Connect()
	defer client.Disconnect()

	result := client.Producer()

	assert.True(suite.T(), result)
}

func TestClientTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(ClientTestSuite))
}
