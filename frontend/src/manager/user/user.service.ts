import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { UserDetail } from './user.model';
import { ReplaceDate } from '../common/date';
import { SearchParams } from '../common/url';

// 用户服务
@Injectable()
export class UserService {
    constructor(private http: Http) {

    }
    // ListNum 返回列表总页数
    ListNum() {
        return this.http.post("/user/listnum", {}).toPromise().then(resp => {
            if (resp.ok) {
                var obj = resp.json()
                if (obj.Code == 0) {
                    return <number>obj.Data.Count
                }
            }
            return 0
        })
    }
    // List 返回指定页的用户列表
    List(page: number) {
        return this.http.post("/user/list", SearchParams({
            Page: page
        })).toPromise().then(resp => {
            if (resp.ok) {
                var obj = resp.json()
                if (obj.Code == 0) {
                    return ReplaceDate(<UserDetail[]>obj.Data, "CreateTime")
                }
            }
            return []
        })
    }

    // Ban 禁止用户登陆
    Ban(userId: number) {
        return this.http.post("/user/ban", SearchParams({
            User: userId
        })).toPromise().then(resp => {
            if (resp.ok) {
                var obj = resp.json()
                return obj.Code == 0
            }
            return false
        })
    }

    // Unban 允许用户登陆
    Unban(userId: number) {
        return this.http.post("/user/unban", SearchParams({
            User: userId
        })).toPromise().then(resp => {
            if (resp.ok) {
                var obj = resp.json()
                return obj.Code == 0
            }
            return false
        })
    }
}