import { NgModule } from "@angular/core"
import { FormsModule } from '@angular/forms';
import { BrowserModule } from "@angular/platform-browser"
import { HttpModule } from "@angular/http"
import { AppComponent } from "./app.component"
import { UserComponent } from "./user/user.component"

@NgModule({
    imports: [BrowserModule, FormsModule],
    declarations: [AppComponent, UserComponent],
    bootstrap: [AppComponent]
})
export class AppModule {

}