'use client';
import { usePostHog } from 'posthog-js/react';
import React, { FormEvent, useState } from 'react';
import {
    Dialog,
    DialogClose,
    DialogContent,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger
} from "~/components/ui/dialog";
import { Button } from "~/components/ui/button";
import { MessageCircleWarning } from "lucide-react";
import { Textarea } from "~/components/ui/textarea";
import { Input } from "~/components/ui/input";
import { Label } from "~/components/ui/label";
import { toast } from "sonner"


interface Question {
    type: string;
    question: string;
    buttonText: string;
    choices?: string[];
}

const survey = {
    id: "0191ae1c-9d39-0000-d3af-52e486af0dc9",
    questions: [
        {
            type: "open",
            question: "What can we do to improve our product?",
            buttonText: "Next"
        },
        {
            type: "multiple_choice",
            choices: [
                "Great Experience",
                "Bad Experience",
                "Bug",
                "Idea"
            ],
            question: "What are you trying to tell us?",
            buttonText: "Submit"
        }
    ]
};

interface FeedbackProps {
    className?: string;
}

export default function Feedback(props: FeedbackProps) {
    const posthog = usePostHog();
    const [isOpen, setIsOpen] = useState(false);
    const [currentQuestionIndex, setCurrentQuestionIndex] = useState(0);
    const [responses, setResponses] = useState<string[]>([]);

    const handleFeedbackSubmit = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        // Debugging logs
        console.log("Form Submitted");

        const formElements = e.currentTarget.elements as typeof e.currentTarget.elements & {
            feedback: HTMLTextAreaElement,
            choice: RadioNodeList
        };

        const response = survey.questions[currentQuestionIndex]?.type === 'open'
            ? formElements.feedback.value
            : (formElements.choice.value || '');

        console.log("Response:", response);

        setResponses([...responses, response]);

        if (currentQuestionIndex < survey.questions.length - 1) {
            setCurrentQuestionIndex(currentQuestionIndex + 1);
        } else {
            console.log("Sending responses to PostHog...");

            posthog.capture("survey sent", {
                $survey_id: survey.id,
                $survey_responses: responses.concat(response)
            });

            toast.success("Thank you for your feedback! :)")

            closeModal();
        }
    }

    const handleBack = () => {
        if (currentQuestionIndex > 0) {
            setCurrentQuestionIndex(currentQuestionIndex - 1);
            setResponses(responses.slice(0, -1));
        }
    }

    const currentQuestion = survey.questions[currentQuestionIndex];

    const openModal = () => setIsOpen(true);
    const closeModal = () => {
        setIsOpen(false);
        setCurrentQuestionIndex(0); // Reset the question index to the first question
        setResponses([]); // Clear responses
    }

    return (
        <div className={props.className}>
            <Dialog open={isOpen} onOpenChange={setIsOpen}>
                <DialogTrigger asChild>
                    <Button variant="ghost"><MessageCircleWarning className={"mr-2"} /> Feedback</Button>
                </DialogTrigger>
                <DialogContent className="sm:max-w-md">
                    <DialogHeader>
                        <DialogTitle>Page Feedback</DialogTitle>
                    </DialogHeader>
                    <form onSubmit={handleFeedbackSubmit}>
                        {currentQuestion?.type === 'open' ? (
                            <Textarea
                                id="feedbackInput"
                                name="feedback"
                                placeholder={currentQuestion.question}
                                required
                            ></Textarea>
                        ) : (
                            <>
                                <Label htmlFor={`question`} className="">
                                    {currentQuestion?.question}
                                </Label>
                            {currentQuestion?.choices && currentQuestion.choices.map((choice, index) => (
                                <div key={index} className={"flex flex-row justify-left pt-4"}>
                                    <Input type="radio" id={`choice${index}`} name="choice" value={choice} required className={"w-4 h-4"} />
                                    <Label htmlFor={`choice${index}`} className="cursor-pointer ml-2">
                                        {choice}
                                    </Label>
                                </div>
                            ))}
                            </>
                        )}
                        <div className="flex justify-between mt-4">
                            {currentQuestionIndex > 0 && (
                                <Button type="button" onClick={handleBack}>
                                    Back
                                </Button>
                            )}
                            <Button type="submit">
                                {currentQuestion?.buttonText}
                            </Button>
                        </div>
                    </form>
                </DialogContent>
            </Dialog>
        </div>
    );
}