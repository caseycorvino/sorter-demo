$(document).ready(function () {
    $(".show-not-available").click((ev)=>{
        const alrt =  $('.status-alert');
        alrt.text("ACTION NOT AVAILABLE IN DEMO");
        alrt.slideDown().delay(1000).slideUp();
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
            const alrt =  $('.status-alert');
            alrt.text("CONFIRMATION SENT TO " + email.toUpperCase());
            alrt.slideDown().delay(1000).slideUp();
        }).fail(function() {
            const alrt =  $('.status-alert');
            alrt.text("UNKNOWN ERROR SENDING EMAIL");
            alrt.slideDown().delay(1000).slideUp();
        });
    })
});
