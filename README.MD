# TODOGO
An example to do list application written in GO using Google OAuth 2.0.

### Install
````bash
dep ensure
````

#### Install - Go Path
##### Add the following to ~/.zshrc
````bash
export GOPATH=$HOME/{{Your base go code dir}}
export PATH=$PATH:$GOPATH/bin

## CONFIG below is used as fallback but needed on server
#export PORT=8000 # Port number
#export REDIRECT_URL=http://127.0.0.1:8000/api/v1/auth # Google API redirect URL
````
##### Create the following file
(creds.json) Stores google Oauth client and secret
````json
{
  "cid": "",
  "csecret": ""
}
````


### Build
````bash
bash ./build.sh
````

### Run
````bash
todogo
````

##### Notes
````bash
# Install DEP
brew install dep

# View Database
sqlite3 storage.db
````
