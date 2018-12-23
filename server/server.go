package main

import (
    "fmt"
    "io"
    "net"
    "os"
    "strings"
)

func main()  {
    argv := os.Args
    argc := len(argv)

    if argc != 3 {
        fmt.Println("Error: Proper Usage -", argv[0], "<host> <port>")
    } else {
        create_server("0.0.0.0:8080")
    }
}

func create_server(network string)  {
    /* returns a Listener */
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
            go handle_connection(conn)  // start a go routine to handle conn
        }
    }

    ln.Close()  // gotta close the listener
    }
}

/*
*
*
*
*
*/
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


func test_parsing(index *int, str string) []string {
    slice := str[*index:]
    crlfPos := strings.Index(slice, "\r\n")
    fmt.Println(crlfPos, "is what we got")
    newSlice := str[*index:crlfPos]
    *index += crlfPos + 2
    return newSlice
}

func join_by_reference(sb *strings.Builder, strs ...string) {
    for _, str := range strs  {
        (*sb).WriteString(str)
    }
}

func join(strs ...string) string {
    var sb strings.Builder
    for _, str := range strs  {
        sb.WriteString(str)
    }
    return sb.String()
}
