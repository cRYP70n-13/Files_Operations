# Everything is a File

One of the fundamental aspects of UNIX is that everything is a file.
We don't necessarily know what the file descriptor maps to, that is abstracted by
the operating system's device drivers. The operating system provides us an
interface to the device in the form of a file.

The reader and writer interfaces in Go are similar abstractions. We simply read
and write bytes, without the need to understand where or how the reader gets its
data or where the writer is sending the data. Look in /dev to find available
devices. Some will require elevated privileges to access.

## Creating an Empty File

```GoLang
package main

import (
    "log"
    "os"
)

var (
    newFile *os.File
    err     error
)

func main() {
    newFile, err = os.Create("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    log.Println(newFile)
    newFile.Close()
}
```

## Truncate a File

```GoLang
package main

import (
    "log"
    "os"
)

func main() {
    // Truncate a file to 100 bytes. If file
    // is less than 100 bytes the original contents will remain
    // at the beginning, and the rest of the space is
    // filled will null bytes. If it is over 100 bytes,
    // Everything past 100 bytes will be lost. Either way
    // we will end up with exactly 100 bytes.
    // Pass in 0 to truncate to a completely empty file

    err := os.Truncate("test.txt", 100)
    if err != nil {
        log.Fatal(err)
    }
}
```

## Get file Info

```GoLang
package main

import (
    "fmt"
    "log"
    "os"
)

var (
    fileInfo os.FileInfo
    err      error
)

func main() {
    // Stat returns file info. It will return
    // an error if there is no file.
    fileInfo, err = os.Stat("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("File name:", fileInfo.Name())
    fmt.Println("Size in bytes:", fileInfo.Size())
    fmt.Println("Permissions:", fileInfo.Mode())
    fmt.Println("Last modified:", fileInfo.ModTime())
    fmt.Println("Is Directory: ", fileInfo.IsDir())
    fmt.Printf("System interface type: %T\n", fileInfo.Sys())
    fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
}
```

## Rename and Move a File

```GoLang
package main

import (
    "log"
    "os"
)

func main() {
    originalPath := "test.txt"
    newPath := "test2.txt"
    err := os.Rename(originalPath, newPath)
    if err != nil {
        log.Fatal(err)
    }
}
```

## Delete a file

```GoLang
package main

import (
    "log"
    "os"
)

func main() {
    err := os.Remove("test.txt")
    if err != nil {
        log.Fatal(err)
    }
}
```

## Check if File exists

```GoLang
package main

import (
    "log"
    "os"
)

var (
    fileInfo *os.FileInfo
    err      error
)

func main() {
    // Stat returns file info. It will return
    // an error if there is no file.
    fileInfo, err := os.Stat("test.txt")
    if err != nil {
        if os.IsNotExist(err) {
            log.Fatal("File does not exist.")
        }
    }
    log.Println("File does exist. File information:")
    log.Println(fileInfo)
}
```

# Quick Write to File

The ioutil package has a useful function called WriteFile() that will handle
creating/opening, writing a slice of bytes, and closing. It is useful if you just
need a quick way to dump a slice of bytes to a file.

```GoLang
package main

import (
    "io/ioutil"
    "log"
)

func main() {
    err := ioutil.WriteFile("test.txt", []byte("Hi\n"), 0666)
    if err != nil {
        log.Fatal(err)
    }
}
```

## Use Buffered Writer

The **bufio** package lets you create a buffered writer so you can work with a
buffer in memory before writing it to disk. This is useful if you need to do a
lot manipulation on the data before writing it to disk to save time from disk IO.
It is also useful if you only write one byte at a time and want to store a large
number in memory before dumping it to file at once, otherwise you would be
performing disk IO for every byte. That puts wear and tear on your disk as well
as slows down the process.

```GoLang
package main

import (
    "log"
    "os"
    "bufio"
)

func main() {
    // Open file for writing
    file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Create a buffered writer from the file
    bufferedWriter := bufio.NewWriter(file)

    // Write bytes to buffer
    bytesWritten, err := bufferedWriter.Write(
        []byte{65, 66, 67},
    )
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Bytes written: %d\n", bytesWritten)

    // Write string to buffer
    // Also available are WriteRune() and WriteByte()   
    bytesWritten, err = bufferedWriter.WriteString(
        "Buffered string\n",
    )
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Bytes written: %d\n", bytesWritten)

    // Check how much is stored in buffer waiting
    unflushedBufferSize := bufferedWriter.Buffered()
    log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

    // See how much buffer is available
    bytesAvailable := bufferedWriter.Available()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Available buffer: %d\n", bytesAvailable)

    // Write memory buffer to disk
    bufferedWriter.Flush()

    // Revert any changes done to buffer that have
    // not yet been written to file with Flush()
    // We just flushed, so there are no changes to revert
    // The writer that you pass as an argument
    // is where the buffer will output to, if you want
    // to change to a new writer
    bufferedWriter.Reset(bufferedWriter) 

    // See how much buffer is available
    bytesAvailable = bufferedWriter.Available()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Available buffer: %d\n", bytesAvailable)

    // Resize buffer. The first argument is a writer
    // where the buffer should output to. In this case
    // we are using the same buffer. If we chose a number
    // that was smaller than the existing buffer, like 10
    // we would not get back a buffer of size 10, we will
    // get back a buffer the size of the original since
    // it was already large enough (default 4096)
    bufferedWriter = bufio.NewWriterSize(
        bufferedWriter,
        8000,
    )

    // Check available buffer size after resizing
    bytesAvailable = bufferedWriter.Available()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Available buffer: %d\n", bytesAvailable)
}
```
