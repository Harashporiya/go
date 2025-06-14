package main

import (
	// "fmt"
	// "image/color/palette"
	// "fmt"
	// "bufio"
	"fmt"
	"os"
)


func main(){

//   f, err := os.Open("example.txt")

//   if err != nil {
// 	// log the error
// 	panic(err)
//   }

//   fileInfo, err := f.Stat()

//   if err != nil {
// 	// log the error
// 	panic(err)
//   }

//   fmt.Println("File name",fileInfo.Name())
//   fmt.Println("File folder",fileInfo.IsDir())
//   fmt.Println("File size",fileInfo.Size())
//   fmt.Println("File modifed at",fileInfo.ModTime()) 
//   fmt.Println("File premissions",fileInfo.Mode())

  //read file
//   f, err := os.Open("example.txt")

//   if err != nil {
// 	panic(err)
//   }

//   defer f.Close()
  
//   buf := make([]byte, 13)

//   d, err := f.Read(buf)

//   if err != nil{
// 	panic(err)
//   }

//   for i:=0;i<len(buf);i++{
// 	 println("data",d,string(buf[i]))
//   }

//  f, err := os.ReadFile("example.txt")

//  if err!=nil{
// 	panic(err)


//  }

//  fmt.Print(string(f))


// read folder

// dir, err := os.Open("../")

// if err != nil {
// 	panic(err)
// }

// defer dir.Close()

//   fileInfo, err :=dir.ReadDir(-1)

//   for _, fi := range fileInfo {
// 	fmt.Println(fi.Name(),fi.IsDir())
//   }
   

   // create a file

//    f, err := os.Create("h.txt")

//    if err!=nil{
// 	panic(err)
//    }

//    defer f.Close()

//    f.WriteString("hi go")
//    f.WriteString("Nice language")

//    bytes := []byte("Hello Golang")
//    f.Write(bytes)


   // read and write to another file (streaming fashion)

//    sourceFile, err := os.Open("example.txt")

//    if err != nil {
// 	panic(err)

//    }

//    defer sourceFile.Close()

//    destfile, err := os.Create("t.txt")
//    if err != nil {
// 	panic(err)

//    }

//    defer destfile.Close()

//    reader := bufio.NewReader(sourceFile)
//    writer := bufio.NewWriter(destfile)

//    for {

// 	b, err := reader.ReadByte()


// 	if err != nil{

// 		if err.Error() != "EOF"{
// 			panic(err)
// 		}

// 		break;

// 	}
//      e:= writer.WriteByte(b)
     
// 	 if e != nil{
// 		panic(e)
// 	 }
//    }

//    writer.Flush()

//    fmt.Println("written to new file succesfully")

   // delete a file
   err:= os.Remove("t.txt")

   if err!=nil{
	panic(err)
   }
   fmt.Println("delete  file succesfully")
   
}