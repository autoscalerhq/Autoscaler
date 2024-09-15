import {cn} from "~/lib/utils";

interface KbdProps {
    className?: string;
    children: React.ReactNode;
}

export function Kbd(props: KbdProps) {
    return (
        <kbd className={cn("border border-gray-500 pl-1 pr-1 pt-0.5 pb-0 rounded", props.className)}>
            {props.children}
        </kbd>
    );
}