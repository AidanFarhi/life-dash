(async function() {
  const data = [
    { category: '🍎', count: 450 },
    { category: '🐾', count: 320 },
    { category: '🚗', count: 210 },
    { category: '🏠', count: 980 },
    { category: '🎬', count: 275 },
    { category: '🏥', count: 430 },
    { category: '💡', count: 360 },
    { category: '📚', count: 150 },
    { category: '👕', count: 190 },
    { category: '📦', count: 120 }
  ];

  const resp = await fetch('/api/expenses')
  const jsonData = await resp.json()
  console.log(jsonData)

  new Chart(
    document.getElementById('expenses-canvas'),
    {
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
        labels: data.map(row => row.category),
        datasets: [
          {
            data: data.map(row => row.count)
          }
        ]
      }
    }
  );
})();