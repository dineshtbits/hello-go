package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
  cases := []struct {in, want string} {{"Helo", "oleh"}}
  for _, c := range cases {
  got := ReverseRunes(c.in)
  if got != c.want {
    t.Errorf("Failed --->")
  }
}

}
