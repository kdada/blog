
// md转换器
var converter = new showdown.Converter()
// MarkdownString 从content生成md
export function MarkdownString(content: string): HTMLElement {
    var wall = document.createElement("div")
    wall.innerHTML = converter.makeHtml(content)
    var codes = wall.getElementsByTagName("code")
    for (var i = 0; i < codes.length; i++) {
        var code = codes.item(i)
        hljs.highlightBlock(code)
        code.className += " code-block"
        var rowNum = code.innerHTML.split('\n').length - 1
        var rowDiv = document.createElement('div')
        rowDiv.className = 'code-line'
        for (var j = 1; j <= rowNum; j++) {
            rowDiv.innerHTML += '<span>' + j.toString() + '</span>\n'
        }
        code.parentElement.insertBefore(rowDiv, code)
    }
    return wall
}