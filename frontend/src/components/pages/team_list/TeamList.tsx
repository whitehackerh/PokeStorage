import { useEffect, useState } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import { Title } from '../../../entity/Title';
import { SVTeam } from '../../../entity/SVTeam';
import { getTeams } from '../../../api/Teams';
import { TitleEnum } from '../../../enum/Title';
import "./TeamList.css";

const TeamList = () => {
    const location = useLocation();
    const navicate  = useNavigate();
    const [title, setTitle] = useState<Title | null>(null);
    const [teams, setTeams] = useState<SVTeam[]>([]);

    useEffect(() => {
        if (location.state && location.state.title) {
            setTitle(location.state.title);
        }
    }, [location.state]);

    useEffect(() => {
        const fetch = async (title: Title) => {
            try {
                if (title.id === TitleEnum.SV) {
                    setTeams(await getTeams()); 
                }
            } catch (error) {
                console.error(error);
            }
        }
        if (title) {
            fetch(title);
        }
    }, [title])

    const handleCardClick = (team: any) => {
        navicate('/team-detail', { state: { title: title, team: team } });
    }

    if (!title || !teams) {
        return <div>Loading...</div>;
    }

    return (
        <div className="team-list">
            {teams.map((team) => (
                <div key={team.team.id} className="team-card" onClick={() => handleCardClick(team)}>
                    <div className="team-header">
                        <h3>{team.team.name}</h3>
                    </div>
                    <div className="bred-pokemons">
                        {team.bredPokemons.map((pokemon) => (
                            pokemon ? (
                                <div key={pokemon.bredPokemon.id} className="pokemon-card">
                                    <div className="pokemon-header">
                                        <h3>{pokemon.bredPokemon.name} {pokemon.bredPokemon.formeName && (
                                        `(${pokemon.bredPokemon.formeName})`
                                        )}</h3>
                                    </div>
                                    {pokemon.bredPokemon.heldItem && (
                                            <p className="held-item">@{pokemon.bredPokemon.heldItem.name}</p>
                                        )}
                                    <p>{pokemon.bredPokemon.ability.name}</p>
                                    <div className="moves">
                                        <ul>
                                            {pokemon.bredPokemon.moves.map((move) => (
                                                <li key={move.name}>{move.name}</li>
                                            ))}
                                        </ul>
                                    </div>
                                </div>
                            ) : (<div className="pokemon-card"></div>)
                        ))}
                    </div>
                    <p className="note">{team.team.note}</p>
                </div>
            ))}
        </div>
    )
}

export default TeamList;