import { withTokenRequest, getRequestHeaders } from '../http';
import { Pokemon } from '../entity/Pokemon';
import { Type } from '../entity/Type';
import { Ability } from '../entity/Ability';
import { BaseStats } from '../entity/BaseStats';
import { Item } from '../entity/Item';

export const getPokemons = async (titleId: number): Promise<Pokemon[]> => {
    try {
        const response = await withTokenRequest.get(`/pokemons/${titleId}`, {
            headers: getRequestHeaders()
        });
        const pokemons: Pokemon[] = response.data.data.pokemons.map((item: any) => {
            const types: Type[] = item.types.map((type: any) => {
                return {
                    id: type.id,
                    name: type.name
                };
            });
            const abilities: Ability[] = item.abilities.map((ability: any) => {
                return {
                    id: ability.id,
                    name: ability.name
                };
            });
            const baseStats: BaseStats = {
                id: item.base_stats.id,
                hitPoints: item.base_stats.hit_points,
                attack: item.base_stats.attack,
                defense: item.base_stats.defense,
                specialAttack: item.base_stats.special_attack,
                specialDefense: item.base_stats.special_defense,
                speed: item.base_stats.speed,
                average: item.base_stats.average,
                total: item.base_stats.total,
            }
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
                baseStats: baseStats,
                presetHeldItem: presetHeldItem,
            }
        });
        return pokemons;
    } catch (error) {
        console.log('error');
        throw error;
    }
}