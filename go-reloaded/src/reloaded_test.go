package src

import (
	"testing"
)

func TestText(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.", "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair."},
		{"Simply add 42 (hex) and 10 (bin) and you will see the result is 68.", "Simply add 66 and 2 and you will see the result is 68."},
		{"There is no greater agony than bearing a untold story inside you.", "There is no greater agony than bearing an untold story inside you."},
		{"Punctuation tests are ... kinda boring ,don't you think !?", "Punctuation tests are... kinda boring, don't you think!?"},
		{"1E (hex) files were added", "30 files were added"},
		{"It has been 10 (bin) years", "It has been 2 years"},
		{"I should stop SHOUTING (low)", "I should stop shouting"},
		{"Ready, set, go (up) !", "Ready, set, GO!"},
		{"If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?", "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"},
		{"I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure", "I have to pack 5 outfits. Packed 26 just to be sure"},
		{"Don not be sad ,because sad backwards is das . And das not good", "Don not be sad, because sad backwards is das. And das not good"},
		{"harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '", "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"},
	}
	for _, test := range tests {
		result := CorrectText(test.input)
		if result != test.expected {
			t.Errorf("Expected %s, but got %s for input %s", test.expected, result, test.input)
		}
	}
}
