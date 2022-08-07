package main


const (
	KINGDOM_CHU         		= "chu"
	KINGDOM_HAN        			= "han"
	KINGDOM_QI         			= "qi"	
	KINGDOM_QIN          		= "qin"
	KINGDOM_WEI        			= "wei"
	KINGDOM_YAN       			= "yan"
	KINGDOM_ZHAO						= "zhao"
	KINGDOM_NONE          	= "none"
	KINGDOM_OUTLAW        	= "outlaw"
	KINGDOM_BARBARIAN     	= "barbarian"
)


type Kingdom struct {
	Name									string      `json:"name"`
	Capital               bool        `json:"capital"`
	MaleRatio							int					`json:"maleRatio"`
	MedianAge							int					`json:"medianAge"`
	BirthRate             int					`json:"birthRate"`
	DeathRate      				int					`json:"deathRate"`
	Population            int					`json:"population"`
	TaxRate               int         `json:"taxRate"`
	ConscriptAge    			int         `json:"conscriptAge"`
}
