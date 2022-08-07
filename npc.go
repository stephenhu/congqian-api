package main

import (
	"fmt"
	"log"
)


func npcId(kingdom string, municipal string, family string, name string) string {

	h := hash(fmt.Sprintf("%s%s%s%s%s%d", kingdom, municipal, family, name,
	  initRating(DEFAULT_BIG)), DEFAULT_HASH_LENGTH)

	return fmt.Sprintf("%s:%s", KEY_NPC, h)

} // npcId


func generateKingdom(k *Kingdom) {

	for i := 0; i < k.Population; i++ {

		family 			:= initFamily()
		name        := initName()
		municipal   := initMunicipal(k.Name)
	
		key := npcId(k.Name, municipal, family, name)
	
		err := rds.HSet(ctx, key,
			ATTR_AGE, initRangeRating(DEFAULT_NPC_AGE_MIN, DEFAULT_NPC_AGE_MAX),
			ATTR_KINGDOM, k.Name,
			ATTR_FEMALE, initGender(),
			ATTR_HEIGHT, initRangeRating(DEFAULT_HEIGHT_MIN, DEFAULT_HEIGHT_MAX),
			ATTR_FAMILY, family,
			ATTR_NAME, name,
			ATTR_EXPERIENCE, initRating(DEFAULT_SKILL_MAX),
			ATTR_LUCK, initRating(DEFAULT_SKILL_MAX),
			ATTR_VIRTUE, initRating(DEFAULT_SKILL_MAX),
			ATTR_LOYALTY, initRating(DEFAULT_SKILL_MAX),
			ATTR_POPULARITY, initRating(DEFAULT_SKILL_MAX),
			ATTR_HEALTH, initRating(DEFAULT_SKILL_MAX),
			ATTR_ENERGY, initRating(DEFAULT_SKILL_MAX),
			ATTR_INTELLIGENCE, initRating(DEFAULT_SKILL_MAX),
			ATTR_WISDOM, initRating(DEFAULT_SKILL_MAX),
			ATTR_INTEGRITY, initRating(DEFAULT_SKILL_MAX),
			ATTR_DEXTERITY, initRating(DEFAULT_SKILL_MAX),
			ATTR_CHARISMA, initRating(DEFAULT_SKILL_MAX),
			ATTR_ATTRACTIVENESS, initRating(DEFAULT_SKILL_MAX),
			ATTR_STRENGTH, initRating(DEFAULT_SKILL_MAX),
			ATTR_STAMINA, initRating(DEFAULT_SKILL_MAX),
			ATTR_BALANCE, initRating(DEFAULT_SKILL_MAX),
			ATTR_TEMPERANCE, initRating(DEFAULT_SKILL_MAX),
			ATTR_RECOVERY, initRating(DEFAULT_SKILL_MAX),
			ATTR_INNOVATION, initRating(DEFAULT_SKILL_MAX),
			ATTR_DAGGER, initRating(DEFAULT_SKILL_MAX),
			ATTR_SWORD, initRating(DEFAULT_SKILL_MAX),
			ATTR_SCIMITAR, initRating(DEFAULT_SKILL_MAX),
			ATTR_SPEAR, initRating(DEFAULT_SKILL_MAX),
			ATTR_HOOKED_SPEAR, initRating(DEFAULT_SKILL_MAX),
			ATTR_STAFF, initRating(DEFAULT_SKILL_MAX),
			ATTR_DRAGON_BLADE, initRating(DEFAULT_SKILL_MAX),
			ATTR_CLUB, initRating(DEFAULT_SKILL_MAX),
			ATTR_HAMMER, initRating(DEFAULT_SKILL_MAX),
			ATTR_AXE, initRating(DEFAULT_SKILL_MAX),
			ATTR_HALBERD, initRating(DEFAULT_SKILL_MAX),
			ATTR_BOW, initRating(DEFAULT_SKILL_MAX),
			ATTR_CROSSBOW, initRating(DEFAULT_SKILL_MAX),
			ATTR_HORSE_BOW, initRating(DEFAULT_SKILL_MAX),
			ATTR_HORSE, initRating(DEFAULT_SKILL_MAX),
			ATTR_CHARIOT, initRating(DEFAULT_SKILL_MAX),
			ATTR_SHIELD, initRating(DEFAULT_SKILL_MAX),
			ATTR_BEAR, initRating(DEFAULT_SKILL_MAX),
			ATTR_CRANE, initRating(DEFAULT_SKILL_MAX),
			ATTR_DRUNKEN, initRating(DEFAULT_SKILL_MAX),
			ATTR_IRONFIST, initRating(DEFAULT_SKILL_MAX),
			ATTR_LEOPARD, initRating(DEFAULT_SKILL_MAX),
			ATTR_MANTIS, initRating(DEFAULT_SKILL_MAX),
			ATTR_MONKEY, initRating(DEFAULT_SKILL_MAX),
			ATTR_SNAKE, initRating(DEFAULT_SKILL_MAX),
			ATTR_TIGER, initRating(DEFAULT_SKILL_MAX),
		).Err()
	
		if err != nil {
			log.Println(err)
		}

	}

} // generateKingdom
