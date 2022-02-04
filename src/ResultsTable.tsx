import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import { IRacer } from './App';


interface IHeader {
 id: string
 label: string
}

const tableHeaders = [
    {
        id: 'bib',
        label: 'Bib #',
    },
    {
        id: "name",
        label: "Name",
    },
    {
        id: "club",
        label: "Club",
    },
    {
        id: "class",
        label: "Class"
    },
    {
        id: "run1",
        label: "First Run",
    },
    {
        id: "run2",
        label: "Second Run",
    },
    {
        id: "totalTime",
        label: "Total Time"
    }
]

type Props = {
    racers: IRacer[]
}


function ResultsTable( {racers}: Props ) {

    return (
        <Paper sx={{ width: '100%', overflow: 'hidden' }}>
        <TableContainer sx={{ maxHeight: 1000 }}>
            <Table stickyHeader aria-label="sticky table">
                <TableHead>
                    <TableRow>
                        {
                            tableHeaders.map((header: IHeader) => (
                                <TableCell
                                 key={header.id} 
                                align="right">{header.label}
                                </TableCell>
                        ))}
                    </TableRow>
                </TableHead>
                <TableBody>
                    {racers.sort((a: IRacer, b: IRacer) => (a.Bib < b.Bib ? -1 : 1)).map((row) => (
                        <TableRow
                            key={row.Bib}
                            sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                        >
                            <TableCell component="th" scope="row">{row.Bib} </TableCell>
                            <TableCell align="right">{row.Name}</TableCell>
                            <TableCell align="right">{row.Club}</TableCell>
                            <TableCell align="right">{row.Class}</TableCell>
                            <TableCell align="right">{row.Run1}</TableCell>
                            <TableCell align="right">{row.Run2}</TableCell>
                            <TableCell align="right">{row.TotalTime}</TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </TableContainer>
        </Paper>
    );
}


  export default ResultsTable