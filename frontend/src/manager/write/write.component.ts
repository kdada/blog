import { Component, ElementRef, ViewChild } from "@angular/core";
import { Router, ActivatedRoute, Params } from "@angular/router";
import { CategoryDetail } from '../category/category.model';
import { ArticleDetail } from '../article/article.model';
import { WriteService } from './write.service';
import { FileService } from '../file/file.service';
import { MarkdownString } from '../../common/markdown';

@Component({
    templateUrl: "/tmpls/write.html",
    providers: [WriteService, FileService]
})
export class WriteComponent {
    private writable: boolean
    private title: string
    private category: number
    private content: string
    private categories: CategoryDetail[]
    private article: number
    private detail: ArticleDetail
    private info: string
    private status: string
    private height: number
    private converter: showdown.Converter
    private dialog: boolean

    @ViewChild("editor")
    private editor: ElementRef
    constructor(private route: ActivatedRoute, private router: Router, private writeService: WriteService, private fileService: FileService) {
        this.writable = true
        route.params.forEach(p => {
            this.article = p["article"]
            this.writeService.Categories().then(v => {
                this.categories = v
                if (v.length > 0 && this.article == 0) {
                    this.category = v[0].Id
                }
            })
            if (this.article > 0) {
                this.writable = false
                this.writeService.Article(this.article).then(v => {
                    this.detail = v
                    if (v) {
                        this.title = v.Title
                        this.content = v.Content
                        this.category = v.Category
                        this.writable = true
                        setTimeout(() => {
                            this.Changed(<HTMLTextAreaElement>this.editor.nativeElement)
                        }, 50);
                    }
                })
            }
        })
        this.title = ""
        this.category = 0
        this.content = ""
        this.status = "editor"
        this.converter = new showdown.Converter()
        this.dialog = false
    }

    // Save 保存文章,根据情况选择创建或更新
    Save() {
        if (this.title.length < 2 || this.title.length > 20) {
            this.ShowMessage("标题长度必须在2-20之间")
            return
        }
        if (this.category <= 0) {
            this.ShowMessage("请选择文章分类")
            return
        }
        if (this.content.length < 2) {
            this.ShowMessage("文章内容至少2个字符")
            return
        }
        if (this.article > 0) {
            this.Update()
        } else {
            this.Create()
        }
    }

    // Create 创建文章
    Create() {
        this.writeService.Create(this.category, this.title, this.content).then(v => {
            if (v > 0) {
                this.router.navigateByUrl("/write/" + v)
            } else {
                this.ShowMessage("创建失败")
            }
        })
    }

    // Update 更新文章
    Update() {
        this.writeService.Update(this.article, this.category, this.title, this.content).then(v => {
            this.ShowMessage(v ? "更新成功" : "更新失败")
        })
    }

    // ShowMessage 显示消息(5秒钟)
    ShowMessage(msg: string) {
        this.info = msg
        setTimeout(() => {
            this.info = ""
        }, 5000);
    }

    // Changed 当编辑框文本发生变化时触发
    Changed(editor: HTMLTextAreaElement) {
        this.height = editor.scrollHeight
    }

    // Markdown 从content生成md
    Markdown(wall: HTMLDivElement) {
        wall.innerHTML = ""
        wall.appendChild(MarkdownString(this.content))
    }

    // Upload 上传文件
    Upload(form: HTMLFormElement) {
        this.fileService.Upload(form).then(v => {
            if (v) {
                this.dialog = false
                var ext = v.substr(v.lastIndexOf(".") + 1).toLowerCase()
                var format: any = { "jpg": 1, "jpeg": 1, "bmp": 1, "gif": 1, "png": 1, "svg": 1 }
                var str = "[](/files/" + v + ")"
                var inPos = 1
                if (format[ext]) {
                    str = "!" + str
                    inPos++
                }
                var ele = <HTMLTextAreaElement>this.editor.nativeElement
                var pos = ele.selectionStart
                ele.value = ele.value.substr(0, pos) + str + ele.value.substr(pos)
                ele.selectionStart = pos + inPos
                ele.focus()
            }
        })
    }
}