import { Box, Container, List, ListItemText, Typography } from "@mui/material"
import { IRace } from "./App"
import ResultsTable, { SortComparison } from "./ResultsTable"

type Props =
    {
        race: IRace
    }

function RaceInformation({ race }: Props) {
    const racersOnCourse = race.Racers.filter(racer => racer.OnCourse === true)

    return (
        <Container>
            <Box sx={{ paddingTop: '15px', marginBottom: '20px', borderBottom: "1px solid #808080" }}>
                <Typography variant="h4" align="left" component="div" gutterBottom>
                    {race.Name}
                </Typography>
                <Typography variant="subtitle1" align="left" gutterBottom component="div">
                    {race.Gender} {race.Technique} {race.StartTime}
                </Typography>
            </Box>
            <Box>
                <Typography variant="h5" align="left" component="div">
                    Racers On Course
                </Typography>
                <List>
                    {
                        racersOnCourse.map((racer) => {
                            return <ListItemText key={racer.Bib} primary={racer.Name + " (" + racer.Bib + ")"} />
                        })
                    }
                </List>
            </Box>
            <Box sx={{ paddingTop: '20px' }} >
                <Typography align="left" variant="h5" component="div">
                    Results
                </Typography>
                <ResultsTable sortFunc={SortComparison} racers={race.Racers}></ResultsTable>
            </Box>
        </Container >
    )
}

export default RaceInformation
