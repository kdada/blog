import { Component } from "@angular/core";
import { FileService } from "./file.service";
import { FileDetail } from "./file.model";


@Component({
    templateUrl: "/tmpls/file.html",
    providers: [FileService]
})
export class FileComponent {
    private pages: number
    private pagesNum: number[]
    private currentNum: number
    private files: FileDetail[]
    private dialog: boolean
    constructor(private fileService: FileService) {
        this.dialog = false
        fileService.ListNum().then(v => {
            this.pages = v
            if (v > 0) {
                this.pagesNum = []
                for (var i = 1; i <= v; i++) {
                    this.pagesNum.push(i)
                }
                this.ShowPage(1)
            }
        })
    }

    // ShowPage 显示指定页面
    ShowPage(page: number) {
        if (page > 0) {
            this.currentNum = page
            this.fileService.List(page).then(data => {
                this.files = data
            })
        }
    }
    // Upload 上传
    Upload(form: HTMLFormElement) {
        this.fileService.Upload(form).then(s => {
            this.dialog = false
        })
    }

    // Delete 删除
    Delete(id: number) {
        this.fileService.Delete(id).then(v => {
            this.ShowPage(this.currentNum)
        })
    }
}