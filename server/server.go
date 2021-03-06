package main

import (
  "bytes"
  "fmt"
  "net"
  "os"
  "strings"
)

func main()  {
    argv := os.Args
    argc := len(argv)

    // GOTTA FIX THIS JUNK AND MAKE IT WORK A LITTLE NICER
    if argc != 3 {
        fmt.Println("Error: Proper Usage -", argv[0], "<host> <port>")
    } else {
        create_server("0.0.0.0:8080")
    }
}

func create_server(network string)  {
  // returns a Listener
  ln, err := net.Listen("tcp", network)

  if err != nil {
    fmt.Println("Error creating listener!")
  } else {

    for {
      // listener waits to accept a connection
      // pretty sure its not a busy wait which is cool af
      conn, err := ln.Accept()

      if err != nil {
        fmt.Println("Error starting connection!")
      } else {
        //fmt.Printf("%T is the type of conn\n", conn)
        go handle_connection(conn)  // start a go routine to handle conn
      }
    }

    ln.Close()  // gotta close the listener
  }
}

/*
NOTES:
1) handle connection will
    - read in the request
      - first it needs the headers
      - then decides whether or not there is a body
      - if so get the body
    - parse the request
      - determine which page it is asking for
      - pull any query parameters out of the url
    - send the response
*/
func handle_connection(conn net.Conn)  {

  // First we need to get the data from the request
  var headers_index []int
  var headers_length []int
  header_bytes := parse_request(conn, &headers_index, &headers_length)

  //for each starting index of the headers
  for i, start := range headers_index {
    length := headers_length[i]
    //fmt.Println(start, "is the start and is length", length)
    fmt.Printf("%q\n", buf[start:start+length])
  }

  fmt.Println("DONE")
}

// conn: current connection that has been opened(maybe needs to be reference)
// headers_index: position in the byte array where the header starts
// headers_length: length(of course) of the corresponding header
//                starting at headers_index[i]
// returns string of bytes that is at least the header section of the request
//         and on failure it returns empty array(probably not)
func parse_request(conn net.Conn, headers_index *[]int, headers_length *[]int) []byte {
  buf := make([]byte, 0)
  tmp := make([]byte, 64)     // using small tmo buffer for demonstrating
  result := 0

  for {
    // read exactly n <= 64 bytes into temp
    n, err := conn.Read(tmp)
    if err != nil {
      fmt.Println("read error:", err)
      result = 0 // flag to break
    } else {
      buf = append(buf, tmp[:n]...)
      // concat tmp and keep reading headers, result is the length
      result = get_headers(headers_length, headers_index, &buf, result)
    }
    // not proud of the break but golang has forced my hand
    // apparently this is how you make a 'dowhile'
    if result == 0 { break }
  }
  return buf
}

// headers_index: position in the byte array where the header starts
// headers_length: length(of course) of the corresponding header
//                starting at headers_index[i]
// header_bytes: slice of the request string, maybe not complete
// index: starting index on the byte string
//
// returns 0 on complete headers, index of last header on failure
//
func get_headers(headers_length *[]int, headers_index *[]int, header_bytes *[]byte, index int) int {
  // while there hasn't been a crlf and index is positive(flags an error)
  for index > -1 && !check_crlf(header_bytes, index) {
    *headers_index = append(*headers_index, index)
    //read from current index to the next crlf
    length := header_parsing(&index, header_bytes)
    *headers_length = append(*headers_length, length)
  }
  // we dont want the last stuff in our headers info
  if index < 0  {
    pop_rear_reference_int(headers_length)
    return pop_rear_reference_int(headers_index)
  }
  return 0
}

// index: pointer used to slice the underlying array of bytes
//        gets updated by reference to the next header's starting index
// str: ptr to full string of bytes that the HTTP request resides in
//
// returns the length of the partitioned header
//
func header_parsing(index *int, str *[]byte) int {
  // just your typical error checking
  if *index > -1 && len(*str) != 0 {
    // slice from the specific index till the end
    slice := (*str)[*index:]
    // make sure the slice still has length
    if len(slice) != 0 {
      // grabbing the end of a header
      // bytes.Index returns -1 on miss
      i := bytes.Index(slice, []byte("\x0d\x0a"))
      if i > 0  {
        // index of element before \r\n so +2 offset
        i += 2
        *index += i
        return i
      }
    }
  }
  // fall through if anything goes wrong
  *index = -1
  return 0
}

func string_to_bytes(s string) []byte {
  var bytes = make([]byte, len(s))
  for i := 0; i < len(s); i++ {
    bytes[i] = byte(s[i])
  }
  return bytes
}

func pop_rear_reference_int(slice *[]int) int {
  len := len(*slice)
  val := (*slice)[len - 1]
  (*slice) = (*slice)[0:len - 1]
  return val
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

// check if the current slice starts with '\r\n'
func check_crlf(s *[]byte, n int) bool {
  if len(*s) != 0 && len(*s) >= n + 2 {
    if (*s)[n + 0] == '\x0d' && (*s)[n + 1] == '\x0a'  {
      return true
    }
  }
  return false
}

func newline()  {
  fmt.Println("")
}
