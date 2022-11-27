package ngo

type Action func()

type ActionT1[T any] func(t T)

type ActionT2[T1,T2 any] func(t1 T1,t2 T2)