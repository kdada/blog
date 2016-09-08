import { Routes, RouterModule } from '@angular/router';
import { UserComponent } from './user/user.component';
import { CategoryComponent } from './category/category.component';
import { FileComponent } from './file/file.component';
import { ArticleComponent } from './article/article.component';

//路由
const routers: Routes = [
    { path: "", component: UserComponent },
    { path: "category", component: CategoryComponent },
    { path: "file", component: FileComponent },
    { path: "article", component: ArticleComponent }
]

//路由模块(with providers)
export const RouterMWP = RouterModule.forRoot(routers) 