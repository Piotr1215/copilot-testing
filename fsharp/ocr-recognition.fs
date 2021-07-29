module OCR.Recognition

let private convertDigit (input: string list) =
    // Here we know that input is the form of a single digit:
    (*
        [
        " _ "
        "| |"
        "|_|"
        "   "
        ]
    *)

    // Flat out the list to a single string, for easier match:
    let oneLinerDigit = input |> String.concat " "

    match oneLinerDigit with
    | " _  | | |_|    " -> "0"
    | "      |   |    " -> "1"
    | " _   _| |_     " -> "2"
    | " _   _|  _|    " -> "3"
    | "    |_|   |    " -> "4"
    | " _  |_   _|    " -> "5"
    | " _  |_  |_|    " -> "6"
    | " _    |   |    " -> "7"
    | " _  |_| |_|    " -> "8"
    | " _  |_|  _|    " -> "9"
    | _ -> "?"

let private convertLine (input: string list) =
    // input is in the form (i.e. all digits in a line of 3):
    (*
        [
        " _    "
        "| |  |"
        "|_|  |"
        "      "
        ]
    *)
    // If not length of 3, then return none.
    if input.[0].Length % 3 <> 0 then
        None
    else
        let numberOfDigits = (input.[0].Length / 3)

        let digits =
            [ for i in [ 0 .. numberOfDigits - 1 ] do
                  input
                  |> List.map (fun s -> s.[i * 3..(i * 3 + 2)]) ] // Pick substrings (i.e. each digit in the "line")
            |> List.map convertDigit
            |> String.concat ""

        Some digits

let convert (input: string list) =
    // Verify that rows are multiple of 4
    if input.Length % 4 <> 0 then
        None
    // Verify that all rows have a length that is multiple of 3
    elif input
         |> List.exists (fun str -> str.Length % 3 <> 0) then
        None
    else
        let lines =
            input
            |> List.chunkBySize 4
            |> List.map convertLine
            |> List.choose id
            |> String.concat ","

        if lines.Length = 0 then
            None
        else
            Some lines
