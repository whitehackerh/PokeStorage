import { NewTypeFromJson, Type } from './Type';
import { MoveCategory, NewMoveCategoryFromJson } from './MoveCategory';

export interface Move {
    id: number;
    name: string;
    type: Type;
    moveCategory: MoveCategory;
    power: number;
    accuracy: number;
}

export const NewMoveFromJson = (move: any): Move => {
    return {
        id: move.id,
        name: move.name,
        type: NewTypeFromJson(move.type),
        moveCategory: NewMoveCategoryFromJson(move.move_category),
        power: move.power,
        accuracy: move.accuracy,
    };
}