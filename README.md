# jrpg-gang
Golang project of turn based multiplayer RPG game server

## Unit base attributes:
* `health`  - the hit points unit can take, till die
* `stamina` - a weapon may require stamina points to perform action
* `mana`    - a weapon or spell may require mana points to perform action
* `curse`   - accumulative, reduces the chance to perform action
* `fear`    - accumulative, increases the chance to retreat

## Unit attributes:
* `strength`     - enhances stabbing, cutting, crushing, exhaustion and bleeding damage
* `physique`     - affects stun chance
* `agility`      - affects attack chance
* `endurance`    - stamina recovery
* `intelligence` - enhances fire, cold, lighting, manaDrain, fear and curse damage
* `initiative`   - affects turn order
* `luck`         - affects critical chance

## Damage, Resistance, Modification:
* `stabbing`   - affects health attribute
* `cutting`    - affects health attribute
* `crushing`   - affects health attribute
* `fire`       - affects health attribute
* `cold`       - affects health attribute
* `lighting`   - affects health attribute
* `poison`     - affects health attribute
* `bleeding`   - affects health attribute
* `exhaustion` - affects stamina attribute
* `manaDrain`  - affects mana attribute
* `fear`       - affects fear attribute
* `curse`      - affects curse attribute

## Math:
* `Attack chance`: (`unit luck` - `unit curse`) - (`target luck` - `target curse`) | minimum `1`
* `Critical attack chance`: (`unit agility` - `unit curse`) - (`target agility` - `target curse`) + `base chance` | minimum `1`
* `Modification chance`: (`unit intelligence` - `unit curse`) + `base chance` | minimum `1`
* `Stun Chance`: (`physical damage` - `target curse`) - (`unit physique` - `unit curse`) | minimum `1`
* `Retreat Chance`: `unit fear` | minimum `0`

## Mechanics:
* `Critical Chance` - Doubles the damage.
* `Stun` - When stunned, a unit loses its turn in the current round and appears at the end of turn queue in the next round. If a unit is hit while stunned, critical damage is dealt and the stun is reset.
* `Fear` - If attribute is more than zero, before each unit turn the `Retreat` check is performed. Unit may miss its turn by moving to a corner of the battlefield.
* `Equipment Wearout` - Each success action followed by damage increases the `Wearout` of unit weapon, as well as `Wearout` of target equipment. If `Wearout` of an item reaches its `Durability`, the item can no longer be used and becomes unequipped.
