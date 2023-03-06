# go-canary

The point of this go program is to "override" any command the user would like to be used as an alarm when invoked.

## how it works
The program will run the command with the arguments it received but with the "_" added in the begining of the name.

## How to use it

1. add the token, server, port or canary token to it
2. compile it in the folder using ```go build .```
3. copy the command you'd like to "_" on the begining of it's name
4. copy the generated executable to the (from step 2) to the command path

### requires

1. Linux OS
2. go compiler(v1.16 or higher)
3. UPD listener/canary token
4. root to replace commands

If you want to run in another system that does not have go installed

5. install gccgo and compile with:  
```bash
gcc go-canary.go -static-libgo -o go-canary
```

### example for command **whoami** 

After build
```bash
sudo cp /usr/bin/whoami /usr/bin/_whoami
sudo cp go-canary /usr/bin/whoami
```

Don't forget to test it.
```shell
whoami
```
You can always execute the original
```bash
_whoami
```

All set.
You can do this to as many commands as you'd like. :)

## logs for stack executing and who

The execution will generate a file that is the result of the execution of:
```bash
ps fauxe
```

Into:

> LOG_FILE_PATH + **command executed** + . + 10 random characters + .log

And who's connected to:
```bash
w
```
Into:
```bash
LOG_FILE_PATH + w- + . + the same random characters + .log
```
The above example would generate the files:
```bash
/tmp/whoami.XXXXXXXXXX.log
/tmp/w-XXXXXXXXXX.log
```

## canary tokens

I do not recommend canary tokens, it's too slow, would alert the attacker.

## why?

Well UDP is faster and I want it all to run without alerting the attacker.

## Also, very important

Use at your own descretion, careful what you replace

### TODO

* Make a version for older versions of go and gccgo
