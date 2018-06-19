### Install
````bash
go get github.com/labstack/echo
go get github.com/mattn/go-sqlite3
````

#### Install - Go Path
Add the following to ~/.zshrc
````bash
export GOPATH=$HOME/code/todogo
export PATH=$PATH:$GOPATH/bin
````

### Build
````bash
bash ./build.sh
````

### Run
````bash
./bin/todo
````


##### Notes
````bash
# View Database
sqlite3 storage.db
````
