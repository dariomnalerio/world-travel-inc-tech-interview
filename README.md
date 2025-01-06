# PawPics

A full-stack application that allows users to discover, view and save dog images using the Dog API. Built as part of the World Travel, Inc. Technical Interview requirements.

## Architecture Overview

- #### N-Tier application architecture consisting of:

  - Presentation Layer: React front-end.
  - Business Logic Layer: Go back-end server.
  - Data Layer: PostgreSQL database.

## Technical Stack

One of the aims of this project was to use as little external libraries as possible. So the technical stack itself is small.

### Frontend

- [Vite](https://vite.dev/ "Vite Page") 6.05
- [React.js](https://react.dev/ "React Page") 19.0.0
- React Context for global stage management
- [CSS Modules](https://github.com/css-modules/css-modules "CSS Modules Page")

### Backend

- [Go](https://go.dev/ "Go Page") 1.23.1
- [Gin web framework](https://gin-gonic.com/ "Gin Page") for REST API
- Raw SQL queries

### Database

- [PostgreSQL](https://www.postgresql.org/ "PostgreSQL Page") 17
- Direct SQL queries for CRUD operations

### External API Integration

- [Dog API](https://dog.ceo/dog-api/ "Dog API Page") for random dog images

#### Infrastructure

- [Docker](https://www.docker.com/ "Docker Page") containers for development, production and testing environments
- Network made via [docker-compose](https://docs.docker.com/compose/ "Docker Compose Page") which consists of a client, server and database image working together in isolation
- [GitHub Actions](https://github.com/features/actions "GitHub Actions Page") CI/CD pipeline that runs tests for both the client and the server when a commit or pull request is made to the main branch, and then deploys the application to AWS
- Automatic [AWS](https://aws.amazon.com/ "AWS Page") deployment on EC2 instance

## Features

- User authentication and authorization using JWT
- CRUD operations for user profiles and saved images
- Random dog image fetching and display
- Image favoriting system
- Persistent storage of liked images

## Project Structure

```bash
├── client # Contains the React front-end of the application
|  ├── public # Static assets
|  └── src
|     ├── api # Manages API requests to the server
|     ├── components
|     |  ├── Auth # Auth View
|     |  ├── Error # Common Error Boundary Component
|     |  ├── Landing # Landing View
|     |  ├── Layout # Common Layouts
|     |  ├── Profile # Profile View
|     |  └── ui # Common, reusable UI elements
|     ├── contexts # Holds React Context providers
|     ├── helpers # Utility functions shared across the application
|     ├── hooks # Custom React hooks for encapsulating reusable logic
|     ├── styles
|     |  ├── abstracts # CSS Files for Abstract Variables (colors, font weight, etc)
|     |  └── base # CSS Reset, global styles
|     ├── types # Common typescript types
|     └── __tests__
├── scripts # Holds utility scripts for building, deploying, or managing the project
└── server # Contains the back-end server logic, written in Go
   ├── cmd
   |  └── api # Entrypoint for the Go server
   ├── config # Global config file
   ├── db
   |  ├── queries # SQL Query functions
   |  └── schema # DB Schema
   ├── docs # Swagger documentation for the API
   └── internal # Core functionality that is not meant to be exported to external repositories
      ├── api
      |  ├── handlers # Process incoming requests and produce responses
      |  ├── middleware # Middleware functions for request/response processing
      |  ├── repositories # Data access layer ofr interacting with the database
      |  └── services # Business logic layer
      ├── errors # Custom error types and error-handling logic.
      ├── models # Data models
      ├── server # Core server setup and initialization code
      ├── testing # Testing utilities
      └── utils # Shared utility functions

```

## Getting Started

### Prerequisites

- Git
- Docker and Docker Compose

### Local Development Setup

1. Clone the repository:

```bash
git clone https://github.com/dariomnalerio/world-travel-inc-tech-interview
cd world-travel-inc-tech-interview
```

2. Set `.env `file in root folder. The `.env.example` file is provided as a template
3. Build and start Docker

```bash
docker-compose -f docker-compose.dev.yml up
```

### Test Development Setup

1. Clone the repository:

```bash
   git clone https://github.com/dariomnalerio/world-travel-inc-tech-interview
   cd world-travel-inc-tech-interview
```

2. Set `.env `file in root folder. The `.env.example` file is provided as a template
3. Build and start Docker

   ```bash
    docker-compose -f docker-compose.test.ymlup
   ```

Test environment is not configured yet. So running tests there is the same as doing it in the development environment.

### Production Development Setup

Production builds are automatically updated with GitHub Actions on commit or pull request to the main branch.

The process includes DockerHub, GitHub Actions and AWS services.

## API Documentation

For API documentation, refer to [swagger](http://ec2-18-216-189-146.us-east-2.compute.amazonaws.com:8080/api/v1/swagger/index.html "Swagger Documentation"). Note this is for documentation purposes only. The API can only be used in local development.

## Testing

You can run tests either locally or on the test environment

```bash
# Back-end tests
cd server
go test ./...

# Front-end tests
cd client
npm run test
```
