

function Login() {
    $.ajax({
        url: "/user/login",
        type: "post",
        data: {
            "Email": $("#loginEmail").val(),
            "Password": $("#loginPassword").val(),
        },
        success: function (data) {
            if (data && data.Code == 0) {
                window.location.reload(true)
            } else {
                var loginError = $("#loginError")
                loginError.text(data.Message)
                loginError.removeClass("hidden")
            }
        },
    })
}

function Register() {
    $.ajax({
        url: "/user/register",
        type: "post",
        data: {
            "Email": $("#regEmail").val(),
            "Name": $("#regName").val(),
            "Password": $("#regPassword").val(),
        },
        success: function (data) {
            if (data && data.Code == 0) {
                $("#loginEmail").val($("#regEmail").val())
                $("#loginPassword").val($("#regPassword").val())
                Login()
            }else {
                var regError = $("#regError")
                regError.text(data.Message)
                regError.removeClass("hidden")
            }
        },
    })
}

$(document).ready(function () {
    $("#loginButton").click(Login)
    $("#regButton").click(Register)
})