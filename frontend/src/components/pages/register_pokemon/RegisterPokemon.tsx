import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useLocation } from 'react-router-dom';
import { TextField, Box } from '@mui/material';
import { getPokemons } from '../../../api/Pokemons';
import { getTeraTypes } from '../../../api/TeraTypes';
import { Title } from '../../../entity/Title';
import { Pokemon } from '../../../entity/Pokemon';
import { Ability } from '../../../entity/Ability';
import { TeraType } from '../../../entity/TeraType';
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
    const [teraTypes, setTeraTypes] = useState<TeraType[]>([]);
    const [selectedPokemon, setSelectedPokemon] = useState<Pokemon | null>(null);
    const [selectedAbility, setSelectedAbility] = useState<Ability | null>(null);
    const [selectedTeraType, setSelectedTeraType] = useState<TeraType | null>(null);
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
        const fetchTeraTypes = async () => {
            try {
                setTeraTypes(await getTeraTypes());
            } catch (error) {
                console.error(error);
            }
        }
        fetchTeraTypes();
    }, []);

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

    const handleTeraTypeChange = (_: any, value: Ability | null) => {
        setSelectedTeraType(value);
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
                </Box>
            )}
            <Autocomplete
                id="abilities"
                options={abilitiesOptions}
                value={selectedAbility}
                getOptionLabel={(option) => option.name}
                onChange={handleAbilityChange}
                renderInput={(params) => (
                    <TextField {...params} label="Ability" variant="outlined" />
                )}
                style={{ marginTop: '16px' }}
            />
            <Autocomplete
                id="abilities"
                options={teraTypes}
                value={selectedTeraType}
                getOptionLabel={(option) => option.name}
                onChange={handleTeraTypeChange}
                renderInput={(params) => (
                    <TextField {...params} label="TeraType" variant="outlined" />
                )}
                style={{ marginTop: '16px' }}
            />
        </>
    );
};

export default RegisterPokemon;
