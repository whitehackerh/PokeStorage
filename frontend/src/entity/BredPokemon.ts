import { Gender } from "./Gender";
import { Type } from "./Type";
import { Ability } from "./Ability";
import { Nature } from "./Nature";
import { Item } from "./Item";
import { BaseStats } from "./BaseStats";
import { Move } from './Move'; 
import { IndividualValues } from "./IndividualValues";
import { BasePoints } from "./BasePoints";
import { ActualValues } from "./ActualValues";

export interface BredPokemon {
    id: string;
    pokemonId: number;
    nationalPokedexNo: number;
    formeNo: number;
    name: string;
    formeName: string | null;
    gender: Gender;
    level: number;
    types: Type[];
    ability: Ability;
    nature: Nature;
    heldItem: Item | null;
    baseStats: BaseStats;
    individualValues: IndividualValues;
    basePoints: BasePoints;
    actualValues: ActualValues;
    moves: Move[];
    note: string | null;
}