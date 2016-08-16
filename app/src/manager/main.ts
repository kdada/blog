import {enableProdMode} from "@angular/core"
import {browserDynamicPlatform} from "@angular/platform-browser-dynamic"
import {AppModule} from "./app.module"

enableProdMode();
browserDynamicPlatform().bootstrapModule(AppModule)