package lib

import (
	"strconv"

	"github.com/google/uuid"
)

func ParseInt32Val(a string) int32 {
	if a == "" {
		return 0
	}
	v, err := strconv.Atoi(a)
	if err != nil {
		return 0
	}
	i32 := int32(v)
	return i32
}

func ParseInt32Ptr(a string) *int32 {
	if a == "" {
		return nil
	}
	v, err := strconv.Atoi(a)
	if err != nil {
		return nil
	}
	i32 := int32(v)
	return &i32
}

func ParseIntPtr(a string) *int {
	if a == "" {
		return nil
	}
	v, err := strconv.Atoi(a)
	if err != nil {
		return nil
	}
	return &v
}

func ParseUUID(a string) uuid.UUID {
	uid, err := uuid.Parse(a)
	if err != nil {
		return uuid.Nil
	}
	return uid
}