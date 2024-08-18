export type SsrValue = number | string | null;

export type BaseComponent = {
    getSsrValue(value: SsrValue): SsrValue;
};

export const baseComponent: BaseComponent = {
    getSsrValue(value: SsrValue): SsrValue {
        if (value === "NULL") {
            value = null;
        }

        return value;
    },
} as const;
Object.freeze(baseComponent);
