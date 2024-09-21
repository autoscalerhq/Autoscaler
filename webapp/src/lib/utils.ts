import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"
import posthog from "posthog-js";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function event(name: string, properties: object ) {
  posthog.capture(name, properties )
}

/**
 * Useful utility to wait n number of milliseconds.
 * It is usually useful for testing purposes, but can also be useful in production scenarios.
 * @param ms The milliseconds to wait
 */
export function delay(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

