import { useEffect, useState } from 'react';
import { Container, Typography, List, ListItem } from '@mui/material';
import { useLocation } from 'react-router-dom';
import { Title } from '../../../entity/title';

const Menu = () => {
    const location = useLocation();
    const [title, setTitle] = useState<Title | null>(null);
    
    useEffect(() => {
        if (location.state && location.state.title) {
          setTitle(location.state.title);
        }
      }, [location.state]);
    
      if (!title) {
        return <div>Loading...</div>;
      }
    
    return (
        <>
            <Container maxWidth="sm" style={{ textAlign: 'center', marginTop: '50px' }}>
                <Typography variant="h4" component="h1" gutterBottom>
                    {title.title}
                </Typography>
                <List>
                    <ListItem style={{ justifyContent: 'center', backgroundColor: '#f0f0f0', margin: '10px 0', borderRadius: '8px' }}>
                        <Typography variant="body1" style={{ fontSize: '1.5rem' }}>
                            Registered Pokémon List
                        </Typography>
                    </ListItem>
                    <ListItem style={{ justifyContent: 'center', backgroundColor: '#f0f0f0', margin: '10px 0', borderRadius: '8px' }}>
                        <Typography variant="body1" style={{ fontSize: '1.5rem' }}>
                            Register Pokémon
                        </Typography>
                    </ListItem>
                    <ListItem style={{ justifyContent: 'center', backgroundColor: '#f0f0f0', margin: '10px 0', borderRadius: '8px' }}>
                        <Typography variant="body1" style={{ fontSize: '1.5rem' }}>
                            Team List
                        </Typography>
                    </ListItem>
                    <ListItem style={{ justifyContent: 'center', backgroundColor: '#f0f0f0', margin: '10px 0', borderRadius: '8px' }}>
                        <Typography variant="body1" style={{ fontSize: '1.5rem' }}>
                            Register Team
                        </Typography>
                    </ListItem>
                </List>
            </Container>
        </>
    )
}

export default Menu;