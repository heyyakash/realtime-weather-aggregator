const eventSource = new EventSource("http://localhost:8080/events");
eventSource.onmessage = (e) => {
    console.log("received something")
    console.log(e)
}
