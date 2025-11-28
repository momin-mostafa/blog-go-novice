# My Go Journeys

I've decided to keep a note or blog documenting my weekend explorations into different technologies. This space will serve as a personal log of how I learn something new, tackle challenges, and strive to improve my skills. By sharing my progress, I hope to reflect better on my learning journey and maybe even help others along the way.

## Go & Backend Developement

- [Getting Started with Go: Building a RESTful API Server with PostgreSQL Database Integration - A Beginner's Journey](blog_1.md)

> In this journey, I built a RESTful API server from scratch using Go, integrated it with a PostgreSQL database, and implemented routing with `ServeMux`. I explored core Go concepts like handlers, package management, and database connectivity—progressing from a beginner to building a working backend in just a few hours.

- [Continuation of backend development](blog_2.md)

> In this follow-up, I modularized the project by separating request and DB logic into packages. I replaced raw SQL with GORM for easier database handling and implemented GET/POST methods for user data. I also introduced an adapter-like pattern for request models and began laying the groundwork for Test-Driven Development (TDD) as I scale the backend.

## how to run the code 
To create executable command :  
```bash 
go build .
``` 
To run the project command : 
```bash 
go run main.go 
``` 