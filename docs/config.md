# config

defaults configurations should be persisted.  there are several types of configs:

config type | key | description
---|---|---
global | configs:global | prefixed with limit_ means maximum
keys | configs:keys | key namespaces
init | configs:init | these are limitations imposed for character creation
kingdom | configs:kingdom:[NAME] | these are per kingdom and should supercede `configs:init`

## global

`configs:global`

attribute | value | description
---|---|---
hash:length | 20 |
female | 0 |
male | 1 |
gender | 2 |
skill:max | 5000 |


## init

`configs:init`

attribute | value | description
---|---|---
age:min | 17 |
age:max | 26 |
area:min | 1 |
area:max | 200 |
population:tiny | 500 |
population:small | 2000 |
population:medium | 35000 |
population:large | 75000 |
population:xlarge | 200000 |
population:massive | 500000 |
population:metropolis | 1000000 |
skill:max | 200 |
wealth:max | 200 |


## kingdom

`configs:kingdom:[NAME]`
