import { Component } from "@angular/core";
import { Router, ActivatedRoute, Params } from "@angular/router";
import { ArticleService } from './article.service';
import { ArticleDetail } from './article.model';

@Component({
    templateUrl: "/tmpls/article.html",
    providers: [ArticleService]
})
export class ArticleComponent {
    private category: number
    private pages: number
    private pagesNum: number[]
    private currentNum: number
    private articles: ArticleDetail[]
    constructor(private route: ActivatedRoute, private articleService: ArticleService) {
        this.currentNum = 1
        route.params.forEach(p => {
            this.category = p["category"]
        })
        articleService.ListNum(this.category).then(v => {
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
            this.articleService.List(this.category,page).then(data => {
                this.articles = data
            })
        }
    }

    // Hide 隐藏
    Hide(id: number) {
        this.articleService.Hide(id).then(v => {
            this.ShowPage(this.currentNum)
        })
    }

    // Unban 显示
    Show(id: number) {
        this.articleService.Show(id).then(v => {
            this.ShowPage(this.currentNum)
        })
    }

    // Delete 删除
    Delete(id: number) {
        this.articleService.Delete(id).then(v => {
            this.ShowPage(this.currentNum)
        })
    }

    // Top 置顶
    Top(id: number) {
        this.articleService.Top(id).then(v => {
            this.ShowPage(this.currentNum)
        })
    }

    // Untop 取消置顶
    Untop(id: number) {
        this.articleService.Untop(id).then(v => {
            this.ShowPage(this.currentNum)
        })
    }

}