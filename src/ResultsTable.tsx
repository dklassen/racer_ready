import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import { useEffect, useState } from 'react';


interface IRacer {
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

// interface IRacerJSON {
//     Bib: number,
//     CheckedAt: string,
//     Class: string,
//     Club: string,
//     Name: string,
//     Rank: string,
//     Run1: string,
//     Run2: string,
//     TotalTime: string,
// }

function decodeRacer(json: IRacer): IRacer {
    return Object.assign({}, json, {

    }
 )
}

export default function ResultsTable() {
    const [racers, setRacers] = useState<IRacer[]>([]);

    useEffect(() => {
        var race_url = "http://localhost:8080?r=227639&&m=1&u=30"

        const request_headers = new Headers();
        const results_request = new Request(race_url, {
            method: 'GET',
            headers: request_headers,
            cache: 'default',
        });

        fetch(results_request)
            .then(response => {
                return response.json()
            })
            .then(data => {
                console.log(data)
                data = data.sort((a: IRacer, b: IRacer) => (a.Bib < b.Bib ? -1 : 1));
                setRacers(data)
            }
            );
    }, [])

    return (
        <Paper sx={{ width: '100%', overflow: 'hidden' }}>
        <TableContainer sx={{ maxHeight: 1000 }}>
            <Table stickyHeader aria-label="sticky table">
                <TableHead>
                    <TableRow>
                        <TableCell>Bib</TableCell>
                        <TableCell align="right">Name</TableCell>
                        <TableCell align="right">Club</TableCell>
                        <TableCell align="right">Class</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {racers.map((row) => (
                        <TableRow
                            key={row.Bib}
                            sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                        >
                            <TableCell component="th" scope="row">
                                {row.Bib}
                            </TableCell>
                            <TableCell align="right">{row.Name}</TableCell>
                            <TableCell align="right">{row.Club}</TableCell>
                            <TableCell align="right">{row.Class}</TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </TableContainer>
        </Paper>
    );
}