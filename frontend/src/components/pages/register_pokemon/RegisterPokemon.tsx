import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useLocation } from 'react-router-dom';
import { Container, Typography, List, ListItem, TextField, Box } from '@mui/material';
import { getPokemons } from '../../../api/Pokemons';
import { Title } from '../../../entity/Title';
import { Pokemon } from '../../../entity/Pokemon';
import { Ability } from '../../../entity/Ability';
import electricIcon from '../../../assets/img/type/electric.png';
import fairyIcon from '../../../assets/img/type/fairy.png';
import steelIcon from '../../../assets/img/type/steel.png';
import Autocomplete from '@mui/material/Autocomplete';

const typeIcons: Record<string, string> = {
    electric: electricIcon,
    fairy: fairyIcon,
    steel: steelIcon,
};

const RegisterPokemon = () => {
    const location = useLocation();
    const [title, setTitle] = useState<Title | null>(null);
    const [pokemons, setPokemons] = useState<Pokemon[]>([]);
    const [selectedPokemon, setSelectedPokemon] = useState<Pokemon | null>(null);
    const [selectedAbility, setSelectedAbility] = useState<Ability | null>(null);
    const [abilitiesOptions, setAbilitiesOptions] = useState<Ability[]>([]);
    const navigate = useNavigate();

    useEffect(() => {
        if (location.state && location.state.title) {
            setTitle(location.state.title);
        }
    }, [location.state]);

    useEffect(() => {
        const fetchPokemons = async (title: Title) => {
            try {
                setPokemons(await getPokemons(title.id));
            } catch (error) {
                console.error(error);
            }
        };
        if (title) {
            fetchPokemons(title);
        }
    }, [title]);

    useEffect(() => {
        if (selectedPokemon) {
            setAbilitiesOptions(selectedPokemon.abilities);
            setSelectedAbility(null);
        } else {
            setAbilitiesOptions([]);
        }
    }, [selectedPokemon]);

    const handlePokemonChange = (_: any, value: Pokemon | null) => {
        setSelectedPokemon(value);
    };

    const handleAbilityChange = (_: any, value: Ability | null) => {
        setSelectedAbility(value);
    };

    if (!title) {
        return <div>Loading...</div>;
    }

    return (
        <>
            <Autocomplete
                id="pokemon"
                options={pokemons}
                getOptionLabel={(option) =>
                    option.formeName ? `${option.name} (${option.formeName})` : option.name
                }
                onChange={handlePokemonChange}
                renderInput={(params) => (
                    <TextField {...params} label="Pokemon" variant="outlined" />
                )}
            />
            {selectedPokemon && (
                <Box mt={2}>
                    <Box display="flex" alignItems="center">
                        {selectedPokemon.types.map((type, index) => (
                            <img
                                key={index}
                                src={typeIcons[type.name.toLowerCase()]}
                                alt={type.name}
                                style={{ height: '40px', marginRight: '8px' }}
                            />
                        ))}
                    </Box>
                    <Autocomplete
                        id="abilities"
                        options={abilitiesOptions}
                        value={selectedAbility}
                        getOptionLabel={(option) => option.name}
                        onChange={handleAbilityChange}
                        renderInput={(params) => (
                            <TextField {...params} label="Select an Ability" variant="outlined" />
                        )}
                        style={{ marginTop: '16px' }}
                    />
                </Box>

            )}
        </>
    );
};

export default RegisterPokemon;
