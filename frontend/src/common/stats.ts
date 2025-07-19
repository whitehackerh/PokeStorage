import { IndividualValues } from "../entity/IndividualValues";

export const stats: (keyof Omit<IndividualValues, 'id'>)[] = [
    'hitPoints',
    'attack',
    'defense',
    'specialAttack',
    'specialDefense',
    'speed',
];
export const statDisplayNames: Record<keyof Omit<IndividualValues, 'id'>, string> = {
    hitPoints: 'Hit Points',
    attack: 'Attack',
    defense: 'Defense',
    specialAttack: 'Special Attack',
    specialDefense: 'Special Defense',
    speed: 'Speed',
};