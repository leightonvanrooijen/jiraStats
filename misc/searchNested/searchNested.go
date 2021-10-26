package searchNested

// TODO change this to accept a *nested list to find 
// Finds first instance of the key input and returns its value
func Search(obj interface{}, key string) (interface{}, bool) {
	switch objtype := obj.(type) {
	case map[string]interface{}:
		if value, ok := objtype[key]; ok {
			return value, ok
		}
		for _, value := range objtype {
			if result, ok := Search(value, key); ok {
				return result, ok
			}
		}
	case []interface{}:
		for _, value := range objtype {
			if result, ok := Search(value, key); ok {
				return result, ok
			}
		}
	}

	// key not found
	return nil, false
}
