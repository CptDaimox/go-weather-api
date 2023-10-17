# go-weather-api
Simple Weather API using open-meteo.com

## Project Structure
```
go-weather-api/
|-- db/
|-- pkg/
|-- controllers/
|-- main.go
|-- go.mod
|-- go.sum
```

This project has 3 Layers. Each layer depends on the previous layer within the project, starting with the Data(base) Layer which has no internal dependencies.

### `db/`
Database Layer. Typically it would fetch data from a database e.g. via SQL-Queries. In this application the data comes from 
an external API. It has no internal dependendencies to another layer.

### `pkg/`
Domain/Service layer. This layer depends on the Database Layer. It executes the functions to fetch the data and can manipulate or sort the data if necessary.

### `controllers/`
This Layer depends on the Service Layer and controlls the Response and extracts and/or validates the Request Data to pass it to the Service Layer functions.