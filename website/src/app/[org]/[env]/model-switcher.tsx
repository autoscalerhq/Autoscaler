"use client"

import * as React from "react"
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "~/components/ui/select";
import {cn} from "~/lib/utils";

interface ModelSwitcherProps {
    isCollapsed: boolean,
    className?: string,
    items: {
        name: string,
        label: string,
        shortname?: string,
        icon?: React.ReactNode,
    }[],
    defaultItemName: string,
    ariaLabel: string,
    label?: string,
}

export function ModelSwitcher({
                                  isCollapsed,
                                  items,
                                  className,
                                  defaultItemName,
                                  ariaLabel,
                                  label
                              }: ModelSwitcherProps) {

    const [selectedItem, setSelectedItem] = React.useState<string>(
        items.length > 0 ? defaultItemName : "Unknown"
    )

    return (
        <div className={className}>
        { label && <p className={"text-grey-200 text-xs ml-2"}>{label}</p>}
        <Select
            defaultValue={selectedItem}
            onValueChange={setSelectedItem}
        >
            <SelectTrigger
                className={cn(
                    "flex items-center gap-2 [&>span]:line-clamp-1 [&>span]:flex [&>span]:w-full [&>span]:items-center [&>span]:gap-1 [&>span]:truncate [&_svg]:h-4 [&_svg]:w-4 [&_svg]:shrink-0",
                    isCollapsed && "flex h-9 w-9 shrink-0 items-center justify-center p-0 [&>span]:w-auto [&>svg]:hidden"
                )}
                aria-label={ariaLabel}
            >
                <SelectValue placeholder="Select an item">
                    {items.find((item) => item.name === selectedItem)?.icon}
                    <span className={cn("ml-2", isCollapsed && "hidden")}>
                        {
                            items.find((item) => item.name === selectedItem)?.label
                        }
                    </span>
                </SelectValue>
            </SelectTrigger>
            <SelectContent>
                {items.map((item) => (
                    <SelectItem key={item.name} value={item.name}>
                        <div className="flex items-center gap-3 [&_svg]:h-4 [&_svg]:w-4 [&_svg]:shrink-0 [&_svg]:text-foreground">
                            {item.icon}
                            {item.name}
                        </div>
                    </SelectItem>
                ))}
            </SelectContent>
        </Select>
        </div>
    )
}