package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	//error handeling in golang log the err
	// 	panic(err)
	// }

	// fileinfo, err := f.Stat()
	// if err != nil {
	// 	//error handeling in golang log the err
	// 	panic(err)
	// }

	// fmt.Println("file Name", fileinfo.Name())
	// fmt.Println("file modified time", fileinfo.ModTime())
	// fmt.Println("file current fir", fileinfo.IsDir())
	// fmt.Println("file permissions", fileinfo.Mode())
	// fmt.Println("file size", fileinfo.Size())

	//read file

	// f, err := os.Open("example.txt") //open file example.txt
	// if err != nil {
	// 	panic(err)
	// }
	// //must close the file you open
	// defer f.Close()

	// buf := make([]byte, 15)
	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }
	// for i := 0; i < len(buf); i++ {
	// 	println("data", d, string(buf[i]))
	// }

	// f, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(f))

	//read folders

	// dir, err := os.Open("..")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()

	// fileinfo, err := dir.ReadDir(-1)
	// for _, fi := range fileinfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	//write file

	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()
	// f.WriteString("hi go")
	// f.WriteString("lang")

	// bytes := []byte("Hello Sumit")

	// f.Write(bytes)

	//read and write another file(streming fashion)

	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	destfile, err := os.Create("example2.txt")

	if err != nil {
		panic(err)
	}
	defer destfile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destfile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)

			}
			break
		}

		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}

	}
	writer.Flush()
	fmt.Println("Written to new file succesfully")
}
