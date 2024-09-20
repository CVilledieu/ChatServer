package main

type Hub struct {
	HubId uint32
	Name  string
	Users User
}

type ListOfHubs []*Hub
