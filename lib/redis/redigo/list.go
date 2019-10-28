package redigo

import (
	"context"

	redigo "github.com/gomodule/redigo/redis"
)

// LLen get the length of the list
func (rdg *Redigo) LLen(ctx context.Context, key string) (int, error) {
	result, err := redigo.Int(rdg.do(ctx, "LLEN", key))
	if err != nil && !IsErrNil(err) {
		return 0, err
	}
	return result, err
}

// LIndex to get value from a certain list index
func (rdg *Redigo) LIndex(ctx context.Context, key string, index int) (string, error) {
	result, err := redigo.String(rdg.do(ctx, "LINDEX", key, index))
	if err != nil && !IsErrNil(err) {
		return "", err
	}
	return result, err
}

// LSet to set value to some index
func (rdg *Redigo) LSet(ctx context.Context, key string, index int, value string) (int, error) {
	result, err := redigo.Int(rdg.do(ctx, "LSET", index, value))
	if err != nil && !IsErrNil(err) {
		return 0, err
	}
	return result, err
}

// LPush prepend values to the list
func (rdg *Redigo) LPush(ctx context.Context, key string, values ...string) (int, error) {
	args := make([]interface{}, len(values)+1)
	args[0] = key
	for i, value := range values {
		args[i+1] = value
	}

	result, err := redigo.Int(rdg.do(ctx, "LPUSH", args...))
	if err != nil && !IsErrNil(err) {
		return 0, err
	}
	return result, err
}

// LPushX prepend values to the list
func (rdg *Redigo) LPushX(ctx context.Context, key string, values ...string) (int, error) {
	args := make([]interface{}, len(values)+1)
	args[0] = key
	for i, value := range values {
		args[i+1] = value
	}

	result, err := redigo.Int(rdg.do(ctx, "LPUSHX", args...))
	if err != nil && !IsErrNil(err) {
		return 0, err
	}
	return result, err
}

// LPop removes and get the first element in the list
func (rdg *Redigo) LPop(ctx context.Context, key string) (string, error) {
	result, err := redigo.String(rdg.do(ctx, "LPOP", key))
	if err != nil && !IsErrNil(err) {
		return "", err
	}
	return result, err
}

// LRem command
func (rdg *Redigo) LRem(ctx context.Context, count int, value string) (int, error) {
	result, err := redigo.Int(rdg.do(ctx, "LREM", count, value))
	if err != nil && !IsErrNil(err) {
		return 0, err
	}
	return result, err
}