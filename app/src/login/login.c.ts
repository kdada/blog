import { Component } from '@angular/core';
@Component({
  selector: 'my-app',
  template: '<h1>{{name}}</h1>'
})
export class AppComponent { 
	name:string
  constructor(){
    this.name = "TestTest" + "测试测试";
  }
}