import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Container, Typography, List, ListItem } from '@mui/material';
import { getTitles } from '../../../api/Titles';
import { Title } from '../../../entity/Title';

const Home = () => {
    const [titles, setTitles] = useState<Title[]>([]);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchTitles = async () => {
            try {
                setTitles(await getTitles()); 
            } catch (error) {
                console.error(error);
            }
        };
        fetchTitles();
    }, []);

    const handleItemClick = (title: Title) => {
        navigate('/menu', { state: { title: title }});
    }

    return (
        <>
            <Container maxWidth="sm" style={{ textAlign: 'center', marginTop: '50px' }}>
                <Typography variant="h4" component="h1" gutterBottom>
                    Select Titles
                </Typography>
                <List>
                    {titles.map((title) => (
                        <ListItem onClick={() => handleItemClick(title)} key={title.id} style={{ justifyContent: 'center', backgroundColor: '#f0f0f0', margin: '10px 0', borderRadius: '8px' }}>
                            <Typography variant="body1" style={{ fontSize: '1.5rem' }}>
                                {title.title}
                            </Typography>
                        </ListItem>
                    ))}
                </List>
            </Container>
        </>
    );
};
export default Home;