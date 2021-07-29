/// <summary>
/// String operations, such as toUpper, toLower, etc.
/// </summary>
module StringOperations

open System

let (|UpperCase|LowerCase|TitleCase|) (inp: string) =
    if inp = inp.ToUpper() then UpperCase
    elif inp = inp.ToLower() then LowerCase
    else TitleCase

let detectCase input =
    match input with
    | UpperCase -> "All was upper case"
    | LowerCase -> "All was lower case"
    | _ -> "All was mixed case"

let printer (inp: string) = printfn "%s" inp

let detectCaseAndPrint = detectCase >> printer
