package channel

import "github.com/sslime336/iter"

func Iter[T any](channel chan T) iter.ChannelIter[T] {
	return &wrapper[T]{
		inner: channel,
	}
}
