package collections

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	category   string
	difficulty string
)

var wordCollections = map[string]map[string][]string{
	"Animals": {
		"Easy":   {"cat", "dog", "rat", "cow", "ant", "bird", "fish", "frog", "deer", "bear", "fox", "lion", "tiger", "goat", "sheep", "hamster", "rabbit", "snake", "eagle", "whale"},
		"Medium": {"elephant", "giraffe", "dolphin", "penguin", "zebra", "kangaroo", "penguin", "hippo", "buffalo", "raccoon", "otter", "cheetah", "panda", "moose", "squirrel", "chinchilla", "platypus", "komodo", "sloth", "wolverine"},
		"Hard":   {"hippopotamus", "caterpillar", "chameleon", "rhinoceros", "ostrich", "arapaima", "axolotl", "anemone", "oryx", "quokka", "tarsier", "manatee", "narwhal", "tapir", "vampire bat", "peregrine falcon", "capybara", "macaw", "okapi", "dung beetle"},
	},
	"Names": {
		"Easy":   {"John", "Alice", "Bob", "Emma", "Liam", "Noah", "Olivia", "Mia", "Ethan", "Ava", "Sophia", "Lily", "James", "Jack", "Emily", "Ella", "Lucas", "Amelia", "Sophie", "Henry"},
		"Medium": {"Sophia", "William", "Oliver", "Isabella", "Mason", "Charlotte", "Evelyn", "Jackson", "Avery", "Harper", "Wyatt", "Scarlett", "Sebastian", "Grace", "Levi", "Luna", "Fiona", "Zoey", "Gabriel", "Chloe"},
		"Hard":   {"Alexander", "Charlotte", "Nathaniel", "Elizabeth", "Benjamin", "Theodore", "Vivienne", "Anastasia", "Maximilian", "Seraphina", "Wilhelmina", "Isadora", "Bernadette", "Genevieve", "Gwendolyn", "Montgomery", "Octavius", "Demetrius", "Arabella", "Kensington"},
	},
	"Fruits": {
		"Easy":   {"apple", "banana", "pear", "kiwi", "plum", "grape", "orange", "peach", "raspberry", "melon", "mango", "date", "fig", "lemon", "lime", "cherry", "nectarine", "coconut", "blackberry", "papaya"},
		"Medium": {"mango", "pineapple", "blueberry", "cherry", "dragonfruit", "pomegranate", "apricot", "kiwifruit", "tangerine", "clementine", "passionfruit", "starfruit", "nectarine", "paprika", "gooseberry", "persimmon", "jackfruit", "blood orange", "raspberry", "soursop"},
		"Hard":   {"rhubarb", "clementine", "durian", "physalis", "ugu", "buddha's hand", "jackfruit", "finger lime", "monk fruit", "jambolan", "nashi pear", "salak", "sapodilla", "purple yam", "huckleberry", "soursop", "longan", "rambutan", "kiwano", "mulberry"},
	},
	"Colors": {
		"Easy":   {"red", "blue", "green", "yellow", "pink", "orange", "purple", "brown", "black", "white", "gray", "violet", "indigo", "cyan", "magenta", "gold", "silver", "beige", "navy", "teal"},
		"Medium": {"cyan", "magenta", "violet", "emerald", "turquoise", "lavender", "peach", "crimson", "charcoal", "olive", "plum", "coral", "salmon", "eggplant", "ivory", "amber", "fuchsia", "mint", "sky blue", "periwinkle"},
		"Hard":   {"lavender", "periwinkle", "auburn", "chartreuse", "vermilion", "cerulean", "sepia", "umber", "celadon", "malachite", "alabaster", "carmine", "taupe", "vermillion", "cobalt", "russet", "honeydew", "zinnia", "onion skin", "pansy"},
	},
	"Countries": {
		"Easy":   {"Brazil", "China", "France", "Japan", "India", "Italy", "Spain", "Canada", "Mexico", "Germany", "Egypt", "Chile", "Russia", "Ireland", "Greece", "Portugal", "Turkey", "Argentina", "Iran", "South Africa"},
		"Medium": {"Australia", "Canada", "Germany", "Mexico", "Italy", "Sweden", "Norway", "Finland", "Belgium", "Netherlands", "Switzerland", "Thailand", "Indonesia", "Vietnam", "Philippines", "Saudi Arabia", "Nigeria", "Argentina", "Pakistan", "Colombia"},
		"Hard":   {"Kazakhstan", "Uzbekistan", "Liechtenstein", "Malawi", "Mauritius", "Tajikistan", "Seychelles", "Nicaragua", "Kyrgyzstan", "Bhutan", "Tonga", "Fiji", "Vanuatu", "Equatorial Guinea", "Saint Kitts and Nevis", "Sao Tome and Principe", "Maldives", "Solomon Islands", "Papua New Guinea", "Montserrat"},
	},
}

func SetWordCollection() []string {
	setCategory()
	setDifficulty()

	titleCaser := cases.Title(language.English)

	fmt.Println("\033[1;33m" + "====================================" + "\033[0m")
	fmt.Println("\033[1;33m" + "          Category:   " + titleCaser.String(category) + "\033[0m")
	fmt.Println("\033[1;33m" + "          Difficulty: " + titleCaser.String(difficulty) + "\033[0m")
	fmt.Println("\033[1;33m" + "====================================" + "\033[0m")

	return wordCollections[category][difficulty]
}

func setCategory() {
	titleCaser := cases.Title(language.English)

	fmt.Println("\033[1;34m" + "====================================" + "\033[0m")
	fmt.Println("\033[1;34m" + "        Choose Game Category        " + "\033[0m")
	fmt.Println("\033[1;34m" + "====================================" + "\033[0m")

	ctgNum := 1
	categories := make([]string, 0, len(wordCollections)) // Create a slice to hold category names

	// Print available categories and store them in a slice
	for ctg := range wordCollections {
		fmt.Printf("\033[1;34m"+"  %d. %s\n"+"\033[0m", ctgNum, ctg)
		categories = append(categories, ctg) // Store category in the slice
		ctgNum++
	}

	fmt.Print("\033[1;34m" + "Please select the game category: " + "\033[0m")
	fmt.Scanln(&category)
	category = titleCaser.String(strings.ToLower(category))

	// Check if the input is a number
	if num, err := strconv.Atoi(category); err == nil && num > 0 && num <= len(categories) {
		category = categories[num-1] // Update category to the selected category name
	} else {
		category = titleCaser.String(strings.ToLower(category))
	}

	// Validate the chosen category
	_, ok := wordCollections[category]
	for !ok {
		fmt.Print("\033[1;34m" + "Invalid category. Please choose a valid category: " + "\033[0m")
		fmt.Scanln(&category)
		category = titleCaser.String(strings.ToLower(category))

		// Recheck if the input is a number and map it if valid
		if num, err := strconv.Atoi(category); err == nil && num > 0 && num <= len(categories) {
			category = categories[num-1] // Update category to the selected category name
		}
		_, ok = wordCollections[category] // Validate again
	}
	clear()
}

func setDifficulty() {
	titleCaser := cases.Title(language.English)

	fmt.Println("\033[1;34m" + "======================================" + "\033[0m")
	fmt.Println("\033[1;34m" + "        Choose Game Difficulty        " + "\033[0m")
	fmt.Println("\033[1;34m" + "======================================" + "\033[0m")
	fmt.Println("\033[1;34m" + "  1. Easy                             " + "\033[0m")
	fmt.Println("\033[1;34m" + "  2. Medium                           " + "\033[0m")
	fmt.Println("\033[1;34m" + "  3. Hard                             " + "\033[0m")
	fmt.Print("\033[1;34m" + "Please select the game difficulty(easy/medium/hard): " + "\033[0m")

	fmt.Scanln(&difficulty)
	difficulty = titleCaser.String(strings.ToLower(difficulty))

	switch difficulty {
	case "1", "Easy":
		difficulty = "Easy"
	case "2", "Medium":
		difficulty = "Medium"
	case "3", "Hard":
		difficulty = "Hard"
	}

	for difficulty != "Easy" && difficulty != "Medium" && difficulty != "Hard" {
		fmt.Print("\033[1;34m" + "Invalid choice. Please choose 'easy', 'medium', or 'hard': " + "\033[0m")
		fmt.Scanln(&difficulty)

		switch difficulty {
		case "1":
			difficulty = "Easy"
		case "2":
			difficulty = "Medium"
		case "3":
			difficulty = "Hard"
		default:
			difficulty = titleCaser.String(strings.ToLower(difficulty))
		}
	}
	clear()
}

func clear() {
	fmt.Print("\033[2J\033[H")
}
