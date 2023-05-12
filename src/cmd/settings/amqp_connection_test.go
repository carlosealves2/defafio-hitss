package settings

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAmqpConn_ConnectWithDocker(t *testing.T) {
	amqpConn, err := NewAmqpConn()
	assert.NoError(t, err)
	defer amqpConn.Conn.Close()
	assert.False(t, amqpConn.Conn.IsClosed())
}

func TestAmqpConn_GivenAnError_WhenConnect_UsingInvalidCredentials(t *testing.T) {
	os.Setenv("RABBITMQ_DSN", "amqp://test:test@localhost:5674/")

	_, err := NewAmqpConn()
	assert.Error(t, err)
}
