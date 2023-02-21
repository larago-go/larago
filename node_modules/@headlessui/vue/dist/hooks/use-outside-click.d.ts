import { Ref, ComputedRef } from 'vue';
declare type Container = Ref<HTMLElement | null> | HTMLElement | null;
declare type ContainerCollection = Container[] | Set<Container>;
declare type ContainerInput = Container | ContainerCollection;
export declare function useOutsideClick(containers: ContainerInput | (() => ContainerInput), cb: (event: MouseEvent | PointerEvent | FocusEvent, target: HTMLElement) => void, enabled?: ComputedRef<boolean>): void;
export {};
