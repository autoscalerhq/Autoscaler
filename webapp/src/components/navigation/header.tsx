import Feedback from "~/components/feedback/PageExperience";
import {
    DropdownMenu,
    DropdownMenuContent, DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger
} from "~/components/ui/dropdown-menu";
import {Button} from "~/components/ui/button";
import {Avatar, AvatarFallback, AvatarImage} from "~/components/ui/avatar";


export default function Header() {
    return (
        <header
            className="sticky top-0 z-30 flex h-16 items-center gap-4 border-b bg-background px-4 border-gray-200">

            <Feedback className={"ml-auto"}/>

            <DropdownMenu>
                <DropdownMenuTrigger asChild>
                    <Button
                        variant="outline"
                        size="icon"
                        className="overflow-hidden rounded-full"
                    >
                        <Avatar>
                            <AvatarImage src="https://github.com/shadcn.png"/>
                            <AvatarFallback>CN</AvatarFallback>
                        </Avatar>
                    </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent align="end">
                    <DropdownMenuLabel>My Account</DropdownMenuLabel>
                    <DropdownMenuSeparator/>
                    <DropdownMenuItem>Settings</DropdownMenuItem>
                    <DropdownMenuItem>Support</DropdownMenuItem>
                    <DropdownMenuSeparator/>
                    <DropdownMenuItem>Logout</DropdownMenuItem>
                </DropdownMenuContent>
            </DropdownMenu>
        </header>
    )
}