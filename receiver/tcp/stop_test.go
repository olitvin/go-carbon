package tcp

import (
	"net"
	"testing"
	"time"

	"github.com/olitvin/go-carbon/points"
	"github.com/olitvin/go-carbon/receiver"
	"github.com/stretchr/testify/assert"
)

func TestStopTCP(t *testing.T) {
	assert := assert.New(t)

	addr, err := net.ResolveTCPAddr("tcp", ":0")
	assert.NoError(err)

	for i := 0; i < 10; i++ {
		r, err := receiver.New("tcp", map[string]interface{}{
			"protocol": "tcp",
			"listen":   addr.String(),
		},
			nil,
		)
		assert.NoError(err)
		addr = r.(*TCP).Addr().(*net.TCPAddr) // listen same port in next iteration
		r.Stop()
	}
}

func TestStopPickle(t *testing.T) {
	assert := assert.New(t)

	addr, err := net.ResolveTCPAddr("tcp", ":0")
	assert.NoError(err)

	for i := 0; i < 10; i++ {
		r, err := receiver.New("pickle", map[string]interface{}{
			"protocol": "pickle",
			"listen":   addr.String(),
		},
			nil,
		)
		assert.NoError(err)
		addr = r.(*TCP).Addr().(*net.TCPAddr) // listen same port in next iteration
		r.Stop()
	}
}

func TestStopConnectedTCP(t *testing.T) {
	test := newTCPTestCase(t, "tcp")
	defer test.Finish()

	ch := test.rcvChan
	test.Send("hello.world 42.15 1422698155\n")
	time.Sleep(10 * time.Millisecond)

	select {
	case msg := <-ch:
		test.Eq(msg, points.OnePoint("hello.world", 42.15, 1422698155))
	default:
		t.Fatalf("Message #0 not received")
	}

	test.receiver.Stop()
	test.receiver = nil
	time.Sleep(10 * time.Millisecond)

	test.Send("metric.name -72.11 1422698155\n")
	time.Sleep(10 * time.Millisecond)

	select {
	case <-ch:
		t.Fatalf("Message #0 received")
	default:
	}
}
