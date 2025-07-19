import React from 'react';
import {
  Dialog, DialogTitle, DialogContent, Grid, Paper, Typography, Box, Table, TableBody, TableHead, TableRow,
  TableCell
} from '@mui/material';
import { SVBredPokemon } from '../../../entity/SVBredPokemon';
import { stats, statDisplayNames } from '../../../common/stats';

type Props = {
    open: boolean;
    onClose: () => void;
    bredPokemons: SVBredPokemon[];
    onSelect: (pokemon: SVBredPokemon) => void;
}

const SelectPokemonDialog: React.FC<Props> = ({ open, onClose, bredPokemons, onSelect }) => {
    return (
       <Dialog open={open} onClose={onClose} fullWidth maxWidth="lg">
            <DialogTitle>Select Pokemon</DialogTitle>
            <DialogContent>
            <Grid container spacing={2}>
                {bredPokemons.map((p) => (
                <Grid item xs={6} key={p.bredPokemon.id}>
                    <Paper sx={{ p: 2, cursor: 'pointer' }} onClick={() => onSelect(p)}>
                        <Typography variant="h5" align="center">
                        {p.bredPokemon.name} {p.bredPokemon.formeName && `(${p.bredPokemon.formeName})`}
                        </Typography><br />
                        <Typography><strong>Gender:</strong> {p.bredPokemon.gender.name}</Typography>
                        <Typography><strong>Level:</strong> {p.bredPokemon.level}</Typography>
                        <Box sx={{ display: 'flex', alignItems: 'flex-start', mt: 0 }}>
                        <Typography sx={{ fontWeight: 'bold', mr: 1 }}>Types:</Typography>
                        <Typography>{p.bredPokemon.types[0].name}</Typography>
                        </Box>
                        <Box sx={{ ml: 7 }}>
                        {p.bredPokemon.types.slice(1).map((type, index) => (
                            <Typography key={index}>{type.name}</Typography>
                        ))}
                        </Box>
                        <Typography><strong>Tera Type:</strong> {p.teraType.name}</Typography>
                        <Typography><strong>Ability:</strong> {p.bredPokemon.ability.name}</Typography>
                        <Typography><strong>Nature:</strong> {p.bredPokemon.nature.name}</Typography>
                        <Typography><strong>Held Item:</strong> {p.bredPokemon.heldItem?.name ?? 'None'}</Typography>

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
                                <TableCell>{p.bredPokemon.baseStats[stat]}</TableCell>
                                <TableCell>{p.bredPokemon.individualValues[stat]}</TableCell>
                                <TableCell>{p.bredPokemon.basePoints[stat]}</TableCell>
                                <TableCell>{p.bredPokemon.actualValues[stat]}</TableCell>
                                </TableRow>
                            ))}
                            </TableBody>
                        </Table>
                        </Paper>

                        <Box sx={{ display: 'flex', alignItems: 'flex-start', mt: 2 }}>
                        <Typography sx={{ fontWeight: 'bold', mr: 1 }}>Moves:</Typography>
                        <Typography>{p.bredPokemon.moves[0].name}</Typography>
                        </Box>
                        <Box sx={{ ml: 7.6 }}>
                        {p.bredPokemon.moves.slice(1).map((move, index) => (
                            <Typography key={index}>{move.name}</Typography>
                        ))}
                        </Box>
                        <Box sx={{ display: 'flex', alignItems: 'flex-start', mt: 1 }}>
                        <Typography sx={{ fontWeight: 'bold', mr: 1 }}>Note:</Typography>
                        <Typography sx={{ whiteSpace: 'pre-wrap' }}>
                            {p.bredPokemon.note || ''}
                        </Typography>
                        </Box>
                    </Paper>
                </Grid>
                ))}
            </Grid>
            </DialogContent>
        </Dialog>
    );
};

export default SelectPokemonDialog;