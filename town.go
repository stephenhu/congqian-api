package main

type Demography struct {
	U10							int
	U20							int
	U30							int
	U40							int
	U50							int
	U60							int
	U70							int
	U80							int
	U90							int
	U100						int	
}

type Town struct {
	Males						Demography
	Females					Demography
	Latitude				int
	Longitude       int
	Mayor           int
	Wealth          int
	Supply          map[string] Resource
}
