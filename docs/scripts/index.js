"use strict";

function setTheme(theme) {
    localStorage.setItem("theme", theme);

    let isDark = theme === "dark";
    document.getElementById("theme-toggle").setAttribute("checked", isDark);

    document.body.setAttribute("data-theme", theme)
}

window.addEventListener("load", (ev) => {
    let theme = localStorage.getItem("theme") || "light";
    setTheme(theme);

    let sidebarToggleBtn = document.getElementById("sidebar-toggle");
    sidebarToggleBtn.addEventListener("click", (ev) => {
        document.getElementById("sidebar-container").classList.toggle("active");
    });

    let themeToggleBtn = document.getElementById("theme-toggle");
    themeToggleBtn.addEventListener("click", (ev) => {
        theme = ev.target.checked ? "dark" : "light";
        setTheme(theme);
    });
});
