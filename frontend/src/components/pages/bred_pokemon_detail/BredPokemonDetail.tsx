import { useEffect, useState } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import { deleteBredPokemons } from '../../../api/BredPokemons';
import { Title } from '../../../entity/Title';
import { IndividualValues } from '../../../entity/IndividualValues';
import { SVBredPokemon } from '../../../entity/SVBredPokemon';
import { Box, Typography, Table, TableHead, TableRow, TableCell, TableBody, Paper, Button } from '@mui/material';
import ConfirmDialog from '../../modules/dialogs/ConfirmDialog';

const BredPokemonDetail = () => {
    const location = useLocation();
    const navigate = useNavigate();
    const [title, setTitle] = useState<Title | null>(null);
    const [bredPokemon, setBredPokemon] = useState<SVBredPokemon | null>(null);
    const [isDialogOpen, setIsDialogOpen] = useState(false);

    useEffect(() => {
        if (location.state && location.state.title && location.state.bredPokemon) {
            setTitle(location.state.title);
            setBredPokemon(location.state.bredPokemon);
        }
    }, [location.state]);

    const handleEditButtonClick = () => {
        navigate('/register-pokemon', { state: { title: title, bredPokemon: bredPokemon } });
    }

    const handleDeleteButtonClick = () => {
        setIsDialogOpen(true);
    };

    const handleDialogClose = async (isConfirmed: boolean) => {
        setIsDialogOpen(false);
        if (isConfirmed && title && bredPokemon) {
            try {
                await deleteBredPokemons(bredPokemon.bredPokemon.id);
                navigate('/bred-pokemon-list', { state: { title: title } });
            } catch (error) {
                console.error('Failed to delete bred Pokémon:', error);
            }
        }
    };

    if (!title || !bredPokemon) {
        return <Typography>Loading...</Typography>;
    }

    const stats: (keyof Omit<IndividualValues, 'id'>)[] = [
        'hitPoints',
        'attack',
        'defense',
        'specialAttack',
        'specialDefense',
        'speed',
    ];

    const statDisplayNames: Record<keyof Omit<IndividualValues, 'id'>, string> = {
        hitPoints: 'Hit Points',
        attack: 'Attack',
        defense: 'Defense',
        specialAttack: 'Special Attack',
        specialDefense: 'Special Defense',
        speed: 'Speed',
    };

    return (
        <>
            <Box sx={{ display: 'flex', justifyContent: 'center', mt: 2 }}>
                <Button variant="contained" sx={{ marginRight: 2 }} onClick={handleEditButtonClick}>
                    Edit
                </Button>
                <Button
                    variant="contained"
                    sx={{ backgroundColor: 'red', '&:hover': { backgroundColor: 'darkred' } }}
                    onClick={handleDeleteButtonClick}
                >
                    Delete
                </Button>
            </Box>
            <Box sx={{ mt: 2, mx: 'auto', maxWidth: 600, p: 2, border: 1, borderRadius: 2, boxShadow: 2 }}>
                <Typography variant="h4" align="center">
                    {bredPokemon.bredPokemon.name} {bredPokemon.bredPokemon.formeName && `(${bredPokemon.bredPokemon.formeName})`}
                </Typography><br />
                <Typography><strong>Gender:</strong> {bredPokemon.bredPokemon.gender.name}</Typography>
                <Typography><strong>Level:</strong> {bredPokemon.bredPokemon.level}</Typography>
                <Box sx={{ display: 'flex', alignItems: 'flex-start', mt: 0 }}>
                    <Typography sx={{ fontWeight: 'bold', mr: 1 }}>Types:</Typography>
                    <Typography>{bredPokemon.bredPokemon.types[0].name}</Typography>
                </Box>
                <Box sx={{ ml: 7 }}>
                    {bredPokemon.bredPokemon.types.slice(1).map((type, index) => (
                        <Typography key={index}>{type.name}</Typography>
                    ))}
                </Box>
                <Typography><strong>Tera Type:</strong> {bredPokemon.teraType.name}</Typography>
                <Typography><strong>Ability:</strong> {bredPokemon.bredPokemon.ability.name}</Typography>
                <Typography><strong>Nature:</strong> {bredPokemon.bredPokemon.nature.name}</Typography>
                <Typography><strong>Held Item:</strong> {bredPokemon.bredPokemon.heldItem == null ? 'None' : bredPokemon.bredPokemon.heldItem.name}</Typography>

                <Paper elevation={3} sx={{ mt: 2 }}>
                    <Table>
                        <TableHead>
                            <TableRow>
                                <TableCell></TableCell>
                                <TableCell><strong>Base Stats</strong></TableCell>
                                <TableCell><strong>Individual Values</strong></TableCell>
                                <TableCell><strong>Base Points</strong></TableCell>
                                <TableCell><strong>Actual Values</strong></TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                        {stats.map((stat) => (
                                <TableRow key={stat}>
                                    <TableCell>{statDisplayNames[stat]}</TableCell>
                                    <TableCell>{bredPokemon.bredPokemon.baseStats[stat]}</TableCell>
                                    <TableCell>{bredPokemon.bredPokemon.individualValues[stat]}</TableCell>
                                    <TableCell>{bredPokemon.bredPokemon.basePoints[stat]}</TableCell>
                                    <TableCell>{bredPokemon.bredPokemon.actualValues[stat]}</TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </Paper>

                <Box sx={{ display: 'flex', alignItems: 'flex-start', mt: 2 }}>
                    <Typography sx={{ fontWeight: 'bold', mr: 1 }}>Moves:</Typography>
                    <Typography>{bredPokemon.bredPokemon.moves[0].name}</Typography>
                </Box>
                <Box sx={{ ml: 7.6 }}>
                    {bredPokemon.bredPokemon.moves.slice(1).map((move, index) => (
                        <Typography key={index}>{move.name}</Typography>
                    ))}
                </Box>
                <Box sx={{ display: 'flex', alignItems: 'flex-start', mt: 1 }}>
                    <Typography sx={{ fontWeight: 'bold', mr: 1 }}>Note:</Typography>
                    <Typography sx={{ whiteSpace: 'pre-wrap' }}>
                        {bredPokemon.bredPokemon.note || ''}
                    </Typography>
                </Box>
            </Box>

            <ConfirmDialog
                text=""
                message="Are you sure you want to delete this bred Pokémon?"
                callback={handleDialogClose}
                open={isDialogOpen}
            />
        </>
    );
};

export default BredPokemonDetail;
