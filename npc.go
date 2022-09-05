package main

import (
	"fmt"
	"log"
)


func npcId(kingdom string, municipal string, family string, name string) string {

	h := hash(fmt.Sprintf("%s%s%s%s%s%d", kingdom, municipal, family, name,
	  initRating(confInit.PopulationSmall)), confGlobal.HashLength)

	return fmt.Sprintf("%s:%s", KEY_NPC, h)

} // npcId


func generateKingdom(k *Kingdom) {

	for i := 0; i < k.Population; i++ {

		family 			:= initFromSet(KEY_FAMILIES)
		name        := initFromSet(KEY_NAMES)
		municipal   := initMunicipal(k.Name)
	
		key := npcId(k.Name, municipal, family, name)
	
		err := rds.HSet(ctx, key,
			ATTR_AGE, initRangeRating(confInit.NpcAgeMin, confInit.NpcAgeMax),
			ATTR_KINGDOM, k.Name,
			ATTR_MUNICIPAL, municipal,
			ATTR_FEMALE, initGender(),
			ATTR_HEIGHT, initRangeRating(confInit.HeightMin, confInit.HeightMax),
			ATTR_FAMILY, family,
			ATTR_NAME, name,
			ATTR_EXPERIENCE, initRating(confInit.SkillMax),
			ATTR_WEALTH, initRating(confInit.WealthMax),
			ATTR_LUCK, initRating(confInit.SkillMax),
			ATTR_VIRTUE, initRating(confInit.SkillMax),
			ATTR_LOYALTY, initRating(confInit.SkillMax),
			ATTR_POPULARITY, initRating(confInit.SkillMax),
			ATTR_HEALTH, initRating(confInit.SkillMax),
			ATTR_ENERGY, initRating(confInit.SkillMax),
			ATTR_INTELLIGENCE, initRating(confInit.SkillMax),
			ATTR_WISDOM, initRating(confInit.SkillMax),
			ATTR_INTEGRITY, initRating(confInit.SkillMax),
			ATTR_DEXTERITY, initRating(confInit.SkillMax),
			ATTR_CHARISMA, initRating(confInit.SkillMax),
			ATTR_ATTRACTIVENESS, initRating(confInit.SkillMax),
			ATTR_STRENGTH, initRating(confInit.SkillMax),
			ATTR_STAMINA, initRating(confInit.SkillMax),
			ATTR_BALANCE, initRating(confInit.SkillMax),
			ATTR_TEMPERANCE, initRating(confInit.SkillMax),
			ATTR_RECOVERY, initRating(confInit.SkillMax),
			ATTR_INNOVATION, initRating(confInit.SkillMax),
			ATTR_DAGGER, initRating(confInit.SkillMax),
			ATTR_SWORD, initRating(confInit.SkillMax),
			ATTR_SCIMITAR, initRating(confInit.SkillMax),
			ATTR_SPEAR, initRating(confInit.SkillMax),
			ATTR_HOOKED_SPEAR, initRating(confInit.SkillMax),
			ATTR_STAFF, initRating(confInit.SkillMax),
			ATTR_DRAGON_BLADE, initRating(confInit.SkillMax),
			ATTR_CLUB, initRating(confInit.SkillMax),
			ATTR_HAMMER, initRating(confInit.SkillMax),
			ATTR_AXE, initRating(confInit.SkillMax),
			ATTR_HALBERD, initRating(confInit.SkillMax),
			ATTR_BOW, initRating(confInit.SkillMax),
			ATTR_CROSSBOW, initRating(confInit.SkillMax),
			ATTR_HORSE_BOW, initRating(confInit.SkillMax),
			ATTR_HORSE, initRating(confInit.SkillMax),
			ATTR_CHARIOT, initRating(confInit.SkillMax),
			ATTR_SHIELD, initRating(confInit.SkillMax),
			ATTR_BEAR, initRating(confInit.SkillMax),
			ATTR_CRANE, initRating(confInit.SkillMax),
			ATTR_DRUNKEN, initRating(confInit.SkillMax),
			ATTR_IRONFIST, initRating(confInit.SkillMax),
			ATTR_LEOPARD, initRating(confInit.SkillMax),
			ATTR_MANTIS, initRating(confInit.SkillMax),
			ATTR_MONKEY, initRating(confInit.SkillMax),
			ATTR_SNAKE, initRating(confInit.SkillMax),
			ATTR_TIGER, initRating(confInit.SkillMax),
		).Err()
	
		if err != nil {
			log.Println(err)
		}

	}

} // generateKingdom
