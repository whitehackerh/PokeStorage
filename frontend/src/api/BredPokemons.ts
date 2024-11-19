import { withTokenRequest, getRequestHeaders } from '../http';

export const postBredPokemons = async (titleId: number, bredPokemon: any) : Promise<void> => {
    try {
        await withTokenRequest.post(`/bred-pokemons/${titleId}`, bredPokemon, {
            headers: getRequestHeaders()
        });
    } catch (error) {
        console.log('error');
        throw error;
    }
}