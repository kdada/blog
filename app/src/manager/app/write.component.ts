import {Component, ElementRef, ViewChild} from "@angular/core"
import {Category, CategoryService} from "./category.service";


@Component({
    templateUrl: `/tmpl/write.html`,
    viewProviders: [CategoryService]
})
export class WriteComponent {
    public Title: string
    public Content: string
    public Status: string
    public Error: string
    public categories: Category[]
    private preStatus: string
    @ViewChild("wall")
    private wall: ElementRef
    private converter: showdown.Converter
    private change: () => void
    public constructor(private categoryService: CategoryService) {
        this.Status = ""
        this.Error = ""
        this.Title = ""
        this.Content = ""
        this.converter = new showdown.Converter()
        this.change = _.debounce(() => {
            this.Markdown()
        }, 100)
        categoryService.List().then(c => this.categories = c)
    }
    // 文章内容改变时自动重新生成markdown
    public Change() {
        this.change()
    }
    // 重新生成markdown
    public Markdown() {
        var ele = <HTMLElement>this.wall.nativeElement
        ele.innerHTML = this.converter.makeHtml(this.Content)
        var codes = ele.getElementsByTagName("code")
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
    // 提交
    public Submit() {
        if (this.Title.length <= 0) {
            this.Error = "请填写文章标题"
            return
        }
        if (this.Content.length <= 0) {
            this.Error = "请填写文章内容"
            return
        }
        this.Error = ""
        console.log(this.Title)
        console.log(this.Content)
    }
    public ngAfterViewChecked() {
        if (this.wall && this.Content && this.Status != this.preStatus) {
            this.Markdown()
        }
        this.preStatus = this.Status
    }
}