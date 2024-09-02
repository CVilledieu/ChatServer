package server

import (
	"fmt"
	"log"
	"net"
)

func Server() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	request := make([]byte, 1024)
	_, err := conn.Read(request)
	if err != nil {
		log.Println(err)
		return
	}

	handleInput(request)

	response := []byte("Hello, World!")
	_, err = conn.Write(response)
	if err != nil {
		log.Println(err)
		return
	}
}

package server

import (
    "crypto/sha1"
    "encoding/base64"
    "fmt"
    "io"
    "log"
    "net"
    "net/http"
    "strings"
)

func Server() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatal(err)
    }
    defer listener.Close()

    fmt.Println("Server is listening on port 8080")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error: ", err)
            continue
        }

        go handleClient(conn)
    }
}

func handleClient(conn net.Conn) {
    defer conn.Close()

    // Perform WebSocket handshake
    buf := make([]byte, 1024)
    n, err := conn.Read(buf)
    if err != nil {
        log.Println("Read error:", err)
        return
    }

    req, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(buf[:n])))
    if err != nil {
        log.Println("Request error:", err)
        return
    }

    if req.Header.Get("Upgrade") != "websocket" {
        log.Println("Not a WebSocket connection")
        return
    }

    key := req.Header.Get("Sec-WebSocket-Key")
    acceptKey := computeAcceptKey(key)

    response := "HTTP/1.1 101 Switching Protocols\r\n" +
        "Upgrade: websocket\r\n" +
        "Connection: Upgrade\r\n" +
        "Sec-WebSocket-Accept: " + acceptKey + "\r\n\r\n"

    _, err = conn.Write([]byte(response))
    if err != nil {
        log.Println("Write error:", err)
        return
    }

    // Handle WebSocket frames
    for {
        message, err := readMessage(conn)
        if err != nil {
            log.Println("Read message error:", err)
            break
        }

        log.Printf("Received: %s", message)

        err = writeMessage(conn, message)
        if err != nil {
            log.Println("Write message error:", err)
            break
        }
    }
}

func computeAcceptKey(key string) string {
    h := sha1.New()
    io.WriteString(h, key+"258EAFA5-E914-47DA-95CA-C5AB0DC85B11")
    return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func readMessage(conn net.Conn) ([]byte, error) {
    header := make([]byte, 2)
    _, err := io.ReadFull(conn, header)
    if err != nil {
        return nil, err
    }

    fin := header[0] & 0x80
    opcode := header[0] & 0x0F
    if fin == 0 || opcode != 1 {
        return nil, fmt.Errorf("unsupported frame")
    }

    mask := header[1] & 0x80
    payloadLen := int(header[1] & 0x7F)

    if payloadLen == 126 {
        extended := make([]byte, 2)
        _, err = io.ReadFull(conn, extended)
        if err != nil {
            return nil, err
        }
        payloadLen = int(extended[0])<<8 | int(extended[1])
    } else if payloadLen == 127 {
        extended := make([]byte, 8)
        _, err = io.ReadFull(conn, extended)
        if err != nil {
            return nil, err
        }
        payloadLen = int(extended[0])<<56 | int(extended[1])<<48 | int(extended[2])<<40 | int(extended[3])<<32 |
            int(extended[4])<<24 | int(extended[5])<<16 | int(extended[6])<<8 | int(extended[7])
    }

    maskKey := make([]byte, 4)
    if mask != 0 {
        _, err = io.ReadFull(conn, maskKey)
        if err != nil {
            return nil, err
        }
    }

    payload := make([]byte, payloadLen)
    _, err = io.ReadFull(conn, payload)
    if err != nil {
        return nil, err
    }

    if mask != 0 {
        for i := 0; i < payloadLen; i++ {
            payload[i] ^= maskKey[i%4]
        }
    }

    return payload, nil
}

func writeMessage(conn net.Conn, message []byte) error {
    header := []byte{0x81, byte(len(message))}
    _, err := conn.Write(header)
    if err != nil {
        return err
	}

    _, err = conn.Write(message)
    return err
}