// nolint:wsl
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
}

func (suite *ClientTestSuite) SetupSuite() {
	go server.ListenAndServe()
	time.Sleep(1 * time.Second)
}

func (suite *ClientTestSuite) TearDownSuite() {
	server.Close()
}

func (suite *ClientTestSuite) TestNewClient() {
	client := NewClient()

	newClient := &Client{
		connection: nil,
	}
	assert.Exactly(suite.T(), newClient, client)
}

func (suite *ClientTestSuite) TestConnect() {
	client := NewClient()

	client.Connect()

	assert.NotNil(suite.T(), client.connection)
}

func (suite *ClientTestSuite) TestDisconnect() {
	client := NewClient()
	client.Connect()

	err := client.Disconnect()

	assert.NoError(suite.T(), err)
}

func (suite *ClientTestSuite) TestPublish() {
	client := NewClient()
	client.Connect()
	defer client.Disconnect()

	message := "Na feira de ontem não tinha tangerina."
	result := client.Publish(message)

	assert.True(suite.T(), result)
}

func (suite *ClientTestSuite) TestSend() {
	client := NewClient()
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
	client := NewClient()
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
	client := NewClient()
	client.Connect()
	defer client.Disconnect()

	result := client.Consumer()

	assert.True(suite.T(), result)
}

func (suite *ClientTestSuite) TestProducer() {
	client := NewClient()
	client.Connect()
	defer client.Disconnect()

	result := client.Producer()

	assert.True(suite.T(), result)
}

func TestClientTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(ClientTestSuite))
}
