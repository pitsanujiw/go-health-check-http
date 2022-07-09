# Health checker

## How to setup

please install nodeJs and Golang

and this I use go module
and use makefile to implement command

## install dependencies

for service please run command this

```bash
    make setup-service
```

and then for client please run command this

```bash
    make install-client
```

after this you can start service and client
please split terminal for run 2 commands

a 1st terminal

```bash
    make start-service
```

a 2nd terminal

```
    make start-client
```

and a file for test have on `test_file` folder
