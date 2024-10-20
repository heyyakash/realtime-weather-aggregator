# Realtime Weather Aggregator Application
## Introduction
This application serves a Realtime Weather data Aggregator for Indian metro cities (Delhi,Mumbai, Chennai, Bangalore, Kolkata, Hyderabad). It is written in Golang and utilizes SSE (Server Side Events) to stream weather updates as well as alerts to users.

## Tools used
1. ChartJS : To visualize weather data
2. OpenWeatherMap API :  To fetch weather data on regular intervals
3. MongoDB : To store json data

## Why use MongoDB?
1. It is a NOSQL database, providing flexible schema to dump irregular weather data.
2. It provides ETL functionalities out of the box.
3. MongoDB's powerful aggregation pipeline supports complex data processing, like calculating averages, trends, or anomalies in weather patterns, without needing to extract and transform the data externally.

## How is weather data is being streamed in realtime?
The latest weather update is pushed to frontend through SSE (Server Side Events). Few advantages of SSE include
- SSE supports automatic reconnection natively if the connection is lost, without the need for custom logic or handling reconnection attempts.
- SSE is built on HTTP/1.1 and is much simpler to set up compared to WebSockets, which require a full-duplex connection.
- SSE is optimized for unidirectional communication, making it an excellent fit for streaming real-time weather updates where the client needs to receive data continuously from the server

## Requirements to run the application
1. Go [How to install go?](https://go.dev/doc/install)
2. Docker (Optional) [How to install Docker?](https://docs.docker.com/engine/install/)
3. MongoDB Atlas cluster [How to create MongoDb Altas Cluster?](https://www.mongodb.com/docs/guides/atlas/cluster/)
4. MongoDB Connection string [How to get MongoDB connection string?](https://www.geeksforgeeks.org/how-to-get-the-database-url-in-mongodb/)
5. OpenWeatherMap API KEY 

## How to run the application through go compiler
1. Clone this repository
2. Open the project directory locally
3. Create a `.env` file.
4. Enter the following details in the .env file
    ``` 
    API_KEY=<Your API Key>
    MONGO_URL=<Your MongoDB connection url>
    INTERVAL=<Time Interval between updates in seconds>
    ```
    Example 
    ``` 
    API_KEY=iYSPEtZ1nWGBlOg5MuY0PA==
    MONGO_URL=mongodb://username:password@host1:port1,host2:port2/database?option1=value1&option2=value2
    INTERVAL=120
    ```
5. Run the following command to generate a go executable
    ``` bash
    go build -o main .
    ```
4. A go executable called `main` must have been generated, to run the executable, run the following on your termninal
    ``` bash
    ./main
    ```
5. If the application has started successfully, you'll get following message. By default, the server will run on port 8080
    ``` bash
    2024/10/20 09:32:11 Pinged your deployment. You successfully connected to MongoDB!
    2024/10/20 09:32:11 Server Started at port 8080
    ```
6. Head over to `http://localhost:8080/static/` to launch the web interface, which instantly subscribes to the weather updates

## How to run the application through Docker
1. Clone this repository
2. Open the project directory locally
3. Create a `.env` file.
4. Enter the following details in the .env file
    ``` 
    API_KEY=<Your API Key>
    MONGO_URL=<Your MongoDB connection url>
    INTERVAL=<Time Interval between updates in seconds>
    ```
    Example 
    ``` 
    API_KEY=iYSPEtZ1nWGBlOg5MuY0PA==
    MONGO_URL=mongodb://username:password@host1:port1,host2:port2/database?option1=value1&option2=value2
    INTERVAL=120
    ```
5. Run the following command to build a docker image
    ``` bash
    docker build -t realtime-app .
    ```
6. Once the docker image is created run the following command to start the server at portt 8080
    ``` bash
    docker run -it -p 8080:8080 --rm --name realtime-app realtime-app
    ```
7. Once the server starts successfully, you'll get following response
    ``` bash
    2024/10/20 04:09:56 Pinged your deployment. You successfully connected to MongoDB!
    2024/10/20 04:09:56 Server Started at port 8080
    ```
8. Head over to `http://localhost:8080/static/` to launch the web interface, which instantly subscribes to the weather updates

## Screenshots
### Realtime weather updates
![Screenshot from 2024-10-20 09-15-04](https://github.com/user-attachments/assets/7d270208-fabc-40b4-af54-f5b15347954c)
### Realtime Alerts
![Screenshot from 2024-10-20 09-15-13](https://github.com/user-attachments/assets/7d66f410-3927-4aca-bc91-a86433b3e7a7)
### Historical Data
![Screenshot from 2024-10-20 09-52-53](https://github.com/user-attachments/assets/b59e3cb3-ff5e-4816-8cbb-150e8fa695d8)


## Functionalities
1. Realtime Weather Data updates visualization.
2. Visual Alerts issued, when the temperature of a city exceeds user defined threshold.
3. Daily aggregates such as `min_temp`,`max_temp,`feels_like`, `dominant_weather`, calculated and displayed on the fly.
4. Historical data can be accessed and visualized.
