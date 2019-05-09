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
    $('.waitlist-form').submit((ev)=>{
        ev.preventDefault();
        const email = $('.waitlist-form .email-input').val();
        $.ajax({
            type: "POST",
            url: "/apply",
            data: { email : email}
        }).done(function() {
            alert( "Confirmation sent to " + email);
            $('.waitlist-form').reset()
        }).fail(function() {
            alert( "error" );
        });
    })
});
