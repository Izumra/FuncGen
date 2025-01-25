# FuncGen
Function generator for Swift programming language writin on pure Go
## Description
This project may be useful for rowers which rovs on whichever galley for fast generating simple stump functions on the **Swift** programming language and inserting it for making up uniquess at the code 
## Usage
You can use it whatever you want. For generating functions manually by using cli with specified flags which definition provided below. Or you can use my package by importing it and use in your solutions.
### Examples
> [!NOTE]
> These examples was generated with standart values of CLI flags
<details>
<summary>**Example 1**</summary>

```
 func HialeahMonkeyStack ()  -> Void {
        var onHers: Int8? = -105
        var outsideIn = Int8(62)
        var nothingFairly: Int8 {
                get {
                        return Int8(95)
                }
                set {
                        print(nothingFairly)
                }
        }
        var instanceOf: Int8? = -69
        var tripBeyond: Int8 = -56
        var crewNormally = Int8(72)
        var alwaysOut: Int8? = -10
        var purseAfterwards: () -> Int8 = {
                return Int8(49)
        }
        var couldWhen: Int8
        couldWhen = 99
        nothingFairly = Int8(65) + nothingFairly
        if (alwaysOut != nil) {
                alwaysOut = Int8(5)
        }
        alwaysOut = 94
        nothingFairly = (Int8(-110) << Int8(-110))
        while alwaysOut! > (-56) {
                for i: Int8 in (-115)...27 {
                        outsideIn = 45
                        switch nothingFairly {
                        default:
                                crewNormally = 37
                                couldWhen = 55

                        }

                }
                if ((Int8(116) < onHers!) || (Int8(116) >= Int8(78))) {
                        alwaysOut = -122

                }
                else if ((Int8(-117) > Int8(99))) {
                        alwaysOut = (alwaysOut ?? Int8(34))

                }
                else{

                }

        }
        for p: Int8 in (-110)...(-92) {
                nothingFairly = 81
                if (instanceOf != nil) {
                        instanceOf = Int8(-21) * Int8(19) + instanceOf! / instanceOf!
                }
                outsideIn = 9
                outsideIn = Int8(-78) - Int8(-93)
                instanceOf = -50

        }
        alwaysOut = -122

}
```

</details>

<details>
<summary>**Example 2**</summary>

```
 func SanFranciscoArmadilloSmell (_ besidesWhose: Int8?, _ downThere: Int8?) -> Void {
        var teamMoney: () -> Int8 = {
                return Int8(4)
        }
        var themselvesSpotted = Int8(-24)
        var withoutFrantic: () -> Int8 = {
                return Int8(69)
        }
        var fewTheirs = Int8(112)
        if ((themselvesSpotted < themselvesSpotted) || (themselvesSpotted >= themselvesSpotted) && (themselvesSpotted < themselvesSpotted)) {
                themselvesSpotted = Int8(-92) + themselvesSpotted / Int8(-3) + themselvesSpotted
                themselvesSpotted = (Int8(16) << themselvesSpotted) >> themselvesSpotted
                fewTheirs = Int8(122) - fewTheirs * fewTheirs / Int8(17) + Int8(58)

        }
        else{
                themselvesSpotted = -20
                fewTheirs = 58
                fewTheirs = Int8(-91) * Int8(74) + fewTheirs * Int8(17) / Int8(99)
                fewTheirs = Int8(-50)
                for s: Int8 in 2...31 {
                        if ((Int8(-41) < Int8(-31))) {
                                themselvesSpotted = themselvesSpotted * Int8(-95)

                        }
                        else{
                                fewTheirs = 104

                        }
                        if ((fewTheirs >= Int8(97))) {

                        }
                        else{

                        }

                }
                if ((Int8(83) > Int8(27))) {
                        themselvesSpotted = (Int8(-22) >> Int8(-29)) << Int8(-78)

                }
                else if ((fewTheirs < Int8(17))) {
                        themselvesSpotted = themselvesSpotted ^ (themselvesSpotted >> Int8(71))

                }
                else if ((fewTheirs == fewTheirs) && (Int8(25) == Int8(16))) {
                        fewTheirs = 75

                }

        }
        fewTheirs = (fewTheirs >> Int8(74))
        switch fewTheirs {
        default:
                while themselvesSpotted < (-12) {

                }

        }

}
```

</details>
### CLI
CLI for defferent platform presented in the folder **build**
There is presented description for each flag that would be needed for working with CLI with specified default values
```
$ ./build/linux/swiftgenerator -h
Usage of ./build/linux/swiftgenerator:
  -checkSyntax
        Check generated function syntax or not.
        -------
        Проверять синтаксис сгенерированной функции или нет
  -countFunctions uint
        Count of the functions that would be generated.
        -------
        Количество функций, которые будет сгенерированы (default 1)
  -countOperationsPerFunc uint
        Count of operations that would be presented for the different perems in the generated function.
        -------
        Количество операций над переменными и элементами коллекций в сгенерированной функции (default 20)
  -maxConditionsPerFunc uint
        Max count of conditions with perems that would be presented in the generated function.
        -------        Максимальное количество сгенерированных условий с переменными в сгенерированной функции (default 5)
  -maxCyclesPerFunc uint
        Max count of the cycles that would be presented in the generated function.
        -------
        Максимальное количество сгенерированных циклов с переменными и коллекциями в сгенерированной функции (default 3)
```

