# kumgo-stack
Kumparan Stack Framework for Go Lang

## Project Structure
1. **buff** - This folder is used to store the *.proto files and the compilation result.
2. **client** - Client folder is used for other services clients.
3. **config** - This folder used for the configuration setting.
4. **console** - All the terminal commands goes here.
5. **deploy** - Deployment scripts.
6. **{{YOUR_SERVICE}}svc** - Your service handler.
7. **repository** - Service model, query and logic.
8. **example** - contain client example for tutorial purpose only. Remove when you use the boilerplate in your service.

## Installing Go
### MacOS
1. You can follow the steps [here](https://golang.org/doc/install#osx)
2. Make sure you set your `$GOPATH` right. Check how to set your `$GOPATH` [here](https://github.com/golang/go/wiki/SettingGOPATH)

Some plugin will need to be installed in the `$GOBIN`. Default `$GOBIN` is `$GOPATH/bin`. It must be in your `$PATH` for the plugin to find it.
1. Add this line in your .zshrc or .bash_profile
```shell
$ export PATH=$PATH:$GOPATH/bin
```
2. Then reload the profile
```shell
$ source ~/.zshrc #if you are using zsh
$ source ~/.bash_profile #if you are using bash
```

## Installing Protobuf Compiler
### MacOS
1. Download the appropriate release [here](https://github.com/google/protobuf/releases)
2. Unzip the folder
3. Enter the folder and run `./autogen.sh && ./configure && make`
...If you run into this `error: autoreconf: failed to run aclocal: No such file or directory`, run `brew install libtool && brew install autoconf && brew install automake`. And run the command from step 3 again.
4. Then run these other commands. They should run without issues
```shell
$ make check
$ sudo make install
$ which protoc
$ protoc --version
```
5. Install the protoc plugin for go.
```shell
$ go get -u github.com/golang/protobuf/protoc-gen-go
```
### Installing Dep (Go Dependency Manager)
You can follow the steps [here](https://github.com/golang/dep)
## Story Service Example (with ArangoDB Connection)
### Run the server in your local
1. Make sure you already clone this repository in your `$GOPATH/src` folder.
2. Install the dependencies by running `dep ensure` in your terminal.
3. Compile the protobuf for the server example using `make proto`.
4. Create a configuration file called `config.yml` in the root directory.
5. You can copy the configuration in `config.yml.example` to your `config.yml`, or you also can change the config.
6. Run the service using `make run`.
7. You should see something like this:
```shell
INFO : Connection to Arango Server success...
INFO : Connection to database success...
INFO : Listening on 9001
```
### Test the client
1. With the server running, open new terminal.
2. Type `make story-client-example`. This command will create a new gRPC client of the server run in port 9001. The client will send a GetStories request to the server, and get the response from the server.
3. You should see something like this:
```shell
INFO[0000] Echo result: i'm screaming
```
## Echo Example (without ArangoDB Connection)
### Run the server in your local
1. Make sure you already clone this repository in your `$GOPATH/src` folder.
2. Install the dependencies by running `dep ensure` in your terminal.
3. Compile the protobuf for the server example using `make echo-example-proto`.
4. Run the service using `make echo-server-example`.
5. You should see something like this:
```shell
INFO[0000] Listening on 9001
```
### Test the client
1. With the server running, open new terminal.
2. Type `make echo-client-example`. This command will create a new gRPC client of the server run in port 9001. The client will send a Echo request to the server, and get the response from the server.
3. You should see something like this:
```shell
INFO[0000] Echo result: i'm screaming
```