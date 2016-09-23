import { NgModule } from "@angular/core";
import { FormsModule } from '@angular/forms';
import { BrowserModule } from "@angular/platform-browser";
import { HttpModule } from "@angular/http";
import { RouterMWP } from './app.routing';
import { AppComponent } from "./app.component";
import { UserComponent } from "./user/user.component";
import { CategoryComponent } from './category/category.component';
import { FileComponent } from './file/file.component';
import { ArticleComponent } from './article/article.component';

@NgModule({
    imports: [BrowserModule, FormsModule, HttpModule, RouterMWP],
    declarations: [AppComponent, UserComponent, CategoryComponent, FileComponent, ArticleComponent],
    bootstrap: [AppComponent]
})
export class AppModule {

}