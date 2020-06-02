package gouid

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"time"
)

var (
	DefaultNode  = 1
	DefaultEpoch = "2020-06-02"
)

type Options struct {
	Node  int    // 范围：0 - 1023
	Epoch string // 格式：2006-01-02

	epoch time.Time
}

type Option func(o *Options)

func Node(i int) Option {
	return func(o *Options) {
		o.Node = i
	}
}

func Epoch(s string) Option {
	return func(o *Options) {
		o.Epoch = s
	}
}

func NewSnowflake(options ...Option) *snowflake.Node {
	return newSnowflake(options...)
}

func newSnowflake(options ...Option) *snowflake.Node {
	o := newOptions(options...)
	snowflake.Epoch = o.epoch.Unix() * 1000
	s, err := snowflake.NewNode(int64(o.Node))
	if err != nil {
		panic(err)
	}
	return s
}

func newOptions(options ...Option) Options {
	o := Options{
		Node:  DefaultNode,
		Epoch: DefaultEpoch,
	}

	for _, v := range options {
		v(&o)
	}

	if o.Node < 0 || o.Node > 1023 {
		panic(fmt.Sprintf("invalid node: %d, support range: [0,1023]", o.Node))
	}

	epoch, err := time.Parse("2006-01-02", o.Epoch)
	if err != nil {
		panic(fmt.Sprintf("invalid epoch: %s", o.Epoch))
	}
	o.epoch = epoch

	return o
}
