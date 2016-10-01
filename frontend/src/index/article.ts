import { MarkdownString } from '../common/markdown';

$(document).ready(function () {
    // 对文章的markdown进行处理    
    var height = $(document).scrollTop()
    var md = $(".markdown-body")
    var oldHeight = md.outerHeight()
    var ele = MarkdownString(md.text())
    md.html("")
    md.append(ele)
    var dH = md.outerHeight() - oldHeight
    if (height > 0) {
        // 跳至正确的位置,通常在定位评论时使用
        $(document).scrollTop(height + dH)
    }
})