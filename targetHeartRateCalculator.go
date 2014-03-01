package main

import "fmt"
import "flag"
import "os"

// command line arguments
var age = flag.Int("age", 0, "The age for calculate the heart rate targets")
var maxHeartRate = flag.Int("maxHeartRate", 0, "The maximal heart rate")

type rate struct {
	age int
	maxHeartRate float32
	fifty float32
	fiftyfive float32
	sixty float32
	sixtyfive float32
	seventy float32
	seventyfive float32
	eighty float32
	eightyfive float32
	ninty float32
	nintyfive float32
}

func (r *rate) Print() {
	if 0 != r.age {
		fmt.Printf("For the age of %v:\n", r.age)
	} else {
		fmt.Printf("for the max heart rate: %v\n", r.maxHeartRate)
	}

	fmt.Printf("\n")
	fmt.Printf("%7v | %10v\n", "Percent", "Heart Rate")
	fmt.Printf("%7v | %10v\n", "50", int(r.fifty))
	fmt.Printf("%7v | %10v\n", "55", int(r.fiftyfive))
	fmt.Printf("%7v | %10v\n", "60", int(r.sixty))
	fmt.Printf("%7v | %10v\n", "65", int(r.sixtyfive))
	fmt.Printf("%7v | %10v\n", "70", int(r.seventy))
	fmt.Printf("%7v | %10v\n", "75", int(r.seventyfive))
	fmt.Printf("%7v | %10v\n", "80", int(r.eighty))
	fmt.Printf("%7v | %10v\n", "85", int(r.eightyfive))
	fmt.Printf("%7v | %10v\n", "90", int(r.ninty))
	fmt.Printf("%7v | %10v\n", "95", int(r.nintyfive))
	fmt.Printf("%7v | %10v\n", "100", r.maxHeartRate)
}

func (r *rate) Calc() {
	if(0 != r.age) {
		r.maxHeartRate = float32(220)
		r.maxHeartRate -= float32(r.age)
	}

	r.fifty = r.maxHeartRate * 0.5
	r.fiftyfive = r.maxHeartRate * 0.55
	r.sixty = r.maxHeartRate * 0.6
	r.sixtyfive = r.maxHeartRate * 0.65
	r.seventy = r.maxHeartRate * 0.7
	r.seventyfive = r.maxHeartRate * 0.75
	r.eighty = r.maxHeartRate * 0.8
	r.eightyfive = r.maxHeartRate * 0.85
	r.ninty = r.maxHeartRate * 0.9
	r.nintyfive = r.maxHeartRate * 0.95
}

func main() {
	flag.Parse()

	var r rate
	r.age = *age
	r.maxHeartRate = float32(*maxHeartRate)

	if 0 == *age && 0 == *maxHeartRate {
		fmt.Printf("At least --age must be set.\n")
		fmt.Println("targetHeartRateCalculator --age --maxHeartRate")
		fmt.Println("--age = the age for the heart rate calclation")
		fmt.Println("--maxHeartRate = set the max heart rate without the age")
		os.Exit(1)
	}

	r.Calc()
	r.Print()
}