import { useEffect, useState } from 'react';
import { getTitles } from '../../../api/Titles';
import { Title } from '../../../entity/title';
import './Home.css';

const Home = () => {
    const [titles, setTitles] = useState<Title[]>([]);

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

    return (
        <>
            <div className="container">
                <h1>Select Titles</h1>
                <ul>
                    {titles.map((title) => (
                        <li key={title.id}>
                            <a href='' style={{ fontSize: '20px' }}>
                            {title.title}
                            </a>
                        </li>
                    ))}
                </ul>
            </div>
        </>
    );
};
export default Home;