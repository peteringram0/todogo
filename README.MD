### Install
````bash
dep ensure
````

#### Install - Go Path
Add the following to ~/.zshrc
````bash
export GOPATH=$HOME/gocode
export PATH=$PATH:$GOPATH/bin
export PORT=8000 # Port number for running the app locally or Heroku will fill this
````

### Build
````bash
bash ./build.sh
````

### Run
````bash
todo
````


##### Notes
````bash
# Install DEP
brew install dep

# View Database
sqlite3 storage.db
````
