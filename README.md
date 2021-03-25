# RandomApp

A small project to learn more about creating a web app. The aim is to get a web app working which runs inside of a docker container. 
- Involves a http service that GETs responses from other APIs (in this case, a random number generator) and then stores them in a db for user retreival.

### Reminders
To create a go module, run the following:
```
go mod init github.com/{username}/{repo}
```

To build docker image, run:
```
docker build -t <name you call it> <dir>
```
Often, you will just use `.` for the dir.

To run docker image, run:
```
docker run <name you called it>
```

Inside of bash scripts, you should have:
- `#! /bin/sh`
- `set -o errexit` to make sure your script exits when your command fails
- `set -o nounset` to exit the script when it tries to call undeclared vars
- `set -o pipefail` which will show the exit status of the last command that threw a non-zero exit code is returned

The `.gitignore` file makes sure the `bin` directory does not get tracked.

The `bin` directory is probably unecessary, as you will build the go app inside the container.

If VS code stops tracking the import paths for go, just restart it. It seems to cache stuff poorly.