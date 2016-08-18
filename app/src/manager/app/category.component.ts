import {Component,Input, trigger, state, style, transition, animate} from "@angular/core"
import {Category, CategoryService} from "./category.service";
import {FocusDirective} from "./focus.directive";

@Component({
    selector:"msssss",
    templateUrl: "/tmpl/category.html",
    viewProviders: [CategoryService],
    directives: [FocusDirective],
    animations: [
        trigger("switchContainer", [
            state("false", style({
                background: "#eee",
            })),
            state("true", style({
                background: "#00bb00",
            })),
            transition("false <=> true", animate("200ms ease-in-out"))
        ]),
        trigger("switchButton", [
            state("false", style({
                background: "#fff",
                marginLeft: "0px"
            })),
            state("true", style({
                background: "#fff",
                marginLeft: "20px"
            })),
            transition("false <=> true", animate("200ms ease-in-out"))
        ])
    ]
})
export class CategoryComponent {
    @Input("losst") public width:number
    @Input("click") public delegateClick:()=>boolean
    public categories: Category[]
    private request:(c:boolean)=>void
    public constructor(private categoryService: CategoryService) {
        categoryService.List().then(c => this.categories = c)
        this.request = _.debounce((c:boolean)=>{
            console.log("result:",c)
        },1000)
    }
    public Switch(c:Category) {
        console.log(this.width)
        console.log(this.delegateClick())
        console.log("kk:"+c)
        this.request(true)
    }
}