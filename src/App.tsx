import { CircularProgress } from '@mui/material';
import React, { useEffect, useLayoutEffect, useRef, useState } from 'react';
import { useSearchParams } from 'react-router-dom';
import './App.css';
import { useInterval } from './hooks';
import RaceInformation from './RaceInformation';
import './ResultsTable';


export interface IRace {
  Country: String
  Name: string
  Gender: string
  Province: string
  Racers: IRacer[]
  Resort: string
  StartTime: string
  Technique: string
}

export interface IRacer {
  Bib: number,
  CheckedAt: string,
  Class: string,
  Club: string,
  Name: string,
  Rank: string,
  Run1: string,
  Run2: string,
  TotalTime: string,
}

export function BuildAPIEndpointURL(searchParams: URLSearchParams): string {
  const race_url = process.env.REACT_APP_RACER_READY_API as string
  return race_url + "?" + searchParams.toString()
}

const fetchRaceResults = (callback: (data: IRace) => void, searchParams: URLSearchParams) => {
  const endpoint_url = BuildAPIEndpointURL(searchParams);
  const request_headers = new Headers();
  const results_request = new Request(endpoint_url, {
    method: 'GET',
    headers: request_headers,
    cache: 'default',
  });

  fetch(results_request)
    .then(response => {
      if (!response.ok) {
        throw Error(response.statusText);
      }
      return response
    })
    .then(response => {
      return response.json()
    })
    .then(data => {
      callback(data)
    }).catch(function(error) {
      console.log(error);
    });
}

function App() {

  const [race, setRace] = useState<IRace>(
    {
      Country: "Unknown",
      Gender: "Unknown",
      Province: "Unknown",
      Name: "Unknown",
      Technique: "Unknown",
      StartTime: "Unknown",
      Resort: "Unknown",
      Racers: [] as IRacer[]
    });

  let [searchParams, _] = useSearchParams();

  useInterval(() => {
    fetchRaceResults(setRace, searchParams)
  }, 5000)

  if (race.Racers === undefined || race.Racers.length === 0) {
    return (
      <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
        <CircularProgress />
      </div>);
  }

  return (
    <div className="App">
      <header className="App-header">
        Racer Ready
      </header>
      <RaceInformation race={race}></RaceInformation>
    </div>
  );
}

export default App;
