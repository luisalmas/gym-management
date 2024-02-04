# gym-management

Implementation of a service in Go of a gym management

## Features

## Clone project

To clone the project repository run the following command:

    git clone https://github.com/luisalmas/gym-management.git

## Run service

The project already contains the executable file `gym-management.exe` to run the service, otherwise just execute the following command on the root directory:

    go run .

## Run tests

To run the tests execute the following command:

    make test

## Swagger url

The swagger intereface is accessible via the following url:

    http://localhost:8080/docs/index.html

### Load swagger

If the swagger documents are not loaded, the following command will create them:

    swag init
