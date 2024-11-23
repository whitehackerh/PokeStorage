import { useEffect, useState } from 'react';
import { Container, Typography, List, ListItem } from '@mui/material';
import { useLocation, useNavigate } from 'react-router-dom';
import { Title } from '../../../entity/Title';

const Menu = () => {
    const location = useLocation();
    const navigate = useNavigate();
    const [title, setTitle] = useState<Title | null>(null);
    
    useEffect(() => {
        if (location.state && location.state.title) {
          setTitle(location.state.title);
        }
      }, [location.state]);

      
    const handleBredPokemonListClick = (title: Title) => {
        navigate('/bred-pokemon-list', { state: {title: title }});
    }

    const handleRegisterPokemonClick = (title: Title) => {
        navigate('/register-pokemon', { state: { title: title }});
    }
    
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
                    <ListItem onClick={() => handleBredPokemonListClick(title)} style={{ justifyContent: 'center', backgroundColor: '#f0f0f0', margin: '10px 0', borderRadius: '8px' }}>
                        <Typography variant="body1" style={{ fontSize: '1.5rem' }}>
                            Registered Pokémon List
                        </Typography>
                    </ListItem>
                    <ListItem onClick={() => handleRegisterPokemonClick(title)} style={{ justifyContent: 'center', backgroundColor: '#f0f0f0', margin: '10px 0', borderRadius: '8px' }}>
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