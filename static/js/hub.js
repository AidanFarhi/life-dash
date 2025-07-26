
document.querySelectorAll(".hub-item").forEach(hubItem => {
    hubItem.addEventListener('click', () => {
        document.querySelectorAll('.nav-div').forEach(el => el.classList.remove('active'))
        document.getElementById(hubItem.dataset.customId).classList.add('active')
    })
})