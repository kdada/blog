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
        return this.http.post("/categorymanager/listnum", {}).toPromise().then(resp => {
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
        return this.http.post("/categorymanager/list", SearchParams({
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
        return this.http.post("/categorymanager/hide", SearchParams({
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
        return this.http.post("/categorymanager/show", SearchParams({
            Category: id
        })).toPromise().then(resp => {
            if (resp.ok) {
                var obj = resp.json()
                return obj.Code == 0
            }
            return false
        })
    }
}