import { useEffect, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import { TextField, Box, Button } from '@mui/material';
import { getItems } from '../../../api/Items';
import { getMoves } from '../../../api/Moves';
import { getPokemons } from '../../../api/Pokemons';
import { getGenders } from '../../../api/Genders';
import { getTeraTypes } from '../../../api/TeraTypes';
import { getNatures } from '../../../api/Natures';
import { postBredPokemons, putBredPokemons } from '../../../api/BredPokemons';
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
import { toSnakeCase } from '../../../util/convert';
import { SVBredPokemon } from '../../../entity/SVBredPokemon';

const RegisterPokemon = () => {
    const location = useLocation();
    const navigate = useNavigate();
    const [isEditMode, setIsEditMode] = useState<boolean>(false);
    const [bredPokemon, setBredPokemon] = useState<SVBredPokemon | null>(null);
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

    const statFields: ("hitPoints" | "attack" | "defense" | "specialAttack" | "specialDefense" | "speed")[] = [
        "hitPoints", "attack", "defense", "specialAttack", "specialDefense", "speed"
    ];

    useEffect(() => {
        if (location.state && location.state.title) {
            setTitle(location.state.title);
            setToggleTeraType(location.state.title.id == TitleEnum.SV);
        }
        if (location.state && location.state.bredPokemon) {
            setIsEditMode(true);
            setBredPokemon(location.state.bredPokemon);
        }
    }, [location.state]);

    useEffect(() => {
        const fetch = async (title: Title) => {
            try {
                if (title.id == TitleEnum.SV) {
                    setPokemons(await getPokemons());
                    setGenders(await getGenders());
                    setNatures(await getNatures());
                    setItems(await getItems());
                    setMoves(await getMoves());
                }
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
        if (pokemons.length && bredPokemon) {
            const matchedPokemon = pokemons.find(
                (pokemon) => pokemon.id === bredPokemon.bredPokemon.pokemonId
            );
            if (matchedPokemon) {
                handlePokemonChange(null, matchedPokemon);
            }
            setSelectedGender(bredPokemon.bredPokemon.gender);
            setLevel(50);
            setSelectedNature(bredPokemon.bredPokemon.nature);
            setSelectedItem(bredPokemon.bredPokemon.heldItem);
            setIndividualValues({
                id: bredPokemon.bredPokemon.individualValues.id,
                hitPoints: bredPokemon.bredPokemon.individualValues.hitPoints,
                attack: bredPokemon.bredPokemon.individualValues.attack,
                defense: bredPokemon.bredPokemon.individualValues.defense,
                specialAttack: bredPokemon.bredPokemon.individualValues.specialAttack,
                specialDefense: bredPokemon.bredPokemon.individualValues.specialDefense,
                speed: bredPokemon.bredPokemon.individualValues.speed,
            });
            setBasePoints({
                id: bredPokemon.bredPokemon.basePoints.id,
                hitPoints: bredPokemon.bredPokemon.basePoints.hitPoints,
                attack: bredPokemon.bredPokemon.basePoints.attack,
                defense: bredPokemon.bredPokemon.basePoints.defense,
                specialAttack: bredPokemon.bredPokemon.basePoints.specialAttack,
                specialDefense: bredPokemon.bredPokemon.basePoints.specialDefense,
                speed: bredPokemon.bredPokemon.basePoints.speed,
            });
            setActualValues({
                id: bredPokemon.bredPokemon.actualValues.id,
                hitPoints: bredPokemon.bredPokemon.actualValues.hitPoints,
                attack: bredPokemon.bredPokemon.actualValues.attack,
                defense: bredPokemon.bredPokemon.actualValues.defense,
                specialAttack: bredPokemon.bredPokemon.actualValues.specialAttack,
                specialDefense: bredPokemon.bredPokemon.actualValues.specialDefense,
                speed: bredPokemon.bredPokemon.actualValues.speed,
            });
            const updatedMoves = [...selectedMoves];
            bredPokemon.bredPokemon.moves.forEach((move, index) => {
                updatedMoves[index] = move;
            });
            setSelectedMoves(updatedMoves);
            setNote(bredPokemon.bredPokemon.note);
            if (title && title.id == TitleEnum.SV) {
                setSelectedTeraType(bredPokemon.teraType);
            }
        }
    }, [bredPokemon, pokemons, genders, natures, items, moves])

    useEffect(() => {
        if (selectedPokemon) {
            setAbilitiesOptions(selectedPokemon.abilities);
            setSelectedAbility(null);
            if (bredPokemon) {
                const matchedAbility = selectedPokemon.abilities.find(
                    (ability) => ability.id === bredPokemon.bredPokemon.ability.id
                )
                if (matchedAbility) {
                    setSelectedAbility(matchedAbility);
                }
            }
        } else {
            setAbilitiesOptions([]);
        }
    }, [selectedPokemon]);

    const handleRegisterClick = async () => {
        try {
            if (title) {
                if (!isEditMode) {
                    const madeBredPokemon = makeBredPokemon()
                    await postBredPokemons(toSnakeCase(madeBredPokemon));
                } else if (isEditMode && bredPokemon) {
                    const madeBredPokemon = makeBredPokemon(bredPokemon.bredPokemon.id);
                    await putBredPokemons(toSnakeCase(madeBredPokemon));
                }
                navigate('/bred-pokemon-list', { state: { title: title } });
            }
        } catch (error) {
            console.error(error);
        }
    }

    const makeBredPokemon = (id = ''): any | null => {
        if (selectedPokemon
            && selectedGender
            && selectedAbility
            && selectedNature
            && selectedTeraType
        ) {
            return {
                bredPokemon: {
                    id: id,
                    pokemonId: selectedPokemon.id,
                    nationalPokedexNo: selectedPokemon.nationalPokedexNo,
                    formeNo: selectedPokemon.formeNo,
                    name: selectedPokemon.name,
                    formeName: selectedPokemon.formeName,
                    gender: selectedGender,
                    level: level,
                    types: selectedPokemon.types,
                    ability: selectedAbility,
                    nature: selectedNature,
                    heldItem: selectedItem,
                    baseStats: selectedPokemon.baseStats,
                    individualValues: individualValues,
                    basePoints: basePoints,
                    actualValues: actualValues,
                    moves: selectedMoves.filter((move): move is Move => move !== null),
                    note: note,
                    teraType: selectedTeraType,
                },
            }
        }
        return null;
    }

    const handlePokemonChange = (_: any, value: Pokemon | null) => {
        setSelectedPokemon(value);
        if (value?.presetHeldItem) {
            setSelectedItem(value.presetHeldItem);
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
                        value={selectedPokemon}
                        getOptionLabel={(option) =>
                            option.formeName ? `${option.name} (${option.formeName})` : option.name
                        }
                        onChange={handlePokemonChange}
                        renderInput={(params) => (
                            <TextField {...params} label="Pokémon" variant="outlined" />
                        )}
                        style={{'width': '400px'}}
                        disabled={isEditMode}
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
                <Box key={"Individual Values"}>
                    <h3>{"Individual Values"}</h3>
                    <Box sx={{ display: "flex", gap: 2, flexWrap: "wrap" }}>
                        {statFields.map((field) => (
                            <TextField
                                key={field}
                                label={field.charAt(0).toUpperCase() + field.slice(1)}
                                variant="outlined"
                                type="number"
                                value={individualValues[field]}
                                onChange={(event) => 
                                    handleIndividualValuesChange(field, Number(event.target.value) || 0)
                                }
                                style={{ width: 120 }}
                                inputProps={{ min: 0, max: 31 }}
                            />
                        ))}
                    </Box>
                </Box>
                <Box key={"Base Points"}>
                    <h3>{"Base Points"}</h3>
                    <Box sx={{ display: "flex", gap: 2, flexWrap: "wrap" }}>
                        {statFields.map((field) => (
                            <TextField
                                key={field}
                                label={field.charAt(0).toUpperCase() + field.slice(1)}
                                variant="outlined"
                                type="number"
                                value={basePoints[field]}
                                onChange={(event) => 
                                    handleBasePointsChange(field, Number(event.target.value) || 0)
                                }
                                style={{ width: 120 }}
                                inputProps={{ min: 0, max: 252 }}
                            />
                        ))}
                    </Box>
                </Box>
                <Box key={"Actual Values"}>
                    <h3>{"Actual Values"}</h3>
                    <Box sx={{ display: "flex", gap: 2, flexWrap: "wrap" }}>
                        {statFields.map((field) => (
                            <TextField
                                key={field}
                                label={field.charAt(0).toUpperCase() + field.slice(1)}
                                variant="outlined"
                                type="number"
                                value={actualValues[field]}
                                onChange={(event) => 
                                    handleActualValuesChange(field, Number(event.target.value) || 0)
                                }
                                style={{ width: 120 }}
                                inputProps={{ min: 1, max: 999 }}
                            />
                        ))}
                    </Box>
                </Box><br /><br />
                <TextField
                    id="note"
                    label="Note"
                    multiline
                    rows={4}
                    variant="outlined"
                    style={{ width: 600 }}
                    value={note}
                    onChange={handleNoteChange}
                /><br /><br />
                <Button variant="contained" style={{ margin: "10px" }} onClick={handleRegisterClick}>Register</Button>
            </div>
        </>
    );
};

export default RegisterPokemon;
