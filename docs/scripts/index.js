// hamburger toggle
window.addEventListener("load", (ev) => {
    let toggleBtn = document.getElementById("sidebar-toggle-button");
    console.log("LOADED")
    toggleBtn.addEventListener("click", (ev) => {
        console.log("CLICKED")
        document.getElementById("sidebar-container").classList.toggle("active");
    });
});
