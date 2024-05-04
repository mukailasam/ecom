package session

import (
	"gopkg.in/boj/redistore.v1"
)

type Session struct {
	RediStore *redistore.RediStore
}
