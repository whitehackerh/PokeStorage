export interface Ability {
    id: number;
    name: string;
}

export const NewAbilityFromJson = (ability: any) : Ability => {
    return {
        id: ability.id,
        name: ability.name,
    };
}