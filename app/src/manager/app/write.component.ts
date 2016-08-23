import {Component, ElementRef, ViewChild} from "@angular/core"


@Component({
    templateUrl: `/tmpl/write.html`
})
export class WriteComponent {
    public Title: string
    public Content: string
    public Status: string
    private preStatus: string
    @ViewChild("wall")
    private wall: ElementRef
    private converter: showdown.Converter
    private change: () => void
    public constructor() {
        this.Status = ""
        this.converter = new showdown.Converter()
        this.change = _.debounce(() => {
            this.wall.nativeElement.innerHTML = this.converter.makeHtml(this.Content)
        }, 100)
    }
    public Change() {
        this.change()
    }
    public ngAfterViewChecked() {
        if (this.wall && this.Content && this.Status != this.preStatus) {
            this.preStatus = this.Status
            this.wall.nativeElement.innerHTML = this.converter.makeHtml(this.Content)
        }
    }
}