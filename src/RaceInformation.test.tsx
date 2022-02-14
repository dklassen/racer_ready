import React from 'react';
import { render } from '@testing-library/react';
import RaceInformation from "./RaceInformation";
import { IRace } from "./App"


describe("<RaceInformation />", () => {
    it("Renders racer on course in list section", () => {
        const inputRace: IRace = {
            Country: "",
            Name: "",
            Gender: "",
            Province: "",
            Racers: [
                {
                    Bib: 1,
                    CheckedAt: "string",
                    Class: "string",
                    Club: "string",
                    Name: "Super Racer",
                    Rank: "string",
                    OnCourse: true,
                    Run1: "string",
                    Run1ms: 0,
                    Run2: "string",
                    Run2ms: 0,
                    TotalTime: "string",
                    TotalTimems: 0,
                },
            ],
            Resort: "string",
            StartTime: "string",
            Technique: "string",
        }

        const { getByText } = render(<RaceInformation race={inputRace} />);

        getByText("Super Racer (1)")
    })
})
