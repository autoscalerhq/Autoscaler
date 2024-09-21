"use client"
import {Copy} from "lucide-react";
import {className} from "postcss-selector-parser";
import {cn} from "~/lib/utils";
import { useClipboard } from '@mantine/hooks';
import {Button} from "~/components/ui/button";

interface CopyBoxProps {
    className?: string;
    content: string;
}

export default function CopyBox(props: CopyBoxProps) {
    const clipboard = useClipboard({ timeout: 500 });

    return (
        <div className={cn("inline-flex items-center bg-gray-200 h-6 rounded", className)}>
            <p className="text-gray-600 font-mono rounded px-2 text-xs">
                {props.content}
            </p>
            <Button variant={"ghost"} className="ml-2 rounded p-2 h-6 w-auto " onClick={() => clipboard.copy(props.content)}>
                <Copy className={"w-3 h-3"}/>
            </Button>
        </div>
    )
}