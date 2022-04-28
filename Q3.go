package main

import (
        "fmt"
        "strings"
        "os"
        "bytes"

        shell "github.com/ipfs/go-ipfs-api"
)

func main() {
        printFile("QmP8jTG1m9GSDJLCbeWhVSVgEzCPPwXRdCRuJtQ5Tz9Kc9")

        myfirststringcid := storeString("My first string in ipfs")
        printFile(myfirststringcid)
}

//Store string
func storeString(mystr string) string {
        // Where your local node is running on localhost:5001
        sh := shell.NewShell("localhost:5001")

        cid, err := sh.Add(strings.NewReader(mystr))
        if err != nil {
        fmt.Fprintf(os.Stderr, "error: %s", err)
        os.Exit(1)
        }
        fmt.Printf("Added \"%s\" in cid \"%s\" \n", mystr, cid)
        return cid
}

//To print contents of the cid
func printFile(cid string) {
        // Where your local node is running on localhost:5001
        sh := shell.NewShell("localhost:5001")

        data, err := sh.Cat(cid)
        if err != nil {
        fmt.Fprintf(os.Stderr, "error: %s", err)
        os.Exit(1)
        }

        buf := new(bytes.Buffer)
        buf.ReadFrom(data)
        newStr := buf.String()
        fmt.Printf("Contents of cid \"%s\":\n%s \n", cid, newStr)
}
