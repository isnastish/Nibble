# Overview 

# Running Nibble service
To boot up a `nibble` service you should have `docker` installed on your machine. When that is done, run `sudo docker compose up --build --force-recreate --no-deps`. That will start one docker container with an actual API service and another one with Postgres database. All the api keys are cofigured inside `compose.yaml`, so you don't need to do anything. If everything went right, you should see the logs:
![Screenshot from 2024-12-10 20-44-08](https://github.com/user-attachments/assets/488e4864-c894-4b81-a290-bdb8a030431e)


# Testing the API
Once `nibble` service is up and running, we can make queries to its API using any client. In my example I will be utilizing [curl](https://en.wikipedia.org/wiki/CURL).
In order to add a new user to a database, put the following command into 
your terminal, make sure that curl is installed: 
```curl
curl -X POST http://localhost:3030/signup -d 
    '{
        "first_name":"Alexey", 
        "last_name":"Yevtushenko",
        "password":"2341234@sdfss", 
        "email":"isnastish@gmail.com" 
    }' 
    -H "X-Forwarded-For: 34.21.9.50"  // Washington, United States
```
**NOTE**: `X-Forwarded-For` header could be omitted, in that case the IP would be deducted from the request URL, and, what's important, a query to an external service for detecting geolocation will fail with an error `RESERVED_IP_ADDRESS`, since `127.0.0.1` is reserved. You can specify any IP address you want with earlier mentioned header.

Now, you can go to your service and see, from the logs, that a corresponding user with its geolocation data was successfully added to a database.![Screenshot from 2024-12-10 20-23-41](https://github.com/user-attachments/assets/60982ffc-c078-4069-8633-36ac41a3d91c)
