package redis

// Get - gets the value at a key
func Get(key string) (string, error) {
	return RedisClient.Get(Context, key).Result()
}

// GetKeys - get a list of keys that match the provided pattern
func GetKeys(pattern string) ([]string, error) {
	keys := []string{}
	iter := RedisClient.Scan(Context, 0, pattern, 0).Iterator()

	for iter.Next(Context) {
		keys = append(keys, iter.Val())
	}

	if err := iter.Err(); err != nil {
		return nil, err
	}

	return keys, nil
}
