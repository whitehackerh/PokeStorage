import { Team } from "./Team";
import { SVBredPokemon } from "./SVBredPokemon";

export interface SVTeam {
    team: Team;
    bredPokemons: (SVBredPokemon | null)[];
}