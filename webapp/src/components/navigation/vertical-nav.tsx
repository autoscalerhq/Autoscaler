import Link from "next/link"
import {LucideIcon} from "lucide-react"
import {Tooltip, TooltipContent, TooltipTrigger} from "~/components/ui/tooltip";
import {Button, buttonVariants} from "~/components/ui/button";
import {cn} from "~/lib/utils";

export interface NavLinks {
    title: string
    label?: string | React.ReactNode
    href: string
    icon?: LucideIcon
    variant: "default" | "ghost"
    type?: "link" | "action"
    action?: () => void
}

interface NavProps {
    isCollapsed: boolean
    className?: string,
    links: NavLinks[]
}

export function /**/Nav({links, isCollapsed, className}: NavProps) {

    return (
        <div
            data-collapsed={isCollapsed}
            className={cn("group flex flex-col gap-4 py-2 data-[collapsed=true]:py-2", className)}
        >
            <nav
                className="grid gap-1 px-2 group-[[data-collapsed=true]]:justify-center group-[[data-collapsed=true]]:px-2">
                {links.map((link, index) =>

                        isCollapsed ? (
                            <Tooltip key={index} delayDuration={0}>
                                <TooltipTrigger asChild>
                                    <Link
                                        href={link.href}
                                        className={cn(
                                            buttonVariants({variant: link.variant, size: "icon"}),
                                            "h-9 w-9",
                                            link.variant === "default" &&
                                            "dark:bg-muted dark:text-muted-foreground dark:hover:bg-muted dark:hover:text-white"
                                        )}
                                    >
                                        {link.icon && <link.icon className="h-4 w-4"/> }
                                        <span className="sr-only">{link.title}</span>
                                    </Link>
                                </TooltipTrigger>
                                <TooltipContent side="right" className="flex items-center gap-4">
                                    {link.title}
                                    {typeof link.label === "string" ? (
                                            <span className="ml-auto text-muted-foreground">
                                                {link.label}
                                            </span>
                                        ) :
                                        <span className="ml-auto text-muted-foreground">
                                            {link.label}
                                        </span>
                                    }
                                </TooltipContent>
                            </Tooltip>
                        ) : (
                            link.type === undefined || link.type === "link" ?
                                <Link
                                    key={index}
                                    href={link.href}
                                    className={cn(
                                        buttonVariants({variant: link.variant, size: "sm"}),
                                        link.variant === "default" &&
                                        "dark:bg-muted dark:text-white dark:hover:bg-muted dark:hover:text-white",
                                        "justify-start"
                                    )}
                                >
                                    {link.icon && <link.icon className="mr-2 h-4 w-4"/>}
                                    {link.title}
                                    {link.label && (
                                        <span
                                            className={cn(
                                                "ml-auto",
                                                link.variant === "default" &&
                                                "text-background dark:text-white"
                                            )}
                                        >
                                      {link.label}
                                    </span>
                                    )}
                                </Link>:
                                <Button
                                    key={index}
                                    onClick={() => {
                                        link.action !== undefined ?
                                            link.action() : null
                                    }}
                                    className={cn(
                                        buttonVariants({variant: link.variant, size: "sm"}),
                                        link.variant === "default" &&
                                        "dark:bg-muted dark:text-white dark:hover:bg-muted dark:hover:text-white",
                                        "justify-start text-black bg-transparent"
                                    )}
                                >
                                    {link.icon && <link.icon className="mr-2 h-4 w-4"/>}
                                    {link.title}
                                    {link.label && (
                                        <span
                                            className={cn(
                                                "ml-auto",
                                                link.variant === "default" &&
                                                "text-background dark:text-white"
                                            )}
                                        >
                                      {link.label}
                                    </span>
                                    )}
                                </Button>
                        )
                )}
            </nav>
        </div>
    )
}