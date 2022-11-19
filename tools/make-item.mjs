import * as readline from 'node:readline/promises';
import { stdin, stdout } from 'node:process';

const rl = readline.createInterface(stdin, stdout);
stdin.isTTY && stdin.setRawMode(true);

async function makeList(value, values) {
    return rl.question(`${value}: ${values}: `);
}

async function makeString(value, defaultValue = '') {
    const result = defaultValue ?
        await rl.question(`${value}: (default: ${defaultValue}) `) :
        await rl.question(`${value}: `);
    return result || defaultValue;
}

async function makeNumber(value, defaultValue = 0) {
    const result = await rl.question(`${value}: number: (default: ${defaultValue}) `);
    return result || defaultValue;
}

async function chooseItemType() {
    const r = await rl.question(
        `Choose type:
        1 - armor
        2 - weapon
        3 - magic
        4 - disposable
        5 - ammunition\r\n`);
    const i = Number.parseInt(r);
    switch (i) {
        case 1:
            console.log(await makeArmor());
            break;
    }
}

async function makeArmor() {
    const result = {
        ...await makeItem('armor'),
        ...await makeEquipment(),
    }
    return result;
}

async function makeUnitBooty(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        coins: await makeNumber('coins'),
        ruby: await makeNumber('ruby'),
    }
    return result;
}

async function makeItem(type, header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        name: await makeString('name'),
        description: await makeString('description'),
        code: await makeString('code'),
        type: type,
        price: await makeUnitBooty('price'),
    }
    return result;
}

async function makeEquipment(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        wearout: 0,
        durability: await makeNumber('durability', 300),
        slot: await makeList('slot', ['weapon', 'head', 'neck', 'body', 'hand', 'leg']),
        slotsNumber: await makeNumber('slotsNumber', 1),
        equipped: false,
        requirements: await makeUnitAttributes('requirements'),
    }
    while (await rl.question('add modification: y/n?') === 'y') {
        result.modification = result.modification || [];
        result.modification.push(await makeUnitModification('modification'));
    }
    return result;
}

async function makeUnitAttributes(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        strength: await makeNumber('strength'),
        physique: await makeNumber('physique'),
        agility: await makeNumber('agility'),
        endurance: await makeNumber('endurance'),
        intelligence: await makeNumber('intelligence'),
        initiative: await makeNumber('initiative'),
        luck: await makeNumber('luck'),
    }
    return result;
}

async function makeUnitBaseAttributes(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        health: await makeNumber('health'),
        stamina: await makeNumber('stamina'),
        mana: await makeNumber('mana'),
    }
    return result;
}

async function makeDamage(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        stabbing: await makeNumber('stabbing'),
        cutting: await makeNumber('cutting'),
        crushing: await makeNumber('crushing'),
        fire: await makeNumber('fire'),
        cold: await makeNumber('cold'),
        lighting: await makeNumber('lighting'),
        poison: await makeNumber('poison'),
        exhaustion: await makeNumber('exhaustion'),
        manaDrain: await makeNumber('manaDrain'),
        bleeding: await makeNumber('bleeding'),
        morale: await makeNumber('morale'),
        fortune: await makeNumber('fortune'),
        isCritical: await makeNumber('isCritical'),
        withStun: await makeNumber('withStun'),
    }
    return result;
}

async function makeUnitState(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        health: await makeNumber('health'),
        stamina: await makeNumber('stamina'),
        mana: await makeNumber('mana'),
        fear: await makeNumber('fear'),
        curse: await makeNumber('curse'),
    }
    return result;
}

async function makeUnitRecovery(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        ...await makeUnitState(),
        ...await makeDamage(),
    }
    return result;
}

async function makeUnitModification() {
    const result = {};
    if (await rl.question('add baseAttributes: y/n?') === 'y') {
        result.baseAttributes = await makeUnitBaseAttributes('baseAttributes');
    }
    if (await rl.question('add attributes: y/n?') === 'y') {
        result.attributes = await makeUnitAttributes('attributes');
    }
    if (await rl.question('add resistance: y/n?') === 'y') {
        result.resistance = await makeDamage('resistance');
    }
    if (await rl.question('add damage: y/n?') === 'y') {
        result.damage = await makeDamage('damage');
    }
    if (await rl.question('add recovery: y/n?') === 'y') {
        result.recovery = await makeUnitRecovery('recovery');
    }
    return result;
}

chooseItemType();
