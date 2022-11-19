import * as readline from 'node:readline/promises';
import * as fs from 'node:fs';
import { stdin, stdout } from 'node:process';

const rl = readline.createInterface(stdin, stdout);
stdin.isTTY && stdin.setRawMode(true);

async function yesNo(prompt) {
    const result = await rl.question(`* add ${prompt}: y/n? `) === 'y';
    stdout.moveCursor(0, -1);
    stdout.clearLine();
    return result;
}

async function makeList(value, values) {
    return await rl.question(`-> ${value}: [${values.join(', ')}]: `) || values[0];
}

async function makeString(value, defaultValue = '') {
    const result = defaultValue ?
        await rl.question(`-> ${value}: (default: ${defaultValue}) `) :
        await rl.question(`-> ${value}: `);
    return result || defaultValue || undefined;
}

async function makeNumber(value, defaultValue = 0) {
    const result = await rl.question(`-> ${value}: (default: 0) `);
    return result || defaultValue || undefined;
}

async function run() {
    while (true) {
        const item = await createItem();
        if (!item) { continue; }
        const resultJson = JSON.stringify(item, null, 2);
        stdout.write(`\r\n************* ${item.name} *************\r\n\r\n${resultJson}\r\n\r\n`);
        storeNewItem(item);
    }
}

function storeNewItem(item) {
    const configPath = 'private/items_config.json';
    let itemsConfigStr = fs.readFileSync(configPath);
    const itemsConfig = JSON.parse(itemsConfigStr);
    switch (item.type) {
        case 'armor': itemsConfig.armor.push(item); break;
        case 'weapon': itemsConfig.weapon.push(item); break;
        case 'magic': itemsConfig.magic.push(item); break;
        case 'disposable': itemsConfig.disposable.push(item); break;
        case 'ammunition': itemsConfig.ammunition.push(item); break;
        default: return;
    }
    itemsConfigStr = JSON.stringify(itemsConfig, null, 4);
    fs.writeFileSync(configPath, itemsConfigStr);
}

async function createItem() {
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
            return await makeArmor();
            break;
        case 2:
            return await makeWeapon();
            break;
        case 3:
            return await makeMagic();
            break;
        case 4:
            return await makeDisposable();
            break;
        case 5:
            return await makeAmmunition();
            break;
        default:
            return null;
    }
}

async function makeArmor() {
    stdout.write('armor:\r\n');
    const result = {
        ...await makeItem('armor'),
        ...await makeEquipment(),
    }
    return result;
}

async function makeWeapon() {
    stdout.write('weapon:\r\n');
    const result = {
        ...await makeItem('weapon'),
        ammunitionKind: await makeString('ammunitionKind'),
        range: await makeActionRange('range'),
        useCost: await makeUnitBaseAttributes('useCost'),
        ...await makeEquipment(),
        damage: [],
    }
    while (await yesNo('damage')) {
        result.damage.push(await makeDamageImpact('damage'));
    }
    return result;
}

async function makeMagic() {
    stdout.write('magic:\r\n');
    const result = {
        ...await makeItem('magic'),
        requirements: await makeUnitAttributes('requirements'),
        range: await makeActionRange('range'),
        useCost: await makeUnitBaseAttributes('useCost'),
        damage: [],
        modification: [],
    }
    while (await yesNo('damage')) {
        result.damage.push(await makeDamageImpact('damage'));
    }
    while (await yesNo('modification')) {
        result.modification.push(await makeUnitModificationImpact('modification'));
    }
    return result;
}

async function makeDisposable() {
    stdout.write('disposable:\r\n');
    const result = {
        ...await makeItem('disposable'),
        range: await makeActionRange('range'),
        damage: [],
        modification: [],
    }
    while (await yesNo('damage')) {
        result.damage.push(await makeDamageImpact('damage'));
    }
    while (await yesNo('modification')) {
        result.modification.push(await makeUnitModificationImpact('modification'));
    }
    return result;
}

async function makeAmmunition() {
    stdout.write('ammunition:\r\n');
    const result = {
        ...await makeItem('ammunition'),
        kind: await makeString('kind'),
        range: await makeActionRange('range'),
        damage: [],
    }
    while (await yesNo('damage')) {
        result.damage.push(await makeDamageImpact('damage'));
    }
    return result;
}

async function makeActionRange(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        minimumX: await makeNumber('minimumX'),
        maximumX: await makeNumber('maximumX'),
        minimumY: await makeNumber('minimumY'),
        maximumY: await makeNumber('maximumY'),
        radius: await makeNumber('radius'),
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
        durability: await makeNumber('durability', 300),
        slot: await makeList('slot', ['weapon', 'head', 'neck', 'body', 'hand', 'leg']),
        slotsNumber: await makeNumber('slotsNumber', 1),
        requirements: await makeUnitAttributes('requirements'),
        modification: [],
    }
    while (await yesNo('modification')) {
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

async function makeImpact(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        duration: await makeNumber('duration'),
        chance: await makeNumber('chance'),
    }
    return result;
}

async function makeDamageImpact(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        ...await makeImpact(),
        ...await makeDamage(),
    }
    return result;
}

async function makeUnitModificationImpact(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        ...await makeImpact(),
        ...await makeUnitModification(),
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

async function makeUnitModification(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {};
    if (await yesNo('baseAttributes')) {
        result.baseAttributes = await makeUnitBaseAttributes('baseAttributes');
    }
    if (await yesNo('attributes')) {
        result.attributes = await makeUnitAttributes('attributes');
    }
    if (await yesNo('resistance')) {
        result.resistance = await makeDamage('resistance');
    }
    if (await yesNo('damage')) {
        result.damage = await makeDamage('damage');
    }
    if (await yesNo('recovery')) {
        result.recovery = await makeUnitRecovery('recovery');
    }
    return result;
}

run();
