import { Component, ElementRef, ViewChild } from '@angular/core';

class Artical {
  title: string
  content: string
}



@Component({
  selector: 'my-app',
  template: '<article><div #markdown class="markdown-body" [innerHtml]="at.content"></div></article>'
})
export class AppComponent {
  @ViewChild("markdown")
  view: ElementRef
  at: Artical
  constructor() {
    var converter = new showdown.Converter()
    this.at = {
      title: "无标题",
      content: converter.makeHtml(`
#中问打算  测试测试
##asdasd  
###asdsad  
1. asdasd  
2. sadsad2

换行测试  
结尾必须双空格
换行测试

换行测试

> asdasd  
\`\`\`go
package main

func main() {
  var i = 0
  return 0
  asdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsdasdsd

}
\`\`\`
[百度](http://www.baidu.com)  
![Logo](http://www.baidu.com/img/baidu_jgylogo3.gif)  
**asdasd**  
*asdasd*  
  
\`\`\`go
package main

func main() {
  var i = 0
  return 0
}
\`\`\`

      `)
    };

  }
  ngAfterViewInit() {
    var codes = (<HTMLDivElement>(this.view.nativeElement)).getElementsByTagName("code")
    for (var i = 0; i < codes.length; i++) {
      var code = codes.item(i)
      hljs.highlightBlock(code)
      var rowNum = code.innerHTML.split('\n').length - 1
      var rowDiv = document.createElement('div')
      rowDiv.className = 'code-row-space'
      for (var j = 1; j <= rowNum; j++) {
        rowDiv.innerHTML += '<span>' + j.toString() + '</span>\n'
      }
      code.parentElement.insertBefore(rowDiv, code)
    }
  }
}