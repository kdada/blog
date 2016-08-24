import {Injectable} from "@angular/core"
import {Http, Response, Headers, URLSearchParams, QueryEncoder} from "@angular/http"
import {Observable} from "rxjs/Rx";

export class Result {
    public constructor(public Code: number, public Message: string, public Data: any) {

    }
}

// 文章服务
@Injectable()
export class ArticleService {
    constructor(private http: Http) {
    }
    // 创建文章
    public New(title: string, content: string, category: number): Promise<Result> {
        var s = "Title="+encodeURIComponent(title)
        s += "&Content="+encodeURIComponent(content)
        s += "&Category="+category.toString()
        return this.http.post("/article/new", s, {
            headers: new Headers({
                "Content-Type": "application/x-www-form-urlencoded"
            })
        }).toPromise().then((resp: Response): Result => {
            var data = resp.json()
            return new Result(data.Code, data.Message, data.Data)
        }).catch(reason => {
            return new Result(reason.status, reason.statusText, null)
        })
    }
}