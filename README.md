# jrpg-gang
Golang project of turn based multiplayer RPG game server

## BaseAttributes:
* `health`  - the **hit points** unit can take, till die
* `stamina` - a weapon may require **stamina points** to perform action
* `mana`    - a weapon or spell may require **mana points** to perform action

## Attributes:
* `strength`     - enhances **stabbing**, **cutting**, **crushing** and **bleeding** damage
* `physique`     - affects **stun chance**, for each 10 points adds 1 point to all **physical resistance**
* `agility`      - affects **attack / dodge chance**
* `endurance`    - **stamina** recovery
* `intelligence` - enhances **fire**, **cold**, **lightning**, **exhaustion**, **manaDrain**, **fear**, **curse** and **madness** damage, multiplies by 1% all **modification** points
* `initiative`   - affects **turn order**
* `luck`         - affects **critical chance**

## State:
* `health`  - current **hit points** unit have
* `stamina` - current **stamina point** unit have
* `mana`    - current **mana points** unit have
* `stress`  - accumulative, reduces the **chance to perform action**, increases the **chance to retreat** and the **critical miss chance**

## Damage, Resistance:
* `stabbing`   - affects **health** attribute
* `cutting`    - affects **health** attribute
* `crushing`   - affects **health** attribute
* `fire`       - affects **health** attribute
* `cold`       - affects **health** attribute
* `lightning`   - affects **health** attribute
* `poison`     - affects **health** attribute
* `bleeding`   - affects **health** attribute
* `exhaustion` - affects **stamina** attribute
* `manaDrain`  - affects **mana** attribute
* `fear`       - affects **stress** attribute
* `curse`      - affects **stress** attribute
* `madness`    - affects **stress** attribute

## Recovery:
* **State**        - recovers state parameters
* **Damage**       - reduces accumulated damage impact 

## Modification:
* **BaseAttributes** - modifies unit base attributes
* **Attributes**     - modifies unit attributes
* **Resistance**     - modifies unit resistance
* **Damage**         - modifies the damage unit can apply
* **Recovery**       - recovers unit state

## Impact:
* `duration`   - **duration** of effect, immediate if zero
* `chance`     - base **chance**

## DamageImpact:
* **Impact**
* **Damage**

## ModificationImpact:
* Impact
* Modification

## Math:
* `Physical Damage / Resistance`: **Stabbing** + **Cutting** + **Crushing** + **Fire** + **Cold** + **lightning**
* `Attack chance`: (**unit agility** - **unit stress**) - (**target agility** - **target stress**) + **base chance** | minimum `1`
* `Attack chance` when `Stunned`: (**unit agility** - **unit stress**) + **target stress** + **base chance** | minimum `1`
* `Critical attack chance`: (**unit luck** - **unit stress**) - (**target luck** - **target stress**) | minimum `1`
* `Modification chance`: (**unit intelligence** - **unit stress**) + **base chance** | minimum `1`
* `Stun Chance`: (**physical damage** - **unit stress**) - (**target physique** - **target stress**) | minimum `1`
* `Retreat Chance`: **unit stress** | minimum `0`
* `Critical Miss Chance`: **unit stress** | minimum `0`

## Mechanics:
* `Critical Damage` - Doubles the damage.
* `Critical Miss` - If a unit's attack misses and the unit's `Stress` attribute is more than zero then `Critical Miss` check is performed. A unit can damage itself with double damage.
* `Exhaustion` - Any `Instant Physical Damage` absorbed by unit's equipment accumulates to the `exhaustion` damage. But **cannot** be enhanced with opponent's **intelligence** attribute. If a unit is hit while **totally exhausted** (exhaustion is zero), **critical damage** is dealt.
* `Stun` - Any `Instant Physical Damage` can cause stun. When stunned, a unit loses its turn in the current round and appears at the end of turn queue in the next round. If a unit is hit while stunned, **critical damage** is dealt and the **stun** is reset.
* `Retreat` - If `Stress` attribute is more than zero, before each unit turn the `Retreat` check is performed. A unit can miss its turn by moving to a corner of the battlefield.
* `Equipment Wearout` - Each success action followed by damage increases the `Wearout` of unit weapon, as well as `Wearout` of target equipment. If `Wearout` of an item reaches its `Durability`, the item can no longer be used and becomes unequipped.
