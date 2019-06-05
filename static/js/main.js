var menu = $('.all-menu');
var miniMenu = $('.menu-section')
var btnOpen = $('.see-more-btn');

function allMenu() {
    btnOpen.toggle();
    menu.toggle("slow");
}
function closeMenu() {
    btnOpen.toggle();
    menu.toggle("slow");
    topDist = miniMenu.offset().top;
    $('body,html').animate({scrollTop: topDist}, 600);
}

