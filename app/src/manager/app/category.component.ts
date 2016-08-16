import {Component} from "@angular/core"
import { Category,CategoryService } from './category.service';


@Component({
    template:`
    <table class="table">
        <tr *ngFor="let c of categories">
            <td>{{c.Id}}</td>
            <td>{{c.Name}}</td>
        </tr>
    </table>
    `,
    viewProviders:[CategoryService]
})
export class CategoryComponent {
    private categories:Category[]
    public constructor(private categoryService:CategoryService){
        categoryService.List().then(c=>this.categories=c)
    }
}