package main

import (
	"fmt"
)

func main() {

	exercises()

}

func examples() {
	iMap := map[string]string{"home_range": "192.168.2.0/24", "company_range": "192.168.1.0/24", "lab_range": "192.168.102.0/24"}
	nMap := make(map[int]int)

	// Looping over map keys and values
	for k, v := range iMap {
		fmt.Printf("K: %v\tV: %v\n", k, v)
	}

	// Looping over map keys
	fmt.Println("-----------------------------------------------------")
	for k := range iMap {
		fmt.Println(k)
	}

	// Looping over map values
	fmt.Println("-----------------------------------------------------")
	for _, k := range iMap {
		fmt.Println(k)
	}

	// Assigning a values to maps
	fmt.Println("-----------------------------------------------------")
	for i := 1; i <= 10; i++ {
		nMap[i] = 10 * i
	}
	fmt.Println(nMap)

	// Deleting item from map
	fmt.Println("-----------------------------------------------------")
	delete(iMap, "company_range") // if the key does not exist, the compiler won't panic.
	fmt.Println(iMap)

	// Making sure the item really exists since the default value would return in case of
	// accessing a key which does not exist
	if v, ok := iMap["Akbar"]; !ok {
		fmt.Println("Queried value does not exist.")
	} else {
		fmt.Printf("exits %v\n", v)
	}
}
func exercises() {
	fmt.Println("************************************************************")
	fmt.Println("*************************Exercises**************************")
	fmt.Println("************************************************************")
	// initializing a map and range over it
	map49 := make(map[string][]string)
	map49[`bond_jame`] = []string{`Say my name`, `what's your name`, `she's a here`}
	map49[`zwsq`] = []string{`danger`, `road runner`, `migmig`}

	fmt.Println("#### After initializing a map ####")
	for k, v := range map49 {
		fmt.Printf("values for the key: %v ==> ", k)
		for i, v2 := range v {
			fmt.Printf("%v %v ", i, v2)
		}
		fmt.Println()
	}

	// deleteing a record from a map
	fmt.Println("#### After deleting a recodrd from map ####")
	delete(map49, "bond_jame")
	for k, v := range map49 {
		fmt.Printf("values for the key: %v ==> ", k)
		for i, v2 := range v {
			fmt.Printf("%v %v ", i, v2)
		}
		fmt.Println()
	}

	xs := []string{"in", "my", "younger", "and", "more", "vulnerable", "years", "my", "father", "gave", "me", "some", "advice", "that", "i’ve", "been", "turning", "over", "in", "my", "mind", "ever", "since.", "whenever", "you", "feel", "like", "criticizing", "anyone,", "he", "told", "me,", "just", "remember", "that", "all", "the", "people", "in", "this", "world", "haven’t", "had", "the", "advantages", "that", "you’ve", "had.", "he", "didn’t", "say", "any", "more,", "but", "we’ve", "always", "been", "unusually", "communicative", "in", "a", "reserved", "way,", "and", "i", "understood", "that", "he", "meant", "a", "great", "deal", "more", "than", "that.", "in", "consequence,", "i’m", "inclined", "to", "reserve", "all", "judgements,", "a", "habit", "that", "has", "opened", "up", "many", "curious", "natures", "to", "me", "and", "also", "made", "me", "the", "victim", "of", "not", "a", "few", "veteran", "bores.", "the", "abnormal", "mind", "is", "quick", "to", "detect", "and", "attach", "itself", "to", "this", "quality", "when", "it", "appears", "in", "a", "normal", "person,", "and", "so", "it", "came", "about", "that", "in", "college", "i", "was", "unjustly", "accused", "of", "being", "a", "politician,", "because", "i", "was", "privy", "to", "the", "secret", "griefs", "of", "wild,", "unknown", "men.", "most", "of", "the", "confidences", "were", "unsought—frequently", "i", "have", "feigned", "sleep,", "preoccupation,", "or", "a", "hostile", "levity", "when", "i", "realized", "by", "some", "unmistakable", "sign", "that", "an", "intimate", "revelation", "was", "quivering", "on", "the", "horizon;", "for", "the", "intimate", "revelations", "of", "young", "men,", "or", "at", "least", "the", "terms", "in", "which", "they", "express", "them,", "are", "usually", "plagiaristic", "and", "marred", "by", "obvious", "suppressions.", "reserving", "judgements", "is", "a", "matter", "of", "infinite", "hope.", "i", "am", "still", "a", "little", "afraid", "of", "missing", "something", "if", "i", "forget", "that,", "as", "my", "father", "snobbishly", "suggested,", "and", "i", "snobbishly", "repeat,", "a", "sense", "of", "the", "fundamental", "decencies", "is", "parcelled", "out", "unequally", "at", "birth.", "and,", "after", "boasting", "this", "way", "of", "my", "tolerance,", "i", "come", "to", "the", "admission", "that", "it", "has", "a", "limit.", "conduct", "may", "be", "founded", "on", "the", "hard", "rock", "or", "the", "wet", "marshes,", "but", "after", "a", "certain", "point", "i", "don’t", "care", "what", "it’s", "founded", "on.", "when", "i", "came", "back", "from", "the", "east", "last", "autumn", "i", "felt", "that", "i", "wanted", "the", "world", "to", "be", "in", "uniform", "and", "at", "a", "sort", "of", "moral", "attention", "forever;", "i", "wanted", "no", "more", "riotous", "excursions", "with", "privileged", "glimpses", "into", "the", "human", "heart.", "only", "gatsby,", "the", "man", "who", "gives", "his", "name", "to", "this", "book,", "was", "exempt", "from", "my", "reaction—gatsby,", "who", "represented", "everything", "for", "which", "i", "have", "an", "unaffected", "scorn.", "if", "personality", "is", "an", "unbroken", "series", "of", "successful", "gestures,", "then", "there", "was", "something", "gorgeous", "about", "him,", "some", "heightened", "sensitivity", "to", "the", "promises", "of", "life,", "as", "if", "he", "were", "related", "to", "one", "of", "those", "intricate", "machines", "that", "register", "earthquakes", "ten", "thousand", "miles", "away.", "this", "responsiveness", "had", "nothing", "to", "do", "with", "that", "flabby", "impressionability", "which", "is", "dignified", "under", "the", "name", "of", "the", "creative", "temperament—it", "was", "an", "extraordinary", "gift", "for", "hope,", "a", "romantic", "readiness", "such", "as", "i", "have", "never", "found", "in", "any", "other", "person", "and", "which", "it", "is", "not", "likely", "i", "shall", "ever", "find", "again.", "no—gatsby", "turned", "out", "all", "right", "at", "the", "end;", "it", "is", "what", "preyed", "on", "gatsby,", "what", "foul", "dust", "floated", "in", "the", "wake", "of", "his", "dreams", "that", "temporarily", "closed", "out", "my", "interest", "in", "the", "abortive", "sorrows", "and", "short-winded", "elations", "of", "men.", "my", "family", "have", "been", "prominent,", "well-to-do", "people", "in", "this", "middle", "western", "city", "for", "three", "generations.", "the", "carraways", "are", "something", "of", "a", "clan,", "and", "we", "have", "a", "tradition", "that", "we’re", "descended", "from", "the", "dukes", "of", "buccleuch,", "but", "the", "actual", "founder", "of", "my", "line", "was", "my", "grandfather’s", "brother,", "who", "came", "here", "in", "fifty-one,", "sent", "a", "substitute", "to", "the", "civil", "war,", "and", "started", "the", "wholesale", "hardware", "business", "that", "my", "father", "carries", "on", "today.", "i", "never", "saw", "this", "great-uncle,", "but", "i’m", "supposed", "to", "look", "like", "him—with", "special", "reference", "to", "the", "rather", "hard-boiled", "painting", "that", "hangs", "in", "father’s", "office.", "i", "graduated", "from", "new", "haven", "in", "1915,", "just", "a", "quarter", "of", "a", "century", "after", "my", "father,", "and", "a", "little", "later", "i", "participated", "in", "that", "delayed", "teutonic", "migration", "known", "as", "the", "great", "war.", "i", "enjoyed", "the", "counter-raid", "so", "thoroughly", "that", "i", "came", "back", "restless.", "instead", "of", "being", "the", "warm", "centre", "of", "the", "world,", "the", "middle", "west", "now", "seemed", "like", "the", "ragged", "edge", "of", "the", "universe—so", "i", "decided", "to", "go", "east", "and", "learn", "the", "bond", "business.", "everybody", "i", "knew", "was", "in", "the", "bond", "business,", "so", "i", "supposed", "it", "could", "support", "one", "more", "single", "man.", "all", "my", "aunts", "and", "uncles", "talked", "it", "over", "as", "if", "they", "were", "choosing", "a", "prep", "school", "for", "me,", "and", "finally", "said,", "why—ye-es,", "with", "very", "grave,", "hesitant", "faces.", "father", "agreed", "to", "finance", "me", "for", "a", "year,", "and", "after", "various", "delays", "i", "came", "east,", "permanently,", "i", "thought,", "in", "the", "spring", "of", "twenty-two.", "the", "practical", "thing", "was", "to", "find", "rooms", "in", "the", "city,", "but", "it", "was", "a", "warm", "season,", "and", "i", "had", "just", "left", "a", "country", "of", "wide", "lawns", "and", "friendly", "trees,", "so", "when", "a", "young", "man", "at", "the", "office", "suggested", "that", "we", "take", "a", "house", "together", "in", "a", "commuting", "town,", "it", "sounded", "like", "a", "great", "idea.", "he", "found", "the", "house,", "a", "weather-beaten", "cardboard", "bungalow", "at", "eighty", "a", "month,", "but", "at", "the", "last", "minute", "the", "firm", "ordered", "him", "to", "washington,", "and", "i", "went", "out", "to", "the", "country", "alone.", "i", "had", "a", "dog—at", "least", "i", "had", "him", "for", "a", "few", "days", "until", "he", "ran", "away—and", "an", "old", "dodge", "and", "a", "finnish", "woman,", "who", "made", "my", "bed", "and", "cooked", "breakfast", "and", "muttered", "finnish", "wisdom", "to", "herself", "over", "the", "electric", "stove.", "it", "was", "lonely", "for", "a", "day", "or", "so", "until", "one", "morning", "some", "man,", "more", "recently", "arrived", "than", "i,", "stopped", "me", "on", "the", "road.", "how", "do", "you", "get", "to", "west", "egg", "village?", "he", "asked", "helplessly.", "i", "told", "him.", "and", "as", "i", "walked", "on", "i", "was", "lonely", "no", "longer.", "i", "was", "a", "guide,", "a", "pathfinder,", "an", "original", "settler.", "he", "had", "casually", "conferred", "on", "me", "the", "freedom", "of", "the", "neighbourhood.", "and", "so", "with", "the", "sunshine", "and", "the", "great", "bursts", "of", "leaves", "growing", "on", "the", "trees,", "just", "as", "things", "grow", "in", "fast", "movies,", "i", "had", "that", "familiar", "conviction", "that", "life", "was", "beginning", "over", "again", "with", "the", "summer.", "there", "was", "so", "much", "to", "read,", "for", "one", "thing,", "and", "so", "much", "fine", "health", "to", "be", "pulled", "down", "out", "of", "the", "young", "breath-giving", "air.", "i", "bought", "a", "dozen", "volumes", "on", "banking", "and", "credit", "and", "investment", "securities,", "and", "they", "stood", "on", "my", "shelf", "in", "red"}
	frequency_calculator(xs)

}

func frequency_calculator(inputSlice []string) {
	dictionary := make(map[string]int)
	for _, v := range inputSlice {
		dictionary[v]++
	}
	for k, v := range dictionary {
		fmt.Printf("%v occurance of\t%v\n", v, k)
	}
}
