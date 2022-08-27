package main

type Resource struct {
	Name						string			`json:"name"`
	Description			string			`json:"description"`
	Quality					int					`json:"quality"`
	Quantity				int					`json:"quantity"`
	IsEdible        bool				`json:"isEdible"`
}
