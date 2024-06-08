package proxy

import (
	"bytes"
	"fmt"
	"github.com/rs/zerolog/log"
	"net"
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

	// TODO -> customize chunk size
	bytesRead := 1024
	for bytesRead == 1024 {
		log.Info().Msg(fmt.Sprintf("BYTES READ: %d", bytesRead))
		content, read, err := c.readChunkFromBuffer()
		if err != nil {
			log.Error().Msg("Cannot read content from buffer due to chunk reading error")
			return "", err
		}
		buffer.Write(content)
		bytesRead = read
	}

	return buffer.String(), nil
}

func (c *ConnectionBuffer) readChunkFromBuffer() ([]byte, int, error) {
	buf := make([]byte, 1024)
	bytesRead, err := c.conn.Read(buf)
	if err != nil {
		log.Error().Msg("Error reading content chunk from buffer")
		return buf, 0, err
	}
	return bytes.Trim(buf, "\u0000"), bytesRead, nil
}
