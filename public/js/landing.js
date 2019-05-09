$(document).ready(function () {
    $(".show-not-available").click((ev)=>{
         $('.status-alert').text("Action Not Available in Demo");
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
        const inputF = $('.waitlist-form .email-input');
        const email = inputF.val();
        inputF.val("");
        $.ajax({
            type: "POST",
            url: "/apply",
            data: { email : email}
        }).done(function() {
            $('.status-alert').text("Confirmation sent to " + email);
        }).fail(function() {
            $('.status-alert').text("Unknown Error Sending Email");
        });
    })
});
