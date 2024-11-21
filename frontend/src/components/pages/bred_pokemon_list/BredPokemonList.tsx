import { useEffect, useState } from 'react';
import { useLocation } from 'react-router-dom';
import { Title } from '../../../entity/Title';
import { TitleEnum } from '../../../enum/Title';
import { SVBredPokemon } from '../../../entity/SVBredPokemon';
import { getBredPokemons } from '../../../api/BredPokemons';

const BredPokemonListPage = () => {
    const location = useLocation();
    const [title, setTitle] = useState<Title | null>(null);
    const [toggleTeraType, setToggleTeraType] = useState<boolean>(false);
    const [bredPokemons, setBredPokemons] = useState<SVBredPokemon[]>([]);

    useEffect(() => {
        if (location.state && location.state.title) {
            setTitle(location.state.title);
            setToggleTeraType(location.state.title.id == TitleEnum.SV);
        }
    }, [location.state]);

    useEffect(() => {
        const fetch = async (title: Title) => {
            try {
                setBredPokemons(await getBredPokemons(title.id));
            } catch (error) {
                console.error(error);
            }
        };
        if (title) {
            fetch(title);
        }
    }, [title]);

    if (!title) {
        return <div>Loading...</div>;
    }

    return (
        <></>
    )
}

export default BredPokemonListPage;