package main

const (
	APP_NAME      			= "congqian"
	APP_VERSION					= "1.0"
	SALT                = "better not try this at home"
)


const (
	DEFAULT_REDIS_HOST					= "localhost"
	DEFAULT_REDIS_PORT        	= "6379"
)


const (
	STRING_EMPTY				= ""
)


const (
	KEY_FAMILIES                = "families"
	KEY_FAMILY                	= "family"
	KEY_KINGDOM                 = "kingdom"
	KEY_KINGDOMS                = "kingdoms"	
	KEY_MUNICIPAL          			= "municipal"
	KEY_MUNICIPALS          		= "municipals"
	KEY_NAMES          					=	"names"
	KEY_NPC                     = "npc"
	KEY_NPCS                    = "npcs"
	KEY_USER                    = "user"
	KEY_USERS                   = "users"
)


const (
	DEFAULT_AGE_MAX           	= 26
	DEFAULT_AGE_MIN           	= 17
	DEFAULT_BIG           			= 35295
	DEFAULT_HASH_LENGTH         = 20
	DEFAULT_HEIGHT_MAX        	= 220
	DEFAULT_HEIGHT_MIN        	= 145
	DEFAULT_FEMALE              = 0
	DEFAULT_MALE                = 1
	DEFAULT_GENDER              = 2
	DEFAULT_SKILL_MAX						= 200
	DEFAULT_WEALTH_MAX          = 200
	SKILL_MAX										= 5000
)


const (
	DEFAULT_NPC_AGE_MAX           	= 70
	DEFAULT_NPC_AGE_MIN           	= 1
)


const (
	ERR_HASH_INVALID_LENGTH			= "Invalid hash length, using default"
	ERR_NO_FAMILIES_FOUND				= "No families found."
	ERR_NO_KINGDOMS_FOUND				= "No kingdoms found."
	ERR_NO_MUNICIPALS_FOUND     = "No municipals found"
	ERR_NO_NAMES_FOUND     			= "No names found"
	ERR_SET_NOT_FOUND						= "Set not found."
)
