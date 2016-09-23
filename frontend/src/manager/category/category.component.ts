import { Component } from "@angular/core";
import { CategoryDetail } from './category.model';
import { CategoryService } from './category.service';

@Component({
    templateUrl: "/tmpls/category.html",
    providers: [CategoryService]
})
export class CategoryComponent {
    private pages: number
    private pagesNum: number[]
    private currentNum: number
    private categories: CategoryDetail[]
    constructor(private categoryService: CategoryService) {
        categoryService.ListNum().then(v => {
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

    // 显示指定页面
    ShowPage(page: number) {
        if (page > 0) {
            this.currentNum = page
            this.categoryService.List(page).then(data => {
                this.categories = data
            })
        }
    }

    // Hide 隐藏
    Hide(id: number) {
        this.categoryService.Hide(id).then(v => {
            this.ShowPage(this.currentNum)
        })
    }

    // Unban 显示
    Show(id: number) {
        this.categoryService.Show(id).then(v => {
            this.ShowPage(this.currentNum)
        })
    }
}