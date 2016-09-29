import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { ArticleDetail } from './article.model';
import { ReplaceDate } from '../common/date';
import { SearchParams } from '../common/url';

// 文章服务
@Injectable()
export class ArticleService {
    constructor(private http: Http) {

    }
    // ListNum 返回列表总页数
    ListNum(category: number) {
        return this.http.post("/article/listnum", SearchParams({
            Category: category
        })).toPromise().then(resp => {
            if (resp.ok) {
                var obj = resp.json()
                if (obj.Code == 0) {
                    return <number>obj.Data.Count
                }
            }
            return 0
        })
    }
    // List 返回指定页的列表
    List(category: number, page: number) {
        return this.http.post("/article/list", SearchParams({
            Category: category,
            Page: page
        })).toPromise().then(resp => {
            if (resp.ok) {
                var obj = resp.json()
                if (obj.Code == 0) {
                    return ReplaceDate(ReplaceDate(<ArticleDetail[]>obj.Data, "CreateTime"), "UpdateTime")
                }
            }
            return []
        })
    }

    // Hide 隐藏
    Hide(id: number) {
        return this.http.post("/article/hide", SearchParams({
            Article: id
        })).toPromise().then(resp => {
            if (resp.ok) {
                var obj = resp.json()
                return obj.Code == 0
            }
            return false
        })
    }

    // Show 显示
    Show(id: number) {
        return this.http.post("/article/show", SearchParams({
            Article: id
        })).toPromise().then(resp => {
            if (resp.ok) {
                var obj = resp.json()
                return obj.Code == 0
            }
            return false
        })
    }

    // Delete 删除
    Delete(id: number) {
        return this.http.post("/article/delete", SearchParams({
            Article: id
        })).toPromise().then(resp => {
            if (resp.ok) {
                var obj = resp.json()
                return obj.Code == 0
            }
            return false
        })
    }

    // Top 置顶
    Top(id: number) {
        return this.http.post("/article/top", SearchParams({
            Article: id
        })).toPromise().then(resp => {
            if (resp.ok) {
                var obj = resp.json()
                return obj.Code == 0
            }
            return false
        })
    }
    // Untop 取消置顶
    Untop(id: number) {
        return this.http.post("/article/untop", SearchParams({
            Article: id
        })).toPromise().then(resp => {
            if (resp.ok) {
                var obj = resp.json()
                return obj.Code == 0
            }
            return false
        })
    }
}