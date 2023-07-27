# FinGreat: Fintech Application Backend

Welcome to the backend section of the FinGreat fintech application. This repository houses all the server-side code for our application, built using GoLang.

## Introduction

FinGreat is a comprehensive fintech application aiming to simplify and secure financial transactions. This backend handles everything from user management to account handling and money transfers. It's built with GoLang, ensuring efficient execution and powerful performance.

Our backend is divided into several modules, each focusing on a specific feature:

- **User Module**: Implements CRUD operations (Create, Read, Update, and Delete) for managing user data and handles user authentication and authorization.

## Setup and Installation

Before you begin, make sure you have GoLang, Docker, docker-compose, and GNU-make installed in your local environment.

Next, clone the repository to your local machine:

```bash
git clone https://github.com/yourgithubusername/FinGreat-backend.git
```
Navigate into the project directory:

```bash
cd fingreat-backend
```

To set up the project environment, execute the following command:

```bash
make setup
```

This command will start up a Postgres server, create a database named 'fingreat_db', and run the migrations to the database.

Before starting the server, install compileDaemon:

```bash
go install github.com/githubnemo/CompileDaemon
```

Now, you can start the server with:

```bash
make start
```

You can then access the server at localhost:8000 (or whichever port you've configured it to run on).
