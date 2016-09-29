import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { ArticleDetail } from '../article/article.model';
import { CategoryDetail } from '../category/category.model';
import { ReplaceDate } from '../common/date';
import { SearchParams } from '../common/url';

// 文章写服务
@Injectable()
export class WriteService {
    constructor(private http: Http) {

    }
    // Categories 获取所有分类
    Categories() {
        return this.http.post("/category/categories", {}).toPromise().then(resp => {
            if (resp.ok) {
                var obj = resp.json()
                if (obj.Code == 0) {
                    return ReplaceDate(<CategoryDetail[]>obj.Data, "CreateTime")
                }
            }
            return []
        })
    }

    // Article 获取文章
    Article(article: number) {
        return this.http.post("/article/article", SearchParams({
            Article: article
        })).toPromise().then(resp => {
            if (resp.ok) {
                var obj = resp.json()
                if (obj.Code == 0) {
                    return ReplaceDate(ReplaceDate([<ArticleDetail>obj.Data], "CreateTime"), "UpdateTime")[0]
                }
            }
            return null
        })
    }

    // Create 创建
    Create(category: number, title: string, content: string) {
        return this.http.post("/article/create", SearchParams({
            Category: category,
            Title: title,
            Content: content
        })).toPromise().then(resp => {
            if (resp.ok) {
                var obj = resp.json()
                if (obj.Code == 0) {
                    return <number>obj.Data.Id
                }
            }
            return 0
        })
    }

    // Update 更新
    Update(id: number, category: number, title: string, content: string) {
        return this.http.post("/article/update", SearchParams({
            Id: id,
            Category: category,
            Title: title,
            Content: content
        })).toPromise().then(resp => {
            if (resp.ok) {
                var obj = resp.json()
                return obj.Code == 0
            }
            return false
        })
    }
}