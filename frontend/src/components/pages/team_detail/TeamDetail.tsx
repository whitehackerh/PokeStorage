import { useEffect, useState } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import { deleteTeams } from '../../../api/Teams';
import { Title } from '../../../entity/Title';
import { SVTeam } from '../../../entity/SVTeam';
import {
  Box, Grid, Paper, Table, TableBody, TableHead, TableRow,
  TableCell, Typography, Button
} from '@mui/material';
import { stats, statDisplayNames } from '../../../common/stats';
import ConfirmDialog from '../../modules/dialogs/ConfirmDialog';

const TeamDetail = () => {
    const location = useLocation();
    const navigate = useNavigate();
    const [title, setTitle] = useState<Title | null>(null);
    const [team, setTeam] = useState<SVTeam | null>(null);
    const [isDialogOpen, setIsDialogOpen] = useState(false);
    
    useEffect(() => {
        if (location.state && location.state.title && location.state.team) {
            setTitle(location.state.title);
            setTeam(location.state.team);
        }
    }, [location.state]);

    const handleEditButtonClick = () => {
        navigate('/register-team', { state: { title: title, team: team } });
    }

    const handleDeleteButtonClick = () => {
        setIsDialogOpen(true);
    };

    const handleDialogClose = async (isConfirmed: boolean) => {
        setIsDialogOpen(false);
        if (isConfirmed && title && team) {
            try {
                await deleteTeams(team.team.id);
                navigate('/team-list', { state: { title: title } });
            } catch (error) {
                console.error('Failed to delete Team:', error);
            }
        }
    };

    if (!title || !team) {
        return <Typography>Loading...</Typography>
    }

    return (
        <>
            <div>
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
                <Box sx={{ maxWidth: 1800, mx: 'auto', mt: 4 }}>
                    <div className="team-header">
                        <h1>{team.team.name}</h1>
                    </div>
                    <Grid container spacing={2}>
                        {Array.from({ length: 6 }).map((_, i) => (
                            <Grid item xs={6} key={i}>
                            <Paper
                                sx={{
                                height: team.bredPokemons[i] ? 'auto' : 150,
                                display: 'flex',
                                justifyContent: 'center',
                                alignItems: team.bredPokemons[i] ? 'stretch' : 'center',
                                backgroundColor: team.bredPokemons[i] ? '#e0f2f1' : '#f9f9f9',
                                cursor: 'pointer',
                                position: 'relative',
                                }}
                            >
                                {team.bredPokemons[i] ? (
                                <Box sx={{ mt: 2, mx: 'auto', maxWidth: 600, p: 2 }}>
                                    <Typography variant="h4" align="center">
                                    {team.bredPokemons[i]?.bredPokemon.name} {team.bredPokemons[i]?.bredPokemon.formeName && `(${team.bredPokemons[i]?.bredPokemon.formeName})`}
                                    </Typography><br />
                                    <Typography><strong>Gender:</strong> {team.bredPokemons[i]?.bredPokemon.gender.name}</Typography>
                                    <Typography><strong>Level:</strong> {team.bredPokemons[i]?.bredPokemon.level}</Typography>
                                    <Box sx={{ display: 'flex', alignItems: 'flex-start', mt: 0 }}>
                                    <Typography sx={{ fontWeight: 'bold', mr: 1 }}>Types:</Typography>
                                    <Typography>{team.bredPokemons[i]?.bredPokemon.types[0].name}</Typography>
                                    </Box>
                                    <Box sx={{ ml: 7 }}>
                                    {team.bredPokemons[i]?.bredPokemon.types.slice(1).map((type, index) => (
                                        <Typography key={index}>{type.name}</Typography>
                                    ))}
                                    </Box>
                                    <Typography><strong>Tera Type:</strong> {team.bredPokemons[i]?.teraType.name}</Typography>
                                    <Typography><strong>Ability:</strong> {team.bredPokemons[i]?.bredPokemon.ability.name}</Typography>
                                    <Typography><strong>Nature:</strong> {team.bredPokemons[i]?.bredPokemon.nature.name}</Typography>
                                    <Typography><strong>Held Item:</strong> {team.bredPokemons[i]?.bredPokemon.heldItem?.name ?? 'None'}</Typography>

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
                                            <TableCell>{team.bredPokemons[i]?.bredPokemon.baseStats[stat]}</TableCell>
                                            <TableCell>{team.bredPokemons[i]?.bredPokemon.individualValues[stat]}</TableCell>
                                            <TableCell>{team.bredPokemons[i]?.bredPokemon.basePoints[stat]}</TableCell>
                                            <TableCell>{team.bredPokemons[i]?.bredPokemon.actualValues[stat]}</TableCell>
                                            </TableRow>
                                        ))}
                                        </TableBody>
                                    </Table>
                                    </Paper>

                                    <Box sx={{ display: 'flex', alignItems: 'flex-start', mt: 2 }}>
                                    <Typography sx={{ fontWeight: 'bold', mr: 1 }}>Moves:</Typography>
                                    <Typography>{team.bredPokemons[i]?.bredPokemon.moves[0].name}</Typography>
                                    </Box>
                                    <Box sx={{ ml: 7.6 }}>
                                    {team.bredPokemons[i]?.bredPokemon.moves.slice(1).map((move, index) => (
                                        <Typography key={index}>{move.name}</Typography>
                                    ))}
                                    </Box>
                                    <Box sx={{ display: 'flex', alignItems: 'flex-start', mt: 1 }}>
                                    <Typography sx={{ fontWeight: 'bold', mr: 1 }}>Note:</Typography>
                                    <Typography sx={{ whiteSpace: 'pre-wrap' }}>
                                        {team.bredPokemons[i]?.bredPokemon.note || ''}
                                    </Typography>
                                    </Box>
                                </Box>
                                ) : (
                                <Typography></Typography>
                                )}
                            </Paper>
                            </Grid>
                        ))}
                    </Grid><br />
                    <Box sx={{ display: 'flex', alignItems: 'flex-start', mt: 1 }}>
                        <Typography sx={{ fontWeight: 'bold', mr: 1 }}>Note:</Typography>
                        <Typography sx={{ whiteSpace: 'pre-wrap' }}>
                            {team.team.note || ''}
                        </Typography>
                    </Box>
                </Box>
            </div>

            <ConfirmDialog
                text=""
                message="Are you sure you want to delete this team?"
                callback={handleDialogClose}
                open={isDialogOpen}
            />
        </>
    )
}

export default TeamDetail;