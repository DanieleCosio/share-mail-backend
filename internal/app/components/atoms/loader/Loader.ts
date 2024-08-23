import "@webcomponents/custom-elements";
import LoaderHtml from "./Loader.html";
import { BaseComponent } from "../../types";

class Loader extends HTMLElement implements BaseComponent {
    public static name = "sm-loader";
    loader: HTMLElement | null;

    constructor() {
        super();

        this.active = true;
        this.loader = null;
    }

    connectedCallback() {
        const template = document.createElement("template");
        template.innerHTML = LoaderHtml;

        const shadow = this.attachShadow({ mode: "open" });
        shadow.appendChild(template.content.cloneNode(true));

        this.loader = shadow.querySelector(".sm_loader");
        this.active = this.getAttribute("active") === "true";

        if (this.selector && this.loader) {
            this.loader.setAttribute("part", this.selector);
        }
    }

    get selector(): string {
        return this.getAttribute("selector") || "";
    }

    get active(): boolean {
        return this.loader?.classList.contains("active") ?? false;
    }

    set active(value: boolean) {
        if (value) {
            this.loader?.classList.add("active");
            return;
        }

        this.loader?.classList.remove("active");
    }
}

export default Loader;
