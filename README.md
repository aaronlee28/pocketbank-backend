# E-Wallet Backend

## Getting Started

1. Initialize the project
    - ```go mod init```
2. Get all the dependencies required
    - ```go get ./...```

Login Info for test and explore purposes: 
- ```email: aaronlee@gmail.com```
```password: password1```
- ```email: peterparker@gmail.com```
```password: password2```
- ```email: admin@admin.com```
  ```password: admin```
- ```email: markruffalo@gmail.com```
  ```password: mark1```
- ```email: johndoe@gmail.com```
  ```password: johnny```
## About the Project
This is part of a full-stack E-Wallet project. Go Languange is solely used to build the application, using Gin as the HTTP web framework and GORM as the ORM library. Postgres is the DB of choice for this project.  

All endpoints could be found in the API Documentation linked below. 

API Documentation:
https://pocket-bank-backend.herokuapp.com/docs/

## Deployed project can be accessed here:
1. Frontend: https://pocket-bank-frontend.herokuapp.com/
2. Backend: https://pocket-bank-backend.herokuapp.com/

The PG_DUMP file of the database could be found in the main package. 



## Database ERD 
![ERD](./image/entity_relationship_diagram.png)

## Test
![test](./image/test_coverage.png)
