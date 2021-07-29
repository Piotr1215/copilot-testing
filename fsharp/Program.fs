open System

// Given a year, report if it is a leap year.
let leapYear (year: int) : bool =
    year % 4 = 0 && year % 100 <> 0 || year % 400 = 0

let leapYear' (year: int) : bool =
    match (year % 4 = 0, year % 100 = 0, year % 400 = 0) with
    | (true, false, true) -> true
    | (_, _, _) -> false

[<EntryPoint>]
let main argv =
    StringOperations.detectCaseAndPrint "asd"
    StringOperations.detectCaseAndPrint "ASD"
    StringOperations.detectCaseAndPrint "Asd"

    printfn "%s" (BMI.bmi 125.0 196.0)
    0 // return an integer exit code
