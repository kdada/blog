import {NgModule} from "@angular/core"
import {FormsModule} from '@angular/forms';
import {BrowserModule} from "@angular/platform-browser"
import {HTTP_PROVIDERS} from "@angular/http"
import {AppComponent} from "./app.component"
import {RouterMWP} from "./app.routing"
import {CategoryComponent} from "./app/category.component"
import {ArticleComponent} from "./app/article.component"
import {WriteComponent} from "./app/write.component"

@NgModule({
    imports: [BrowserModule, FormsModule, RouterMWP],
    declarations: [AppComponent, CategoryComponent, ArticleComponent, WriteComponent],
    providers: [HTTP_PROVIDERS],
    bootstrap: [AppComponent]
})
export class AppModule {

}