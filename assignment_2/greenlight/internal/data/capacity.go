package data

import (
	"errors" // New import
	"fmt"
	"strconv"
	"strings" // New import
)

var ErrInvalidCapacityFormat = errors.New("invalid capacity format")

// Declare a custom Capacity type, which has the underlying type int32 (the same as our
// Blender struct field).
type Capacity int32

// Implement a MarshalJSON() method on the Capacity type so that it satisfies the
// json.Marshaler interface. This should return the JSON-encoded value for the blender
// capacity (in our case, it will return a string in the format "<capacity> litres").
func (r Capacity) MarshalJSON() ([]byte, error) {
	// Generate a string containing the blender capacity in the required format.
	jsonValue := fmt.Sprintf("%d litres", r)

	// Use the strconv.Quote() function on the string to wrap it in double quotes. It
	// needs to be surrounded by double quotes in order to be a valid *JSON string*.
	quotedJSONValue := strconv.Quote(jsonValue)

	// Convert the quoted string value to a byte slice and return it.
	return []byte(quotedJSONValue), nil
}

// Implement a UnmarshalJSON() method on the Capacity type so that it satisfies the
// json.Unmarshaler interface. IMPORTANT: Because UnmarshalJSON() needs to modify the
// receiver (our Capacity type), we must use a pointer receiver for this to work
// correctly. Otherwise, we will only be modifying a copy (which is then discarded when
// this method returns).
func (r *Capacity) UnmarshalJSON(jsonValue []byte) error {
	// We expect that the incoming JSON value will be a string in the format
	// "<Capacity> litres", and the first thing we need to do is remove the surrounding
	// double-quotes from this string. If we can't unquote it, then we return the
	// ErrInvalidCapacityFormat error.
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidCapacityFormat
	}

	// Split the string to isolate the part containing the number.
	parts := strings.Split(unquotedJSONValue, " ")

	// Sanity check the parts of the string to make sure it was in the expected format.
	// If it isn't, we return the ErrInvalidCapacityFormat error again.
	if len(parts) != 2 || parts[1] != "litres" {
		return ErrInvalidCapacityFormat
	}

	// Otherwise, parse the string containing the number into an int32. Again, if this
	// fails return the ErrInvalidCapacityFormat error.
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidCapacityFormat
	}
	
	// Convert the int32 to a Capacity type and assign this to the receiver. Note that we
	// use the * operator to deference the receiver (which is a pointer to a Capacity
	// type) in order to set the underlying value of the pointer.
	*r = Capacity(i)
	return nil
}
	