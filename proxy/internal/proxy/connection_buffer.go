package proxy

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"net"
	"time"
)

type ConnectionBuffer struct {
	conn net.Conn
}

func (c *ConnectionBuffer) Close() error {
	err := c.conn.Close()
	if err != nil {
		log.Error().Msg("Error closing connection buffer")
		return err
	}
	return nil
}

func (c *ConnectionBuffer) Write(content []byte) (int, error) {
	bytesWritten, err := c.conn.Write(content)

	if err != nil {
		log.Error().Msg("Error writing content to connection buffer")
		return bytesWritten, err
	}

	log.Info().Msg(fmt.Sprintf("Wrote %d bytes to the buffer", bytesWritten))
	return bytesWritten, nil
}

func (c *ConnectionBuffer) ReadAll() (string, error) {
	var buffer bytes.Buffer
	rdr := bufio.NewWriter(&buffer)

	err := c.conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
	if err != nil {
		log.Error().Msg("Error setting read deadline for connection")
		return "", err
	}

	b := make([]byte, 10024)
	_, err = io.CopyBuffer(rdr, c.conn, b)
	if err != nil {
		log.Error().Msg("Error copying data from connection buffer")
	}

	return buffer.String(), nil
}
