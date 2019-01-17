package main
import "strings"
import "bytes"
import "fmt"

type MyStruct struct {
  id int
	text string
}


func main() {
	fmt.Println("Hello, World!")

  h1 := "GET / HTTP/1.1\r\n"
  h2 := "cache-control: no-cache\r\n"
  h3 := "Postman-Token: b86db5e0-316f-4b76-8e14-6309004dd961\r\n"
  h4 := "User-Agent: PostmanRuntime/7.4.0\r\n"
  h5 := "Accept: */*\r\n"
  h6 := "Host: www.hgking.xyz:8080\r\n"
  h7 := "accept-encoding: gzip, deflate\r\n"
  h8 := "Connection: keep-alive\r\n"
  h9 := "content-length: 129"

  header_string := join(h1, h2, h3, h4, h5, h6, h7, h8, h9)
  header_bytes := string_to_bytes(header_string)

  // these two together will be used to define each header
  //
  // header_index: position in the byte array where the header starts
  // header_length: length(of course) of the corresponding header
  //                starting at header_index[i]
  // result:
  var headers_index []int
  var headers_length []int
  var result int

  //for result != 0 {
    // read()
    result = get_headers(&headers_length, &headers_index, &header_bytes, result)
  //}


  if result > 0 {
    fmt.Println("We have not yet finished the request!")
    fmt.Printf("Starting at %v\n", result)

    //read more!!
    var str = "\r\nthis is just some random text to test if this works\r\n\r\n"
    header_string = join(header_string, str)
    header_bytes = string_to_bytes(header_string)

    result = get_headers(&headers_length, &headers_index, &header_bytes, result)
    if result > 0 {
      fmt.Println("We have not yet finished we must keep going!!")
    } else {
      fmt.Println("This bitch empty")
    }
  } else {
    fmt.Println("Our job here is done, time to check for the content-length")
  }

  //for each starting index of the headers
  for i, start := range headers_index {
    length := headers_length[i]
    fmt.Println(start, "is the start and is length", length)
    fmt.Printf("%q\n", header_bytes[start:start+length])
  }

  fmt.Println("Done!")
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

    //fmt.Printf("% x with length %d\n", slice, len(slice))

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
    // fall through if anything goes wrong
  }
  *index = -1
  return 0
}

func string_to_bytes(s string) []byte {
  var bytes = make([]byte, len(s))
  fmt.Printf("the type of header_bytes is %T\n", bytes)
  fmt.Printf("the type of header_string is %T\n", s)

  for i := 0; i < len(s); i++ {
    bytes[i] = byte(s[i])
  }

  return bytes
}


func newline()  {
  fmt.Println("")
}

func pop_rear_reference_byte(slice *[]byte) byte {
  len := len(*slice)
  val := (*slice)[len - 1]
  (*slice) = (*slice)[0:len - 1]
  return val
}

func pop_rear_value_byte(slice []byte) byte {
  len := len(slice)
  val := slice[len - 1]
  slice = slice[0:len - 1]
  return val
}

func pop_rear_value_int(slice []int) int {
  len := len(slice)
  val := slice[len - 1]
  slice = slice[0:len - 1]
  return val
}

func pop_rear_reference_int(slice *[]int) int {
  len := len(*slice)
  val := (*slice)[len - 1]
  (*slice) = (*slice)[0:len - 1]
  return val
}

func AddOneToEachElement(slice []byte) {
  for i := range slice {
    slice[i]++
  }
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

func check_crlf(s *[]byte, n int) bool {
  if len(*s) != 0  {
    if (*s)[n + 0] == '\x0d' && (*s)[n + 1] == '\x0a'  {
      //fmt.Println("This appears to be what we're looking for y'all")
      return true
    }
  }
  return false
}

//func function_name( [parameter list] ) [return_types]
func foo(i *int, j string) bool {
	fmt.Println("We are in the subroutine y'all")

	*i++

	fmt.Println(i)
	fmt.Println(j)

	return true
}
