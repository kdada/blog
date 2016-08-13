

// Validate 修改ele元素的验证状态
function Validate(ele: HTMLInputElement, v: boolean) {
    if (v) {
        $(ele).data("valid", true)
        $(ele).parent().removeClass("has-error")
        $(ele).parent().addClass("has-success")
    } else {
        $(ele).data("valid", false)
        $(ele).parent().removeClass("has-success")
        $(ele).parent().addClass("has-error")
    }
}

// EmailChaned 邮箱校验
var emailRegexp = /\w+[\.\w]*@\w+(\.\w+)+/
function EmailChaned() {
    var v: string = $(this).val()
    Validate(this, emailRegexp.test(v) && v.length <= 100)
}

// PasswordChaned 密码校验
function PasswordChaned() {
    var v: string = $(this).val()
    Validate(this, v.length >= 6 && v.length <= 15)
}

// NameChaned 昵称校验
function NameChaned() {
    var v: string = $(this).val()
    Validate(this, v.length >= 2 && v.length <= 10)
}

// CnfChaned 重复密码校验
function CnfChaned() {
    var v: string = $(this).val()
    var y: string = $("#regPassword").val()
    Validate(this, v == y)
}


function CheckValid(ids: [string]): boolean {
    for (var i = 0; i < ids.length; i++) {
        var ele = $("#" + ids[i])
        if (ele.data("valid") != true) {
            ele.focus()
            return false
        }
    }
    return true
}

// LoginRequest 登录请求
function LoginRequest(email: string, password: string, success: (data: any) => void) {
    $.ajax({
        url: "/user/login",
        type: "post",
        data: {
            "Email": email,
            "Password": password,
        },
        success: success,
    })
}
// Login 登录
function Login() {
    if (CheckValid(["loginEmail", "loginPassword"])) {
        LoginRequest($("#loginEmail").val(), $("#loginPassword").val(), (data: any) => {
            if (data && data.Code == 0) {
                window.location.reload(true)
            } else {
                var loginError = $("#loginError")
                loginError.text(data.Message)
                loginError.removeClass("hidden")
            }
        })
    }
}

// Logout 登出
function Logout() {
    $.ajax({
        url: "/user/logout",
        type: "post",
        success: function() {
            window.location.reload(true)
        },
    })
}

// Register 注册
function Register() {
    if (CheckValid(["regEmail", "regName", "regPassword", "regCnfPassword"])) {
        var email:string = $("#regEmail").val()
        var password:string = $("#regPassword").val()
        $.ajax({
            url: "/user/register",
            type: "post",
            data: {
                "Email": email,
                "Name": $("#regName").val(),
                "Password": password,
            },
            success: function (data) {
                if (data && data.Code == 0) {
                    LoginRequest(email,password,(data:any)=>{
                        window.location.reload(true)
                    })
                } else {
                    var regError = $("#regError")
                    regError.text(data.Message)
                    regError.removeClass("hidden")
                }
            },
        })
    }
}

$(document).ready(function () {
    $("#loginEmail").on("input", EmailChaned)
    $("#loginPassword").on("input", PasswordChaned)
    $("#regEmail").on("input", EmailChaned)
    $("#regName").on("input", NameChaned)
    $("#regPassword").on("input", PasswordChaned)
    $("#regCnfPassword").on("input", CnfChaned)

    $("#loginButton").click(Login)
    $("#regButton").click(Register)
    $("#logout").click(Logout)
})