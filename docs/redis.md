# redis

## kingdoms (hash map)

kingdoms:name:values

qin, qi, chu, zhao, yan, wei, han, zhou, none, barbarians, outlaws

## families (hash map)

families:name:values

bai, cao, chen, fan, gao, gu, guo, he, hong, huang, hu, jiang, kong, li, liu, mao, pan, ren, sun, wang, xi, xu, yang, yu, yuan, zhang, zhao, zhou, zhu


## users (hash map)

users:name:fields

field | type | description
--- | --- | ---
name | string | unique
email | string | unique, encrypted
password | string | salted and hashed
mobile | string | unique, encrypted

### attriburtes

field | type | description
--- | --- | ---
gender | int | 0 - F, 1 - M
age | int | initial range 18-21
height | int | (cm)
family | string | family surname
