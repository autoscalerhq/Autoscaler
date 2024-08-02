import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"
import posthog from "posthog-js";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function event(name: string, properties: object ) {
  posthog.capture(name, properties )
}

