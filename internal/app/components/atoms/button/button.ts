import { baseComponent, BaseComponent, SsrValue } from "../../types";

export interface Button extends BaseComponent {
    onClick: SsrValue;
    init(onClick: SsrValue): void;
}

export const button: Button = {
    onClick: null,
    init(onClick) {
        this.onClick = onClick;
    },
    ...baseComponent,
};

export default () => button;
