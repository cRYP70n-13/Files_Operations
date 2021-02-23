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
