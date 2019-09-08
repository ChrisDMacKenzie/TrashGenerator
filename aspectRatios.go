package main

var aspectRatios = map[string][]int{
	"1:1":    []int{240, 240},
	"4:3":    []int{320, 240},
	"16:9":   []int{320, 180},
	"1.85:1": []int{370, 200},
	"2.4:1":  []int{360, 150},
}

var ratioNames = []string{"1:1", "4:3", "16:9", "1.85:1", "2.4:1"}

func ratioSetter(s string) {
	rChoice = s
}
