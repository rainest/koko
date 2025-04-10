package tests

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"testing"

	internalJSON "github.com/kong/koko/internal/json"
	"github.com/kong/koko/internal/persistence"
	"github.com/kong/koko/internal/test/util"
	"github.com/stretchr/testify/require"
)

type jsonWrapper struct {
	Value string `json:"value"`
}

func json(value string) []byte {
	res, err := internalJSON.ProtoJSONMarshal(jsonWrapper{value})
	if err != nil {
		panic(fmt.Sprintf("marshal json: %v", err))
	}
	return res
}

func equalJSON(t *testing.T, expected, actual []byte) {
	require.JSONEq(t, string(expected), string(actual))
}

func TestPersister(t *testing.T) {
	p, err := util.GetPersister(t)
	require.Nil(t, err)

	t.Run("Get()", func(t *testing.T) {
		t.Run("returns an existing value", func(t *testing.T) {
			// put
			value := json("value1")
			require.Nil(t, p.Put(context.Background(), "key1", value))
			// get
			gotValue, err := p.Get(context.Background(), "key1")
			require.Nil(t, err)
			equalJSON(t, value, gotValue)
		})
		t.Run("returns ErrNotFound when value doesn't exist", func(t *testing.T) {
			gotValue, err := p.Get(context.Background(),
				"key-does-not-exist")
			require.Equal(t, persistence.ErrNotFound{Key: "key-does-not-exist"}, err)
			require.Nil(t, gotValue)
		})
	})
	t.Run("Put()", func(t *testing.T) {
		t.Run("stores a value", func(t *testing.T) {
			// put
			value := json("value2")
			require.Nil(t, p.Put(context.Background(), "key2", value))
			// get
			gotValue, err := p.Get(context.Background(), "key2")
			require.Nil(t, err)
			equalJSON(t, value, gotValue)
		})
		t.Run("overwrites a value", func(t *testing.T) {
			// put
			value := json("value3")
			require.Nil(t, p.Put(context.Background(), "key3", value))
			// get
			gotValue, err := p.Get(context.Background(), "key3")
			require.Nil(t, err)
			equalJSON(t, value, gotValue)

			value = json("value3-new")
			require.Nil(t, p.Put(context.Background(), "key3", value))
			// get
			gotValue, err = p.Get(context.Background(), "key3")
			require.Nil(t, err)
			equalJSON(t, value, gotValue)
		})
	})
	t.Run("Delete()", func(t *testing.T) {
		t.Run("deletes when key exist", func(t *testing.T) {
			value := json("value4")
			require.Nil(t, p.Put(context.Background(), "key4", value))
			require.Nil(t, p.Delete(context.Background(), "key4"))
			// get after delete
			gotValue, err := p.Get(context.Background(), "key4")
			require.ErrorAs(t, err,
				&persistence.ErrNotFound{Key: "key4"})
			require.Nil(t, gotValue)
		})
		t.Run("deletes fails when a key does not exist",
			func(t *testing.T) {
				require.Equal(t,
					persistence.ErrNotFound{Key: "key-no-exist"}, p.Delete(
						context.Background(), "key-no-exist"))
			})
	})
	t.Run("List()", func(t *testing.T) {
		t.Run("lists all keys with prefix", func(t *testing.T) {
			var expectedValues, expectedKeys []string
			for i := 0; i < 100; i++ {
				value := json(fmt.Sprintf("prefix-value-%d", i))
				key := fmt.Sprintf("prefix/key%d", i)
				require.Nil(t, p.Put(context.Background(), key, value))
				expectedKeys = append(expectedKeys, key)
				expectedValues = append(expectedValues, string(value))
			}
			listResult, err := p.List(context.Background(), "prefix/", persistence.NewDefaultListOpts())
			require.Nil(t, err)
			require.Len(t, listResult.KVList, 100)

			var valuesAsStrings []string
			var keysAsStrings []string
			for _, kv := range listResult.KVList {
				key := string(kv.Key)
				value := string(kv.Value)
				keysAsStrings = append(keysAsStrings, key)
				value = strings.ReplaceAll(value, " ", "")
				valuesAsStrings = append(valuesAsStrings, value)
			}
			sort.Strings(keysAsStrings)
			sort.Strings(expectedKeys)
			sort.Strings(valuesAsStrings)
			sort.Strings(expectedValues)

			require.Equal(t, expectedKeys, keysAsStrings)
			require.Equal(t, expectedValues, valuesAsStrings)
		})
		t.Run("other prefixes are left as is", func(t *testing.T) {
			var expected []string
			for i := 0; i < 100; i++ {
				value := json(fmt.Sprintf("prefix-value-%d", i))
				require.Nil(t, p.Put(
					context.Background(),
					fmt.Sprintf("prefix/key%d", i),
					value))
				expected = append(expected, string(value))
				// other prefixes
				value = json(fmt.Sprintf("other-prefix-value-%d", i))
				require.Nil(t, p.Put(
					context.Background(),
					fmt.Sprintf("ix/prefix/key%d", i),
					value))
			}
			listResult, err := p.List(context.Background(), "prefix/", persistence.NewDefaultListOpts())
			require.Nil(t, err)
			require.Len(t, listResult.KVList, 100)
			var valuesAsStrings []string
			for _, value := range listResult.KVList {
				valuesAsStrings = append(valuesAsStrings, string(value.Value))
			}
			sort.Strings(valuesAsStrings)
			sort.Strings(expected)
			// require.Equal(t, expected, valuesAsStrings)
		})
		t.Run("list with non default pagination", func(t *testing.T) {
			var expectedValuesBatchOne, expectedKeysBatchOne, expectedValuesBatchTwo, expectedKeysBatchTwo []string
			for i := 0; i < 200; i++ {
				value := json(fmt.Sprintf("prefix-value-%06d", i))
				key := fmt.Sprintf("myprefix/key%06d", i)
				require.Nil(t, p.Put(context.Background(), key, value))
				if i < 100 {
					expectedKeysBatchOne = append(expectedKeysBatchOne, key)
					expectedValuesBatchOne = append(expectedValuesBatchOne, string(value))
				} else {
					expectedKeysBatchTwo = append(expectedKeysBatchTwo, key)
					expectedValuesBatchTwo = append(expectedValuesBatchTwo, string(value))
				}
			}
			listResult, err := p.List(context.Background(), "myprefix/", &persistence.ListOpts{
				Offset: 0,
				Limit:  100,
			})
			require.Nil(t, err)
			require.Len(t, listResult.KVList, 100)

			var valuesAsStrings []string
			var keysAsStrings []string
			for _, kv := range listResult.KVList {
				key := string(kv.Key)
				value := string(kv.Value)
				keysAsStrings = append(keysAsStrings, key)
				value = strings.ReplaceAll(value, " ", "")
				valuesAsStrings = append(valuesAsStrings, value)
			}
			sort.Strings(keysAsStrings)
			sort.Strings(expectedKeysBatchOne)
			sort.Strings(valuesAsStrings)
			sort.Strings(expectedValuesBatchOne)
			require.Equal(t, expectedKeysBatchOne, keysAsStrings)
			require.Equal(t, expectedValuesBatchOne, valuesAsStrings)

			// Now get the Second batch
			listResult, err = p.List(context.Background(), "myprefix/", &persistence.ListOpts{
				Offset: 100,
				Limit:  100,
			})
			require.Nil(t, err)
			require.Len(t, listResult.KVList, 100)

			var valuesAsStringsTwo []string
			var keysAsStringsTwo []string
			for _, kv := range listResult.KVList {
				key := string(kv.Key)
				value := string(kv.Value)
				keysAsStringsTwo = append(keysAsStringsTwo, key)
				value = strings.ReplaceAll(value, " ", "")
				valuesAsStringsTwo = append(valuesAsStringsTwo, value)
			}
			sort.Strings(keysAsStringsTwo)
			sort.Strings(expectedKeysBatchTwo)
			sort.Strings(valuesAsStringsTwo)
			sort.Strings(expectedValuesBatchTwo)
			require.Equal(t, expectedKeysBatchTwo, keysAsStringsTwo)
			require.Equal(t, expectedValuesBatchTwo, valuesAsStringsTwo)

			// See if there is more
			listResult, err = p.List(context.Background(), "myprefix/", &persistence.ListOpts{
				Offset: 200,
				Limit:  100,
			})
			require.Nil(t, err)
			require.Len(t, listResult.KVList, 0)
		})
	})
	t.Run("Tx()", func(t *testing.T) {
		t.Run("transaction rollbacks correctly", func(t *testing.T) {
			ctx := context.Background()
			tx, err := p.Tx(ctx)
			require.Nil(t, err)
			value := json("value5")
			err = tx.Put(ctx, "key5", value)
			require.Nil(t, err)
			gotValue, err := tx.Get(ctx, "key5")
			require.Nil(t, err)
			equalJSON(t, value, gotValue)
			err = tx.Rollback()
			require.Nil(t, err)
			value, err = p.Get(ctx, "key5")
			require.Equal(t, persistence.ErrNotFound{Key: "key5"}, err)
			require.Nil(t, value)
		})
	})
}
