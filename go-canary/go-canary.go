package main

import (
    "time"
    "math/rand"
    "io"
    "fmt"
    "os/exec"
    "os"
    "os/user"
    "net"
    "path/filepath"
    "strings"
)

//DO NOT FORGET TO ADD THESE VALUES OR IT WILL DO NOTHING
const TOKEN = ""//this is a simple string for the UDP server to flag/allow this message
const UDP_SERVER = ""
const PORT = ""
//FOR CANARY TOKENS, WILL BE SLOWER THEN A SIMPLE UDP PACKET
const CANARY_TOKEN = ""
//path where the log for ps will be stored
const LOG_FILE_PATH = "/tmp/"

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

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

//process > LOG_FILE_PATH + command executed + . + 10 random characters + .log
//who's connected > LOG_FILE_PATH + w- + . + 10 random characters + .log
func logInfo(arg string){
    rand.Seed(time.Now().UnixNano())
    rnum := RandString(10)
    command_line := "ps faux > '" + LOG_FILE_PATH + arg + "." + rnum + ".log" + "'"
    c2 := "w > '" + LOG_FILE_PATH + "w-" + rnum + ".log" + "'"
    cmd := exec.Command("bash", "-c",command_line + ";" + c2)

    cmd.Start()
}

func main() {
    //get the current executable and its arguments
    args := os.Args[1:]
    _, file := filepath.Split(os.Args[0])
    app := "_" + file
    //logs run first
    logInfo(file)
    //alerts
    canary_token()
    message(strings.Join(os.Args," "))
    //execution
    cmd := exec.Command(app, args...)
    //this is better then run because it pipes stdout and stderr to stderr and stdout
    stdout, _ := cmd.StdoutPipe()
    stderr, _ := cmd.StderrPipe()
    cmd.Start();
    serr, _ := io.ReadAll(stderr)
    sout, _ := io.ReadAll(stdout)
    //print outs
    if serr != nil{
        fmt.Fprintf(os.Stderr,string(serr))
    }
    if sout != nil{
        fmt.Fprintf(os.Stdout,string(sout))
    }
    err := cmd.Wait()
    //if it failed to get an error return error
    if err != nil{
        os.Exit(1)
    }
}
