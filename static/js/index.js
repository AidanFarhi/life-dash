const navDivs = document.querySelectorAll('.nav-div')

navDivs.forEach(navDiv => {
    navDiv.addEventListener('click', () => {
        navDivs.forEach(el => el.classList.remove('active'))
        navDiv.classList.add('active')
    })
})