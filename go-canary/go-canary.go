package main

import (
    "fmt"
    "os/exec"
    "os"
    "path/filepath"
    "net"
)

//DO NOT FORGET TO ADD THESE VALUES OR IT WILL DO NOTHING
const TOKEN = ""//this is a simple string for the UDP server to flag/allow this message
const UDP_SERVER = ""
const PORT = ""
//FOR CANARY TOKENS, WILL BE SLOWER THEN A SIMPLE UDP PACKET
const CANARY_TOKEN = ""

func message(command string) {
    if TOKEN == ""{
        return
    }
    conn, err := net.Dial("udp", UDP_SERVER + ":" + PORT)
    if err != nil {
        return
    }
    prompt := "$"
    path, err := os.Getwd()
    if err == nil{
        prompt = path + prompt
    }
    user, err := user.Current()
    if err == nil{
        prompt = user.Username + ":" + prompt
    }
    fmt.Fprintf(conn, TOKEN + prompt + command)
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
    //get the current executable and its arguments
    args := os.Args[1:]
    _, file := filepath.Split(os.Args[0])
    app := "_" + file
    
    cmd := exec.Command(app, args...)
    stdout, err := cmd.Output()
    //send errors to stderr
    if err != nil {
        fmt.Fprintf(os.Stderr,err.Error())
    }
    if stdout != nil && stdout != ""{
        fmt.Fprintf(os.Stdout,string(stdout))
    }
}
