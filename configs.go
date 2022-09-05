package main

import (

)


type GlobalConf struct {
	HashLength 					int					`redis:"hash:length"`
	Female 							int					`redis:"female"`
	Male 								int					`redis:"male"`
	SkillMax 						int					`redis:"skill:max"`
}


type InitConf struct {
	AgeMin 									int					`redis:"age:min"`
	AgeMax 									int					`redis:"age:max"`
	AreaMin 								int					`redis:"area:min"`
	AreaMax 								int					`redis:"area:max"`
	HeightMin               int         `redis:"height:min"`
	HeightMax               int         `redis:"height:max"`
	PopulationTiny 					int					`redis:"population:tiny"`
	PopulationSmall 				int					`redis:"population:small"`
	PopulationMedium 				int					`redis:"population:medium"`
	PopulationLarge 				int					`redis:"population:large"`
	PopulationXLarge 				int					`redis:"population:xlarge"`
	PopulationMassive 			int					`redis:"population:assive"`
	PopulationMetropolis 		int					`redis:"population:metropolis"`
	SkillMax 								int					`redis:"skill:max"`
	WealthMax 							int					`redis:"wealth:max"`
	NpcAgeMin 							int					`redis:"npc:age:min"`
	NpcAgeMax 							int					`redis:"npc:age:max"`
}
