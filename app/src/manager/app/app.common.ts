import {trigger, state, style, transition, animate} from "@angular/core"

export const AnimationSwitchContainer = trigger("switchContainer", [
    state("false", style({
        background: "#eee",
    })),
    state("true", style({
        background: "#0b0",
    })),
    transition("false <=> true", animate("200ms ease-in-out"))
])
export const AnimationSwitchButton = trigger("switchButton", [
    state("false", style({
        marginLeft: "0px"
    })),
    state("true", style({})),
    transition("false <=> true", animate("200ms ease-in-out"))
])