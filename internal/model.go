package main

type Services map[string]Service

type Service struct {
	Repository interface{}
	Service    interface{}
	Handler    interface{}
}
