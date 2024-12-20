# Overview 
`Nibble` service written in Golang for the following assignment [Nibble assignment](https://makeheadway.notion.site/Backend-Engineer-Nibble-15316bdf9fd28052b718d2a1ea8358a5)

In a nutshell, it's a simple API server which accepts user data, makes a request to an external service for detecting user's geolocation based on its IP address, and stores all the information into a Postgres database. I have purchased an API key for this assignment with a limit of `30,000` requests per month, so that should be more than enough for our purposes. For more technical overview and design documentation please refer to [api_design document](/api_design.md)  

# Running Nibble service
To boot up a `nibble` service you should have `docker` installed on your machine. When that is done, run `sudo docker compose up`. That will start one docker container with an actual API service, and another one with Postgres database. All the api keys are cofigured inside `compose.yaml`, so you don't need to do anything. If everything went right, you should see the following logs:
![Screenshot from 2024-12-10 20-44-08](https://github.com/user-attachments/assets/488e4864-c894-4b81-a290-bdb8a030431e)


# Testing the API
Once `nibble` service is up and running, we can make queries to its API using any client. In my example I will be utilizing [curl](https://en.wikipedia.org/wiki/CURL).
In order to add a new user to a database, put the following command into 
your terminal, make sure that curl is installed: 
```
curl -X POST http://localhost:3030/signup -d \
'{
 "first_name": "Alexey",
 "last_name": "Yevtusnenko",
 "password": "234324@sd22ad",
 "email": "isnastish@gmail.com"
}' \
-H "X-Forwarded-For: 34.21.9.50"
```
**NOTE**: `X-Forwarded-For` header could be omitted, in that case the IP would be deducted from the request URL, and, what's important, a query to an external service for detecting geolocation will fail with an error `RESERVED_IP_ADDRESS`, since `127.0.0.1` is reserved. You can specify any IP address you want with an earlier mentioned header to simulate a real-world example, where clients are geographically distributed all over the world.

Now, you can go to your service and see, from the logs, that a corresponding user with its geolocation data was successfully added to a database.![Screenshot from 2024-12-10 20-23-41](https://github.com/user-attachments/assets/60982ffc-c078-4069-8633-36ac41a3d91c)
