export interface Nature {
    id: number;
    name: string;
}

export const NewNatureFromJson = (nature: any): Nature => {
    return {
        id: nature.id,
        name: nature.name,
    };
}