import { Type } from './Type';
import { MoveCategory } from './MoveCategory';

export interface Move {
    id: number;
    name: string;
    type: Type;
    moveCategory: MoveCategory;
    power: number;
    accuracy: number;
}