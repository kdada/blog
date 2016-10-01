(function (document: HTMLDocument) {
    $(document).ready(function () {
        RegisterReply()
        RegisterPage()
    })
    // RegisterPage 给页码按钮注册事件
    function RegisterPage() {
        $(".pagination").find("a").click(function () {
            RefreshPage($(this).data("page"))
        })
    }

    // RefreshPage 显示指定页的评论
    function RefreshPage(page: number) {
        $.ajax({
            url: "/reply/view",
            type: "post",
            data: {
                "Article": $(".reply").data("article"),
                "Page": page,
            },
            success: function (data) {
                $(".reply").html(data)
                RegisterPage()
            },
        })
    }

    // RegisterReply 注册评论按钮事件
    function RegisterReply() {
        $('#replyDialog').on('shown.bs.modal', function (e) {
            $("#replyContent").focus()
        })
        $("#replyButton").click(function () {
            //评论
            $.ajax({
                url: "/reply/create",
                type: "post",
                data: {
                    "Article": $("#replyArticle").val(),
                    "Reply": $("#replyUser").val(),
                    "Content": $("#replyContent").val(),
                },
                success: function (data) {
                    $("#replyDialog").modal("hide")
                    $("#replyContent").val("")
                    if (data && data.Code == 0) {
                        RefreshPage(1)
                    }
                },
            })
        })
    }

})(document)