package classes

import (
  // "fmt"
)

type classroom struct {
  room_number int
}

func newClassroom(room_number int) *classroom {
  
  cr := classroom { room_number: room_number }
  return &cr
}