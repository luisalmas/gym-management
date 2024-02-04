# gym-management

Implementation of a service in Go of a gym management

## Clone project

    git clone https://github.com/luisalmas/gym-management.git

## Run service

    go run .

## Run tests

    go test ./...

## Implemented endpoints

### Classes

#### Get classes

    GET /api/classes

    returns: [{"classId": int, "capacity": int, "endDate": date RFC3339, "name": string, "startDate": date RFC3339}]

#### Get class

    GET /api/classes/:id

    returns: {"classId": int, "capacity": int, "endDate": date RFC3339, "name": string, "startDate": date RFC3339}

#### Insert class

    POST /api/classes

    body: {"capacity": int, "endDate": date RFC3339, "name": string, "startDate": date RFC3339}

    returns: {"classId": int, "capacity": int, "endDate": date RFC3339, "name": string, "startDate": date RFC3339}

#### Update class

    PUT /api/classes/:id

    body: {"capacity": int, "endDate": date RFC3339, "name": string, "startDate": date RFC3339}

    returns: {"classId": int, "capacity": int, "endDate": date RFC3339, "name": string, "startDate": date RFC3339}

#### Delete class

    DELETE /api/classes/:id

    returns: {"classId": int, "capacity": int, "endDate": date RFC3339, "name": string, "startDate": date RFC3339}

#### Get bookings of class

    GET /api/classes/:id/bookings
    GET /api/classes/:id/bookings?date=2024-02-04T00:00:00Z

### Bookings

#### Get bookings

    GET /api/bookings

    returns: body: [{bookingId: int, "classId": int, "date": date RFC3339, "name": string}]

#### Get booking

    GET /api/bookings/:id

    returns: body: {bookingId: int, "classId": int, "date": date RFC3339, "name": string}

#### Insert booking

    POST /api/bookings

    body: {"classId": int, "date": date RFC3339, "name": string}

    returns: body: {bookingId: int, "classId": int, "date": date RFC3339, "name": string}

#### Update booking

    PUT /api/bookings/:id

    body: {"classId": int, "date": date RFC3339, "name": string}

    returns: body: {bookingId: int, "classId": int, "date": date RFC3339, "name": string}

#### Delete booking

    DELETE /api/bookings/:id

    returns: body: {bookingId: int, "classId": int, "date": date RFC3339, "name": string}

#### Todo

    - Simplify dates in inputs
    - Errors when parsing dates

    - More search parameters (if there is time)

    https://chat.openai.com/c/e06e57fc-ba39-4c3b-ba70-a5c17322b140 (dto validations)
