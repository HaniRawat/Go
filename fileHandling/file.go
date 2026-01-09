package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")	//cd into the fileHandling folder so the file can be opened
	// // fmt.Println(f, err)

	// if err != nil {
	// 	//log the error
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	//log the error
	// 	panic(err)
	// }
	// fmt.Println("file name -",fileInfo.Name())
	// fmt.Println("file or folder -",fileInfo.IsDir())
	// fmt.Println("file size -",fileInfo.Size())
	// fmt.Println("file permissions -",fileInfo.Mode())
	// fmt.Println("file modified at-",fileInfo.ModTime())

	//read file
	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// buf := make([]byte, 10)

	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	println("data", d, string(buf[i]))

	// }

	//another way, preferred for smaller files, as it loads the data all at once
	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	//to read folders
	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)
	// if err != nil {
	// 	panic(err)
	// }

	// for _,fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	//create a file
	// f, err := os.Create("example2.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("hi go")
	// f.WriteString(" nice language") //appends the text to file, not overwrite

	//another way
	// bytes := []byte ("Hello golang")
	// f.Write(bytes)

	//read and write to another file (streaming version)

	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }

	// writer.Flush()
	// fmt.Println("written to new file successfully")

	//delete a file
	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("file deleted successfully")

}
