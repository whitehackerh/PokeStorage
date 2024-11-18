import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useLocation } from 'react-router-dom';
import { TextField, Box } from '@mui/material';
import { getItems } from '../../../api/Items';
import { getMoves } from '../../../api/Moves';
import { getPokemons } from '../../../api/Pokemons';
import { getGenders } from '../../../api/Genders';
import { getTeraTypes } from '../../../api/TeraTypes';
import { getNatures } from '../../../api/Natures';
import { Title } from '../../../entity/Title';
import { Pokemon } from '../../../entity/Pokemon';
import { Gender } from '../../../entity/Gender';
import { Ability } from '../../../entity/Ability';
import { Move } from '../../../entity/Move';
import { Item } from '../../../entity/Item';
import { Nature } from '../../../entity/Nature';
import { TeraType } from '../../../entity/TeraType';
import { IndividualValues } from '../../../entity/IndividualValues';
import { BasePoints } from '../../../entity/BasePoints';
import { ActualValues} from '../../../entity/ActualValues';
import { TitleEnum } from '../../../enum/Title';
import { typeIcons } from '../../Icons/type';
import { moveCategoryIcons } from '../../Icons/move_category';
import Autocomplete from '@mui/material/Autocomplete';

const RegisterPokemon = () => {
    const location = useLocation();
    const [title, setTitle] = useState<Title | null>(null);
    const [pokemons, setPokemons] = useState<Pokemon[]>([]);
    const [genders, setGenders] = useState<Gender[]>([]);
    const [level, setLevel] = useState<number>(50);
    const [teraTypes, setTeraTypes] = useState<TeraType[]>([]);
    const [natures, setNatures] = useState<Nature[]>([]);
    const [items, setItems] = useState<Item[]>([]);
    const [moves, setMoves] = useState<Move[]>([]);
    const [selectedPokemon, setSelectedPokemon] = useState<Pokemon | null>(null);
    const [selectedGender, setSelectedGender] = useState<Gender | null>(null);
    const [selectedAbility, setSelectedAbility] = useState<Ability | null>(null);
    const [selectedItem, setSelectedItem] = useState<Item | null>(null);
    const [selectedMoves, setSelectedMoves] = useState<(Move | null)[]>([null, null, null, null]);
    const [toggleTeraType, setToggleTeraType] = useState<boolean>(false);
    const [selectedTeraType, setSelectedTeraType] = useState<TeraType | null>(null);
    const [selectedNature, setSelectedNature] = useState<Nature | null>(null);
    const [abilitiesOptions, setAbilitiesOptions] = useState<Ability[]>([]);
    const [individualValues, setIndividualValues] = useState<IndividualValues>({
        id: "",
        hitPoints: 0,
        attack: 0,
        defense: 0,
        specialAttack: 0,
        specialDefense: 0,
        speed: 0,
    });
    const [basePoints, setBasePoints] = useState<BasePoints>({
        id: "",
        hitPoints: 0,
        attack: 0,
        defense: 0,
        specialAttack: 0,
        specialDefense: 0,
        speed: 0,
    });
    const [actualValues, setActualValues] = useState<ActualValues>({
        id: "",
        hitPoints: 1,
        attack: 1,
        defense: 1,
        specialAttack: 1,
        specialDefense: 1,
        speed: 1,
    });
    const [note, setNote] = useState<string | null>(null);
    const navigate = useNavigate();

    const statFields: ("hitPoints" | "attack" | "defense" | "specialAttack" | "specialDefense" | "speed")[] = [
        "hitPoints", "attack", "defense", "specialAttack", "specialDefense", "speed"
    ];

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
                setGenders(await getGenders());
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

    const handleGenderChange = (_: any, value: Gender | null) => {
        setSelectedGender(value);
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

    const handleIndividualValuesChange = ( field: keyof IndividualValues, value: number) => {
        setIndividualValues((prev) => ({ ...prev, [field]: value }));
    }

    const handleBasePointsChange = ( field: keyof BasePoints, value: number) => {
        setBasePoints((prev) => ({ ...prev, [field]: value }));
    }

    const handleActualValuesChange = ( field: keyof ActualValues, value: number) => {
        setActualValues((prev) => ({ ...prev, [field]: value }));
    }

    const handleNoteChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setNote(event.target.value);
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
                    id="genders"
                    options={genders}
                    value={selectedGender}
                    getOptionLabel={(option) => option.name}
                    onChange={handleGenderChange}
                    renderInput={(params) => (
                        <TextField {...params} label="Gender" variant="outlined" />
                    )}
                    style={styles.autocomplete}
                />
                <Autocomplete
                    id="genders"
                    options={Array.from({ length: 100 }, (_, i) => i + 1)}
                    value={50}
                    renderInput={(params) => (
                        <TextField {...params} label="Level" variant="outlined" />
                    )}
                    style={styles.autocomplete}
                    disabled
                />
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
                        <TextField {...params} label="Held Item" variant="outlined" />
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
                <Box sx={{ display: "flex", flexDirection: "column", gap: 2 }}>
                    <Box key={"Individual Values"}>
                        <h3>{"Individual Values"}</h3>
                        <Box sx={{ display: "flex", gap: 2, flexWrap: "wrap" }}>
                            {statFields.map((field) => (
                                <Autocomplete
                                    key={field}
                                    options={Array.from({ length: 32 }, (_, i) => i + 0)}
                                    value={individualValues[field]}
                                    onChange={(event, newValue) => handleIndividualValuesChange(field, newValue ?? 0)}
                                    getOptionLabel={(option) => option.toString()} 
                                    renderInput={(params) => (
                                        <TextField
                                            {...params}
                                            label={field.charAt(0).toUpperCase() + field.slice(1)}
                                            variant="outlined"
                                        />
                                    )}
                                    style={{ width: 120 }}
                                />
                            ))}
                        </Box>
                    </Box>
                    <Box key={"Base Points"}>
                        <h3>{"Base Points"}</h3>
                        <Box sx={{ display: "flex", gap: 2, flexWrap: "wrap" }}>
                            {statFields.map((field) => (
                                <Autocomplete
                                    key={field}
                                    options={Array.from({ length: 253 }, (_, i) => i + 0)}
                                    value={basePoints[field]}
                                    onChange={(event, newValue) => handleBasePointsChange(field, newValue ?? 0)}
                                    getOptionLabel={(option) => option.toString()} 
                                    renderInput={(params) => (
                                        <TextField
                                            {...params}
                                            label={field.charAt(0).toUpperCase() + field.slice(1)}
                                            variant="outlined"
                                        />
                                    )}
                                    style={{ width: 120 }}
                                />
                            ))}
                        </Box>
                    </Box>
                    <Box key={"Actual Values"}>
                        <h3>{"Actual Values"}</h3>
                        <Box sx={{ display: "flex", gap: 2, flexWrap: "wrap" }}>
                            {statFields.map((field) => (
                                <Autocomplete
                                    key={field}
                                    options={Array.from({ length: 1000 }, (_, i) => i + 0)}
                                    value={actualValues[field]}
                                    onChange={(event, newValue) => handleActualValuesChange(field, newValue ?? 0)}
                                    getOptionLabel={(option) => option.toString()} 
                                    renderInput={(params) => (
                                        <TextField
                                            {...params}
                                            label={field.charAt(0).toUpperCase() + field.slice(1)}
                                            variant="outlined"
                                        />
                                    )}
                                    style={{ width: 120 }}
                                />
                            ))}
                        </Box>
                    </Box>
                </Box><br /><br />
                <TextField
                    id="note"
                    label="Note"
                    multiline
                    rows={4}
                    variant="outlined"
                    fullWidth
                    value={note}
                    onChange={handleNoteChange}
                />
            </div>
        </>
    );
};

export default RegisterPokemon;
