// hamburger toggle
window.addEventListener("load", (ev) => {
    let toggleBtn = document.getElementById("sidebar-toggle-button");
    toggleBtn.addEventListener("click", (ev) => {
        document.getElementById("sidebar-container").classList.toggle("active");
    });
});
