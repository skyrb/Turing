package main  
import (  
   "fmt"
   "sort"
)

// Function to count votes for each option within each region
func countVotesByRegion(surveyData map[string]map[string][]string) map[string]map[string]int {  
   regionVotes := make(map[string]map[string]int)  
   for region, userData := range surveyData {  
       regionVotes[region] = make(map[string]int)  
       for _, options := range userData {  
           for _, option := range options {  
               regionVotes[region][option]++
           }
       }
   }
   return regionVotes
}

// Function to find the most popular option by region
func mostPopularOptionByRegion(regionVotes map[string]map[string]int) map[string]string {  
   mostPopularOptions := make(map[string]string)  
   for region, votes := range regionVotes {  
       maxCount := 0  
       var mostPopularOption string  
       for option, count := range votes {  
           if count > maxCount {  
               mostPopularOption = option  
               maxCount = count  
           }  
       }  
       mostPopularOptions[region] = mostPopularOption  
   }  
   return mostPopularOptions  
}  
// Function to find the least popular option overall
func leastPopularOptionOverall(regionVotes map[string]map[string]int) string {  
   allVotes := make(map[string]int)  
   for _, votes := range regionVotes {  
       for option, count := range votes {  
           allVotes[option] += count  
       }  
   }
   minCount := 99999999
   var leastPopularOption string
   for option, count := range allVotes {
       if count < minCount {
           leastPopularOption = option
           minCount = count
       }
   }
   return leastPopularOption
}

func main() {  
   // Sample survey data from multiple regions
   surveyData := map[string]map[string][]string{  
       "Region1": {  
           "User1": {"Option A", "Option B"},  
           "User2": {"Option B", "Option C"},  
           "User3": {"Option A", "Option C"},
       },  
       "Region2": {  
           "User4": {"Option A", "Option B"},  
           "User5": {"Option B", "Option A", "Option C"},  
           "User6": {"Option C", "Option D"},
       },  
   }  
   //Count votes by region
   regionVotes := countVotesByRegion(surveyData)  
   fmt.Println("Vote counts by region:")  
   for region, votes := range regionVotes {  
       fmt.Println("\n", region)  
       for option, count := range votes {  
           fmt.Printf("%s: %d votes\n", option, count)  
       }  
   }  

   //Find the most popular option by region
   mostPopularOptions := mostPopularOptionByRegion(regionVotes)  
   fmt.Println("\nMost popular option by region:")  
   for region, option := range mostPopularOptions {  
       fmt.Printf("%s: %s\n", region, option)  
   }  
 
   // Find the least popular option overall
   leastPopularOption := leastPopularOptionOverall(regionVotes)
   fmt.Println("\nLeast popular option overall:", leastPopularOption)

   // **Optional**: Sort regions based on their total votes
   fmt.Println("\nRegions sorted by total votes:")
   type regionVoteCount struct {
       region string
       voteCount int
   }
   var regionCounts []regionVoteCount
   for region, votes := range regionVotes {
       totalVotes := 0
       for _, count := range votes {