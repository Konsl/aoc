package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/bits"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Machine struct {
	n_lamps int
	goal uint
	buttons []uint
	joltage []uint
}

func parse_lamp_list_bitset(str string) (goal uint, n_lamps int) {
	n_lamps = len(str)
	goal = uint(0)

	for i, v := range str {
		bit := uint(0)
		if v == '#' {
			bit = uint(1)
		}
		goal |= bit << i
	}

	return
}

func parse_int_list_bitset(str string) uint {
	parts := strings.Split(str, ",")
	val := uint(0)

	for _, v := range parts {
		iv, _ := strconv.Atoi(v)
		val |= 1 << iv
	}

	return val
}

func parse_int_list(str string) []uint {
	parts := strings.Split(str, ",")
	val := make([]uint, 0)

	for _, v := range parts {
		iv, _ := strconv.ParseUint(v, 10, 0)
		val = append(val, uint(iv))
	}

	return val
}

func parse_machine(str string) Machine {
	lamps_re := regexp.MustCompile(`\[([\.#]+)\]`)
	button_re := regexp.MustCompile(`\(([^)]+)\)`)
	joltage_re := regexp.MustCompile(`\{([^)]+)\}`)

	lamps_str := lamps_re.FindStringSubmatch(str)[1]
	buttons_str := button_re.FindAllStringSubmatch(str, -1)
	joltage_str := joltage_re.FindStringSubmatch(str)[1]

	goal, n_lamps := parse_lamp_list_bitset(lamps_str)
	buttons := make([]uint, 0)
	for _, b := range buttons_str {
		buttons = append(buttons, parse_int_list_bitset(b[1]))
	}

	joltage := parse_int_list(joltage_str)

	return Machine{
		n_lamps,
		goal,
		buttons,
		joltage,
	}
}

func solve_machine_1(machine Machine) int {
	best_count := math.MaxInt
	fmt.Println(machine)

	for t := range uint(1) << len(machine.buttons) {
		v := uint(0)
		for i, b := range machine.buttons {
			if t & (uint(1) << i) != 0 {
				v ^= b
			}
		}

		button_count := bits.OnesCount(t)


		if v == machine.goal && button_count < best_count {
			best_count = button_count
			fmt.Println("i", t, "c", best_count)
		}
	}

	return best_count
}

func clone_slice(sl []uint) []uint {
	return append(make([]uint, 0, len(sl)), sl...)
}

func sub_button2(remaining []uint, button uint, count uint) []uint {
	result := clone_slice(remaining)

	for i := range 32 {
		if button & (uint(1) << i) != 0 {
			result[i] -= count
		}
	}

	return result
}

func get_button_max2(remaining []uint, button uint) uint {
	max_thing := uint(math.MaxUint)

	for i := range remaining {
		if button & (uint(1) << i) != 0 {
			max_thing = min(max_thing, remaining[i])
		}
	}

	return max_thing
}

func is_finished2(joltage []uint) bool {
	for _, j := range joltage {
		if j != 0 {
			return false
		}
	}

	return true
}

// this could maybe work but it takes too long
func solve_machine_2_rec(buttons []uint, total_count uint, remaining []uint) uint {
	best_ret := uint(math.MaxUint)
	if len(buttons) == 0 {
		return best_ret
	}

	cur_i := 0
	cur_start_val := uint(0)

	for i := range remaining {
		if remaining[i] == 0 {
			continue
		}

		selecting_buttons := make([]int, 0)
		for ib, b := range buttons {
			if b & (uint(1) << i) != 0 {
				selecting_buttons = append(selecting_buttons, ib)
			}
		}

		if len(selecting_buttons) > 1 {
			continue
		}

		if len(selecting_buttons) == 0 {
			return best_ret
		}

		cur_i = selecting_buttons[0]
		cur_start_val = remaining[i]
	}

	btn_max := get_button_max2(remaining, buttons[cur_i])

	for n := cur_start_val; n <= btn_max; n++ {
		next_arr := sub_button2(remaining, buttons[cur_i], n)

		var ret uint
		if is_finished2(next_arr) {
			ret = total_count + n
		} else {
			buttons_next := clone_slice(buttons)
			buttons_next = slices.Delete(buttons_next, cur_i, cur_i + 1)
			ret = solve_machine_2_rec(buttons_next, total_count + n, next_arr)
		}

		best_ret = min(best_ret, ret)
	}

	return best_ret
}

func solve_machine_2(machine Machine) uint {
	return solve_machine_2_rec(machine.buttons, 0, machine.joltage)
}

func main() {
	file, err := os.Open("../input/10")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	machines := make([]Machine, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		machines = append(machines, parse_machine(scanner.Text()))
	}

	sum1 := 0
	for _, m := range machines {
		btn_c := solve_machine_1(m)
		sum1 += btn_c
	}
	fmt.Println(sum1)

	sum2 := uint(0)
	for _, m := range machines {
		btn_c := solve_machine_2(m)
		fmt.Println("2b", btn_c)
		sum2 += btn_c
	}

	fmt.Println(sum2)
}

