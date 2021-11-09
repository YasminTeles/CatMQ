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

func (suite *ClientTestSuite) TestConnect() {
	result := Connect()

	assert.True(suite.T(), result)
}

func (suite *ClientTestSuite) TestDisconnect() {
	Connect()

	result := Disconnect()

	assert.True(suite.T(), result)
}

func (suite *ClientTestSuite) TestPublish() {
	Connect()
	defer Disconnect()

	message := "Na feira de ontem não tinha tangerina."
	result := Publish(message)

	assert.True(suite.T(), result)
}

func (suite *ClientTestSuite) TestSend() {
	Connect()
	defer Disconnect()

	request := `{"operation": "GET","data":""}
`

	result, err := send(request)

	assert.NoError(suite.T(), err)

	response := `{"operation":"EMP","data":""}
`
	assert.Equal(suite.T(), response, result)
}

func (suite *ClientTestSuite) TestGet() {
	Connect()
	defer Disconnect()

	Producer()
	Publish("Me gusta oír el mar.")
	Publish("Na feira de ontem não tinha tangerina.")
	Consumer()

	result := Get()

	assert.Equal(suite.T(), "Me gusta oír el mar.", result)
}

func (suite *ClientTestSuite) TestConsumer() {
	Connect()
	defer Disconnect()

	result := Consumer()

	assert.True(suite.T(), result)
}

func (suite *ClientTestSuite) TestProducer() {
	Connect()
	defer Disconnect()

	result := Producer()

	assert.True(suite.T(), result)
}

func TestClientTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(ClientTestSuite))
}
