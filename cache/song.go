package cache

import (
	"time"
)

type Song struct {
	path       string
	title      string
	length     time.Duration
	lastPlayed time.Time
	requester  string
	addCount   int
	playCount  int
	id         string
}

func (s *Song) Path() string {
	return s.path
}
func (s *Song) Title() string {
	return s.title
}
func (s *Song) Length() time.Duration {
	return s.length
}
func (s *Song) LastPlayed() time.Time {
	return s.lastPlayed
}
func (s *Song) Play() {
	s.lastPlayed = time.Now()
}
func (s *Song) Requester() string {
	return s.requester
}
func (s *Song) AddCount() int {
	return s.addCount
}
func (s *Song) PlayCount() int {
	return s.playCount
}
func (s *Song) Added() {
	s.lastPlayed = time.Now()
	s.addCount++
	return
}
func (s *Song) Played() {
	s.playCount++
	return
}
func (s *Song) ID() string {
	return s.id
}
