import { EventCallback } from "./types";

export function getNullableSsrValue(
    element: HTMLElement,
    attributeKey: string,
): string | null {
    const value = element.getAttribute(attributeKey);
    if (value === "") {
        return null;
    }

    return value;
}

export function getSsrCallback(value: string): EventCallback {
    const callback = (window as never)[value];
    if (!callback) {
        return null;
    }

    return callback;
}
