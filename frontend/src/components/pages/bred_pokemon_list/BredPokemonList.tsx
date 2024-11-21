import { useEffect, useState } from 'react';
import { useLocation } from 'react-router-dom';
import { Title } from '../../../entity/Title';
import { SVBredPokemon } from '../../../entity/SVBredPokemon';
import { getBredPokemons } from '../../../api/BredPokemons';
import "./BredPokemonListPage.css";

const BredPokemonListPage = () => {
    const location = useLocation();
    const [title, setTitle] = useState<Title | null>(null);
    const [bredPokemons, setBredPokemons] = useState<SVBredPokemon[]>([]);

    useEffect(() => {
        if (location.state && location.state.title) {
            setTitle(location.state.title);
        }
    }, [location.state]);

    useEffect(() => {
        const fetch = async (title: Title) => {
            try {
                setBredPokemons(await getBredPokemons(title.id));
            } catch (error) {
                console.error(error);
            }
        };
        if (title) {
            fetch(title);
        }
    }, [title]);

    const getTypeIconPath = (typeName: string) => {
        try {
            return require(`../../../assets/img/type/${typeName.toLowerCase()}.png`);
        } catch {
            return "";
        }
    };

    if (!title || !bredPokemons) {
        return <div>Loading...</div>;
    }

    return (
        <div className="pokemon-list">
            {bredPokemons.map((pokemon) => (
                <div key={pokemon.bredPokemon.id} className="pokemon-card">
                    <div className="pokemon-header">
                        <h3>{pokemon.bredPokemon.name} {pokemon.bredPokemon.formeName && (
                        `(${pokemon.bredPokemon.formeName})`
                        )}</h3>
                    </div>
                    {pokemon.bredPokemon.heldItem && (
                            <p className="held-item">@{pokemon.bredPokemon.heldItem.name}</p>
                        )}
                    <p>{pokemon.bredPokemon.nature.name}</p>
                    <p>{pokemon.bredPokemon.ability.name}</p>
                    <div className="moves">
                        <ul>
                            {pokemon.bredPokemon.moves.map((move) => (
                                <li key={move.name}>{move.name}</li>
                            ))}
                        </ul>
                    </div>
                    <p className="note">{pokemon.bredPokemon.note}</p>
                </div>
            ))}
        </div>
    );
}

export default BredPokemonListPage;