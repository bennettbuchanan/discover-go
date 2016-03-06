package main
import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  seed := rand.New(rand.NewSource(time.Now().UnixNano()))
  randomNumber := seed.Intn(100)
  school := "Holberton School"
  beautifulWeather := true
  holbertonFounders := [3]string{"Rudy Rigot", "Sylvain Kalache", "Julien Barbier"}

  if randomNumber > 50 {
    fmt.Println("my random number is", randomNumber, "and is greater than 50")
  } else if randomNumber == 0 {
    fmt.Println("my random number is", randomNumber, "and is equal to 50")
  } else {
    fmt.Println("my random number is", randomNumber, "and is less than 50")
  }

  if school == "Holberton School" {
    fmt.Println("I am a student of the Holberton School")
  } else {
    fmt.Println("I am not a student of the Holberton School")
  }

  if beautifulWeather == true {
    fmt.Println("It's a beautiful weather!")
  } else {
    fmt.Println("It's not a beautiful weather!")
  }

  for i := 0; i < 3; i++ {
    fmt.Println(holbertonFounders[i], "is a founder")
  }
}
