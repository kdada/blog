$(document).ready(function () {
    // 对markdown进行处理
    var converter = new showdown.Converter()
    var md = $(".markdown-body")
    md.html(converter.makeHtml(md.text()))
    var codes = md.find("code")
    for (var i = 0; i < codes.length; i++) {
        var code = codes[i]
        hljs.highlightBlock(code)
        var rowNum = code.innerHTML.split('\n').length - 1
        var rowDiv = document.createElement('div')
        rowDiv.className = 'code-row-space'
        for (var j = 1; j <= rowNum; j++) {
            rowDiv.innerHTML += '<span>' + j.toString() + '</span>\n'
        }
        code.parentElement.insertBefore(rowDiv, code)
    }
})