package main
import "strings"
import "fmt"

type MyStruct struct {
    id int
	text string
}

func main() {
	/* this is a comment I hope */

	fmt.Println("Hello, World!")

    // h1 := byte("GET / HTTP/1.1\r\n")
    // h2 := byte("cache-control: no-cache\r\n")
    // h3 := byte("Postman-Token: b86db5e0-316f-4b76-8e14-6309004dd961\r\n")
    // h4 := byte("User-Agent: PostmanRuntime/7.4.0\r\n")
    // h5 := byte("Accept: */*\r\n")
    // h6 := byte("Host: www.hgking.xyz:8080\r\n")
    // h7 := byte("accept-encoding: gzip, deflate\r\n")
    // h8 := byte("Connection: keep-alive\r\n")
    //
    // header_string := join(h1, h2, h3, h4, h5, h6, h7, h8)
    // n := 0
    //
    // for header := test_parsing(&n, header_string); n > 1; {
    //     fmt.Println(header, "so far n is", n)
    //     //headers = append(headers, header)
    // }
    const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

    fmt.Println("Println:")
    fmt.Println(sample)

    fmt.Println("Byte loop:")
    for i := 0; i < len(sample); i++ {
        fmt.Printf("%x ", sample[i])
    }
    fmt.Printf("\n")

    fmt.Println("Printf with %x:")
    fmt.Printf("%x\n", sample)

    fmt.Println("Printf with % x:")
    fmt.Printf("% x\n", sample)

    fmt.Println("Printf with %q:")
    fmt.Printf("%q\n", sample)

    fmt.Println("Printf with %+q:")
    fmt.Printf("%+q\n", sample)

}

// func test_parsing(index *int, str []byte) []byte {
//     slice := str[*index:]
//     crlfPos := strings.Index(slice, "\r\n")
//     fmt.Println(crlfPos, "is what we got")
//     newSlice := str[*index:crlfPos]
//     *index += crlfPos + 2
//     return newSlice
// }


func newline()  {
    fmt.Println("")
}

func pop_rear_reference(slice *[]byte) byte {
  len := len(*slice)
  val := (*slice)[len - 1]
  (*slice) = (*slice)[0:len - 1]
  return val
}

func pop_rear_value(slice []byte) byte {
  len := len(slice)
  val := slice[len - 1]
  slice = slice[0:len - 1]
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

//func function_name( [parameter list] ) [return_types]
func foo(i *int, j string) bool {
	fmt.Println("We are in the subroutine y'all")

	*i++

	fmt.Println(i)
	fmt.Println(j)

	return true
}
