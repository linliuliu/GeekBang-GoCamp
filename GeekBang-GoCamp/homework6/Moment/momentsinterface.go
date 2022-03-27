package main

type MomentsInterface interface {
	CreateMoment() error
	DeleteMoment()
	GetMoment() ([] *Moment)
}

