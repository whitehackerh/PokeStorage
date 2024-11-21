export interface Gender {
    id: number;
    name: string;
}

export const NewGenderFromJson = (gender: any): Gender => {
    return {
        id: gender.id,
        name: gender.name,
    };
}