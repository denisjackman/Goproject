package greetings

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)
// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("Empty Name")
    }
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomFormat(), name)
    // message := fmt.Sprint(randomFormat())
    return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
    // A map to associate names with messages.
    messages := make(map[string]string)
    // Loop through the received slice of names, calling
    // the Hello function to get a message for each name.
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        // In the map, associate the retrieved message with
        // the name.
        messages[name] = message
    }
    return messages, nil
}

func init(){
    rand.Seed(time.Now().UnixNano())
}
func randomFormat() string {
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v.!",
        "Hail, %v. Well met!",
        "Live long and prosper ! %v.",
    }
    return formats[rand.Intn(len(formats))]
}