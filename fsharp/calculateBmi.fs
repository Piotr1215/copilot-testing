module BMI

open System

let bmiFormula (x: float) (y: float) = float x / float y ** 2.0

let (|LessThan|_|) k value = if value <= k then Some() else None
let (|MoreThan|_|) k value = if value >= k then Some() else None

let bmi weight height =
    match bmiFormula weight height with
    | LessThan 18.5 -> "underweight"
    | MoreThan 18.5 & LessThan 25.0 -> "normal"
    | MoreThan 25.0 & LessThan 30.0 -> "overweight"
    | MoreThan 30.0 -> "obese"
    | _ -> "Out of scale"
