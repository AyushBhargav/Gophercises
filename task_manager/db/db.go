package db

import "github.com/boltdb/bolt"
import "strconv"

func openDB() *bolt.DB {
	db, err := bolt.Open("tasks.db", 0777, nil)
	if err != nil {
		panic(err)
	}
	return db
}

// Init bold for storage
func Init() {
	db := openDB()
	defer db.Close()

	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Tasks"))
		if err != nil {
			panic(err)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

// CreateNewTask in Bolt
func CreateNewTask(task string) {
	db := openDB()
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		id, _ := b.NextSequence()
		b.Put([]byte(string(id + '0')), []byte(task))
		return nil
	})
}

// GetIncompleteTasks returns list of pending tasks
func GetIncompleteTasks() map[string]string {
	db := openDB()
	defer db.Close()

	tasks := make(map[string]string)
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		c := b.Cursor()
		for k, task := c.First(); k != nil; k, task = c.Next() {
			tasks[string(k)] = string(task)
		}
		return nil
	})
	return tasks
}

// MarkComplete removes task from database
func MarkComplete(taskKey string) {
	db := openDB()
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		b.Delete([]byte(taskKey))
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		min := []byte(taskKey)
		c := b.Cursor()
		for k, task := c.Seek(min); k != nil; k, task = c.Next() {
			index, e := strconv.Atoi(string(k))
			if e != nil {
				panic(e)
			}
			k = []byte(string(index + '0' - 1))
			b.Put(k, task)
		}
		k, _ := c.Last()
		b.Delete([]byte(k))
		return nil
	})
}
