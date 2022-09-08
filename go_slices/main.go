// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	favCards := []int{2, 6, 9}
    return favCards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	for k, v := range slice {
        if k == index {
            return v
        }
    }
  return -1
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
   if len(slice) == index { // nil or empty slice or after last element
		return append(slice, value)
	}
	//slice = append(slice[:index+1], slice[index:]...) // index < len(slice)

	if index < 0 {

		//slice[len(slice) + index] = value -> overwrites the last element
        slice = append(slice[:], value) // negative index 
		return slice
	}
	slice[index] = value
	return slice
	
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, value ...int) []int {
       slice = append(value[:], slice...)
   
return slice
  

        

	

   
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
   if (index < 0 )  || (index > len(slice)){
       return slice
   }
		
	

   
  
	return append(slice[:index], slice[index+1:]...)
}
