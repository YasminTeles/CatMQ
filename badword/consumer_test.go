package badword

import (
	"testing"
	"time"

	"github.com/YasminTeles/CatMQ/client"
	"github.com/YasminTeles/CatMQ/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConsumerTestSuite struct {
	suite.Suite
}

func (suite *ConsumerTestSuite) SetupSuite() {
	go server.ListenAndServe()
	time.Sleep(1 * time.Second)
}

func (suite *ConsumerTestSuite) TearDownSuite() {
	server.Close()
}

func (suite *ConsumerTestSuite) TestStart() {
	client.Connect()
	// defer client.Disconnect()

	client.Producer()
	coolMessage := "Me gusta oír el mar."
	client.Publish(coolMessage)

	niceMessage := "Arturo y Lucho son mis patas."
	client.Publish(niceMessage)

	badMessage := "O pato feio foi a feira."
	client.Publish(badMessage)

	Start()

	result := client.Get()
	assert.Equal(suite.T(), coolMessage, result)

	result = client.Get()
	assert.Equal(suite.T(), niceMessage, result)

	result = client.Get()
	assert.Equal(suite.T(), "", result)
}

func TestConsumerTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(ConsumerTestSuite))
}
