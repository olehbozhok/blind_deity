package utils

// MoveVect represent relative coordinates
type MoveVect struct {
	H, W int
}

// MMoveVect make move vect, represent relative coordinates
func MMoveVect(h, w int) MoveVect {
	return MoveVect{h, w}
}
