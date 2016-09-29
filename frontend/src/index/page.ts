$(document).ready(function () {
    RefreshPage()
    RegisterReply()
})

// RefreshPage 刷新页码
export function RefreshPage() {
    var pagination = $(".pagination")
    pagination.html("")
    var pages = pagination.data("pages")
    if (pages > 0) {
        var current = pagination.data("current")
        var disabled = ""
        var onclick = `onclick="Jump(1)"`
        if (current == 1) {
            disabled = ` class="disabled" `
            onclick = ""
        }
        // 添加首页按钮
        pagination.append(`<li ` + disabled + `>
        <a href="javascript:void(0);" ` + onclick + ` aria-label="Previous">
          <span aria-hidden="true">&laquo;</span>
        </a>
      </li>`)
        // 添加当前页的附近7页
        var start = current > 3 ? current - 3 : 1
        var end = pages - start >= 6 ? start + 6 : pages
        start = start - 1 > 6 - (end - start) ? end - 6 : 1
        for (var i = start; i <= end; ++i) {
            if (i == current) {
                pagination.append(`<li class="page` + i + ` active"><a href="javascript:void(0);">` + i + `</a></li>`)
            } else {
                pagination.append(`<li class="page` + i + `"><a href="javascript:void(0);"  onclick="Jump(` + i + `)">` + i + `</a></li>`)
            }
        }

        disabled = ""
        onclick = `onclick="Jump(` + pages + `)"`
        if (current == pages) {
            disabled = ` class="disabled" `
            onclick = ""
        }
        //添加尾页按钮
        pagination.append(`<li ` + disabled + `>
        <a href="javascript:void(0);" `+ onclick + ` aria-label="Next">
          <span aria-hidden="true">&raquo;</span>
        </a>
      </li>`)
    }
}

// RegisterReply 注册评论按钮事件
function RegisterReply() {
    $("#ShowReply").click(function () {
        $("#replyUser").val("0")
        $("#replyDialog").modal("show")
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
                if (data && data.Code == 0) {
                    $("#replyDialog").modal("hide")
                    $("#replyContent").val("")
                    Jump(1)
                }
            },
        })
    })
}


// UpdateReplies 更新评论
export function UpdateReplies(replies: any) {
    function resolveNum(num: number): string {
        if (num < 10) {
            return "0" + num
        }
        return "" + num
    }
    var content = $(".reply-content")
    content.html("")
    if (replies != null && replies.length) {
        for (var i = 0; i < replies.length; i++) {
            var r = replies[i]
            var date = new Date(r.CreateTime)
            var dateStr = date.getFullYear() + "-" + resolveNum(date.getMonth()) + "-" + resolveNum(date.getDate()) + " " + resolveNum(date.getHours()) + ":" + resolveNum(date.getMinutes()) + ":" + resolveNum(date.getSeconds())
            content.append(`
            <div class="well">
                <h5 class="media-heading">`+ r.Floor + `楼 ` + dateStr + `</h5>
                <h5 class="media-heading">`+ r.Name + `:</h5>
                `+ r.Content + `
            </div>
            `)
        }
    }
}

// Jump 跳转
export function Jump(page: number) {
    var pagination = $(".pagination")
    var article = pagination.data("article")
    $.ajax({
        url: "/reply/list",
        type: "post",
        data: {
            "Article": article,
            "Page": page,
        },
        success: function (data) {
            if (data && data.Code == 0) {
                pagination.data("current", page)
                RefreshPage()
                UpdateReplies(data.Data)
            } else {
                console.log(data)
            }
        },
    })
}

