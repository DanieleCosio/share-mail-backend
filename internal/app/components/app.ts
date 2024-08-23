import { demoClickCallback } from "./templates/demo/demo";
import { MouseEventCallback } from "./types";
import Btn from "./atoms/btn/Btn";
import Icon from "./atoms/icon/Icon";
import Modal from "./molecules/modal/Modal";
import Loader from "./atoms/loader/Loader";

/* GLOBAL FUNCTIONS */
declare global {
    interface Window {
        demoClickCallback: MouseEventCallback;
    }
}

window.demoClickCallback = demoClickCallback;

/* COMPONENTS */
customElements.define(Btn.name, Btn);
customElements.define(Icon.name, Icon);
customElements.define(Modal.name, Modal);
customElements.define(Loader.name, Loader);
