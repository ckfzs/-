function active_sidebar(menu, submenu) {
    $("#side-menu li").removeClass("active");
    $("#side-menu li .nav-second-level").addClass("collapse");
    if (menu != null) {
        $("#side-menu " + menu).addClass("active");
        $("#side-menu " + menu + " .nav-second-level").removeClass("collapse");
    }
    if (submenu != null) {
        $("#side-menu "+ menu + " " + submenu).addClass("active");
    }
}