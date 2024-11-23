export interface Type {
    id: number;
    name: string;
}

export const NewTypeFromJson = (type: any): Type => {
    return {
        id: type.id,
        name: type.name,
    };
}