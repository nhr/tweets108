package main
          }ckage main
import (
        "os"
        "github.com/codegangsta/cli"
        "github.com/ChimeraCoder/anaconda"
        "fmt"
        "sort"
        "strings"
)
func RemoveDuplicates(strings*[]string)[]string { 
    m := map[string]bool{}
    cleaned := []string{}
    for _, str := range *strings {
        if _, seen := m[str]; !seen {
            cleaned = append(cleaned, str)
            m[str] = true
        }
    }
    return cleaned
}
func FindDuplicates(strings*[]string)[]string {
    m := map[string]bool{}
    duplicate := []string{}
    for _, str := range *strings {
        if _, seen := m[str]; !seen {
           duplicate = append(duplicate, str)   
           m[str] = true
        }
    }
    return duplicate
}


func main() {
    app := cli.NewApp()
    app.Name = "tweets108"
    app.Usage = "Search Twitter y'all!"
    app.Action = func(c *cli.Context){

        //authentication
        anaconda.SetConsumerKey("ofCy0nXTB22KtFKytvhBwtRRL")
        anaconda.SetConsumerSecret("8zdqrZpoRKMqLmvjMnBLdHMZqXctC527Q6XSkEi78TGaGLuqDm")
        api := anaconda.NewTwitterApi("3016393947-8KoV0HCnrCFhan4qs44geLx7TXh77mYNbDkPQr4",
                                      "940YmdBVhZvCO8YZI4GllGgPSbqkHgnoHQsPyEDd6YdyF")

        //search tweets containing Arg[0]
        searchText := c.Args()[0]
        searchTextWhash := fmt.Sprint("#", searchText)
        fmt.Println(searchTextWhash)

        //make list users, sorted alphabetically
        userSlice := make([]string, 30)

        //make list retweeted tweets
        retweetSlice :=make([]string, 30)

        //use GetSearch function to find number of tweets
        searchResult, err := api.GetSearch(searchText, nil)
        if err != nil {
           panic(err)
        }

        count := 0
        countRetweet := 0
        for _ , tweet := range searchResult.Statuses {
            if (tweet.Text != ""){
                count ++
                userSlice = append(userSlice, tweet.User.Name)

                if(strings.Contains(tweet.Text, "RT")){
                    countRetweet ++
                    retweetSlice = append(retweetSlice, tweet.Text)
                }
            }
        }

        //work on list of users
        sort.Strings(userSlice[:])
        RemoveDuplicates(&userSlice)
        countUnique := 0
        //find number of unique users who used keyword in tweet
        fmt.Printf("List of users who tweeted %s:\n", searchText)
        for i := range userSlice {
            countUnique++
            if userSlice[i] != "" {
                fmt.Println(userSlice[i])
            }
        }
        fmt.Printf("\n%d unique users tweeted something that matched %s\n", countUnique, searchText)
        fmt.Printf("%d retweets in search results\n", countRetweet)

        //work on RT list
        //mostRetweet := FindDuplicates(&retweetSlice)
          fmt.Println("Most retweeted tweet:")
          fmt.Println("My FindDuplicates func still needs some work........")
       
        //find tweets that used keyword in a hashtag
          searchResultwHash, err := api.GetSearch(searchTextWhash, nil)
          if err != nil {
            panic(err)
          }
        //find number of tweets that used keyword in a hashtag
        countHash :=0
        for _ , tweet := range searchResultwHash.Statuses {
           if (tweet.Text != "") {
               countHash ++
           }
        }
        fmt.Printf("%d tweets used %s in a hashtag\n", countHash, searchText)
    fmt.Printf("\n")
    }
    app.Run(os.Args)
}

