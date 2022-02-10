import ReactDOM from "react-dom";
import { BrowserRouter } from "react-router-dom";
import ResultsTable, { SortComparison } from "./ResultsTable";

interface ITestObject {
    orderBy: number | null
    redundantField?: number
}

describe("ResultsTable", () => {
    it("renders without crashing", () => {
        const div = document.createElement('div');
        ReactDOM.render(<BrowserRouter><ResultsTable sortFunc={(a, b, c) => 0} racers={[]} /></BrowserRouter>, div);
    });
})

describe("SortComparison behaviour", () => {
    it("sorting numbers defaults to descending", () => {
        const sortingField = 'orderBy'

        const input = [
            { 'orderBy': 2 },
            { 'orderBy': 1 },
            { 'orderBy': 3 }
        ]

        const expected = [
            { 'orderBy': 3 },
            { 'orderBy': 2 },
            { 'orderBy': 1 }
        ]

        const ascendingSort = (a: ITestObject, b: ITestObject) => {
            return SortComparison(a, b, sortingField)
        }

        const output = input.sort(ascendingSort)

        expect(output).toEqual(expected)
    })

    it("When sorting ascending NULL values are sorted to top", () => {
        const sortingField = 'orderBy'

        const input = [
            { 'orderBy': 2 },
            { 'orderBy': 1 },
            { 'orderBy': null }
        ]

        const expected = [
            { 'orderBy': null },
            { 'orderBy': 2 },
            { 'orderBy': 1 }
        ]

        const ascendingSort = (a: ITestObject, b: ITestObject) => {
            return SortComparison(a, b, sortingField)
        }

        const output = input.sort(ascendingSort)

        expect(output).toEqual(expected)
    })

    it("When sorting descending NULL values are sorted to bottom", () => {
        const sortingField = 'orderBy'

        const input = [
            { 'orderBy': 2 },
            { 'orderBy': 1 },
            { 'orderBy': null }
        ]

        const expected = [
            { 'orderBy': 1 },
            { 'orderBy': 2 },
            { 'orderBy': null }
        ]

        const descendingSort = (a: ITestObject, b: ITestObject) => {
            return -SortComparison(a, b, sortingField)
        }

        const output = input.sort(descendingSort)

        expect(output).toEqual(expected)
    })

    it("When sorting, sort by specified field", () => {
        const sortingField = 'orderBy'

        const input = [
            {
                'orderBy': 2,
                'redundantField': 1,
            },
            {
                'orderBy': 1,
                'redundantField': 2,
            },
            {
                'orderBy': 3,
                'redundantField': 3,

            }
        ]

        const expected = [
            {
                'orderBy': 1,
                'redundantField': 2,
            },
            {
                'orderBy': 2,
                'redundantField': 1,
            },
            {
                'orderBy': 3,
                'redundantField': 3,

            }
        ]

        const descendingSort = (a: ITestObject, b: ITestObject) => {
            return -SortComparison(a, b, sortingField)
        }

        const output = input.sort(descendingSort)

        expect(output).toEqual(expected)
    })

    it("When sorting equivalent fields object orders remain unchanged", () => {
        const sortingField = 'orderBy'

        const expected = [
            {
                'orderBy': 1,
                'redundantField': 1,
            },
            {
                'orderBy': 1,
                'redundantField': 2,
            },
            {
                'orderBy': 1,
                'redundantField': 3,
            },
        ]

        const descendingSort = (a: ITestObject, b: ITestObject) => {
            return -SortComparison(a, b, sortingField)
        }

        const output = expected.sort(descendingSort)

        expect(output).toEqual(expected)
    })
})