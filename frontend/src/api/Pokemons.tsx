import { withTokenRequest, getRequestHeaders } from '../http';
import { Pokemon } from '../entity/Pokemon';
import { Type, NewTypeFromJson } from '../entity/Type';
import { Ability, NewAbilityFromJson } from '../entity/Ability';
import { NewBaseStatsFromJson } from '../entity/BaseStats';
import { Item } from '../entity/Item';

export const getPokemons = async (titleId: number): Promise<Pokemon[]> => {
    try {
        const response = await withTokenRequest.get(`/pokemons/${titleId}`, {
            headers: getRequestHeaders()
        });
        const pokemons: Pokemon[] = response.data.data.pokemons.map((item: any) => {
            const types: Type[] = item.types.map((type: any) => {
                return NewTypeFromJson(type);
            });
            const abilities: Ability[] = item.abilities.map((ability: any) => {
                return NewAbilityFromJson(ability);
            });
            const presetHeldItem: Item | null = item.preset_held_item
                ? {
                    id: item.preset_held_item.id,
                    name: item.preset_held_item.name,
                }
                : null;
            return {
                id: item.id,
                nationalPokedexNo: item.national_pokedex_no,
                formeNo: item.forme_no,
                name: item.name,
                formeName: item.forme_name,
                types: types,
                abilities: abilities,
                baseStats: NewBaseStatsFromJson(item.base_stats),
                presetHeldItem: presetHeldItem,
            }
        });
        return pokemons;
    } catch (error) {
        console.log('error');
        throw error;
    }
}