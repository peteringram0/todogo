### Install
````bash
dep ensure
````

#### Install - Go Path
Add the following to ~/.zshrc
````bash
export GOPATH=$HOME/gocode
export PATH=$PATH:$GOPATH/bin

## CONFIG below is used as fallback but needed on server
#export PORT=8000 # Port number for running the app locally or Heroku will fill this
#export REDIRECT_URL=http://127.0.0.1:8000/api/v1/auth # Google API redirect URL
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


#### Questions
* Is there a better way then building and running ?
* I create JTW and when only allow to route when its valid, but i don't know how to confirm the user from the JTW
* Work out how to invalidate the JWT & Session
* There is an issue when your first login, you need to logout and back in for it to work
