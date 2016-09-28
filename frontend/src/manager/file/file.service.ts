import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { FileDetail } from './file.model';
import { ReplaceDate } from '../common/date';
import { SearchParams } from '../common/url';

// 文件服务
@Injectable()
export class FileService {
    constructor(private http: Http) {

    }
    // ListNum 返回列表总页数
    ListNum() {
        return this.http.post("/file/listnum", {}).toPromise().then(resp => {
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
        return this.http.post("/file/list", SearchParams({
            Page: page
        })).toPromise().then(resp => {
            if (resp.ok) {
                var obj = resp.json()
                if (obj.Code == 0) {
                    return ReplaceDate(<FileDetail[]>obj.Data, "UploadTime")
                }
            }
            return []
        })
    }

    // Upload 上传
    Upload(form: HTMLFormElement) {
        return new Promise<string>((resolve, reject) => {
            var formData = new FormData(form)
            console.log(formData)
            var xhr = new XMLHttpRequest()
            xhr.open("post", "/file/upload", true)
            xhr.onreadystatechange = e => {
                if (xhr.readyState == 4) {
                    return <string>xhr.responseBody
                }
                return ""
            }
            xhr.onerror = e => {
                reject(xhr)
            }
            xhr.send(formData)
        })
    }

    // Delete 删除
    Delete(id: number) {
        return this.http.post("/file/delete", SearchParams({
            File: id
        })).toPromise().then(resp => {
            if (resp.ok) {
                var obj = resp.json()
                return obj.Code == 0
            }
            return false
        })
    }

}