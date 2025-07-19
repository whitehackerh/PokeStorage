import React, { useState } from 'react';
import {
  Box, Grid, Paper, Table, TableBody, TableHead, TableRow,
  TableCell, Typography, Button
} from '@mui/material';
import { SVBredPokemon } from '../../../entity/SVBredPokemon';
import { stats, statDisplayNames } from '../../../common/stats';

type Props = {
  selected: (SVBredPokemon | null)[];
  onSlotClick: (index: number) => void;
  onUnselect: (index: number) => void;
};

const SelectPokemonGrid: React.FC<Props> = ({ selected, onSlotClick, onUnselect }) => {
  const [hoverIndex, setHoverIndex] = useState<number | null>(null);

  return (
    <Grid container spacing={2}>
      {Array.from({ length: 6 }).map((_, i) => (
        <Grid item xs={6} key={i}>
          <Paper
            onClick={() => onSlotClick(i)}
            onMouseEnter={() => setHoverIndex(i)}
            onMouseLeave={() => setHoverIndex(null)}
            sx={{
              height: selected[i] ? 'auto' : 150,
              display: 'flex',
              justifyContent: 'center',
              alignItems: selected[i] ? 'stretch' : 'center',
              backgroundColor: selected[i] ? '#e0f2f1' : '#f9f9f9',
              cursor: 'pointer',
              position: 'relative',
            }}
          >
            {selected[i] ? (
              <Box sx={{ mt: 2, mx: 'auto', maxWidth: 600, p: 2 }}>
                {hoverIndex === i && (
                  <Box sx={{ display: 'flex', justifyContent: 'flex-end', mb: 1 }}>
                    <Button
                      variant="outlined"
                      color="error"
                      size="small"
                      onClick={(e) => {
                        e.stopPropagation();
                        onUnselect(i);
                      }}
                    >
                      Unselect
                    </Button>
                  </Box>
                )}

                <Typography variant="h4" align="center">
                  {selected[i]?.bredPokemon.name} {selected[i]?.bredPokemon.formeName && `(${selected[i]?.bredPokemon.formeName})`}
                </Typography><br />
                <Typography><strong>Gender:</strong> {selected[i]?.bredPokemon.gender.name}</Typography>
                <Typography><strong>Level:</strong> {selected[i]?.bredPokemon.level}</Typography>
                <Box sx={{ display: 'flex', alignItems: 'flex-start', mt: 0 }}>
                  <Typography sx={{ fontWeight: 'bold', mr: 1 }}>Types:</Typography>
                  <Typography>{selected[i]?.bredPokemon.types[0].name}</Typography>
                </Box>
                <Box sx={{ ml: 7 }}>
                  {selected[i]?.bredPokemon.types.slice(1).map((type, index) => (
                    <Typography key={index}>{type.name}</Typography>
                  ))}
                </Box>
                <Typography><strong>Tera Type:</strong> {selected[i]?.teraType.name}</Typography>
                <Typography><strong>Ability:</strong> {selected[i]?.bredPokemon.ability.name}</Typography>
                <Typography><strong>Nature:</strong> {selected[i]?.bredPokemon.nature.name}</Typography>
                <Typography><strong>Held Item:</strong> {selected[i]?.bredPokemon.heldItem?.name ?? 'None'}</Typography>

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
                          <TableCell>{selected[i]?.bredPokemon.baseStats[stat]}</TableCell>
                          <TableCell>{selected[i]?.bredPokemon.individualValues[stat]}</TableCell>
                          <TableCell>{selected[i]?.bredPokemon.basePoints[stat]}</TableCell>
                          <TableCell>{selected[i]?.bredPokemon.actualValues[stat]}</TableCell>
                        </TableRow>
                      ))}
                    </TableBody>
                  </Table>
                </Paper>

                <Box sx={{ display: 'flex', alignItems: 'flex-start', mt: 2 }}>
                  <Typography sx={{ fontWeight: 'bold', mr: 1 }}>Moves:</Typography>
                  <Typography>{selected[i]?.bredPokemon.moves[0].name}</Typography>
                </Box>
                <Box sx={{ ml: 7.6 }}>
                  {selected[i]?.bredPokemon.moves.slice(1).map((move, index) => (
                    <Typography key={index}>{move.name}</Typography>
                  ))}
                </Box>
                <Box sx={{ display: 'flex', alignItems: 'flex-start', mt: 1 }}>
                  <Typography sx={{ fontWeight: 'bold', mr: 1 }}>Note:</Typography>
                  <Typography sx={{ whiteSpace: 'pre-wrap' }}>
                    {selected[i]?.bredPokemon.note || ''}
                  </Typography>
                </Box>
              </Box>
            ) : (
              <Typography>{i + 1}</Typography>
            )}
          </Paper>
        </Grid>
      ))}
    </Grid>
  );
};

export default SelectPokemonGrid;