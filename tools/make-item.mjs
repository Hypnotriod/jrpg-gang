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
    return await rl.question(`-> ${value}: [${values.join(', ')}]: (default: ${values[0]}) `) || values[0];
}

async function makeValue(value, defaultValue) {
    const result = defaultValue ?
        await rl.question(`-> ${value}: (default: ${defaultValue}) `) :
        await rl.question(`-> ${value}: `);
    return result || defaultValue;
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
    const configPath = 'items_config.json';
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
        case 2:
            return await makeWeapon();
        case 3:
            return await makeMagic();
        case 4:
            return await makeDisposable();
        case 5:
            return await makeAmmunition();
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
        ammunitionKind: await makeValue('ammunitionKind'),
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
        kind: await makeValue('kind'),
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
        minimumX: await makeValue('minimumX'),
        maximumX: await makeValue('maximumX'),
        minimumY: await makeValue('minimumY'),
        maximumY: await makeValue('maximumY'),
        radius: await makeValue('radius'),
    }
    return result;
}

async function makeUnitBooty(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        coins: await makeValue('coins'),
        ruby: await makeValue('ruby'),
    }
    return result;
}

async function makeItem(type, header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        name: await makeValue('name'),
        description: await makeValue('description'),
        code: await makeValue('code'),
        type: type,
        price: await makeUnitBooty('price'),
    }
    return result;
}

async function makeEquipment(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        durability: await makeValue('durability', 300),
        slot: await makeList('slot', ['weapon', 'head', 'neck', 'body', 'hand', 'leg']),
        slotsNumber: await makeValue('slotsNumber', 1),
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
        strength: await makeValue('strength'),
        physique: await makeValue('physique'),
        agility: await makeValue('agility'),
        endurance: await makeValue('endurance'),
        intelligence: await makeValue('intelligence'),
        initiative: await makeValue('initiative'),
        luck: await makeValue('luck'),
    }
    return result;
}

async function makeUnitBaseAttributes(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        health: await makeValue('health'),
        stamina: await makeValue('stamina'),
        mana: await makeValue('mana'),
    }
    return result;
}

async function makeImpact(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        duration: await makeValue('duration'),
        chance: await makeValue('chance'),
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
        stabbing: await makeValue('stabbing'),
        cutting: await makeValue('cutting'),
        crushing: await makeValue('crushing'),
        fire: await makeValue('fire'),
        cold: await makeValue('cold'),
        lighting: await makeValue('lighting'),
        poison: await makeValue('poison'),
        exhaustion: await makeValue('exhaustion'),
        manaDrain: await makeValue('manaDrain'),
        bleeding: await makeValue('bleeding'),
        morale: await makeValue('morale'),
        fortune: await makeValue('fortune'),
    }
    return result;
}

async function makeUnitState(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        health: await makeValue('health'),
        stamina: await makeValue('stamina'),
        mana: await makeValue('mana'),
        fear: await makeValue('fear'),
        curse: await makeValue('curse'),
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
