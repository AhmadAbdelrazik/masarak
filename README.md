# Masarak - Online Job Search Platform

Masarak is an API for an Online job search platform. written in Golang

## Description

Masarak provides a platform for both freelancers and businesses to find each other easier.

In Masarak, Freelancers can build their profile, search, and apply for jobs. Businesses can post jobs, search for freelancers based on the chosen skills, and review applications.

## Installation

```bash
# download all the needed components
go mod tidy

# run the API
make run/api
```

## Architecture

Masarak architecture is inspired by the amazing work of [ThreeDots](https://threedots.tech/series/modern-business-software-in-go) that introduces [Domain Driven Design](https://en.wikipedia.org/wiki/Domain-driven_design) that is more idiomatic to Golang.

The architecture utilizes the strength of DDD and it's modularity and ease of scalability, but careful not to fall in a lot of boiler plate. the architecture can be described by this 

![application architecture](https://github.com/AhmadAbdelrazik/masarak/blob/main/docs/application%20layers.png)


The architecture lean more towards the [Hexagonal architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)) defining Ports which are entry points to the application, and adapters which implements the necessary infrastructure and services defined by the domain layer.

## Domain Layer

Domain Layer consists of the main business logic and it's invariants. Here The model enforces the rules of the business, making sure that the system is always in a valid state. This is where the most important parts of the system resides. The core logic and the heart of the system.

The Domain layer defines repositories and services interfaces, but not bother themselves with implementing them for the sake of modularity and ease of refactoring. The Variance in the database, or the changing of mail service should not reflect on the domain model that has the business laws and rules.

## Application Layer

Application Layer consists of the application use cases. It uses the domain layer and it's defined repositories along side services (e.g. mailing service) to accomplish the use case. It's main job is to work as an orchestrator between different services.

The Application layer is designed with [CQRS](https://en.wikipedia.org/wiki/Command%E2%80%93query_separation) in mind. In CQRS the commands and the queries can have different databases, so that we can utilize database configured for read operations in our queries, and use the one for write operations in our commands. and the queries and commands can evolve separately.

## Ports & Adapters

For those familiar with the ports and adapters in the Hexagonal architecture you might notice a different naming, what we mean here by ports and adapters is more like entry point adapters and infrastructure adapters. The use of ports does not mean the interfaces defined by the application that acts usually as an anti-corruption layer and to standardize communication with infrastructure.  

Ports are the entry points to the main application, it sanitize the inputs for the application layer use cases. provide an interface (HTTP, MQTT, ...). This method can scale quiet easily especially when trying to scale the architecture to a microservice which have multiple entry points for each service, yet they can also share the same core domain.

Adapters here are the implementation of the interfaces defined by the domain layer. Most of the time it's always the database implementation, any services such as mailing services, external APIs, etc...
