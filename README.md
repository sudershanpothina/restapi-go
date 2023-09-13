# restapi-go
rest api using golang


## refer internal packages
```
import (

    "module name/folder name"
)

module name can be found in go.mod file
```

## import external module
```
go get github.com/gorilla/mux
```
## check in go.mod and go.sum to maintain dependencies


## input validation
```
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer)

```

## http methods
```
    router.HandleFunc("/customers", getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
```