import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import { IRacer } from './App';
import { SetStateAction, useState } from 'react';
import { TableSortLabel } from '@mui/material';

type SortDirection = 'DESCENDING' | 'ASCENDING';
type TableRowAlignment = "inherit" | "left" | "center" | "right" | "justify" | undefined

interface IHeader {
    id: string
    label: string
    sortBy: string
    align: TableRowAlignment
}

const RecordToColumnMap: Record<string, IHeader> = {
    "Bib": {
        id: "Bib",
        label: "Bib #",
        sortBy: "Bib",
        align: "left"
    },
    "Name": {
        id: "Name",
        label: "Name",
        sortBy: "Name",
        align: "left",
    },
    "Club": {
        id: "Club",
        label: "Club",
        sortBy: "Club",
        align: "left",
    },
    "Class": {
        id: "Class",
        label: "Class",
        sortBy: "Class",
        align: "left"
    },
    "Run1": {
        id: "Run1",
        label: "First Run",
        sortBy: "Run1ms",
        align: "right"
    },
    "Run2": {
        id: "Run2",
        label: "Second Run",
        sortBy: "Run2ms",
        align: "right",
    },
    "TotalTime": {
        id: "TotalTime",
        label: "Total Time",
        sortBy: "TotalTimems",
        align: "right"
    }
}

type ObjectSortFunc = <T, K extends keyof T>(a: T, b: T, orderField: K) => number

type Props = {
    racers: IRacer[]
    sortFunc: ObjectSortFunc
}

const SortComparison = <T, K extends keyof T>(a: T, b: T, orderField: K): number => {
    if (a[orderField] === null) {
        return -1
    }

    if (b[orderField] === null) {
        return 1
    }

    if (b[orderField] < a[orderField]) {
        return -1;
    }
    if (b[orderField] > a[orderField]) {
        return 1;
    }
    return 0;
}

function getComparator(order: SortDirection, orderBy: string, descendingComparisonFunc: ObjectSortFunc) {
    return order === 'DESCENDING'
        ? (a: IRacer, b: IRacer) => descendingComparisonFunc(a, b, orderBy as any)
        : (a: IRacer, b: IRacer) => -descendingComparisonFunc(a, b, orderBy as any);
}

function ResultsTable({ racers, sortFunc }: Props) {
    const [order, setOrder] = useState<SortDirection>('ASCENDING');
    const [orderBy, setOrderBy] = useState('bib');

    const handleRequestSort = (event: any, property: SetStateAction<string>) => {
        const isAsc = orderBy === property && order === 'ASCENDING';

        setOrder(isAsc ? 'DESCENDING' : 'ASCENDING');
        setOrderBy(property);
    };

    const createSortHandler = (property: string) => (event: React.MouseEvent<HTMLElement>) => {
        handleRequestSort(event, property);
    };

    return (
        <Paper sx={{ width: '100%', overflow: 'hidden' }}>
            <TableContainer sx={{ maxHeight: 1000 }}>
                <Table stickyHeader aria-label="sticky table">
                    <TableHead>
                        <TableRow>
                            {
                                Object.values(RecordToColumnMap).map((header: IHeader) => (
                                    <TableCell
                                        key={header.id}
                                        align={header.align}>
                                        <TableSortLabel
                                            active={orderBy === header.id}
                                            direction={orderBy === header.id ? order as any : 'asc'}
                                            onClick={createSortHandler(header.sortBy)}
                                        >
                                            {header.label}
                                        </TableSortLabel>
                                    </TableCell>
                                ))}
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {racers.sort(getComparator(order, orderBy, sortFunc)).map((row) => (
                            <TableRow
                                key={row.Bib}
                                sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                            >
                                <TableCell component="th" scope="row">{row.Bib} </TableCell>
                                <TableCell align={RecordToColumnMap.Bib.align}>{row.Name}</TableCell>
                                <TableCell align={RecordToColumnMap.Club.align}>{row.Club}</TableCell>
                                <TableCell align={RecordToColumnMap.Class.align}>{row.Class}</TableCell>
                                <TableCell align={RecordToColumnMap.Run1.align}>{row.Run1}</TableCell>
                                <TableCell align={RecordToColumnMap.Run2.align}>{row.Run2}</TableCell>
                                <TableCell align={RecordToColumnMap.TotalTime.align}>{row.TotalTime}</TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </Paper >
    );
}

export { SortComparison }
export default ResultsTable