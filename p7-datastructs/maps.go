package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) disp() {
	fmt.Println(m)
}

func main() {
	// websites := []string{"https://google.com", "https://aws.com"}
	websites := map[string]string{
		"Google": "https://google.com",
		"Amazon": "https://aws.com",
	}
	userNames := make([]string, 2, 5)
	userNames[0] = "may"
	userNames[1] = "art"
	// userNames[2] = "mal"
	userNames = append(userNames, "mal")
	fmt.Println(userNames)
	for range userNames {
		fmt.Println("Hi")
	}
	for index, value := range userNames {
		fmt.Println("See: ", index, value)
	}
	fmt.Println()
	fmt.Println()

	fmt.Println(websites)
	fmt.Println(websites["Google"])
	websites["vk"] = "https://vk.com"
	fmt.Println(websites)
	delete(websites, "Google")
	fmt.Println(websites)
	fmt.Println()
	fmt.Println()

	courseRatings := make(map[string]float64, 3)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.3
	courseRatings["c"] = 4.6
	fmt.Println(courseRatings)

	fmt.Println()
	fmt.Println()
	speedLimits := make(floatMap, 3)
	speedLimits["road56"] = 56.6
	speedLimits["road32"] = 68.1
	speedLimits["road12"] = 26.2
	speedLimits.disp()
	fmt.Println()
	fmt.Println()

}
