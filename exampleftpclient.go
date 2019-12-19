package main

import (
	ftp	"github.com/jlaffaye/QUICftpClient"
	"log"
	"time"
	"os"
	"fmt"
	"io"
)

func main() {
	c, err := ftp.Dial("172.26.17.76:2121", ftp.DialWithTimeout(5*time.Second))
    if err != nil{
        log.Fatal(err)
    }

    err = c.Login("admin", "123456")
    if err != nil {
        log.Fatal(err)
    }

    r, err := c.Retr("100MB.txt")
    if err != nil{
        panic(err)
    }

    file, err := os.Create("filefromserver.txt")
    log.Println("Receiving from server")
	start := time.Now()
	a, err := io.Copy(file, r)
	fmt.Println(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	notifier := fmt.Sprintf("Finished transfer in %.2f seconds!", time.Since(start).Seconds())
	log.Println(notifier)

    /*file, err = os.Open("filefromserver.txt")
    fi , _ := file.Stat()
    fmt.Printf("The file is %d bytes long\n", fi.Size())
    log.Println("Storing to server")
    start = time.Now()
    err = c.Stor("filefromclient.txt",file)
    if err != nil {
        panic(err)
    }
    notifier = fmt.Sprintf("Finished transfer in %.2f seconds!", time.Since(start).Seconds())
    log.Println(notifier)*/

    if err := c.Quit(); err != nil {
        log.Fatal(err)
    }

}
