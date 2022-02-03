import { Box, Container } from "@mui/material"
import { IRace } from "./App"
import ResultsTable from "./ResultsTable"

type Props = 
{
    race: IRace
}

function RaceInformation( {race}: Props ) {

    return (
        <Container>
            <Box sx={{
                typography: 'h4',
                textAlign: 'left',
                paddingTop: '15px',
            }}>
            {race.Name} - {race.Technique} {race.Gender}
            </Box>

            <Box sx={{
                typography: 'h6',
                textAlign: 'left'
            }}>
            Start Time: {race.StartTime}
            </Box>


        <ResultsTable racers={race.Racers}></ResultsTable>
    
        </Container>
    )
}

export default  RaceInformation
