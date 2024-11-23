import { TitleEnum } from '../enum/Title';
import { SVBredPokemon } from '../entity/SVBredPokemon';
import { NewGenderFromJson } from '../entity/Gender';
import { NewTypeFromJson, Type } from '../entity/Type';
import { NewAbilityFromJson } from '../entity/Ability';
import { NewNatureFromJson } from '../entity/Nature';
import { Item } from '../entity/Item';
import { Move, NewMoveFromJson } from '../entity/Move';
import { withTokenRequest, getRequestHeaders } from '../http';
import { NewBaseStatsFromJson } from '../entity/BaseStats';
import { NewIndividualValuesFromJson } from '../entity/IndividualValues';
import { NewBasePointsFromJson } from '../entity/BasePoints';
import { NewActualValuesFromJson } from '../entity/ActualValues';
import { NewTeraTypeFromJson } from '../entity/TeraType';

export const postBredPokemons = async (bredPokemon: any) : Promise<void> => {
    try {
        await withTokenRequest.post('/sv/bred-pokemons', bredPokemon, {
            headers: getRequestHeaders()
        });
    } catch (error) {
        console.log('error');
        throw error;
    }
}

export const getBredPokemons = async () : Promise<any> => {
    try {
        const response = await withTokenRequest.get('/sv/bred-pokemons', {
            headers: getRequestHeaders()
        });
        const bredPokemons: SVBredPokemon[] = response.data.data.bred_pokemons.map((item: any) => {
            const types: Type[] = item.types.map((type: any) => {
                return NewTypeFromJson(type);
            });
            const heldItem: Item | null = item.held_item
            ? {
                id: item.held_item.id,
                name: item.held_item.name,
            }
            : null;
            const moves: Move[] = item.moves.map((move: any) => {
                return NewMoveFromJson(move);
            })

            return {
                bredPokemon: {
                    id: item.id,
                    pokemonId: item.pokemon_id,
                    nationalPokedexNo: item.national_pokedex_no,
                    formeNo: item.forme_no,
                    name: item.name,
                    formeName: item.forme_name,
                    gender: NewGenderFromJson(item.gender),
                    level: item.level,
                    types: types,
                    ability: NewAbilityFromJson(item.ability),
                    nature: NewNatureFromJson(item.nature),
                    heldItem: heldItem,
                    baseStats: NewBaseStatsFromJson(item.base_stats),
                    individualValues: NewIndividualValuesFromJson(item.individual_values),
                    basePoints: NewBasePointsFromJson(item.base_points),
                    actualValues: NewActualValuesFromJson(item.actual_values),
                    moves: moves,
                    note: item.note,
                },
                teraType: NewTeraTypeFromJson(item.tera_type),
            }
        });
        return bredPokemons;
    } catch (error) {
        console.log('error');
        throw error;
    }
}