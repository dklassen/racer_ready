import ReactDOM from "react-dom";
import { BrowserRouter } from "react-router-dom";
import ResultsTable from "./ResultsTable";

it('renders without crashing', () => {
    const div = document.createElement('div');
    ReactDOM.render(<BrowserRouter><ResultsTable racers={[]}/></BrowserRouter>, div);
});


