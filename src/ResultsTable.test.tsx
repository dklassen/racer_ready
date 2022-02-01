import ReactDOM from "react-dom";
import { BrowserRouter } from "react-router-dom";
import ResultsTable, { BuildAPIEndpointURL } from "./ResultsTable";

it('renders without crashing', () => {
    const div = document.createElement('div');
    ReactDOM.render(<BrowserRouter><ResultsTable /></BrowserRouter>, div);
});

it('buildAPIEndpointURL returns expected url', () => {
    const inputSearchParams = new URLSearchParams("c=12&b=12&a=abc");
    const outcome = BuildAPIEndpointURL(inputSearchParams);
    const expected = process.env.REACT_APP_RACER_READY_API + "?" + inputSearchParams.toString()
    
    expect(outcome).toEqual(expected)
});

