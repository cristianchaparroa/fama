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
cd path/to/project/fama/
docker build -t fama .
docker run -p 8080:8080 fama
```


### Docker-compose

If you have installed `docker` and `docker-compose` you also can deploy the services just with

```shell script
cd path/to/project/fama/
docker-compose up
```
It'll start the API server that will be exhibited in
```
{host}:8080
```
And the swagger documentation will be expose in 
```
{host}:8081
``` 

## Infrastructure deployment

The following are the steps to set up the Digitalocean account and terraform.

1. First of all you need to create a token as is showed on https://www.digitalocean.com/docs/apis-clis/api/create-personal-access-token/
after that should export the token as environment variable to work with it properly.
```
export DO_PAT={YOUR_PERSONAL_ACCESS_TOKEN}
```

2. After that should have at least one ssh key, if you don't have it you should follow the next documentation https://www.digitalocean.com/docs/droplets/how-to/add-ssh-keys/

3. The `fama_rsa` ssh key was added in the second step to digital ocean account. 
Now is need to retrieve the  finger print ssh key.

```shell script
ssh-keygen -E md5 -lf ~/.ssh/fama_rsa.pub | awk '{print $2}'
md5:d0:65:34:b8:42:93:88:45:24:59:cd:a7:e7:90:c2:6d
```
You need to take the part after `md5:` in this case should be `d0:65:34:b8:42:93:88:45:24:59:cd:a7:e7:90:c2:6d`

Also is needed to add all credentials in the `terraform.tfvars`

Before to run the terraform plan is needed to download the terraform dependencies then you need to run the command init.
After it you need to validate it check that all structures you have all correct.
```shell script
cd terraform/production
terraform init
terraform validate
```

Next we should need to generate the plan
```shell script
terraform plan  -var-file=terraform.json
```

Finally we can deploy the plan with the command
```shell script
terraform apply -var-file=terraform.json
```

