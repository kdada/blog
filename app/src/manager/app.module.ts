import {NgModule} from "@angular/core"
import {BrowserModule} from "@angular/platform-browser"
import {HTTP_PROVIDERS} from "@angular/http"
import {AppComponent} from "./app.component"
import {RouterMWP} from "./app.routing"

@NgModule({
    imports:[BrowserModule,RouterMWP],
    declarations:[AppComponent],
    providers:[HTTP_PROVIDERS],
    bootstrap:[AppComponent]
})
export class AppModule {

}