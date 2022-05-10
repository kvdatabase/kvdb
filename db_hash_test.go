package kvdb

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var key = "myhash"

func TestRoseDB_HSet(t *testing.T) {

	t.Run("test1", func(t *testing.T) {
		db := InitDb()
		defer db.Close()

		db.HSet(nil, nil, nil)

		_, _ = db.HSet([]byte(key), []byte("my_name"), []byte("kvdatabase"))
	})

	t.Run("reopen and set", func(t *testing.T) {
		db := ReopenDb()
		defer db.Close()
		_, _ = db.HSet([]byte(key), []byte("my_hobby"), []byte("coding better"))
		_, _ = db.HSet([]byte(key), []byte("my_lang"), []byte("Java and Go"))
	})

	//t.Run("multi data", func(t *testing.T) {
	//	db := ReopenDb()
	//	defer db.Close()
	//
	//	rand.Seed(time.Now().Unix())
	//
	//	fieldPrefix := "hash_field_"
	//	valPrefix := "hash_data_"
	//
	//	var res int
	//	for i := 0; i < 100000; i++ {
	//		field := fieldPrefix + strconv.Itoa(rand.Intn(1000000))
	//		val := valPrefix + strconv.Itoa(rand.Intn(1000000))
	//
	//		res, _ = db.HSet([]byte(key), []byte(field), []byte(val))
	//	}
	//	t.Log(res)
	//})
}

func TestRoseDB_HSetNx(t *testing.T) {
	db := ReopenDb()
	defer db.Close()

	db.HSetNx(nil, nil, nil)

	ok, _ := db.HSetNx([]byte(key), []byte("my_hobby"), []byte("coding better"))
	t.Log(ok)
	ok, _ = db.HSetNx([]byte(key), []byte("my_new_lang"), []byte("Java Go Python"))
	t.Log(ok)

	t.Log(db.HLen([]byte(key)))
}

func TestRoseDB_HGet(t *testing.T) {
	db := ReopenDb()
	defer db.Close()

	_ = db.HGet([]byte(key), []byte("my_name"))
	_ = db.HGet([]byte(key), []byte("not exist"))
	_ = db.HGet([]byte(key), []byte("my_hobby"))

	_ = db.HGet([]byte(key), []byte("hash_field_732328"))
	_ = db.HGet([]byte(key), []byte("hash_field_112243"))
}

func TestRoseDB_HGetAll(t *testing.T) {
	db := ReopenDb()
	defer db.Close()

	values := db.HGetAll([]byte(key))
	for _, v := range values {
		t.Log(string(v))
	}
}

func TestRoseDB_HDel(t *testing.T) {
	db := ReopenDb()
	defer db.Close()

	db.HDel(nil, nil)

	res, _ := db.HDel([]byte(key), []byte("my_name"), []byte("my_name2"), []byte("my_name3"))
	t.Log(res)
}

func TestRoseDB_HExists(t *testing.T) {
	db := ReopenDb()
	defer db.Close()

	db.HExists(nil, nil)

	ok := db.HExists([]byte(key), []byte("my_name"))
	t.Log(ok)

	t.Log(db.HExists([]byte(key), []byte("my_hobby")))
	t.Log(db.HExists([]byte(key), []byte("my_name1")))
	t.Log(db.HExists([]byte(key+"abcd"), []byte("my_hobby")))
}

func TestRoseDB_HKeys(t *testing.T) {
	db := ReopenDb()
	defer db.Close()

	db.HKeys(nil)
	keys := db.HKeys([]byte(key))
	for _, k := range keys {
		t.Log(k)
	}
}

func TestRoseDB_HVals(t *testing.T) {
	db := ReopenDb()
	defer db.Close()

	db.HVals(nil)
	keys := db.HVals([]byte(key))
	for _, k := range keys {
		t.Log(string(k))
	}
}

func TestRoseDB_HLen(t *testing.T) {
	db := ReopenDb()
	defer db.Close()

	db.HLen(nil)

	db.HLen([]byte("11"))
}

func TestRoseDB_HExpire(t *testing.T) {
	db := InitDb()
	defer db.Close()

	key := []byte("hash_key")
	res, err := db.HSet(key, []byte("a"), []byte("hash-val-1"))
	assert.Equal(t, err, nil)
	assert.Equal(t, res, 1)

	err = db.HExpire(key, 10)
	assert.Equal(t, err, nil)
}

func TestRoseDB_HTTL(t *testing.T) {
	db := InitDb()
	defer db.Close()

	key := []byte("hash_key_2")

	res, err := db.HSet(key, []byte("bb"), []byte("hash-val-2"))
	assert.Equal(t, res, 1)
	assert.Equal(t, err, nil)

	err = db.HExpire(key, 30)
	assert.Equal(t, err, nil)

	for i := 0; i < 5; i++ {
		t.Log(db.HTTL(key))
		//time.Sleep(time.Second * 2)
	}
}

func TestRoseDB_HClear(t *testing.T) {
	db := InitDb()
	defer db.Close()

	key := []byte("hash_key_3")

	res, err := db.HSet(key, []byte("bb"), []byte("hash-val-2"))
	assert.Equal(t, res, 1)
	assert.Equal(t, err, nil)

	err = db.HClear(key)
	assert.Equal(t, err, nil)
}

func TestRoseDB_HMSet(t *testing.T) {
	db := InitDb()
	defer db.Close()

	t.Run("wrong number", func(t *testing.T) {
		key := []byte("hash_batch_key")
		err := db.HMSet(key, []byte("field1"))
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, ErrWrongNumberOfArgs)
	})

	t.Run("wrong field", func(t *testing.T) {
		largeValue := strings.Builder{}
		// 9mb
		largeValue.Grow(int(DefaultMaxValueSize + DefaultMaxKeySize))
		for i := 0; i < int(DefaultMaxValueSize+DefaultMaxKeySize); i++ {
			largeValue.WriteByte('A')
		}

		key := []byte("hash_batch_key")
		err := db.HMSet(key, []byte(largeValue.String()), []byte("field"))
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, ErrValueTooLarge)
	})

	t.Run("wrong values", func(t *testing.T) {
		largeValue := strings.Builder{}
		// 9mb
		largeValue.Grow(int(DefaultMaxValueSize + DefaultMaxKeySize))
		for i := 0; i < int(DefaultMaxValueSize+DefaultMaxKeySize); i++ {
			largeValue.WriteByte('A')
		}

		key := []byte("hash_batch_key")
		err := db.HMSet(key, []byte("field1"), []byte(largeValue.String()))
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, ErrValueTooLarge)
	})

	t.Run("success", func(t *testing.T) {
		key := []byte("hash_batch_key")
		err := db.HMSet(key, []byte("field1"), []byte("hello"), []byte("field2"), []byte("world"))
		assert.Nil(t, err)
	})
}

func TestRoseDB_HMGet(t *testing.T) {
	db := InitDb()
	defer db.Close()

	key := []byte("hash_batch_key")

	t.Run("empty", func(t *testing.T) {
		res := db.HMGet(key, []byte("not"), []byte("found"))
		assert.Equal(t, 2, len(res))
		assert.Equal(t, 0, len(res[0]))
		assert.Equal(t, 0, len(res[1]))
	})

	t.Run("success", func(t *testing.T) {
		err := db.HMSet(key, []byte("field1"), []byte("hello"), []byte("field2"), []byte("world"))
		assert.Nil(t, err)
		res := db.HMGet(key, []byte("field1"), []byte("field2"))
		assert.Equal(t, 2, len(res))
		assert.Equal(t, "hello", string(res[0]))
		assert.Equal(t, "world", string(res[1]))
	})
}
