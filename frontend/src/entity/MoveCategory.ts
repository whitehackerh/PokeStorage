export interface MoveCategory {
    id: number;
    name: string;
}

export const NewMoveCategoryFromJson = (moveCategory: any): MoveCategory => {
    return {
        id: moveCategory.id,
        name: moveCategory.name,
    };
}