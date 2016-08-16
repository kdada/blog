import {Routes,RouterModule} from "@angular/router"
import {CategoryComponent} from "./app/category.component"
import {ArticleComponent} from "./app/article.component"
import {WriteComponent} from "./app/write.component"


var routes:Routes = [
    {path:"",component:CategoryComponent},
    {path:"article",component:ArticleComponent},
    {path:"write",component:WriteComponent},
]


export var RouterMWP = RouterModule.forRoot(routes)