$(document).ready(function () {
    // 对markdown进行处理
    var converter = new showdown.Converter()
    var mds = $(".article-summary")
    mds.each(function (index, elem) {
        var md = $(elem)
        md.html(converter.makeHtml(md.text()))
    })
})