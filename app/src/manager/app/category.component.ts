import {Component,Input, trigger, state, style, transition, animate} from "@angular/core"
import {AnimationSwitchContainer,AnimationSwitchButton} from "./app.common";
import {Category, CategoryService} from "./category.service";
import {FocusDirective} from "./focus.directive";

@Component({
    templateUrl: "/tmpl/category.html",
    viewProviders: [CategoryService],
    directives: [FocusDirective],
    animations: [AnimationSwitchContainer,AnimationSwitchButton]
})
export class CategoryComponent {
    public categories: Category[]
    public constructor(private categoryService: CategoryService) {
        categoryService.List().then(c => this.categories = c)
    }
    // 切换隐藏状态
    public Switch(c:Category) {
        c.Switch()
    }
    // 修改分类名称
    public ToggleStatus(c:Category) {
        c.ToggleStatus()
    }
    // 删除分类
    public Delete(c:Category) {
        console.log("删除"+c.Name)
    }
    // 显示添加分类窗口
    public ShowDialog() {

    }
}