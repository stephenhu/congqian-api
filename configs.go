package main

import (

)


type GlobalConf struct {
	HashLength 					int					`json:"hashLength"`
	Female 							int					`json:"female"`
	Male 								int					`json:"male"`
	SkillMax 						int					`json:"skillMax"`
}


type InitConf struct {
	AgeMin 									int					`json:"ageMin"`
	AgeMax 									int					`json:"ageMax"`
	AreaMin 								int					`json:"areaMin"`
	AreaMax 								int					`json:"areaMax"`
	PopulationTiny 					int					`json:"populationTiny"`
	PopulationSmall 				int					`json:"populationSmall"`
	PopulationMedium 				int					`json:"populationMedium"`
	PopulationLarge 				int					`json:"populationLarge"`
	PopulationXLarge 				int					`json:"populationXLarge"`
	PopulationMassive 			int					`json:"populationMassive"`
	PopulationMetropolis 		int					`json:"populationMetropolis"`
	SkillMax 								int					`json:"skillMax"`
	WealthMax 							int					`json:"wealthMax"`
	NpcAgeMin 							int					`json:"npcAgeMin"`
	NpcAgeMax 							int					`json:"npcAgeMax"`
}
