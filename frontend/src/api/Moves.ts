import { withTokenRequest, getRequestHeaders } from '../http';
import { Move } from '../entity/Move';
import { Type } from '../entity/Type';
import { MoveCategory } from '../entity/MoveCategory';

export const getMoves = async (titleId: number): Promise<Move[]> => {
    try {
        const response = await withTokenRequest.get(`/moves/${titleId}`, {
            headers: getRequestHeaders()
        });
        const moves: Move[] = response.data.data.moves.map((item: any) => {
            const type: Type = {
                id: item.type.id,
                name: item.type.name
            };
            const moveCategory: MoveCategory = {
                id: item.move_category.id,
                name: item.move_category.name
            }
            
            return {
                id: item.id,
                name: item.name,
                type: type,
                moveCategory: moveCategory,
                power: item.power,
                accuracy: item.accuracy,
            }
        });
        return moves;
    } catch (error) {
        console.log('error');
        throw error;
    }
}