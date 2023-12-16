# Library Management System

A simple Library Management System implemented in Go (Golang) with GOFR and MySQL.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
- [Usage](#usage)
- [Database Schema](#database-schema)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Overview

The Library Management System is a project implemented in Go and GOFR that allows users to manage books, authors, and student information. It uses GORM as the language and MySQL as the database.

## Features

- Add, update,  delete  and search for entries of books in database 
- Manage author information
- Track student information and book issuance
- View a list of books, authors, and student records

## Getting Started

### Prerequisites

Make sure you have the following installed:

- Go (Golang)
- MySQL

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/ujjwall9837/GOFR_CRUD.git
    cd GOFR_CRUD
    ```

2. Install dependencies:

    ```bash
    go get gofr.dev
    docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=password -p 2001:9092 -d mysql:8.0.30
    ```

### Configuration

1. Set up the MySQL database:

    - Create a new database.
    - Update the database connection details in `config/config.go`.

2. Run the application:

    ```bash
    go run main.go
    ```

## Usage

1. Access the application by visiting [http://localhost:9092](http://localhost:9092) in your web browser.
2. Use the web interface to manage books, authors, and student records.

## Database Schema

The project uses the following database schema:

- `entries`
  - `bookid` (Primary Key)
  - `book_name`
  - `author_name` 
  - `issued_to`
  - `issued_date`

## Testing

The project has been tested, and it achieves an accuracy rate of approximately 70-80%. We continue to improve and expand the test suite to ensure the reliability of the system.

<img src ="/images/post.png" width = "400" height = "200" >
*Figure 1: Create entry in database*

<img src ="/images/put.png" width = "400" height = "200" >

*Figure 2: Update entry in database*

<img src ="/images/get.png" width = "400" height = "200" >
*Figure 3: Get entry from database*

<img src ="/images/delete.png" width = "400" height = "200" >
*Figure 4: Delete entry from database*

## Contributing

Feel free to contribute to the project. Fork the repository, make your changes, and submit a pull request.


