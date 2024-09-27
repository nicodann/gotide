package functions

//COUNTLETTERS FUNCTION

// Create function countletters
// take a string as input adn return object with count of each letter
// skip and not count spaces

func Countletters(str string) int {
	count := 0
	for i := 0; i < len(str); i++ {
		count += 1
	}
	return count
}
