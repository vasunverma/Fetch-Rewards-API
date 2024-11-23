# Fetch Rewards API Service in GO

## Endpoints
- **POST /receipts/process**: Submit a receipt for processing and receive a receipt ID.
- **GET /receipts/{id}/points**: Retrieve calculated points for a receipt by its ID.

## Prerequisites
Before setting up, make sure you have the following installed:
- **Docker**: To build and run the containerized application.
- **Go**: Optional for running locally without Docker

## Setup

### 1. Clone the Repository
Start by cloning this repository to your local machine.
```sh
https://github.com/vasunverma/Fetch-Rewards-API.git
cd Fetch-Rewards-API
```

### 2. Build the Docker Image
To build the Docker image, run the following command:
```sh
docker build -t fetch-rewards-api .
```
This will build the Docker image using the provided Dockerfile.

### 3. Run the Docker Container
Once the image is built, you can run the container using:
```sh
docker run -p 8080:8080 fetch-rewards-api
```

### 4. Accessing the Endpoints
The service will be accessible at http://localhost:8080. I used postman to access it.

- POST /receipts/process: Submit a receipt for processing and receive a receipt ID.
    ```sh
    {
      "retailer": "TargetMart",
      "purchaseDate": "2024-11-23",
      "purchaseTime": "15:24",
      "items": [
        {"shortDescription": "Eggs", "price": "4.99"},
        {"shortDescription": "Milk", "price": "3.25"}
      ],
      "total": "9.25"
    }
    ```
- GET /receipts/{id}/points: Retrieve the calculated points for a receipt by its ID.
    ```sh
    {
      "points": 56
    }
    ```
