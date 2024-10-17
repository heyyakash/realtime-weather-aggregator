const eventSource = new EventSource("http://localhost:8080/events");
const cityCharts = {};

eventSource.onmessage = (e) => {
    try {
        const data = JSON.parse(e.data);
        console.log("Received JSON data:", data);
        updateChart(data);
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

function createChart(city) {
    const chartContainer = document.createElement('div');
    chartContainer.classList.add('chart-container');
    chartContainer.innerHTML = `<h2>${city}</h2><canvas id="${city}-chart"></canvas>`;
    document.getElementById('charts-container').appendChild(chartContainer);

    const ctx = document.getElementById(`${city}-chart`).getContext('2d');
    const chart = new Chart(ctx, {
        type: 'line',
        data: {
            labels: [],
            datasets: [{
                label: 'Temperature (°C)',
                data: [],
                borderColor: 'rgba(75, 192, 192, 1)',
                backgroundColor: 'rgba(75, 192, 192, 0.2)',
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

    return chart;
}