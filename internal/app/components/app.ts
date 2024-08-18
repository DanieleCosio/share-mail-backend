import Alpine from "alpinejs";
import button from "./atoms/button/button";

/* COMPONENTS */
Alpine.data("button", () => button);

// eslint-disable-next-line @typescript-eslint/no-explicit-any
(window as any).Alpine = Alpine;
Alpine.start();
