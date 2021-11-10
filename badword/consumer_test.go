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
	cli := client.NewClient()
	cli.Connect()
	defer cli.Disconnect()

	cli.Producer()
	coolMessage := "Me gusta oír el mar."
	cli.Publish(coolMessage)

	niceMessage := "Arturo y Lucho son mis patas."
	cli.Publish(niceMessage)

	badMessage := "O pato feio foi a feira."
	cli.Publish(badMessage)

	Start()

	result := cli.Get()
	assert.Equal(suite.T(), coolMessage, result)

	result = cli.Get()
	assert.Equal(suite.T(), niceMessage, result)

	result = cli.Get()
	assert.Equal(suite.T(), "", result)
}

func TestConsumerTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(ConsumerTestSuite))
}
