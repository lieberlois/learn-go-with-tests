package dictionary

// Dictionary is a wrapper for a map
type Dictionary map[string]string

const (
	// ErrNotFound is the error for looking up a non-existent key
	ErrNotFound = DictErr("could not find the word you were looking for")
	
	// ErrWordExists is the error for adding a key that already exists
	ErrWordExists = DictErr("key already exists in map")
	
	// ErrWordDoesNotExist is the error for updating a value to an invalid key
	ErrWordDoesNotExist = DictErr("cannot update word because it does not exist")
)

// DictErr is a type for Errors related to Dictionaries
type DictErr string

func (e DictErr) Error() string {
	return string(e)
}


// Search finds a key in a dictionary
func (dictionary Dictionary) Search(word string) (string, error) {
	definition, ok := dictionary[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add adds an element to the Dictionary
func (dictionary Dictionary) Add(key string, value string) (error) {
	if _, ok := dictionary[key]; ok {
		return ErrWordExists
	}
	dictionary[key] = value
	return nil
}

// Update changes the value to a key
func (dictionary Dictionary) Update(key, newDefinition string) error {
	if _, ok := dictionary[key]; !ok {
		return ErrWordDoesNotExist
	}
	dictionary[key] = newDefinition
	return nil
}

// Delete deletes a key-value pair from the dictionary
func (dictionary Dictionary) Delete(key string) {
	delete(dictionary, key)
}