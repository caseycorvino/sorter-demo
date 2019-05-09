$(document).ready(function () {
    $(".show-not-available").click((ev)=>{
        const alrt =  $('.status-alert');
        alrt.text("Action Not Available in Demo");
        alrt.show();
        alrt.fadeOut();
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
            alrt.text("Confirmation sent to " + email);
            alrt.show();
            alrt.fadeOut();
        }).fail(function() {
            const alrt =  $('.status-alert');
            alrt.text("Unknown Error Sending Email");
            alrt.show();
            alrt.fadeOut();
        });
    })
});
