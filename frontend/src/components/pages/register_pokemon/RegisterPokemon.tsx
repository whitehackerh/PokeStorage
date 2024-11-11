import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useLocation } from 'react-router-dom';
import { TextField, Box } from '@mui/material';
import { getItems } from '../../../api/Items';
import { getMoves } from '../../../api/Moves';
import { getPokemons } from '../../../api/Pokemons';
import { getTeraTypes } from '../../../api/TeraTypes';
import { getNatures } from '../../../api/Natures';
import { Title } from '../../../entity/Title';
import { Pokemon } from '../../../entity/Pokemon';
import { Ability } from '../../../entity/Ability';
import { Move } from '../../../entity/Move';
import { Item } from '../../../entity/Item';
import { Nature } from '../../../entity/Nature';
import { TeraType } from '../../../entity/TeraType';
import { TitleEnum } from '../../../enum/Title';
import { typeIcons } from '../../Icons/type';
import { moveCategoryIcons } from '../../Icons/move_category';
import Autocomplete from '@mui/material/Autocomplete';

const RegisterPokemon = () => {
    const location = useLocation();
    const [title, setTitle] = useState<Title | null>(null);
    const [pokemons, setPokemons] = useState<Pokemon[]>([]);
    const [teraTypes, setTeraTypes] = useState<TeraType[]>([]);
    const [natures, setNatures] = useState<Nature[]>([]);
    const [items, setItems] = useState<Item[]>([]);
    const [moves, setMoves] = useState<Move[]>([]);
    const [selectedPokemon, setSelectedPokemon] = useState<Pokemon | null>(null);
    const [selectedAbility, setSelectedAbility] = useState<Ability | null>(null);
    const [selectedItem, setSelectedItem] = useState<Item | null>(null);
    const [selectedMoves, setSelectedMoves] = useState<(Move | null)[]>([null, null, null, null]);
    const [toggleTeraType, setToggleTeraType] = useState<boolean>(false);
    const [selectedTeraType, setSelectedTeraType] = useState<TeraType | null>(null);
    const [selectedNature, setSelectedNature] = useState<Nature | null>(null);
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
                setNatures(await getNatures());
                setItems(await getItems(title.id));
                setMoves(await getMoves(title.id));
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

    const handleNatureChange = (_: any, value: Nature | null) => {
        setSelectedNature(value);
    }

    const handleItemChange = (_: any, value: Ability | null) => {
        setSelectedItem(value);
    };

    const handleMoveChange = (index: number) => (_: any, value: Move | null) => {
        const newSelectedMoves = [...selectedMoves];
        newSelectedMoves[index] = value;
        setSelectedMoves(newSelectedMoves);
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
            width: "200px",
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
                            <TextField {...params} label="Pokémon" variant="outlined" />
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
                    id="natures"
                    options={natures}
                    value={selectedNature}
                    getOptionLabel={(option) => option.name}
                    onChange={handleNatureChange}
                    renderInput={(params) => (
                        <TextField {...params} label="Nature" variant="outlined" />
                    )}
                    style={styles.autocomplete}
                />
                <Autocomplete
                    id="items"
                    options={items}
                    value={selectedItem}
                    getOptionLabel={(option) => option.name}
                    onChange={handleItemChange}
                    renderInput={(params) => (
                        <TextField {...params} label="Item" variant="outlined" />
                    )}
                    disabled={selectedPokemon !== null && selectedPokemon.presetHeldItem !== null}
                    style={styles.autocomplete}
                />
                <br />
                {selectedMoves.map((_, index) => (
                    <Box key={index} style={styles.container}>
                        <Autocomplete
                            options={moves}
                            getOptionLabel={(option) => option.name}
                            value={selectedMoves[index]}
                            onChange={handleMoveChange(index)}
                            renderInput={(params) => (
                                <TextField {...params} label={`Move ${index + 1}`} variant="outlined" />
                            )}
                            style={{ width: '400px', marginBottom: '20px' }}
                        />
                        {selectedMoves[index]?.type && (
                            <Box style={styles.iconContainer}>
                                <img
                                    src={typeIcons[selectedMoves[index]?.type?.name.toLowerCase() || '']}
                                    alt={selectedMoves[index]?.type?.name || ''}
                                    style={styles.icon}
                                />
                                <img
                                    src={moveCategoryIcons[selectedMoves[index]?.moveCategory?.name.toLowerCase() || '']}
                                    alt={selectedMoves[index]?.moveCategory?.name || ''}
                                    style={{ height: '40x', width: '40px', objectFit: 'contain', marginRight: '30px' }}
                                />
                                <h3 style={{ marginRight: '30px' }}>Power: {selectedMoves[index]?.power ? selectedMoves[index]?.power : '---'}</h3>
                                <h3>Accuracy: {selectedMoves[index]?.accuracy ? selectedMoves[index]?.accuracy : '---'}</h3>
                            </Box>
                        )}
                    </Box>
                ))}
            </div>
        </>
    );
};

export default RegisterPokemon;
