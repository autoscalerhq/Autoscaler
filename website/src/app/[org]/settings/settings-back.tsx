"use client"
import {useRouter} from "next/navigation";
import {Button} from "~/components/ui/button";
import {ChevronLeft} from "lucide-react";

interface SettingsBackProps {
    className?: string
    org: string
}

export default function SettingsBack(props: SettingsBackProps) {

    const router = useRouter();

    const handleBack = () => {
        // router.back();
        // TODO change to orgs default env
        router.push(`/${props.org}/Development`)
    };

    return (
        <Button
            variant={"ghost"}
            onClick={handleBack}
            className={props.className}>
                <ChevronLeft/>
        </Button>
    )
}