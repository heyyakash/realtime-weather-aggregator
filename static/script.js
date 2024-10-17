const eventSource = new EventSource("http://localhost:8080/events");
const cityCharts = {};


async function fetchCityData(city) {
    try {
        const response = await fetch(`http://localhost:8080/daily?city=${city}`);
        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }
        return await response.json();
    } catch (error) {
        console.error('Failed to fetch city data:', error);
        return null;
    }
}
async function fetchCityHistoryData(city) {
    try {
        const response = await fetch(`http://localhost:8080/history?city=${city}`);
        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }
        return await response.json();
    } catch (error) {
        console.error('Failed to fetch city data:', error);
        return null;
    }
}

async function AddAlert(alert) {
    console.log(alert)
    container = document.getElementById("alerts-container")
    container.innerHTML += `<div class = "alert-box">Temperature in <b>${alert.city}</b> has exceeded the threshold; Current Temperature : <b>${alert.temperature}</b>`
}


eventSource.onmessage = (e) => {
    try {
        const data = JSON.parse(e.data);
        console.log("Received JSON data:", data);
        if (data.eventType === "weather_data") {
            updateChart(data);
        } else if (data.eventType === "alert_data") {
            AddAlert(data)
        }

    } catch (error) {
        console.error("Error parsing JSON:", error);
    }
};

function updateChart(data) {
    const { city, temperature, dt } = data;

    if (!cityCharts[city]) {
        cityCharts[city] = createChart(city);
    }

    const chart = cityCharts[city];
    chart.data.labels.push(new Date(dt * 1000).toLocaleTimeString())
    chart.data.datasets[0].data.push(temperature);
    chart.update();
}

function getRandomColor() {
    const r = Math.floor(Math.random() * 256);
    const g = Math.floor(Math.random() * 256);
    const b = Math.floor(Math.random() * 256);
    return `rgba(${r}, ${g}, ${b},`;
}

function createChart(city) {
    const chartContainer = document.createElement('div');
    chartContainer.classList.add('chart-container');
    chartContainer.innerHTML = `
        <h2>${city}</h2>
        <canvas id="${city}-chart"></canvas>
        <div class = "button-container">
        <button style = "margin:10px" id="${city}-fetch-button">Fetch Daily Aggregates</button>
        <button style = "margin:10px" id="${city}-history-button">Fetch Historical Data</button>
        </div>
        <div id="${city}-weather-data" class="weather-data"></div>
    `;
    document.getElementById('charts-container').appendChild(chartContainer);
    
    const chart = CreateNewChart(`${city}-chart`)

    // Add event listener for the button
    document.getElementById(`${city}-fetch-button`).addEventListener('click', async () => {
        const data = await fetchCityData(city);
        if (data) {
            displayWeatherData(city, data);
        }
    });

    document.getElementById(`${city}-history-button`).addEventListener('click', async () => {
        const data = await fetchCityHistoryData(city)
        if (data) {
            displayHistoryChart(city, data)
        }
    })

    return chart;
}

function CreateNewChart(id) {
    const randomColor = getRandomColor()
    const ctx = document.getElementById(id).getContext('2d');
    const chart = new Chart(ctx, {
        type: 'line',
        data: {
            labels: [],
            datasets: [{
                label: 'Temperature (°C)',
                data: [],
                borderColor: randomColor + '1)',
                backgroundColor: randomColor + '0.2)',
                borderWidth: 2,
                fill: true,
            }],
        },
        options: {
            responsive: true,
            scales: {
                x: {
                    title: {
                        display: true,
                        text: 'Time',
                    },
                },
                y: {
                    title: {
                        display: true,
                        text: 'Temperature (°C)',
                    },
                    beginAtZero: true,
                },
            },
        },
    });
    return chart
}

function displayWeatherData(city, data) {
    const weatherDataContainer = document.getElementById(`${city}-weather-data`);
    weatherDataContainer.innerHTML = `
        <p><strong>Average Temp:</strong> ${data.avgTemp} °C</p>
        <p><strong>Max Temp:</strong> ${data.maxTemp} °C</p>
        <p><strong>Min Temp:</strong> ${data.minTemp} °C</p>
        <p><strong>Dominant Weather:</strong> ${data.dominantWeather}</p>
        <p><strong>Day:</strong> ${new Date(data.day).toLocaleDateString()}</p>
    `;
}


function displayHistoryChart(city, data) {
    container = document.getElementById("historical-data")
    container.innerHTML = `
            <h2>${city}</h2>
        <canvas id="${city}-history-chart"></canvas>`
    const chart = CreateNewChart(`${city}-history-chart`)
    data.map((d) => {
        chart.data.labels.push(new Date(d.dt * 1000).toLocaleTimeString())
        chart.data.datasets[0].data.push(d.temperature);
        chart.update()
    })
}