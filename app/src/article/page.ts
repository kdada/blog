$(document).ready(function () {
    RefreshPage()
    RegisterReply()
})

// RefreshPage 刷新页码
function RefreshPage() {
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

// UpdateReplies 更新评论
function UpdateReplies(replies: any) {
    var content = $(".reply-content")
    content.html("")
    if (replies != null && replies.length) {
        for (var i = 0; i < replies.length; i++) {
            var r = replies[i]
            var date = new Date(r.CreateTime)
            var dateStr = date.toLocaleString('zh-cn',{
                year:"numeric",
                month:"2-digit",
                day:"2-digit",
                hour:"2-digit",
                minute:"2-digit",
                second:"2-digit",
                hour12:false,
            }).replace(/[年月]/g,"-").replace("日","")
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
function Jump(page: number) {
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

// RegisterReply 注册评论按钮事件
function RegisterReply() {
    $("#ShowReply").click(function(){
        $("#replyId").val("0")
        $("#Reply").modal("show")
    })
    $("#replyButton").click(function () {
        //评论
        $.ajax({
            url: "/reply/new",
            type: "post",
            data: {
                "Article": $("#articleId").val(),
                "Reply": $("#replyId").val(),
                "Content": $("#replyContent").val(),
            },
            success: function (data) {
                $("#Reply").modal("hide")
                $("#replyContent").val("")
                if (data && data.Code == 0) {
                    Jump(1)
                }
            },
        })
    })
}
