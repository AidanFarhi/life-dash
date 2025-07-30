
const fetchExpenseDistributionData = async () => {
    const resp = await fetch('/api/aggregatedExpenses')
    const jsonData = await resp.json()
    return jsonData
}

const renderBarChart = async (canvas, chartData) => {
    let chartConfig = {
        type: 'bar',
        options: {
            animation: true,
            maintainAspectRatio: false,
            plugins: {
                legend: {
                    display: false
                },
                tooltip: {
                    enabled: false
                }
            },
            scales: {
                y: {
                    beginAtZero: true,
                    grid: {
                        display: false
                    }
                },
                x: {
                    ticks: {
                        font: {
                            size: 20
                        }
                    },
                    grid: {
                        display: false
                    }
                }
            }
        },
        data: {
            labels: chartData.map(row => row.category),
            datasets: [
                {
                    data: chartData.map(row => row.amount)
                }
            ]
        }
    }
    new Chart(canvas, chartConfig)
}

const fetchDataAndRenderBarChart = async () => {
    const canvas = document.getElementById('expenses-canvas')
    const data = await fetchExpenseDistributionData()
    const dataSortedByAmountDesc = data.sort((a, b) => b.amount - a.amount)
    await renderBarChart(canvas, dataSortedByAmountDesc)
}

document.body.addEventListener('htmx:afterSwap', event => {
    if (event.detail.pathInfo.requestPath === '/expenses') {
        fetchDataAndRenderBarChart()
    }
})