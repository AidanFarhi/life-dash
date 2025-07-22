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

  new Chart(
    document.getElementById('expenses'),
    {
      type: 'bar',
      options: {
        animation: true,
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
            beginAtZero: true
          },
          x: {
            ticks: {
              font: {
                size: 20
              }
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