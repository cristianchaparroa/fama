# Fama

[![codecov](https://codecov.io/gh/cristianchaparroa/fama/branch/master/graph/badge.svg?token=BTHAPCM0II)](https://codecov.io/gh/cristianchaparroa/fama)

Fama is a project that is in charge to transform numbers into words.

## Local development environment

### Requirements

You should have installed Go in your machine, please follow this guide https://golang.org/doc/install 

### Download the project 

```shell script
git clone https://github.com/cristianchaparroa/fama.git 
```
### Execute the application

You can run the server using the following command

```shell script
cd path/to/project/fama/
go run *.go
``` 

After executes the last commands you can start to using the API. 

### Docker

If is not possible create a local environment, and you have docker installed in 
your machine you can follow the next steps to run the application.

```shell script
docker build -t fama .
docker run -p 8080:8080 fama
```
