$(document).ready(function () {
    $(".show-not-available").click((ev)=>{
         alert("This ability is not in the demo ...");
    });
    $(window).scroll(function() {
        if ($(this).scrollTop() < 16) {
            $('.nav-head').slideUp();
        }else{
            $('.nav-head').slideDown();
        }
    });
});
