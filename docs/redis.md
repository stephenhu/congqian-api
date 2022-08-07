# redis

## kingdoms

qin, qi, chu, zhao, yan, wei, han, zhou, barbarians, outlaws

structure | description
---|---
**type** | set
**key** | kingdoms

### kingdom

structure | description
---|---
**type** | hash map
**key** | kingdom:[NAME]

field | type | description
---|---|---
land | int |
population | int |
capital | string |
wealth | int |
trees | int |
rocks | int |
cows | int |
taxRate | int |

#### kingdom municipality

structure | description
---|---
**type** | hash map
**key** | kingdom:[NAME]:municipalities

field | type | description
---|---|---
magistrate | |
deputy magistrate | |
magistrate protector | |
guard magistrate | |

land | int |
population | int |
wealth | int |
trees | int |
rocks | int |
cows | int |


## families

bai, cai, cao, chen, fan, gao, gu, guo, he, hong, huang, hu, jiang, kong, li, liu, mao, pan, ren, sun, wang, xi, xu, yang, yu, yuan, zhang, zhao, zhou, zhu

structure | description
---|---
**type** | set
**key** | families

### family

structure | description
---|---
**type** | hash map
**key** | family:[NAME]

field | type | description
---|---|---
total | int |


## users

structure | description
---|---
**type** | set
**key** | users

### user

structure | description
---|---
**type** | hash map
**key** | user:[NAME]

field | type | description
---|---|---
name | string | unique
email | string | unique, encrypted
password | string | salted and hashed
mobile | string | unique, encrypted

### attributes

field | type | description
--- | --- | ---
gender | int | 0 - F, 1 - M
age | int | initial range 18-21
height | int | (cm)
family | string | family surname
