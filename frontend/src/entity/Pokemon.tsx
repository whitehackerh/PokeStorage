import { Type } from './Type';
import { Ability } from './Ability';
import { BaseStats } from './BaseStats';
import { Item } from './Item';

export interface Pokemon {
    id: number;
    nationalPokedexNo: number;
    formeNo: number;
    name: string;
    formeName: string | null;
    types: Type[];
    abilities: Ability[];
    baseStats: BaseStats;
    presetHeldItem: Item | null;
}