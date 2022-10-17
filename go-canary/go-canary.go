package main

import (
    "fmt"
    "os/exec"
    "os"
    "strings"
    "net"
)

//DO NOT FORGET TO ADD THESE VALUES OR IT WILL DO NOTHING
const TOKEN = ""//this is a simple string for the UDP server to flag/allow this message
const UDP_SERVER = ""
const PORT = ""
//FOR CANARY TOKENS, WILL BE SLOWER THEN A SIMPLE UDP PACKET
const CANARY_TOKEN = ""

func message(command string) {
    if UDP_SERVER == ""{
        return
    }
    conn, err := net.Dial("udp", UDP_SERVER + ":" + PORT)
    if err != nil {
        return
    }
    fmt.Fprintf(conn, TOKEN+"Cmd:"+command)
    conn.Close()
}

//dns lookups can be too slow, so I'd recommend the usage of an UDP with token server
func canary_token(){
    if CANARY_TOKEN != "" {
        net.LookupIP(CANARY_TOKEN)
    }
}


func main() {
    canary_token()
    message(strings.Join(os.Args," "))
    args := os.Args[1:]
    split := strings.Split(os.Args[0], "/")
    app := "_" + split[len(split)-1]
    cmd := exec.Command(app, args...)
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Fprintf(os.Stderr,err.Error())
        return
    }
    fmt.Fprintf(os.Stdout,string(stdout))
}
