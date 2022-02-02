import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import { IRacer } from './App';


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
                        <TableCell>Bib</TableCell>
                        <TableCell align="right">Name</TableCell>
                        <TableCell align="right">Club</TableCell>
                        <TableCell align="right">Class</TableCell>
                        <TableCell align="right">Run 1</TableCell>
                        <TableCell align="right">Run 2</TableCell>
                        <TableCell align="right">Total Time</TableCell>
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