package main

import (
      "github.com/tarm/serial"
      "log"
)

func main() {
      c := &serial.Config{Name: "/dev/ttyS0", Baud: 115200}
      s, err := serial.OpenPort(c)
      if err != nil {
              log.Fatal(111, err)
      }

	  out := []byte{0xff,20,30}	
	  log.Printf("--> %q", out)

      n, err := s.Write(out)
      if err != nil {
              log.Fatal(222, err)
      }

      buf := make([]byte, 125)
      n, err = s.Read(buf)
      if err != nil {
              log.Fatal(333, err)
      }
      log.Printf("%q <--", buf[:n])
      
      s.Close()
}
