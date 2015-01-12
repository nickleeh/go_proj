package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

// var s_end = [3]string{".", "!", "?"}

func main() {
	b, err := ioutil.ReadFile("i_have_a_dream.txt")
	if err != nil {
		panic(err)
	}

	s := string(b)
	r := regexp.MustCompile("(\r\n)+")
	counter := 1

	repl := func(match string) string {
		p_num := counter
		counter++
		return fmt.Sprintf("%s [%d] ", match, p_num)
	}

	fmt.Println(r.ReplaceAllStringFunc(s, repl))
}

// func ReadLine(filename string) {
//     f, err := os.Open(filename)
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//     defer f.Close()
//     r, err := bufio.NewReaderSize(f, 4*1024)
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//     line, isPrefix, err := r.ReadLine()
//     for err == nil && !isPrefix {
//         s := string(line)
//         fmt.Println(s)
//         line, isPrefix, err = r.ReadLine()
//     }
//     if isPrefix {
//         fmt.Println(os.NewError("buffer size to small"))
//         return
//     }
//     if err != os.EOF {
//         fmt.Println(err)
//         return
//     }
// }

// func ReadString(filename string) {
//     f, err := os.Open(filename)
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//     defer f.Close()
//     r := bufio.NewReader(f)
//     line, err := r.ReadString('\n')
//     for err == nil {
//         fmt.Print(line)
//         line, err = r.ReadString('\n')
//     }
//     if err != os.EOF {
//         fmt.Println(err)
//         return
//     }
// }

// func main() {
//     filename := `i_have_a_dream.txt`
//     ReadLine(filename)
//     // ReadString(filename)
// }
