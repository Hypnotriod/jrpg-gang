# jrpg-gang
Golang project of turn based multiplayer RPG game server

## BaseAttributes:
* `health`        - the **hit points** unit can take, till die
* `stamina`       - a weapon may require **stamina points** to perform action
* `mana`          - a weapon or spell may require **mana points** to perform action
* `action points` - a weapon, spell or disposable require **action points** to perform action

## Attributes:
* `strength`     - enhances **stabbing**, **cutting**, **crushing** and **bleeding** damage
* `physique`     - affects **stun chance**, for each 10 points adds 1 point to all **physical resistance**
* `agility`      - affects **attack / dodge chance**
* `endurance`    - **stamina** recovery
* `intelligence` - enhances **fire**, **cold**, **lightning**, **exhaustion**, **manaDrain**, **fear**, **curse** and **madness** damage, multiplies by 1% all **modification** points
* `initiative`   - affects **turn order**, for each 10 points adds **1 action point**
* `luck`         - affects **critical chance**

## State:
* `health`        - current **hit points** unit have
* `stamina`       - current **stamina point** unit have
* `mana`          - current **mana points** unit have
* `action points` - current **action points** for the current turn unit have
* `stress`        - accumulative, reduces the **chance to perform action**, increases the **chance to retreat** and the **critical miss chance**

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

## Items:
* `Weapon`     - is an `equipment` and represents any **weapon** unit can equip and use against the enemy. **Weapon** may require special `ammunition`. Deal `DamageImpact`.
* `Ammunition` - is an `equipment` and represents the **ammunition** required by the **weapon**. Should be equipped before the **weapon** usage. Can have a quantity.
* `Armor`      - is an `equipment` and represents any **armor** unit can equip.
* `Magic`      - is a *knowledge* that unit can use as a `weapon` against the enemy or as a `modification` to modify/recover stats of friendly units or himself. Deal `DamageImpact` or `ModificationImpact`.
* `Disposable` - is any sort of "single-use" item. Can have a quantity. Deal `DamageImpact` or `ModificationImpact`.

## Math:
* `Physical Damage / Resistance`: **Stabbing** + **Cutting** + **Crushing** + **Fire** + **Cold** + **lightning**
* `Attack chance`: (**unit agility** - **unit stress**) - (**target agility** - **target stress**) + **base chance** | minimum `1`
* `Attack chance` when `Stunned`: (**unit agility** - **unit stress**) + **target stress** + **base chance** | minimum `1`
* `Critical attack chance`: (**unit luck** - **unit stress**) - (**target luck** - **target stress**) | minimum `1`
* `Modification chance`: (**unit intelligence** - **unit stress**) + **base chance** | minimum `1`
* `Stun Chance`: (**physical damage** - **unit stress**) - (**target physique** - **target stress**) | minimum `1`
* `Retreat Chance`: **unit stress** | minimum `0`
* `Critical Miss Chance`: **unit stress** | minimum `0`

## Game Phases
* `Prepare Unit Phase` - Unit can be placed on any available position before battle starts. Only `Equip` or `Unequip` actions can be performed during this phase.
* `Move or Action Phase` - During this phase `Move`, `Use`, `Equip`, `Unequip`, `Wait` or `Skip` actions can be performed while **actions points** remains more or equal to **4**
  * `Use` action - Can be performed with `weapon`, `magic` and `disposable` items. Reduces required **actions points**.
  * `Move` action - Requires **4 action points**
  * `Equip` or `Unequip` action - Can be performed with `equipment` any time and requires no ation points.
  * `Wait` action - moves unit to the back of the queue and keeps all unused action points.

## Mechanics:
* `Critical Damage` - Doubles the damage.
* `Critical Miss` - If a unit's attack misses and the unit's `Stress` attribute is more than zero then `Critical Miss` check is performed. A unit can damage itself with double damage.
* `Exhaustion` - Any `Instant Physical Damage` absorbed by unit's equipment accumulates to the `exhaustion` damage. But **cannot** be enhanced with opponent's **intelligence** attribute. If a unit is hit while **totally exhausted** (stamina is zero), **critical damage** is dealt.
* `Stun` - Any `Instant Physical Damage` can cause stun. When stunned, a unit loses its turn in the current round and appears at the end of turn queue in the next round. If a unit is hit while stunned, **critical damage** is dealt and the **stun** is reset.
* `Retreat` - If `Stress` attribute is more than zero, before each unit turn the `Retreat` check is performed. A unit can miss its turn by moving to a corner of the battlefield.
* `Equipment Wearout` - Each success action followed by damage, or each action that requires ammunition, increases the `Wearout` of unit weapon, as well as `Wearout` of target equipment. If `Wearout` of an item reaches its `Durability`, the item can no longer be used and becomes unequipped.
