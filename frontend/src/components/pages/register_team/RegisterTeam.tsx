import { useEffect, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import { TextField, Box, Button } from '@mui/material';
import { Title } from '../../../entity/Title';
import { TitleEnum } from '../../../enum/Title';
import { SVBredPokemon } from '../../../entity/SVBredPokemon';
import { getBredPokemons } from '../../../api/BredPokemons';
import SelectPokemonGrid from './SelectPokemonGrid';
import SelectPokemonDialog from './SelectPokemonDialog';
import { toSnakeCase } from '../../../util/convert';
import { postTeams } from '../../../api/Teams';

const RegisterTeam = () => {
    const location = useLocation();
    const navigate = useNavigate();
    const [title, setTitle] = useState<Title | null>(null);
    const [bredPokemons, setBredPokemons] = useState<SVBredPokemon[]>([]);
    const [activeSlot, setActiveSlot] = useState<number | null>(null);
    const [selectedPokemons, setSelectedPokemons] = useState<(SVBredPokemon | null)[]>(Array(6).fill(null));
    const [name, setName] = useState<string>('');
    const [note, setNote] = useState<string | null>(null);
    const [isSelectDialogOpen, setSelectDialogOpen] = useState(false);

    useEffect(() => {
        if (location.state && location.state.title) {
            setTitle(location.state.title);
        }
    }, [location.state]);

    useEffect(() => {
        const fetch = async (title: Title) => {
            try {
                if (title.id == TitleEnum.SV) {
                    setBredPokemons(await getBredPokemons());
                }
            } catch (error) {
                console.error(error);
            }
        };
        if (title) {
            fetch(title);
        }
    }, [title]);

    const handleSlotClick = (index: number) => {
        setActiveSlot(index);
        setSelectDialogOpen(true);
    };

    const handleSelect = (pokemon: SVBredPokemon) => {
        if (activeSlot !== null) {
            const updated = [...selectedPokemons];
            updated[activeSlot] = pokemon;
            setSelectedPokemons(updated);
            setSelectDialogOpen(false);
        }
    };

    const handleUnselect = (index: number) => {
        const updated = [...selectedPokemons];
        updated[index] = null;
        setSelectedPokemons(updated);
    };

    const handleRegisterClick = async () => {
        try {
            if (name) {
                const madeTeam = makeTeam();
                await postTeams(toSnakeCase(madeTeam));
            }
            navigate('/team-list', { state: { title: title } });
        } catch (error) {
            console.error(error);
        }
    }

    const makeTeam = (id = ''): any | null => {
        if (name) {
            return {
                team: {
                    id: id,
                    name: name,
                    bredPokemonIds: selectedPokemons.map(p => p ? p.bredPokemon.id : null),
                    note: note,
                }
            }
        }
        return null;
    }

    return (
        <>
            <div>
                <Box sx={{ maxWidth: 1800, mx: 'auto', mt: 4 }}>
                    <TextField
                        label="Name"
                        style={{ width: 600 }}
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                        sx={{ mb: 2 }}
                    />
                    <br />
                    <SelectPokemonGrid
                        selected={selectedPokemons}
                        onSlotClick={handleSlotClick}
                        onUnselect={handleUnselect}
                    /><br /><br />
                    <TextField
                        label="Note"
                        multiline
                        rows={20}
                        variant="outlined"
                        style={{ width: 600 }}
                        value={note}
                        onChange={(e) => setNote(e.target.value)}
                    /><br /><br />
                    <Button onClick={handleRegisterClick} variant="contained" sx={{ mt: 3 }}>Register</Button>

                    <SelectPokemonDialog
                        open={isSelectDialogOpen}
                        onClose={() => setSelectDialogOpen(false)}
                        onSelect={handleSelect}
                        bredPokemons={bredPokemons}
                    />
                </Box>
            </div>
        </>
    )
}

export default RegisterTeam;