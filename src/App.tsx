import React, { useEffect, useState } from 'react';
import { useSearchParams } from 'react-router-dom';
import './App.css';
import RaceInformation from './RaceInformation';
import './ResultsTable';
import ResultsTable from './ResultsTable';


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
  const [racers, setRacers] = useState<IRacer[]>([]);

  let [searchParams, _] = useSearchParams();

  useEffect(() => {
      const endpoint_url = BuildAPIEndpointURL(searchParams);
      const request_headers = new Headers();
      const results_request = new Request(endpoint_url, {
          method: 'GET',
          headers: request_headers,
          cache: 'default',
      });

      fetch(results_request)
          .then(response => {
              return response.json()
          })
          .then(data => {
              data.Racers = data.Racers.sort((a: IRacer, b: IRacer) => (a.Bib < b.Bib ? -1 : 1));
              setRace(data)
              setRacers(data.Racers)
          }
          );
  }, [])


  if (racers.length === 0 ) {
    return (
    <div style={{display: 'flex',  justifyContent:'center', alignItems:'center', height: '100vh'}}>
        <h1> LOADING....</h1>
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
