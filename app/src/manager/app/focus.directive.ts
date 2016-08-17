import {Directive, Input, ElementRef} from '@angular/core'


@Directive({
    selector: "[focus]",
})
export class FocusDirective {
    constructor(private view:ElementRef){

    }
    ngOnInit() {
        (<HTMLInputElement>this.view.nativeElement).focus()
    }
}