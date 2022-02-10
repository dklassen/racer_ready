import { Box, Container, Typography } from "@mui/material"
import { IRace } from "./App"
import ResultsTable, { SortComparison } from "./ResultsTable"

type Props =
    {
        race: IRace
    }

function RaceInformation({ race }: Props) {

    return (
        <Container>
            <Box sx={{ paddingTop: '15px', textAlign: 'left' }}>
                <Typography variant="h4" component="div" gutterBottom>
                    {race.Name} - {race.Technique} {race.Gender}
                </Typography>
                <Typography variant="subtitle1" gutterBottom component="div">
                    Start Time: {race.StartTime}
                </Typography>
            </Box>
            <ResultsTable sortFunc={SortComparison} racers={race.Racers}></ResultsTable>
        </Container>
    )
}

export default RaceInformation
