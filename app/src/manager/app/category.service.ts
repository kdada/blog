import {Injectable} from "@angular/core"
import {Http,Response} from "@angular/http"
import "rxjs/add/operator/toPromise"

// 分类
export class Category {
    constructor(public Id:number,public Name:string) {

    }
}

// 分类服务
@Injectable()
export class CategoryService {
    constructor(private http: Http) {
    }
    public List(): Promise<Category[]> {
        return this.http.post("/category/list",{}).toPromise().then((resp:Response):Category[]=>{
            var result = resp.json()
            if (result.Code != 0) {
                console.log(result.Message)
                return [new Category(1,"aa"),new Category(2,"bb")]
            } else {
                return result.Data
            }
        })
    }
}