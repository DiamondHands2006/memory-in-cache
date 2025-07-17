/*
*
git add .     add to index
git commit -m "Commit for example" create a commit
git push final step,push our new commit
*
*/
package memoryincache

type Cache interface {
	New()
	Get(key string)
	Set(key string, value interface{})
	Delete(key string)
}
