
# File Management System

The File Management System is a simple web application built in Go using the Gin framework and MongoDB for file storage and tracking file transactions.

## Features:
- JWT User authentication for secure access to the system.
- API endpoints for uploading, downloading, and listing files.
- Storage of uploaded files in a local folder.
- Recording file transactions in MongoDB, including upload, download, and update operations.
- Retrieval of file transaction records via APIs.
## Additional features:
- Retrieval of file properties (file format, creation date, etc.).


## Technologies Used

1. [Gin](https://github.com/gin-gonic/gin): A web framework written in Go.
2. [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver): Official MongoDB driver for Go.
3. [Golang](https://go.dev/): Programming language used for backend development.


## Prerequisites

Before running the Log Ingestor, ensure you have the following dependencies installed:

1. Go (Golang)
2. MongoDB

- Linux
```bash
sudo apt update
sudo apt install mongodb
```
- macOS
```bash
brew tap mongodb/brew
brew install mongodb-community
```
- Windows
```bash
Download and install MongoDB from the official website: [MongoDB Downloads](https://www.mongodb.com/try/download/community)
```
## Installation

1. Clone the repository

```bash
git clone https://github.com/swag2716/File-Management-System.git
cd File-Management-System    
```
2. Install necessary Go packages
```bash
go get -u github.com/gin-gonic/gin
go get -u go.mongodb.org/mongo-driver/mongo    
``` 
3. Set up your environment variables:

 Create a .env file in the project root and configure the following variables for MongoDB connection string, jwt authentication and port:
 ```bash
 MONGODB_URL=<your connection string>
 SECRET_KEY=<random hex key>
 PORT=9000
 ```
4. Run the application
```bash
go run main.go
```
5. Access the application at http://localhost:9000
## API Reference

### Authentication:

`/signup`: POST - User signup.

`/login`: POST - User login.

### File Operations:

`/upload`: POST - Upload a file.

`/download/:file_id`: GET - Download a file.

`/upload/:file_id`: DELETE - delete a file.

`/files`: GET - Get all files in order of their sizes.

`/transactions`: GET - get all transactions for files