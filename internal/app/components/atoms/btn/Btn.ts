import "@webcomponents/custom-elements";
import BtnHtml from "./Btn.html";
import { MouseEventCallback } from "../../types";
import { getNullableSsrValue, getSsrCallback } from "../../utils";

class Btn extends HTMLElement {
    private buttonElement: HTMLButtonElement | undefined | null;
    public static name = "sm-btn";

    constructor() {
        super();
        console.debug("sm-button constructor executed...");
    }

    get text(): string {
        return this.getAttribute("text") || "";
    }

    set onClick(callback: MouseEventCallback) {
        if (!this.buttonElement || callback === null) {
            console.debug("Error: ");
            console.debug("Element: ", this.buttonElement);
            console.debug("Callback: ", callback);
            return;
        }

        this.buttonElement.addEventListener("click", callback);
    }

    connectedCallback() {
        const template = document.createElement("template");
        template.innerHTML = BtnHtml;

        const shadow = this.attachShadow({ mode: "open" });
        shadow.appendChild(template.content.cloneNode(true));

        this.buttonElement = shadow.querySelector("button");
        if (!this.buttonElement) {
            throw new Error("Button not found");
        }

        this.buttonElement.textContent = this.text;

        const onClickCallback = getNullableSsrValue(this, "click");
        if (onClickCallback === null) {
            return;
        }

        this.onClick = getSsrCallback(onClickCallback) as MouseEventCallback;
    }
}

export default Btn;
