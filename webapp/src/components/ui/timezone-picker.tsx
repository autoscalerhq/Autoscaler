// components/TimezonePicker.js
'use client';

import { useState } from 'react';
import {Select, SelectContent, SelectGroup, SelectItem, SelectLabel, SelectTrigger} from "~/components/ui/select"; // Or you can use a library that provides timezones
import { useTimezoneSelect, allTimezones } from "react-timezone-select"

export default function TimezonePicker() {

    const labelStyle = "abbrev"

    const timezones = {
        ...allTimezones
    }

    const { options, parseTimezone } = useTimezoneSelect({ labelStyle, timezones });

    const [selectedTimezone, setSelectedTimezone] = useState(parseTimezone(Intl.DateTimeFormat().resolvedOptions().timeZone));

    return (
        <div className="w-72">
            <Select
                value={selectedTimezone.value}
                onValueChange={(e) => {
                    setSelectedTimezone(parseTimezone(e))
                }}
            >
                <SelectTrigger className="w-full">
                    {selectedTimezone.value} ({selectedTimezone.abbrev})
                </SelectTrigger>
                <SelectContent>
                    <SelectGroup>
                        <SelectLabel>Timezones</SelectLabel>
                        {options.map((timezone) => (
                            <SelectItem key={timezone.value} value={timezone.value}>
                                {timezone.label}
                            </SelectItem>
                        ))}
                    </SelectGroup>
                </SelectContent>
            </Select>
        </div>
    );
};
