package gouid

import "testing"

func TestNewSnowflake(t *testing.T) {
	s := NewSnowflake(Node(0), Epoch("2020-05-20"))
	t.Log(s.Generate().Int64())
}
