# REST API in Golang

This is a simple REST API built with Go that allows you to perform CRUD (Create, Read, Update, Delete) operations on a collection of items.

## Features

- **GET All Items**: Retrieve a list of all items.
- **GET Item by ID**: Retrieve a specific item by its unique ID.
- **POST New Item**: Create a new item and add it to the collection.
- **UPDATE Item**: Update an existing item's information.
- **DELETE Item**: Remove an item from the collection.

## Installation

1. Make sure you have Go installed on your system.

2. Clone this repository:

   ```shell
   git clone https://github.com/ARUP-G/REST-API.git

3. Build and run the application:

    ```shell
    go build
    ./REST-API

4. The API server will be running at http://localhost:9000.

## Usage
### GET All Books

- **Endpoint**: /api/books
- **Method**: `GET`
- **Description**: Retrieve a list of all books.
- **Example Response**:
    ```json
    [
    {
        "id": "1",
        "isbn": "86348",
        "title": "Book1",
        "author": {
            "fname": "Ryan",
            "lname": "Gosling"
        }
    },
    {
        "id": "2",
        "isbn": "85578",
        "title": "Book2",
        "author": {
            "fname": "Arup",
            "lname": "Das"
        }
    }
    ]
 
    ```

### GET Books by ID
 - **Endpoint**: /api/books/{id}
 - **Method**: `GET`
 - **Description**: Retrieve a specific book by its unique ID.
 - **Example Request**: /api/books/2
 - **Example Response**:
    ```json
        {
            "id": "2",
            "isbn": "85578",
            "title": "Book2",
            "author": {
                "fname": "Arup",
                "lname": "Das"
                }
        }
 
    ```
###  Create Books
 - **Endpoint**: /api/books
 - **Method**: `POST`
 - **Description**: Create a new book and add it to the collection.
 - **Example Request**: /api/books

 - **Example Request Body**
    ```json
        {
            "isbn": "90078",
            "title": "New Book",
            "author": {
                "fname": "Chilian",
                "lname": "Murphy"
                }
        }
    ```
 - **Example Response**:
    ```json
        {
            "id": "60",
            "isbn": "90078",
            "title": "New Book",
            "author": {
                "fname": "Chilian",
                "lname": "Murphy"
            }
        }
 
    ```

###  Update Book
 - **Endpoint**: /api/books/{id}
 - **Method**: `UPDATE`
 - **Description**: Update existing book and add it to the collection.
 - **Example Request**: /api/books/60
 - **Example Request Body**
    ```json
        {
            "isbn": "90078",
            "title": "Updated Book",
            "author": {
                "fname": "Chilian",
                "lname": "Murphy"
                }
        }
    ```
 - **Example Response**:
     ```json
        {
            "id": "60",
            "isbn": "90078",
            "title": "Updated Book",
            "author": {
                "fname": "Chilian",
                "lname": "Murphy"
            }
        }
 
    ```

###  Delete Book
 - **Endpoint**: /api/books/{id}
 - **Method**: `DELETE`
 - **Description**: Delete a book from the collection.
 - **Example Request**: /api/books/0
- **Example Request Body**
    Deleted