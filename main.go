package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Entry struct {
	Text        string
	Attribution string
}

var entries = []Entry{
	// STEVE JOBS (10)
	{Text: "Stay hungry. Stay foolish.", Attribution: "Steve Jobs"},
	{Text: "Creativity is just connecting things.", Attribution: "Steve Jobs"},
	{Text: "Design is not just what it looks like. Design is how it works.", Attribution: "Steve Jobs"},
	{Text: "Your time is limited, so don't waste it living someone else's life.", Attribution: "Steve Jobs"},
	{Text: "You can't connect the dots looking forward; only backward.", Attribution: "Steve Jobs"},
	{Text: "Innovation distinguishes between a leader and a follower.", Attribution: "Steve Jobs"},
	{Text: "Details matter. It's worth waiting to get it right.", Attribution: "Steve Jobs"},
	{Text: "Quality is more important than quantity.", Attribution: "Steve Jobs"},
	{Text: "Real artists ship.", Attribution: "Steve Jobs"},
	{Text: "I want to put a ding in the universe.", Attribution: "Steve Jobs"},

	// ALAN WATTS (10)
	{Text: "The only way to make sense of change is to plunge into it.", Attribution: "Alan Watts"},
	{Text: "Be completely engaged in the here and now.", Attribution: "Alan Watts"},
	{Text: "Muddy water is best cleared by leaving it alone.", Attribution: "Alan Watts"},
	{Text: "You are an aperture through which the universe looks at itself.", Attribution: "Alan Watts"},
	{Text: "To have faith is to trust yourself to the water.", Attribution: "Alan Watts"},
	{Text: "No valid plans for the future can be made without living now.", Attribution: "Alan Watts"},
	{Text: "Trying to define yourself is like trying to bite your own teeth.", Attribution: "Alan Watts"},
	{Text: "You are something the whole universe is doing.", Attribution: "Alan Watts"},
	{Text: "Man suffers only because he takes seriously what the gods made for fun.", Attribution: "Alan Watts"},
	{Text: "We cannot be sensitive to pleasure without being sensitive to pain.", Attribution: "Alan Watts"},

	// RICK RUBIN (10)
	{Text: "The aim isn't to make art. It's to be in the state where art becomes inevitable.", Attribution: "Rick Rubin"},
	{Text: "Make something you love, as well as you can, right now.", Attribution: "Rick Rubin"},
	{Text: "Amplify what makes your perspective unique.", Attribution: "Rick Rubin"},
	{Text: "The way you do anything is the way you do everything.", Attribution: "Rick Rubin"},
	{Text: "The magic is in what we don't understand.", Attribution: "Rick Rubin"},
	{Text: "Share ideas freely; hoarding dries the river.", Attribution: "Rick Rubin"},
	{Text: "Awareness is tuning in without clinging.", Attribution: "Rick Rubin"},
	{Text: "Self-doubt doesn't disqualify you; keep going.", Attribution: "Rick Rubin"},
	{Text: "Treat your whole life as a work of art.", Attribution: "Rick Rubin"},
	{Text: "Inspiration comes first; the audience last.", Attribution: "Rick Rubin"},

	// HAIKU (10, original)
	{Text: "cold morning sidewalk\nsteam rising from street grates\ncity starts to breathe", Attribution: ""},
	{Text: "empty coffee cup\na ring of light on the desk\nnight's last conversation", Attribution: ""},
	{Text: "slow autumn river\none red leaf keeps changing lanes\nwith no destination", Attribution: ""},
	{Text: "headphones on the train\nstrangers mouthing silent words\neach a private film", Attribution: ""},
	{Text: "first snow, quiet street\nevery car becomes a ghost\nlooking for a home", Attribution: ""},
	{Text: "server fans at night\nhumming like a distant hive\ndata dreams in code", Attribution: ""},
	{Text: "moon above the lot\ngrocery carts in crooked rows\npilgrims of the day", Attribution: ""},
	{Text: "rain against my screen\ncursor waits in blinking prayer\nfor the next true line", Attribution: ""},
	{Text: "laundry in the wind\nt-shirts practice taking flight\nfor a better sky", Attribution: ""},
	{Text: "beach at 3 a.m.\ntide erasing every name\nlike it's nothing new", Attribution: ""},

	// KEROUAC-STYLE AMERICAN HAIKU (10, original)
	{Text: "neon diner sign\nflickers on, flickers back off—\nthe cook keeps whistling", Attribution: ""},
	{Text: "midnight bus station\nkid asleep on his backpack\ndreams a different town", Attribution: ""},
	{Text: "payphone in the rain\nnobody calling tonight\ndial tone holds its breath", Attribution: ""},
	{Text: "desert gas station\ncoke machine humming softly\nstars drink with me", Attribution: ""},
	{Text: "motel bible open\none dead moth between pages—\nresurrection, kid", Attribution: ""},
	{Text: "dawn freight train passing\ncoffee in a paper cup\nwheels finish my thoughts", Attribution: ""},
	{Text: "jukebox in the bar\nquarter stuck, song repeats twice\nnobody complains", Attribution: ""},
	{Text: "rest stop metal sink\ncold water on my face and—\nhey, I'm still here, man", Attribution: ""},
	{Text: "bus window twilight\npower lines and broken signs\nall spelling \"keep going\"", Attribution: ""},
	{Text: "cheap hotel hallway\nice machine coughs in the dark\nsomeone laughs alone", Attribution: ""},

	// KOANS (10)
	{Text: "Two hands clap and there is a sound.\nWhat is the sound of one hand?", Attribution: ""},
	{Text: "Show me your original face\nbefore your parents were born.", Attribution: ""},
	{Text: "A monk asked, \"What is Buddha?\"\nThe master replied, \"Three pounds of flax.\"", Attribution: ""},
	{Text: "A student asked, \"What is the Way?\"\nThe master said, \"Walk.\"", Attribution: ""},
	{Text: "A monk said, \"I am lonely.\"\nThe master asked, \"Who is lonely?\"", Attribution: ""},
	{Text: "\"Where is enlightenment?\"\n\"Who is asking?\"", Attribution: ""},
	{Text: "A man chased his shadow all day\nand grew tired. At dusk, he lit a candle.", Attribution: ""},
	{Text: "The student said, \"I understand.\"\nThe master poured tea until the cup overflowed.", Attribution: ""},
	{Text: "\"What is God?\"\nThe master pointed to the sound of rain.", Attribution: ""},
	{Text: "\"Give me the final truth.\"\nThe teacher closed his eyes and said nothing.", Attribution: ""},
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Check if entries slice is empty
	if len(entries) == 0 {
		fmt.Println("No entries available yet. Please add some wisdom to the entries slice!")
		return
	}

	// Pick a random entry
	randomIndex := rand.Intn(len(entries))
	entry := entries[randomIndex]

	// Print the entry text
	fmt.Println(entry.Text)

	// Print attribution if present
	if entry.Attribution != "" {
		fmt.Printf("— %s\n", entry.Attribution)
	}
}
