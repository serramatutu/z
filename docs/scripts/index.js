"use strict";

let themeToggleBtns;

function setTheme(theme) {
    localStorage.setItem("theme", theme);

    let isDark = theme === "dark";
    for (let i = 0; i < themeToggleBtns.length; i++) {
        const btn = themeToggleBtns.item(i);
        if (isDark) {
            btn.setAttribute("checked", "");
        }
        else {
            btn.removeAttribute("checked");
        }
    }

    document.body.setAttribute("data-theme", theme)
}

window.addEventListener("load", (ev) => {
    themeToggleBtns = document.getElementsByClassName("theme-toggle");
    for (let i = 0; i < themeToggleBtns.length; i++) {
        themeToggleBtns.item(i).addEventListener("click", (ev) => {
            theme = ev.target.checked ? "dark" : "light";
            setTheme(theme);
        })
    }

    let sidebarToggleBtn = document.getElementById("sidebar-toggle");
    sidebarToggleBtn.addEventListener("click", (ev) => {
        document.getElementById("sidebar-container").classList.toggle("active");
    });

    let theme = localStorage.getItem("theme") || "light";
    setTheme(theme);
});
