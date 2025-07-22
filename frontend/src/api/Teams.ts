import { withTokenRequest, getRequestHeaders } from "../http";
import { SVTeam } from "../entity/SVTeam";
import { SVBredPokemon } from "../entity/SVBredPokemon";
import { NewTypeFromJson, Type } from "../entity/Type";
import { Item } from "../entity/Item";
import { Move, NewMoveFromJson } from "../entity/Move";
import { NewGenderFromJson } from "../entity/Gender";
import { NewAbilityFromJson } from "../entity/Ability";
import { NewNatureFromJson } from "../entity/Nature";
import { NewBaseStatsFromJson } from "../entity/BaseStats";
import { NewIndividualValuesFromJson } from "../entity/IndividualValues";
import { NewBasePointsFromJson } from "../entity/BasePoints";
import { NewActualValuesFromJson } from "../entity/ActualValues";
import { NewTeraTypeFromJson } from "../entity/TeraType";

export const postTeams = async (team: any): Promise<void> => {
    try {
        await withTokenRequest.post('/sv/teams', team, {
            headers: getRequestHeaders()
        });
    } catch (error) {
        console.error('Error posting team:', error);
        throw error;
    }
}

export const getTeams = async (): Promise<any> => {
    try {
        const response = await withTokenRequest.get('/sv/teams', {
            headers: getRequestHeaders()
        });
        const teams: SVTeam[] = response.data.data.teams.map((team: any) => {
            const bredPokemons: SVBredPokemon[] | null[]  = team.bred_pokemons.map((bredPokemon: any) => {
                if (bredPokemon === null) {
                    return null;
                }

                const types: Type[] = bredPokemon.types.map((type: any) => {
                    return NewTypeFromJson(type);
                });
                const heldItem: Item | null = bredPokemon.held_item
                ? {
                    id: bredPokemon.held_item.id,
                    name: bredPokemon.held_item.name,
                }
                : null;
                const moves: Move[] = bredPokemon.moves.map((move: any) => {
                    return NewMoveFromJson(move);
                });
    
                return {
                    bredPokemon: {
                        id: bredPokemon.id,
                        pokemonId: bredPokemon.pokemonId,
                        nationalPokedexNo: bredPokemon.nationalPokedexNo,
                        formeNo: bredPokemon.formeNo,
                        name: bredPokemon.name,
                        formeName: bredPokemon.formeName,
                        gender: NewGenderFromJson(bredPokemon.gender),
                        level: bredPokemon.level,
                        types: types,
                        ability: NewAbilityFromJson(bredPokemon.ability),
                        nature: NewNatureFromJson(bredPokemon.nature),
                        heldItem: heldItem,
                        baseStats: NewBaseStatsFromJson(bredPokemon.base_stats),
                        individualValues: NewIndividualValuesFromJson(bredPokemon.individual_values),
                        basePoints: NewBasePointsFromJson(bredPokemon.base_points),
                        actualValues: NewActualValuesFromJson(bredPokemon.actual_values),
                        moves: moves,
                        note: bredPokemon.note,
                    },
                    teraType: NewTeraTypeFromJson(bredPokemon.tera_type),
                };
            })
            return {
                team: {
                    id: team.id,
                    name: team.name,
                    note: team.note
                },
                bredPokemons: bredPokemons
            };
        })
        return teams;
    } catch (error) {
        console.error('Error fetching teams:', error);
        throw error;
    }
}

export const putTeams = async (team: any): Promise<void> => {
    try {
        await withTokenRequest.put('/sv/teams', team, {
            headers: getRequestHeaders()
        });
    } catch (error) {
        console.error('Error posting team:', error);
        throw error;
    }
}