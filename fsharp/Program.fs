open System

// Given a year, report if it is a leap year.
let leapYear (year: int): bool = year % 4 = 0 && year % 100 <> 0 || year % 400 = 0

let leapYear' (year: int): bool =
    match (year % 4 = 0, year % 100 = 0, year % 400 = 0) with
    |(true, false, true) -> true
    |(_, _, _) -> false

[<EntryPoint>]
let main argv =
    let year = int argv.[0]
    let isLeap = leapYear year
    printfn "%d is %b" year isLeap
    0 // return an integer exit code