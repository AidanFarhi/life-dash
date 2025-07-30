import { fetchDataAndRenderBarChart } from './expenses.js'

// toggle active for links
const navDivs = document.querySelectorAll('.nav-div')
navDivs.forEach(navDiv => {
    navDiv.addEventListener('click', () => {
        navDivs.forEach(el => el.classList.remove('active'))
        navDiv.classList.add('active')
    })
})

// trigger chart loading functions based on which canvases are present
document.body.addEventListener('htmx:afterSwap', (event) => {
  if (event.target.id === 'main-content') {
    if (document.getElementById('expenses-canvas')) {
      fetchDataAndRenderBarChart()
    }
  }
})
