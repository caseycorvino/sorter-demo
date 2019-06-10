# Layout 
## EC2
```
.
├── Makefile
├── README.md
├── configs
├── scripts
├── docs
├── vendor
├── build
├── cmd
│   └── Sorter
│       ├── main.go
│       └── main_test.go
└── pkg
│       └── user
│       │    ├── entity.go
│       │    ├── mongodb.go
│       │    ├── interface.go
│       │    ├── test.go
│       │    └── service.go
│       └── profile
│       │     ├── entity.go
│       │     ├── mongodb.go
│       │     ├── interface.go
│       │     ├── test.go
│       │     └── service.go
│       │...
└──app
   └── API
   │    └── tests/
   │    └── controllers/
   └── Pipeline (file upload to trigger lambda)
   │    └── tests/
   │    └── controllers/
   └── Frontend
   │     └── tests/
   │     └── controllers/
   └── UserManagement 
   │     └── tests/
   │     └── controllers/
   |...
```

## LAMBDA
```
.
├── Makefile
├── README.md
├── configs
├── scripts
├── docs
├── vendor
├── build
├── cmd
│   └── Sorter
│       ├── main.go
│       └── main_test.go
└── pkg
│       └── user
│       │    ├── entity.go
│       │    ├── mongodb.go
│       │    ├── interface.go
│       │    ├── test.go
│       │    └── service.go
│       └── profile
│       │     ├── entity.go
│       │     ├── mongodb.go
│       │     ├── interface.go
│       │     ├── test.go
│       │     └── service.go
│       │...
└── Pipeline
│    └── tests/
│    └── controllers/
|...
```


## Configs
Configuration files - parameters and system details.

## Scripts
Helper scripts for building and installing, help the Makefile.

## Docs
Various documentation on the project.

## Vendor
Application dependencies

## Build
Packaging and continous integration files

## cmd
Main files to run.

## pkg
Entities and their services shared by the entire application. Can be imported and used in all projects.

## app
Here is where our controllers and tests lie. They are divided by business service. 
The controllers convert data into useful information. 

Ex: Fetching a user profile and using the data to render a profile page. 
Ex2: Fetching watson personality type and adding to to the user account.

## About
The controllers and test are divided into psuedo "microservices" so that in the future we can easily make them standalone applications. However, the enitities are not because it is important that the entities have the same structure for every service. 

It is designed this way because we may not want to do microservices initially - violates "2 pizzas" rule and more upkeep having a minimum of 3 ec2 instances, not including QA environments. But, once we need to scale, having our CPU intensive services on the same server will be very costly to our other processes.  (*may not be the case with lambda)

https://github.com/golang-standards/project-layout
https://medium.com/@eminetto/clean-architecture-using-golang-b63587aa5e3f

