# Overview 

# Running Nibble service

# Testing the API
Once `nibble` service is up and running, we can make queries to its API using any client. In my example I will be utilizing [curl](https://en.wikipedia.org/wiki/CURL)
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
**NOTE**: `X-Forwarded-For` header could be omitted, in that case the IP would be deducted from the request URL, and, what's important, a query to an external service for detecting geolocation will fail with an rror `RESERVED_IP_ADDRESS`, since 127.0.0.1 is reserved. You can specify any IP address you want with earlier mentioned header.

Now, you can go to your service and see, from the logs, that a corresponding user with its geolocation data was successfully added to a database.