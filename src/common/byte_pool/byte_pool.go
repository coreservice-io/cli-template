package byte_pool

import (
	"errors"
	"sync"
)

type BytePool struct {
	name      string
	sync_p    sync.Pool
	page_size int
}

type BytePage struct {
	content   []byte
	byte_pool *BytePool
}

func (bp *BytePage) Recycle() {
	bp.byte_pool.sync_p.Put(bp.content)
}

func (bp *BytePage) GetBytes() []byte {
	return bp.content
}

// return nil if size > page_size
func (bp *BytePage) GetBytesWithSize(size int) []byte {
	if size <= 0 || size > bp.byte_pool.page_size {
		return nil
	} else {
		return bp.content[:size]
	}
}

var all_pools = make(map[string]*BytePool)
var new_pool_lock sync.Mutex

func NewPool(p_name string, size int) (*BytePool, error) {

	new_pool_lock.Lock()
	defer new_pool_lock.Unlock()

	if all_pools[p_name] != nil {
		return nil, errors.New("NewBytePool pool already exist err")
	}

	all_pools[p_name] = &BytePool{
		name:      p_name,
		page_size: size,
		sync_p: sync.Pool{
			New: func() interface{} {
				buff_ := make([]byte, size)
				return buff_
			},
		}}

	return all_pools[p_name], nil
}

func GetPool(p_name string) *BytePool {
	return all_pools[p_name]
}

func (b_pool *BytePool) Allocate() *BytePage {
	return &BytePage{
		byte_pool: b_pool,
		content:   b_pool.sync_p.Get().([]byte),
	}
}
