package main

import (
	"flag"
	"log"

	"github.com/Izumra/FuncGen/pkg/funcgen/swift"
)

const (
	def_count_functions           uint = 1
	def_max_cycles_per_func       uint = 3
	def_max_conditions_per_func   uint = 5
	def_count_operations_per_func uint = 20
	def_check_syntax              bool = false
)

func main() {
	var (
		count_functions           uint
		max_cycles_per_func       uint
		max_conditions_per_func   uint
		count_operations_per_func uint
		check_syntax              bool
	)

	flag.UintVar(
		&count_functions,
		"countFunctions",
		def_count_functions,
		"Count of the functions that would be generated.\n-------\nКоличество функций, которые будет сгенерированы",
	)
	flag.UintVar(
		&max_cycles_per_func,
		"maxCyclesPerFunc",
		def_max_cycles_per_func,
		"Max count of the cycles that would be presented in the generated function.\n-------\nМаксимальное количество сгенерированных циклов с переменными и коллекциями в сгенерированной функции",
	)
	flag.UintVar(
		&max_conditions_per_func,
		"maxConditionsPerFunc",
		def_max_conditions_per_func,
		"Max count of conditions with perems that would be presented in the generated function.\n-------\nМаксимальное количество сгенерированных условий с переменными в сгенерированной функции",
	)
	flag.UintVar(
		&count_operations_per_func,
		"countOperationsPerFunc",
		def_count_operations_per_func,
		"Count of operations that would be presented for the different perems in the generated function.\n-------\nКоличество операций над переменными и элементами коллекций в сгенерированной функции",
	)
	flag.BoolVar(
		&check_syntax,
		"checkSyntax",
		def_check_syntax,
		"Check generated function syntax or not.\n-------\nПроверять синтаксис сгенерированной функции или нет",
	)
	flag.Parse()

	funcgen := swift.NewFunction(max_conditions_per_func, max_cycles_per_func)
	for i := 0; i < int(count_functions); i++ {
		function := funcgen.Generate(count_operations_per_func)
		if check_syntax {
			err := funcgen.CheckFunction(function)
			if err != nil {
				panic(err)
			}
		}

		log.Println("\n", function)
	}
}
