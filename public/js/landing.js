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
        const alrt =  $('.status-alert');
        $.ajax({
            type: "POST",
            url: "/apply",
            data: { email : email}
        }).done(function() {
            alrt.text("CONFIRMATION SENT TO " + email.toUpperCase());
            alrt.slideDown().delay(1000).slideUp();
        }).fail(function() {
            alrt.text("UNKNOWN ERROR SENDING EMAIL");
            alrt.slideDown().delay(1000).slideUp();
        });
    });
    // $('.upload-csv-form').submit((ev)=>{
    //     ev.preventDefault();
    //     const fData = new FormData(this);
    //     console.log(fData);
    //     const alrt =  $('.status-alert');
    //     $.ajax({
    //         type: "POST",
    //         url: "/upload",
    //         data: fData
    //     }).done(function() {
    //         alrt.text("FILE UPLOADED");
    //         alrt.slideDown().delay(1000).slideUp();
    //     }).fail(function() {
    //         alrt.text("UNKNOWN ERROR UPLOADING FILE");
    //         alrt.slideDown().delay(1000).slideUp();
    //     });
    // });
});
