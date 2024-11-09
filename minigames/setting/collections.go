package setting

var GamesWithCollection = map[string]bool{
	"Hangman":       true,
	"Word Scramble": true,
}

var WordCollections = map[string]map[string][]string{
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
