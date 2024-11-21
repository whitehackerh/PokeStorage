export interface TeraType {
    id: number;
    name: string;
}

export const NewTeraTypeFromJson = (teraType: any): TeraType => {
    return {
        id: teraType.id,
        name: teraType.name,
    };
}