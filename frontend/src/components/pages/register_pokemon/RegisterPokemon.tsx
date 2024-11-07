import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useLocation } from 'react-router-dom';
import { TextField, Box } from '@mui/material';
import { getItems } from '../../../api/Items';
import { getPokemons } from '../../../api/Pokemons';
import { getTeraTypes } from '../../../api/TeraTypes';
import { Title } from '../../../entity/Title';
import { Pokemon } from '../../../entity/Pokemon';
import { Ability } from '../../../entity/Ability';
import { Item } from '../../../entity/Item';
import { TeraType } from '../../../entity/TeraType';
import { TitleEnum } from '../../../enum/Title';
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
    const [items, setItems] = useState<Item[]>([]);
    const [selectedPokemon, setSelectedPokemon] = useState<Pokemon | null>(null);
    const [selectedAbility, setSelectedAbility] = useState<Ability | null>(null);
    const [selectedItem, setSelectedItem] = useState<Item | null>(null);
    const [toggleTeraType, setToggleTeraType] = useState<boolean>(false);
    const [selectedTeraType, setSelectedTeraType] = useState<TeraType | null>(null);
    const [abilitiesOptions, setAbilitiesOptions] = useState<Ability[]>([]);
    const navigate = useNavigate();

    useEffect(() => {
        if (location.state && location.state.title) {
            setTitle(location.state.title);
            setToggleTeraType(location.state.title.id == TitleEnum.SV);
        }
    }, [location.state]);

    useEffect(() => {
        const fetch = async (title: Title) => {
            try {
                setPokemons(await getPokemons(title.id));
                setItems(await getItems(title.id));
            } catch (error) {
                console.error(error);
            }
        };
        if (title) {
            fetch(title);
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
        if (title != null && toggleTeraType) {
            fetchTeraTypes();
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
        if (value?.presetHeldItem) {
            setSelectedItem(value.presetHeldItem);
        } else {

        }
    };

    const handleAbilityChange = (_: any, value: Ability | null) => {
        setSelectedAbility(value);
    };

    const handleTeraTypeChange = (_: any, value: Ability | null) => {
        setSelectedTeraType(value);
    };

    const handleItemChange = (_: any, value: Ability | null) => {
        setSelectedItem(value);
    };

    if (!title) {
        return <div>Loading...</div>;
    }

    const styles: {
        container: React.CSSProperties;
        iconContainer: React.CSSProperties;
        icon: React.CSSProperties;
        autocomplete: React.CSSProperties;
    } = {
        container: {
            display: "flex",
            alignItems: "center",
            gap: "16px",
        },
        iconContainer: {
            display: "flex",
            alignItems: "center",
        },
        icon: {
            height: "40px",
            marginRight: "8px",
        },
        autocomplete: {
            marginTop: "16px",
            width: "400px"
        },
    };

    return (
        <>
            <div style={{"margin": '40px'}}>
                <Box style={styles.container}>
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
                        style={{'width': '400px'}}
                    />
                    {selectedPokemon && (
                        <Box style={styles.iconContainer}>
                            {selectedPokemon.types.map((type, index) => (
                                <img
                                    key={index}
                                    src={typeIcons[type.name.toLowerCase()]}
                                    alt={type.name}
                                    style={styles.icon}
                                />
                            ))}
                        </Box>
                    )}
                </Box>
                <Autocomplete
                    id="abilities"
                    options={abilitiesOptions}
                    value={selectedAbility}
                    getOptionLabel={(option) => option.name}
                    onChange={handleAbilityChange}
                    renderInput={(params) => (
                        <TextField {...params} label="Ability" variant="outlined" />
                    )}
                    style={styles.autocomplete}
                />
                {toggleTeraType && (
                    <Autocomplete
                        id="teraType"
                        options={teraTypes}
                        value={selectedTeraType}
                        getOptionLabel={(option) => option.name}
                        onChange={handleTeraTypeChange}
                        renderInput={(params) => (
                            <TextField {...params} label="Tera Type" variant="outlined" />
                        )}
                        style={styles.autocomplete}
                    />
                )}
                <Autocomplete
                    id="items"
                    options={items}
                    value={selectedItem}
                    getOptionLabel={(option) => option.name}
                    onChange={handleItemChange}
                    renderInput={(params) => (
                        <TextField {...params} label="Item" variant="outlined" />
                    )}
                    disabled={selectedPokemon?.presetHeldItem !== null}
                    style={styles.autocomplete}
                />
            </div>
        </>
    );
};

export default RegisterPokemon;
