# jrpg-gang
Golang project for a turn-based multiplayer RPG game server

[hypnotriod.org/jrpg-gang](https://hypnotriod.org/jrpg-gang)

## BaseAttributes:
* `health`        - the **hit points** a unit can take before dying
* `stamina`       - a weapon may require **stamina points** to perform an action
* `mana`          - a weapon or spell may require **mana points** to perform an action
* `action points` - a weapon, spell, or disposable item may require **action points** to perform an action

## Attributes:
* `strength`     - enhances **stabbing**, **cutting**, **crushing**, and **bleeding** damage
* `physique`     - affects **stun chance**. For every 10 points, adds 1 point to all **physical resistances**
* `agility`      - affects **attack/dodge chance**
* `endurance`    - affects **stamina** recovery
* `intelligence` - enhances **fire**, **cold**, **lightning**, **exhaustion**, **manaDrain**, **fear**, **curse**, and **madness** damage. Multiplies by 1% all the **modification** points
* `initiative`   - affects **turn order**. For every 10 points, adds **1 action point**
* `luck`         - affects **critical chance**

## State:
* `health`        - current **hit points** a unit has
* `stamina`       - current **stamina points** a unit has
* `mana`          - current **mana points** a unit has
* `action points` - current **action points** a unit has for the current turn
* `stress`        - accumulative, reduces the **chance to perform an action**, increases the **chance to retreat**, and the **critical miss chance**

## Damage, Resistance:
* `stabbing`   - affects the **health** attribute
* `cutting`    - affects the **health** attribute
* `crushing`   - affects the **health** attribute
* `fire`       - affects the **health** attribute
* `cold`       - affects the **health** attribute
* `lightning`  - affects the **health** attribute
* `poison`     - affects the **health** attribute
* `bleeding`   - affects the **health** attribute
* `exhaustion` - affects the **stamina** attribute
* `manaDrain`  - affects the **mana** attribute
* `fear`       - affects the **stress** attribute
* `curse`      - affects the **stress** attribute
* `madness`    - affects the **stress** attribute

## Recovery:
* **State**        - recovers state parameters
* **Damage**       - reduces accumulated damage impact 

## Modification:
* **BaseAttributes** - modifies unit base attributes
* **Attributes**     - modifies unit attributes
* **Resistance**     - modifies unit resistance
* **Damage**         - modifies the damage a unit can apply
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
* `Weapon`     - an `equipment` item representing any **weapon** a unit can equip and use against the enemy. **Weapons** may require special `ammunition`. Deals `DamageImpact`. Can only be used in the `Combat Phase`.
* `Ammunition` - an `equipment` item representing the **ammunition** required by a **weapon**. Should be equipped before using the **weapon**. Can have a quantity.
* `Armor`      - an `equipment` item representing any **armor** a unit can equip.
* `Magic`      - a *knowledge* that a unit can use as a `weapon` against the enemy or as a `modification` to modify/recover stats of friendly units or itself. Deals `DamageImpact` or `ModificationImpact`. Can only be used in the `Combat Phase`.
* `Disposable` - any sort of "single-use" item. Can have a quantity. Deals `DamageImpact` or `ModificationImpact`. Can only be used in the `Combat Phase`.
* `Provision`  - a *food* item. Can have a quantity. Deals `Recovery`. Can only be used in the `Resting Phase`.

## Math:
* `Physical Damage / Resistance`: **Stabbing** + **Cutting** + **Crushing** + **Fire** + **Cold** + **Lightning**
* `Attack chance`: (**unit agility** - **unit stress**) - (**target agility** - **target stress**) + **base chance** | minimum `1`
* `Attack chance` when `Stunned`: (**unit agility** - **unit stress**) + **target stress** + **base chance** | minimum `1`
* `Critical attack chance`: (**unit luck** - **unit stress**) - (**target luck** - **target stress**) | minimum `1`
* `Modification chance`: (**unit intelligence** - **unit stress**) + **base chance** | minimum `1`
* `Stun Chance`: (**physical damage** - **unit stress**) - (**target physique** - **target stress**) | minimum `1`
* `Retreat Chance`: **unit stress** | minimum `0`
* `Critical Miss Chance`: **unit stress** | minimum `0`

## Game Phases
* `Prepare Unit Phase` *(prepareUnit)* - Unit can be placed in any available position before the battle starts. Only `Equip` and `Unequip` actions can be performed during this phase.
* `Combat Phase` *(takeAction)*- During this phase, `Move`, `Use`, `Equip`, `Unequip`, `Wait`, or `Skip` actions can be performed while **action points** remain greater than or equal to **4**
  * `Use` action - Can be performed with `weapon`, `magic`, and `disposable` items. Reduces required **action points**.
  * `Move` action - Requires **4 action points**
  * `Equip` or `Unequip` action - Can be performed with `equipment` at any time and requires no action points.
  * `Wait` action - Moves the unit to the back of the queue and retains all unused action points.
* `Resting Phase` *(spotComplete)* - During this phase, the `Use`, `Equip`, and `Unequip` actions can be performed. Players can also leave the session with a `Booty Share`.

## Mechanics:
* `Critical Damage` - Doubles the damage.
* `Critical Miss` - If a unit's attack misses and its `Stress` attribute is greater than zero, a `Critical Miss` check is performed. A unit may damage itself with double damage.
* `Exhaustion` - Any `Instant Physical Damage` absorbed by a unit's equipment accumulates as `exhaustion` damage. However, it **cannot** be enhanced by the opponent's **intelligence** attribute. If a unit is hit while **totally exhausted** (stamina is zero), **critical damage** is dealt.
* `Stun` - Any `Instant Physical Damage` can cause stun. When stunned, a unit loses its turn in the current round and appears at the end of the turn queue in the next round. If a unit is hit while stunned, **critical damage** is dealt and the **stun** is reset.
* `Retreat` - If the `Stress` attribute is greater than zero, before each unit's turn the `Retreat` check is performed. A unit can miss its turn by moving to a corner of the battlefield.
* `Equipment Wearout` - Each successful action followed by damage, or each action that requires ammunition, increases the `Wearout` of the unit's weapon, as well as the `Wearout` of the target's equipment. If the `Wearout` of an item reaches its `Durability`, the item can no longer be used and becomes unequipped.
