// nolint:staticcheck
package server

import (
	"bufio"
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerRun(t *testing.T) {
	t.Parallel()

	address := getAddress()

	conn, err := net.Dial(PROTOCOL, address)
	defer conn.Close()

	assert.NoError(t, err)
}

func TestSendMessage(t *testing.T) {
	t.Parallel()

	address := getAddress()

	conn, err := net.Dial(PROTOCOL, address)
	defer conn.Close()

	assert.NoError(t, err)

	request := `{"operation": "GET","data":""}
`
	fmt.Fprintln(conn, request)

	response, _ := bufio.NewReader(conn).ReadString('\n')

	assert.Equal(t, `{"operation":"EMP","data":""}
`, response)
}

func TestGetAddress(t *testing.T) {
	t.Parallel()

	address := getAddress()

	expected := ":23023"

	assert.Equal(t, expected, address)
}
