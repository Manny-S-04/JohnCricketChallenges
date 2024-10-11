package main

import (
	"bufio"
	"flag"
	"io/fs"
	"os"
	"unicode/utf8"
)

func main(){
    var bytePtr *string = flag.String("c", "", "bytes of file")
    var linesPtr *string = flag.String("l", "", "lines of file")
    var wordsPtr *string = flag.String("w", "", "word count of file")
    var charPtr *string = flag.String("m", "", "char count of file")
    flag.Parse()

    if(len(*bytePtr) == 0 && len(*linesPtr) == 0 && len(*wordsPtr) == 0 && len(*charPtr) == 0){
        return
    }

    if(len(*bytePtr) > 0){
        BytesFlag(bytePtr)    
        return
    }
    if(len(*linesPtr) > 0){
        LinesFlag(linesPtr)
        return
    }
    if(len(*wordsPtr) > 0){
        WordsFlag(wordsPtr)
    }
    if(len(*charPtr) > 0){
        CharFlag(charPtr)
    }
}

func CharFlag(filePtr *string){
    var(
        file *os.File
        fileScanner *bufio.Scanner
    )
    file, fileScanner = ReadFile(*filePtr)
    defer file.Close()

    var count int
    for fileScanner.Scan(){
        var line string = fileScanner.Text()
        count += utf8.RuneCountInString(line)
    }
    println(count, *filePtr)
}

func WordsFlag(filePtr *string){
    var(
        file *os.File
        fileScanner *bufio.Scanner
    )
    file, fileScanner = ReadFile(*filePtr)
    defer file.Close()

    fileScanner.Split(bufio.ScanWords)

    var count int = 0
    for fileScanner.Scan(){
        count++
    }
    println(count, *filePtr)
}

func LinesFlag(filePtr *string){
    var(
        file *os.File
        fileScanner *bufio.Scanner
    )
    file, fileScanner = ReadFile(*filePtr)
    defer file.Close()

    fileScanner.Split(bufio.ScanLines)

    var counter int = 0
    for fileScanner.Scan(){
        counter++
    }
    println(counter)
}   

func BytesFlag(filePtr *string){
    var(
        fileInfo fs.FileInfo
        err error
    )
    fileInfo, err = os.Stat(*filePtr)
    if err != nil{
        println("Error getting file stat")
        return
    }

    var fileSize int64 = fileInfo.Size()
    println(fileSize, *filePtr)
}

func ReadFile(fileToOpen string) (*os.File, *bufio.Scanner){
    var(
        file *os.File
        err error
    )
    file, err  = os.Open(fileToOpen)
    if err != nil{
        println("There was an error opening the file")
        return nil, nil
    }

    var fileScanner *bufio.Scanner = bufio.NewScanner(file)
 
    return file, fileScanner
}
