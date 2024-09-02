"use client"
import {useRouter} from "next/navigation";
import {Button} from "~/components/ui/button";
import {ChevronLeft} from "lucide-react";

interface SettingsBackProps {
    className?: string
}

export default function SettingsBack(props: SettingsBackProps) {

    const router = useRouter();

    const handleBack = () => {
        router.back();
    };

    return (
        // <div className="flex flex-row ">
        //     <div>
                <Button variant={"ghost"} onClick={handleBack} className={props.className}><ChevronLeft/></Button>
    //             <p className={"text-xl"}> Settings</p>
    //         </div>
    //     </div>
    )
}