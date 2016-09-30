import { Component } from "@angular/core";
import { UserService } from './user.service';
import { UserDetail } from './user.model';

// 用户组件
@Component({
    templateUrl: "/tmpls/user.html",
    providers: [UserService]
})
export class UserComponent {
    private pages: number
    private pagesNum: number[]
    private currentNum: number
    private users: UserDetail[]
    constructor(private userService: UserService) {
        this.currentNum = 1
        userService.ListNum().then(v => {
            this.pages = v
            if (v > 0) {
                this.pagesNum = []
                for (var i = 1; i <= v; i++) {
                    this.pagesNum.push(i)
                }
                this.ShowPage(1)
            }
        })
    }

    // 显示指定页面
    ShowPage(page: number) {
        if (page > 0) {
            this.currentNum = page
            this.userService.List(page).then(data => {
                this.users = data
            })
        }
    }

    // Ban 禁用指定用户
    Ban(userId: number) {
        this.userService.Ban(userId).then(v=>{
            this.ShowPage(this.currentNum)
        })
    }

    // Unban 允许指定用户
    Unban(userId: number) {
        this.userService.Unban(userId).then(v=>{
            this.ShowPage(this.currentNum)
        })
    }
}