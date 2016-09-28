import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { CategoryDetail } from './category.model';
import { ReplaceDate } from '../common/date';
import { SearchParams } from '../common/url';

// 分类服务
@Injectable()
export class CategoryService {
    constructor(private http: Http) {

    }
    // ListNum 返回列表总页数
    ListNum() {
        return this.http.post("/category/listnum", {}).toPromise().then(resp => {
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
    List(page: number) {
        return this.http.post("/category/list", SearchParams({
            Page: page
        })).toPromise().then(resp => {
            if (resp.ok) {
                var obj = resp.json()
                if (obj.Code == 0) {
                    return ReplaceDate(<CategoryDetail[]>obj.Data, "CreateTime")
                }
            }
            return []
        })
    }

    // Hide 隐藏分类
    Hide(id: number) {
        return this.http.post("/category/hide", SearchParams({
            Category: id
        })).toPromise().then(resp => {
            if (resp.ok) {
                var obj = resp.json()
                return obj.Code == 0
            }
            return false
        })
    }

    // Show 显示分类
    Show(id: number) {
        return this.http.post("/category/show", SearchParams({
            Category: id
        })).toPromise().then(resp => {
            if (resp.ok) {
                var obj = resp.json()
                return obj.Code == 0
            }
            return false
        })
    }

    // Delete 删除分类
    Delete(id: number) {
        return this.http.post("/category/delete", SearchParams({
            Category: id
        })).toPromise().then(resp => {
            if (resp.ok) {
                var obj = resp.json()
                return obj.Code == 0
            }
            return false
        })
    }

    // Create 创建分类
    Create(name: string) {
        return this.http.post("/category/create", SearchParams({
            Name: name
        })).toPromise().then(resp => {
            if (resp.ok) {
                var obj = resp.json()
                if (obj.Code == 0) {
                    return ""
                }
                return obj.Message
            }
            return "网络错误"
        })
    }
}