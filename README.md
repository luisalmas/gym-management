# gym-management

Implementation of a service in Go of a gym management

## Features

- This project contains the requests POST, PUT, GET and DELETE to insert, update, retrieve and delete respectively for both classes and bookings resourses.
- Validations not only for the data transfer objects (DTOs) but also for the bookings that are inserted outside the range of the respective class or if a class is updated the changes are reflected in the corresponding bookings (ie: changes in the date ranges deletes bookings outside of the new range).
- Tests for all the layers of each resource
- Swagger UI to document the api and endpoint manual tests

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
