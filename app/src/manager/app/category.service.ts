import {Injectable} from "@angular/core"
import {Http, Response} from "@angular/http"
import {Observable} from "rxjs/Rx";
// 分类
export class Category {
    public Modifing: boolean
    public State: string
    constructor(public Id: number, public Name: string) {
        this.Modifing = false
        this.State = "false"
    }
    //切换状态
    public Switch() {
        this.State = this.State == 'true' ? 'false' : 'true';
    }
    //切换修改状态
    public ToggleStatus() {
        this.Modifing = !this.Modifing
    }
}

// 分类服务
@Injectable()
export class CategoryService {
    constructor(private http: Http) {
    }
    public List(): Promise<Category[]> {
        return this.http.post("/category/list", {}).toPromise().then((resp: Response): Category[] => {
            var result = resp.json()
            if (result.Code != 0 || result.Data == null) {
                return []
            } else {
                var categories = new Array<Category>()
                result.Data.forEach(ele => {
                    categories.push(new Category(ele.Id,ele.Name))
                });
                return categories
            }
        })
    }
}