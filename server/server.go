package main

import "fmt"
import "net"
import "io"

func main()  {
  create_server("0.0.0.0:8080")
}


func create_server(network string)  {
  /* ln is a Listener */
  ln, err := net.Listen("tcp", network)
  if err != nil {
    fmt.Println("Error creating listener!")
  } else {
    fmt.Printf("%T is the type\n", ln)
    for {
      conn, err := ln.Accept()
      if err != nil {
        fmt.Println("Error starting connection!")
      } else {
        fmt.Printf("%T is the type\n", conn)
        go handle_connection(conn)
      }
    }
    ln.Close()
  }
}

func handle_connection(conn net.Conn)  {
  buf := make([]byte, 4096)
  tmp := make([]byte, 64)     // using small tmo buffer for demonstrating

  var err error
  n := 1

  for err != io.EOF || n != 0 {
    n, err = conn.Read(tmp)
    if err != nil {
      fmt.Println("read error:", err)
    }
    fmt.Println("got", n, "bytes.")
    buf = append(buf, tmp[:n]...)
  }


  fmt.Println("total size:", len(buf))
  fmt.Println(string(buf))
  fmt.Println("DONE")
}
