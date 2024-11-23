import { withTokenRequest, getRequestHeaders } from '../http';
import { Move, NewMoveFromJson } from '../entity/Move';

export const getMoves = async (): Promise<Move[]> => {
    try {
        const response = await withTokenRequest.get('/sv/moves', {
            headers: getRequestHeaders()
        });
        const moves: Move[] = response.data.data.moves.map((item: any) => {
            return NewMoveFromJson(item);
        });
        return moves;
    } catch (error) {
        console.log('error');
        throw error;
    }
}