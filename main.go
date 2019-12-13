package main

import (
	"log"
	"time"

	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{
		Name:        "/dev/ttyAMA0",
		Baud:        2400,
		ReadTimeout: 5 * time.Second,
		Size:        8,
		Parity:      'N',
		StopBits:    2,
	}
	p, err := serial.OpenPort(c)
	if err != nil {
		log.Println(err)
	}
	for {
		_, err = p.Write([]byte("lasdfb42bri37349hfr3f9hf37fy"))
		buf := make([]byte, 128)
		n, err := p.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%q", buf[:n])
	}
}
